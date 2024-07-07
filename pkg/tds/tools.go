package tds

import (
	"fmt"
	"github.com/gocolly/colly"
	"github.com/ocrosby/soccer/pkg/common"
	"strconv"
	"strings"
)

func linkFromCell(element *colly.HTMLElement, index int) string {
	var link string

	prefix := "http://www.topdrawersoccer.com"

	goquerySelector := fmt.Sprintf("td:nth-child(%d) > a", index)

	links := element.ChildAttrs(goquerySelector, "href")
	if len(links) > 0 {
		link = prefix + links[0]
	} else {
		link = ""
	}

	return link
}

func normalizeText(text string) string {
	text = strings.ReplaceAll(text, "\u00a0", " ")
	text = strings.Trim(text, " ")

	return text
}

func textFromCell(element *colly.HTMLElement, index int) string {
	goquerySelector := fmt.Sprintf("td:nth-child(%d)", index)

	text := element.ChildText(goquerySelector)
	text = normalizeText(text)

	return text
}

func idFromUrl(url string) int {
	index := strings.LastIndex(url, "-")

	if index < 0 {
		return -1
	}

	suffix := url[index+1:]
	id, err := strconv.Atoi(suffix)

	if err != nil {
		return -1
	}

	return id
}

func genderFromUrl(url string) common.Gender {
	if strings.Contains(url, "/men/") {
		return common.Male
	} else if strings.Contains(url, "/women/") {
		return common.Female
	} else {
		return common.UnspecifiedGender
	}
}
