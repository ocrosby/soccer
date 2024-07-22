package tds

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
	"github.com/ocrosby/soccer/internal/tds-coaching-change-service/models"
	"github.com/ocrosby/soccer/pkg/common"
	"net/url"
	"strings"
)

func createChange(s *goquery.Selection, gender common.Gender) (models.Change, error) {
	change := models.Change{}

	if s == nil {
		return change, fmt.Errorf("selection is nil")
	}

	change.Program = strings.TrimSpace(s.Find("td").Eq(0).Text())
	change.ProgramUrl = s.Find("td").Eq(0).Find("a").AttrOr("href", "")
	change.ProgramId = IdentifierFromUrl(change.ProgramUrl)
	change.DepartingCoach = strings.TrimSpace(s.Find("td").Eq(1).Text())
	change.DepartingCoachUrl = s.Find("td").Eq(1).Find("a").AttrOr("href", "")
	change.NewCoach = strings.TrimSpace(s.Find("td").Eq(2).Text())
	change.NewCoachUrl = s.Find("td").Eq(2).Find("a").AttrOr("href", "")
	change.Gender = gender

	// Prepend the base URL to the program URL
	if len(change.ProgramUrl) > 0 {
		baseURL, err := url.Parse("https://www.topdrawersoccer.com")
		if err != nil {
			return change, fmt.Errorf("error parsing base URL: %v", err)
		}
		programURL, err := url.Parse(change.ProgramUrl)
		if err != nil {
			return change, fmt.Errorf("error parsing program URL: %v", err)
		}
		resolvedURL := baseURL.ResolveReference(programURL)
		change.ProgramUrl = resolvedURL.String()
	}

	return change, nil
}

func GetMaleCoachingChanges(url string) ([]models.Change, error) {
	var (
		err     error
		change  models.Change
		changes []models.Change
	)

	c := colly.NewCollector()

	// Handle the MEN's table
	c.OnHTML("div.col:contains('MEN\\'S')", func(e *colly.HTMLElement) {
		e.DOM.Find("table").First().Find("tr").Each(func(i int, s *goquery.Selection) {
			if i == 0 {
				return
			}

			if change, err = createChange(s, common.Male); err != nil {
				fmt.Printf("Error creating change: %v\n", err)
				return
			}

			if len(change.Program) > 0 {
				changes = append(changes, change)
			}
		})
	})

	if err = c.Visit(url); err != nil {
		return nil, err
	}

	return changes, nil
}

func GetFemaleCoachingChanges(url string) ([]models.Change, error) {
	var (
		err     error
		change  models.Change
		changes []models.Change
	)

	c := colly.NewCollector()

	// Handle the WOMEN's table
	c.OnHTML("div.col:contains('WOMEN\\'S')", func(e *colly.HTMLElement) {
		e.DOM.Find("table").First().Find("tr").Each(func(i int, s *goquery.Selection) {
			if i == 0 {
				return
			}

			if change, err = createChange(s, common.Female); err != nil {
				fmt.Printf("Error creating change: %v\n", err)
				return
			}

			if len(change.Program) > 0 {
				changes = append(changes, change)
			}
		})
	})

	if err = c.Visit(url); err != nil {
		return nil, err
	}

	return changes, nil
}

func GetAllCoachingChanges(url string) ([]models.Change, error) {
	var (
		err     error
		change  models.Change
		changes []models.Change
	)

	c := colly.NewCollector()

	// Handle the MEN's table
	c.OnHTML("div.col:contains('MEN\\'S')", func(e *colly.HTMLElement) {
		e.DOM.Find("table").First().Find("tr").Each(func(i int, s *goquery.Selection) {
			if i == 0 {
				return
			}

			if change, err = createChange(s, common.Male); err != nil {
				fmt.Printf("Error creating change: %v\n", err)
				return
			}

			if len(change.Program) > 0 {
				changes = append(changes, change)
			}
		})
	})

	// Handle the WOMEN's table
	c.OnHTML("div.col:contains('WOMEN\\'S')", func(e *colly.HTMLElement) {
		e.DOM.Find("table").First().Find("tr").Each(func(i int, s *goquery.Selection) {
			if i == 0 {
				return
			}

			if change, err = createChange(s, common.Female); err != nil {
				fmt.Printf("Error creating change: %v\n", err)
				return
			}

			if len(change.Program) > 0 {
				changes = append(changes, change)
			}
		})
	})

	if err = c.Visit(url); err != nil {
		return nil, err
	}

	return changes, nil
}
