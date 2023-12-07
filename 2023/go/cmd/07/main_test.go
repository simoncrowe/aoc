package main

import "testing"

func TestRankTypeJoker(t *testing.T) {
	testCases := []struct {
		name             string
		cards            string
		expectedTypeRank int
	}{
		{"FiveJokers", "JJJJJ", 7},
		{"ThreeJokersMixed", "JJJ23", 6},
		{"ThreeJokersTwoOfKind", "JJJQQ", 7},
		{"TwoJokersThreeOfKind", "JJ222", 7},
		{"TwoJokersMixed", "JJ235", 4},
		{"TwoJokersWithPair", "JJ232", 5},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := RankTypeJoker(tc.cards)
			if result != tc.expectedTypeRank {
				t.Errorf("%s: expected %s to get rank %d. Got %d", tc.name, tc.cards, tc.expectedTypeRank, result)
			}
		})
	}
}
