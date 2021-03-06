package parser

import (
	"gostudy/crawler/engine"
	"regexp"
)

const cityRe = `<a href="(http://album.zhenai.com/u/[0-9]+)" [^>]*>([^<]+)</a></th></tr> <tr><td width="180"><span class="grayL">性别：</span>([^<]+)</td>`

func ParseCity(contents []byte) engine.ParserResult {
	re := regexp.MustCompile(cityRe)
	match := re.FindAllSubmatch(contents, -1)
	result := engine.ParserResult{}
	for _, m := range match {
		name := string(m[2])
		gender := string(m[3])
		result.Request = append(result.Request, engine.Request{
			Url: string(m[1]),
			ParserFunc: func(contents []byte) engine.ParserResult {
				param := make(map[string]string)
				param["name"] = name
				param["gender"] = gender
				return ParseProfile(contents, param)
			},
		})
		result.Items = append(result.Items, "User "+string(m[2]))
	}
	return result
}
