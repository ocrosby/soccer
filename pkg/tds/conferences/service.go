package conferences

import (
	"github.com/gocolly/colly"
	"github.com/ocrosby/soccer/pkg/common"
	"github.com/ocrosby/soccer/pkg/tds"
)

type ConferenceServicer interface {
	ConferencesUrl(division common.Division) (string, error)
	Conferences() ([]tds.Conference, error)
	ConferencesByGenderAndDivision(gender common.Gender, division common.Division) ([]tds.Conference, error)
}

type ConferenceService struct {
}

func NewConferenceService() *ConferenceService {
	return &ConferenceService{}
}

func (s ConferenceService) ConferencesUrl(division common.Division) (string, error) {
	switch division {
	case common.DI:
		return "https://www.topdrawersoccer.com/college-soccer/college-conferences/di/divisionid-1", nil
	case common.DII:
		return "https://www.topdrawersoccer.com/college-soccer/college-conferences/dii/divisionid-2", nil
	case common.DIII:
		return "https://www.topdrawersoccer.com/college-soccer/college-conferences/diii/divisionid-3", nil
	case common.NAIA:
		return "https://www.topdrawersoccer.com/college-soccer/college-conferences/naia/divisionid-4", nil
	case common.NJCAA:
		return "https://www.topdrawersoccer.com/college-soccer/college-conferences/njcaa/divisionid-5", nil
	case common.TestDivision:
		return "https://some/test/url", nil
	default:
		return "", common.ErrInvalidDivision
	}
}

func (s ConferenceService) Conferences() ([]tds.Conference, error) {
	var (
		confs []tds.Conference
		err   error
	)

	genders := []common.Gender{common.Male, common.Female}
	divisions := []common.Division{common.DI, common.DII, common.DIII, common.NAIA, common.NJCAA}
	conferences := make([]tds.Conference, 0)

	for _, gender := range genders {
		for _, division := range divisions {
			confs, err = s.ConferencesByGenderAndDivision(gender, division)
			if err != nil {
				return nil, err
			}

			conferences = append(conferences, confs...)
		}
	}

	return conferences, nil
}

func (s ConferenceService) ConferencesByGenderAndDivision(gender common.Gender, division common.Division) ([]tds.Conference, error) {
	var (
		url         string
		err         error
		conferences []tds.Conference
	)

	collector := colly.NewCollector()

	collector.OnHTML("table > tbody", func(h *colly.HTMLElement) {
		h.ForEach("tr", func(_ int, el *colly.HTMLElement) {
			url = tds.LinkFromCell(el, 1)
			currentGender := tds.GenderFromUrl(url)

			if gender == common.UnspecifiedGender || gender == currentGender {
				conference := tds.Conference{
					ID:       tds.IdentifierFromUrl(url),
					Gender:   currentGender,
					Division: division,
					Name:     tds.TextFromCell(el, 1),
					Url:      tds.LinkFromCell(el, 1),
				}

				conferences = append(conferences, conference)
			}
		})
	})

	if url, err = s.ConferencesUrl(division); err != nil {
		return nil, err
	}

	if err = collector.Visit(url); err != nil {
		return nil, err
	}

	return conferences, nil
}
