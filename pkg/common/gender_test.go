package common

import "testing"

func TestStringToGender(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    Gender
		wantErr bool
	}{
		{"Empty", "", UnspecifiedGender, false},
		{"Male Lowercase", "male", Male, false},
		{"Male Uppercase", "MALE", Male, false},
		{"Female Lowercase", "female", Female, false},
		{"Female Uppercase", "FEMALE", Female, false},
		{"Unspecified Lowercase", "unspecified", UnspecifiedGender, false},
		{"Unspecified Uppercase", "UNSPECIFIED", UnspecifiedGender, false},
		{"Invalid", "notAGender", 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := StringToGender(tt.input)

			if (err != nil) != tt.wantErr {
				t.Errorf("StringToGender() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if got != tt.want {
				t.Errorf("StringToGender() = %v, want %v", got, tt.want)
			}
		})
	}
}
