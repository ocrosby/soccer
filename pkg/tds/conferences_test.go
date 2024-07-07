package tds

import (
	"github.com/ocrosby/soccer/pkg/common"
	"testing"
)

func TestConferencesInvalidGender(t *testing.T) {
	t.Run("TestConferences", func(t *testing.T) {
		gotConferences, err := Conferences(999, common.DI)
		if err == nil {
			t.Fatal("expected error, got nil")
		}

		if gotConferences != nil {
			t.Fatalf("expected nil, got %v", gotConferences)
		}
	})
}

func TestConferencesFemaleDI(t *testing.T) {
	t.Run("TestConferences", func(t *testing.T) {
		gotConferences, err := Conferences(common.Female, common.DI)

		if err != nil {
			t.Fatalf("expected nil error, got %v", err)
		}

		// Loop over all the conferences and check that they are all female and DI
		for _, conference := range gotConferences {
			if conference.Gender != common.Female {
				t.Fatalf("expected all Female conferences, got a male one")
			}

			if conference.Division != common.DI {
				t.Fatalf("expected DI, got %v", conference.Division)
			}
		}
	})
}

func TestConferencesFemaleDII(t *testing.T) {
	t.Run("TestConferences", func(t *testing.T) {
		gotConferences, err := Conferences(common.Female, common.DII)

		if err != nil {
			t.Fatalf("expected nil error, got %v", err)
		}

		// Loop over all the conferences and check that they are all female and DI
		for _, conference := range gotConferences {
			if conference.Gender != common.Female {
				t.Fatalf("expected all Female conferences, got a male one")
			}

			if conference.Division != common.DII {
				t.Fatalf("expected DII, got %v", conference.Division)
			}
		}
	})
}

func TestConferencesMaleDII(t *testing.T) {
	t.Run("TestConferences", func(t *testing.T) {
		gotConferences, err := Conferences(common.Male, common.DII)

		if err != nil {
			t.Fatalf("expected nil error, got %v", err)
		}

		// Loop over all the conferences and check that they are all female and DI
		for _, conference := range gotConferences {
			if conference.Gender != common.Male {
				t.Fatalf("expected all Male conferences, got a female one")
			}

			if conference.Division != common.DII {
				t.Fatalf("expected DII, got %v", conference.Division)
			}
		}
	})
}
