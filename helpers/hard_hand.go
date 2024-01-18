package helpers

import "github.com/thoas/go-funk"

func PlayHardHand(playerCards []int, dealerCard int) string {
	playerCardTotal := funk.SumInt(playerCards)

	if playerCardTotal <= 8 {
		return HIT
	}
	if playerCardTotal == 9 {
		if dealerCard >= 3 && dealerCard <= 6 {
			return DOUBLE
		}
		return HIT
	}

	if playerCardTotal == 10 {
		if dealerCard >= 2 && dealerCard <= 9 {
			return DOUBLE
		}
		return HIT
	}

	if playerCardTotal == 11 {
		if dealerCard != 1 {
			return DOUBLE
		}
		return HIT
	}

	if playerCardTotal == 12 {
		if dealerCard >= 4 && dealerCard <= 6 {
			return STAND
		}
		return HIT
	}

	if playerCardTotal >= 13 && playerCardTotal <= 16 {
		if dealerCard >= 2 && dealerCard <= 6 {
			return STAND
		}
		return HIT
	}

	// Stand if 17 or more
	return STAND
}
