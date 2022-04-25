package enum

import (
	"fmt"
	"testing"
)

func TestEnum(t *testing.T) {
	tests := []struct {
		game   GameType
		expect int
	}{
		{game: PS4, expect: 1},
		{game: PS5, expect: 2},
		{game: Wii, expect: 3},
		{game: WiiU, expect: 4},
		{game: NintendoSwitch, expect: 5},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("%v", i+1), func(t *testing.T) {
			if int(tt.game) != tt.expect {
				t.Errorf("actually = %v, expect = %v", tt.game, tt.expect)
			}
		})
	}
}

func TestEnumString(t *testing.T) {
	tests := []struct {
		game   GameType
		expect string
	}{
		{game: PS4, expect: "PS4"},
		{game: PS5, expect: "PS5"},
		{game: Wii, expect: "Wii"},
		{game: WiiU, expect: "WiiU"},
		{game: NintendoSwitch, expect: "NintendoSwitch"},
	}

	for _, tt := range tests {
		t.Run(tt.expect, func(t *testing.T) {
			if tt.game.String() != tt.expect {
				t.Errorf("actually = %v, expect = %v", tt.game, tt.expect)
			}
		})
	}
}
