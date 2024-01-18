package helpers

func DoublePlay(playerCard1 int, playerCard2 int, dealerCard int) string {
	if ShouldSplit(playerCard1, playerCard2, dealerCard) {
		return SPLIT
	}

	return PlayHardHand([]int{playerCard1, playerCard2}, dealerCard)
}

func ShouldSplit(playerCard1 int, playerCard2 int, dealerCard int) bool {
	if playerCard1 == 2 {
		return dealerCard >= 4 && dealerCard <= 7
	}
	if playerCard1 == 3 {
		return dealerCard >= 4 && dealerCard <= 7
	}
	if playerCard1 == 4 {
		return false
	}
	if playerCard1 == 5 {
		return false
	}
	if playerCard1 == 6 {
		return dealerCard >= 3 && dealerCard <= 6
	}
	if playerCard1 == 7 {
		return dealerCard >= 2 && dealerCard <= 7
	}
	if playerCard1 == 8 {
		return true
	}
	if playerCard1 == 9 {
		return (dealerCard >= 2 && dealerCard <= 6) || dealerCard == 8 || dealerCard == 9
	}
	if playerCard1 == 10 {
		// NEVER SPLIT 10s
		return false
	}
	// If playerCard is an ACE
	return true
}
