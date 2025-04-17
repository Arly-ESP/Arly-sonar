package enums

import "testing"

func TestMoodEnumValues(t *testing.T) {
	tests := []struct {
		name     string
		mood     MoodEnum
		expected string
	}{
		{"TestBad", Bad, "bad"},
		{"TestSad", Sad, "sad"},
		{"TestPoker", Poker, "poker"},
		{"TestNice", Nice, "nice"},
		{"TestHappy", Happy, "happy"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if string(tt.mood) != tt.expected {
				t.Errorf("Expected %s, got %s", tt.expected, tt.mood)
			}
		})
	}
}
