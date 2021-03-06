package parser

import (
	"gostudy/crawler/engine"
	"gostudy/crawler/model"
	"regexp"
	"strconv"
)

var ageRe = regexp.MustCompile(`<div class="m-btn purple" data-v-8b1eac0c>([0-9]+)岁</div>`)
var heightRe = regexp.MustCompile(`<div class="m-btn purple" data-v-8b1eac0c>([0-9]+)cm</div>`)
var incomeRe = regexp.MustCompile(`<div class="m-btn purple" data-v-8b1eac0c>(月收入:[^<]+)</div>`)
var marriageRe = regexp.MustCompile(`<div class="purple-btns" data-v-8b1eac0c><div class="m-btn purple" data-v-8b1eac0c>([未婚|离异])</div>`)
var educationRe = regexp.MustCompile(`<div class="m-btn purple" data-v-8b1eac0c>([^<]+)</div></div>`)
var hokouRe = regexp.MustCompile(`<div class="m-btn pink" data-v-8b1eac0c>籍贯:([^<]+)</div>`)
var xinzuoRe = regexp.MustCompile(`<div class="m-btn purple" data-v-8b1eac0c>(.+座)[^<]+<`)
var houseRe = regexp.MustCompile(`<div class="m-btn pink" data-v-8b1eac0c>(.+房)[^<]+<`)
var carRe = regexp.MustCompile(`<div class="m-btn pink" data-v-8b1eac0c>(.+车)[^<]+<`)

func ParseProfile(contents []byte, param map[string]string) engine.ParserResult {
	profile := model.Profile{}
	age, err := strconv.Atoi(extractString(contents, ageRe))
	if err == nil {
		profile.Age = age
	}
	height, err := strconv.Atoi(extractString(contents, heightRe))
	if err == nil {
		profile.Height = height
	}
	income := extractString(contents, incomeRe)
	profile.Income = income
	profile.Marriage = extractString(contents, marriageRe)
	profile.Education = extractString(contents, educationRe)
	profile.Hokou = extractString(contents, hokouRe)
	profile.Xinzuo = extractString(contents, xinzuoRe)
	profile.House = extractString(contents, houseRe)
	profile.Car = extractString(contents, carRe)
	profile.Name = param["name"]
	profile.Gender = param["gender"]
	result := engine.ParserResult{
		Items: []interface{}{profile},
	}
	return result
}

func extractString(contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)
	if len(match) >= 2 {
		return string(match[1])
	} else {
		return ""
	}
}
