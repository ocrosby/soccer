package tds

import (
	"errors"
	"github.com/gocolly/colly"
	"github.com/ocrosby/soccer/pkg/common"
)

type Conference struct {
	ID       int             `json:"id"`
	Name     string          `json:"name"`
	Url      string          `json:"url"`
	Division common.Division `json:"division"`
	Gender   common.Gender   `json:"gender"`
}

// Division URLs mapped to their unique parts
var divisionURLMap = map[common.Division]string{
	common.DI:    "di/divisionid-1",
	common.DII:   "dii/divisionid-2",
	common.DIII:  "diii/divisionid-3",
	common.NAIA:  "naia/divisionid-4",
	common.NJCAA: "njcaa/divisionid-5",
}

// conferencesUrl returns the URL for the conferences page for a given division.
func conferencesUrl(division common.Division) (string, error) {
	base := "https://www.topdrawersoccer.com/college-soccer/college-conferences/"
	if path, ok := divisionURLMap[division]; ok {
		return base + path, nil
	}
	return "", common.ErrInvalidDivision
}

func conferencesFromUrl(url string, division common.Division) ([]Conference, error) {
	var (
		err         error
		conferences []Conference
	)

	collector := colly.NewCollector()

	collector.OnHTML("table > tbody", func(h *colly.HTMLElement) {
		h.ForEach("tr", func(_ int, el *colly.HTMLElement) {
			conference := Conference{
				Division: division,
				Name:     textFromCell(el, 1),
				Url:      linkFromCell(el, 1),
			}

			conference.ID = idFromUrl(conference.Url)
			conference.Gender = genderFromUrl(conference.Url)

			conferences = append(conferences, conference)
		})
	})

	if err = collector.Visit(url); err != nil {
		return nil, err
	}

	return conferences, nil
}

func maleConferencesFromUrl(url string) ([]Conference, error) {
	return nil, errors.New("not implemented")
}

func femaleConferencesFromUrl(url string) ([]Conference, error) {
	return nil, errors.New("not implemented")
}

// conferencesAll returns all conferences for a given division
func conferencesAll(division common.Division) ([]Conference, error) {
	var (
		url                             string
		err                             error
		conferences, currentConferences []Conference
	)

	if division != common.UnspecifiedDivision {
		// If the division is specified return all conferences for that division.
		// Get the URL for the division
		url, err = conferencesUrl(division)
		if err != nil {
			return nil, err
		}

		// Get the conferences for the division
		return conferencesFromUrl(url, division)
	}

	// The division is unspecified return all conferences for all divisions.
	// Get the URL for each division

	for division := common.DI; division <= common.NJCAA; division++ {
		url, err = conferencesUrl(common.Division(division))
		if err != nil {
			return nil, err
		}

		// Get the conferences for each division
		currentConferences, err = conferencesFromUrl(url, common.Division(division))
		if err != nil {
			return nil, err
		}

		// Combine the results
		conferences = append(conferences, currentConferences...)
	}

	return conferences, nil
}

// conferencesMale returns all male conferences for a given division
func conferencesMale(division common.Division) ([]Conference, error) {
	return conferencesSelectedGender(division, common.Male)
}

// conferencesFemale returns all female conferences for a given division
func conferencesFemale(division common.Division) ([]Conference, error) {
	return conferencesSelectedGender(division, common.Female)
}

func conferencesSelectedGender(division common.Division, targetedGender common.Gender) ([]Conference, error) {
	var (
		err                 error
		selectedConferences []Conference
		allConferences      []Conference
	)

	allConferences, err = conferencesAll(division)
	if err != nil {
		return nil, err
	}

	for _, conference := range allConferences {
		if conference.Gender == targetedGender || targetedGender == common.UnspecifiedGender {
			selectedConferences = append(selectedConferences, conference)
		}
	}

	return selectedConferences, nil
}

// Conferences returns all conferences for a given targetedGender and division
func Conferences(targetedGender common.Gender, division common.Division) ([]Conference, error) {
	if targetedGender == common.UnspecifiedGender {
		return conferencesAll(division)
	}

	if targetedGender == common.Male {
		return conferencesMale(division)
	}

	if targetedGender == common.Female {
		return conferencesFemale(division)
	}

	return nil, common.ErrInvalidGender
}
