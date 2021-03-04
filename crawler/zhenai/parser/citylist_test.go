package parser

import (
	"io/ioutil"
	"testing"
)

func TestParseCityList(t *testing.T) {
	contents, err := ioutil.ReadFile("citylist_test_data.html")
	if err != nil {
		panic(err)
	}
	result := ParseCityList(contents)

	const resultSize = 470
	expectedUrl := []string{
		"http://www.zhenai.com/zhenghun/aba", "http://www.zhenai.com/zhenghun/akesu", "http://www.zhenai.com/zhenghun/alashanmeng",
	}
	expectedCity := []string{
		"阿坝", "阿克苏", "阿拉善盟",
	}
	if len(result.Request) != resultSize {
		t.Errorf("result should have %d requests; but had %d", resultSize, len(result.Request))
	}
	if len(result.Items) != resultSize {
		t.Errorf("result should have %d requests; but had %d", resultSize, len(result.Items))
	}
	for i, url := range expectedUrl {
		if result.Request[i].Url != url {
			t.Errorf("expected url #%d: %s; but was %s ", i, result.Request[i].Url, url)
		}
	}
	for i, city := range expectedCity {
		if result.Items[i] != city {
			t.Errorf("expected url #%d: %s; but was %s", i, result.Items[i], city)
		}
	}
}
