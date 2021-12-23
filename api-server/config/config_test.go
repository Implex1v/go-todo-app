package config

import (
	"os"
	"testing"
)

func TestGetConfig(t *testing.T) {
	inputs := []string{
		"5000",
		"",
	}

	expected := []Config{
		{Port: "5000"},
		{Port: "8080"},
	}

	for i, input := range inputs {
		if i == 1 {
			os.Clearenv()
		} else {
			e := os.Setenv("PORT", input)
			if e != nil {
				t.Errorf("Failed to set env '%s' to '%s", e, input)
			}
		}

		got := GetConfig()
		expect := expected[i].Port
		if got.Port != expect {
			t.Errorf("Got '%s' expected '%s'", got, expect)
		}
	}
}
