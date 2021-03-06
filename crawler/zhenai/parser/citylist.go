package parser

import (
	"gostudy/crawler/engine"
	"regexp"
)

const cityListRe = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`

func ParseCityList(contents []byte) engine.ParserResult {
	re := regexp.MustCompile(cityListRe)
	match := re.FindAllSubmatch(contents, -1)
	result := engine.ParserResult{}
	limit := 10
	for _, m := range match {
		result.Request = append(result.Request, engine.Request{
			Url:        string(m[1]),
			ParserFunc: ParseCity,
		})
		result.Items = append(result.Items, "City "+string(m[2]))
		limit--
		if limit == 0 {
			break
		}
	}
	return result
}
