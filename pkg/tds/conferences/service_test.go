package conferences

import (
	"github.com/ocrosby/soccer/pkg/common"
	"testing"
)

func TestNewConferenceService(t *testing.T) {
	// Act
	service := NewConferenceService()

	// Assert
	if service == nil {
		t.Error("Expected NewConferenceService to return a non-nil value")
	}
}

func TestConferenceService_ConferencesURL(t *testing.T) {
	service := NewConferenceService()
	testCases := []struct {
		name     string
		division common.Division
		wantURL  string
		wantErr  bool
	}{
		{
			name:     "DI",
			division: common.DI,
			wantURL:  "https://www.topdrawersoccer.com/college-soccer/college-conferences/di/divisionid-1",
			wantErr:  false,
		},
		{
			name:     "DII",
			division: common.DII,
			wantURL:  "https://www.topdrawersoccer.com/college-soccer/college-conferences/dii/divisionid-2",
			wantErr:  false,
		},
		{
			name:     "DIII",
			division: common.DIII,
			wantURL:  "https://www.topdrawersoccer.com/college-soccer/college-conferences/diii/divisionid-3",
			wantErr:  false,
		},
		{
			name:     "NAIA",
			division: common.NAIA,
			wantURL:  "https://www.topdrawersoccer.com/college-soccer/college-conferences/naia/divisionid-4",
			wantErr:  false,
		},
		{
			name:     "NJCAA",
			division: common.NJCAA,
			wantURL:  "https://www.topdrawersoccer.com/college-soccer/college-conferences/njcaa/divisionid-5",
			wantErr:  false,
		},
		{
			name:     "Unspecified",
			division: common.UnspecifiedDivision,
			wantURL:  "",
			wantErr:  true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			gotURL, err := service.ConferencesUrl(tc.division)
			if (err != nil) != tc.wantErr {
				t.Errorf("ConferenceService.ConferencesUrl() error = '%v', wantErr %v", err, tc.wantErr)
				return
			}
			if gotURL != tc.wantURL {
				t.Errorf("ConferenceService.ConferencesUrl() = '%v', want %v", gotURL, tc.wantURL)
			}
		})
	}
}

func TestConferenceService_ConferencesByGenderAndDivision(t *testing.T) {
	// Arrange
	service := NewConferenceService()
	testCases := []struct {
		name       string
		gender     common.Gender
		division   common.Division
		wantCount  int
		wantErr    bool
		wantGotNil bool
	}{
		// Add test cases here
		{
			name:       "Male DI",
			gender:     common.Male,
			division:   common.DI,
			wantCount:  25,
			wantErr:    false,
			wantGotNil: false,
		},
		{
			name:       "Male DII",
			gender:     common.Male,
			division:   common.DII,
			wantCount:  23,
			wantErr:    false,
			wantGotNil: false,
		},
		{
			name:       "Male DIII",
			gender:     common.Male,
			division:   common.DIII,
			wantCount:  49,
			wantErr:    false,
			wantGotNil: false,
		},
		{
			name:       "Male NAIA",
			gender:     common.Male,
			division:   common.NAIA,
			wantCount:  29,
			wantErr:    false,
			wantGotNil: false,
		},
		{
			name:       "Male NJCAA",
			gender:     common.Male,
			division:   common.NAIA,
			wantCount:  29,
			wantErr:    false,
			wantGotNil: false,
		},
		{
			name:       "Female DI",
			gender:     common.Female,
			division:   common.DI,
			wantCount:  31,
			wantErr:    false,
			wantGotNil: false,
		},
		{
			name:       "Female DII",
			gender:     common.Female,
			division:   common.DII,
			wantCount:  24,
			wantErr:    false,
			wantGotNil: false,
		},
		{
			name:       "Female DIII",
			gender:     common.Female,
			division:   common.DIII,
			wantCount:  50,
			wantErr:    false,
			wantGotNil: false,
		},
		{
			name:       "Female NAIA",
			gender:     common.Female,
			division:   common.NAIA,
			wantCount:  29,
			wantErr:    false,
			wantGotNil: false,
		},
		{
			name:       "Female NJCAA",
			gender:     common.Female,
			division:   common.NAIA,
			wantCount:  29,
			wantErr:    false,
			wantGotNil: false,
		},
		{
			name:       "Female Unspecified Division",
			gender:     common.Female,
			division:   common.UnspecifiedDivision,
			wantCount:  0,
			wantErr:    true,
			wantGotNil: true,
		},
		{
			name:       "Male Test Division",
			gender:     common.Male,
			division:   common.TestDivision,
			wantCount:  0,
			wantErr:    true,
			wantGotNil: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Act
			got, err := service.ConferencesByGenderAndDivision(tc.gender, tc.division)

			// Assert
			if (err != nil) != tc.wantErr {
				t.Errorf("ConferenceService.Conferences() error = %v, wantErr %v", err, tc.wantErr)
				return
			}

			if (got == nil) != tc.wantGotNil {
				if tc.wantGotNil {
					t.Errorf("Expected got to be nil")
				} else {
					t.Errorf("Expected got to be non-nil")
				}
			}

			if len(got) != tc.wantCount {
				t.Errorf("ConferenceService.Conferences() count = %v, want %v", len(got), tc.wantCount)
			}
		})
	}
}

func TestConferenceService_Conferences(t *testing.T) {
	// Arrange
	service := NewConferenceService()

	// Act
	got, err := service.Conferences()

	// Assert
	if err != nil {
		t.Errorf("ConferenceService.Conferences() error = %v, wantErr %v", err, false)
		return
	}

	if got == nil {
		t.Errorf("Expected got to be non-nil")
	}

	if len(got) != 377 {
		t.Errorf("ConferenceService.Conferences() count = %v, want %v", len(got), 156)
	}
}
