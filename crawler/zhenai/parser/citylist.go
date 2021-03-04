package parser

import (
	"gostudy/crawler/engine"
	"regexp"
)

const CityListRe = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`

func ParseCityList(contents []byte) engine.ParserResult {
	re := regexp.MustCompile(CityListRe)
	match := re.FindAllSubmatch(contents, -1)
	result := engine.ParserResult{}
	for _, m := range match {
		result.Request = append(result.Request, engine.Request{
			Url:        string(m[1]),
			ParserFunc: engine.NilParser,
		})
		result.Items = append(result.Items, string(m[2]))
	}
	return result
}
