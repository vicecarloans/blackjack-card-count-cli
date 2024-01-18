package helpers

// Soft hands are hands that contain an ACE
func PlaySoftHand(playerCard1 int, playerCard2 int, dealerCard int) string {
	playerCardTotal := playerCard1 + playerCard2

	if playerCardTotal == 3 || playerCardTotal == 4 {
		if dealerCard >= 5 && dealerCard <= 6 {
			return DOUBLE
		}
		return HIT
	}

	if playerCardTotal == 5 || playerCardTotal == 6 {
		if dealerCard >= 4 && dealerCard <= 6 {
			return DOUBLE
		}
		return HIT
	}

	if playerCardTotal == 7 {
		if dealerCard >= 3 && dealerCard <= 6 {
			return DOUBLE
		}
		return HIT
	}

	if playerCardTotal == 8 {
		if dealerCard >= 2 && dealerCard <= 8 {
			return STAND
		}
		return HIT
	}

	return STAND
}
