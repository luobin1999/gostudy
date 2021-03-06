package fetcher

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

const cookie = "sid=d1a184f7-dd46-435d-be7d-4cb4a7a14dd8; ec=NgAIBngW-1612421170416-d79633c0ed328-1989762046; FSSBBIl1UgzbN7NO=5F8y6b7n6eUrp6nAcRTm5oF6dKNCKa7jPS9qZJjDWFe4Y1CytAIlN95FG99.hssObFFPGEad34qaM6qz2UFk_mA; spliterabparams=1612421250223%3A-5347231936326107893; refreshToken=1833764262.1614849856213.673f52cc410fc43bf187902fc029ecce; _pc_register_validate_sec=58; recommendId=%7B%22main-flow%22%3A%22fm-v16%22%2C%22off-feat%22%3A%22v1%22%2C%22feat-config%22%3A%22v1%22%2C%22model-version%22%3A%22v16%22%7D; Hm_lvt_2c8ad67df9e787ad29dbd54ee608f5d2=1614762228,1614842554,1615005683; Hm_lpvt_2c8ad67df9e787ad29dbd54ee608f5d2=1615014862; _efmdata=OmG0BnU6ghWuFYpHAk%2FBvJNMCj8Yb6RfMzcO%2FbZ3%2BU6wuq4733tFgZejjqa2Uxfnsl%2FZTHBFKOBo%2BCD7MFE6HT4qmQusubO%2B%2F6lHgCnOFFs%3D; _exid=zIHn6rgTBhR5UHQ9vPvuoSWQN9BQdQ3RIJI5qZMfSGV%2FTWkapemZwGESJeMRa0NUiC7oMtBVsj9NBTVj4wTm%2Fw%3D%3D; FSSBBIl1UgzbN7NP=53lW.0brktZ0qqqmgrbY4LA_.946q4ChC45frg05CZ9mRCo5.eNhcNjPJxEY1OpNupWxqRAI2YS0quE2YPGEdwYyDg9Eaqf9QAxgtiaOhOSz5mfsQ_bBkjnW.3w7Kt3RV9c.iVJPzPUiVR1SYjphC3gPMRMjgZoERGqvjUWvdzC6vK7NdXq1gx53GI6zlth64dItfcup91TcVUvWhPcrR.DekdAUWfTWlNdoac1HQ5hapslkeL9kK2GPNUYH7bscw9"

func Fetch(url string) ([]byte, error) {
	client := &http.Client{}
	newUrl := strings.Replace(url, "http://", "https://", 1)
	log.Printf("Replace Url: %s", newUrl)
	request, err := http.NewRequest("GET", newUrl, nil)
	if err != nil {
		return nil, err
	}
	request.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.111 Safari/537.36")
	request.Header.Add("cookie", cookie)
	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	//自动检测编码转成utf8
	e := determineEncoding(resp.Body)
	utf8Reader := transform.NewReader(resp.Body, e.NewDecoder())

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("wrong status code: %d", resp.StatusCode)
	}

	return ioutil.ReadAll(utf8Reader)
}

func determineEncoding(r io.Reader) encoding.Encoding {
	//如果直接使用resp.Body只能读一次，因此使用bufio.NewReader().Peek来读取
	bytes, err := bufio.NewReader(r).Peek(1024)
	if err != nil {
		log.Panicf("Fetcher error: %v", err)
		return unicode.UTF8
	}
	//根据1024字节的数据判断编码返回
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}
