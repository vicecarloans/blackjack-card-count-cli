package services

import (
	"math"
	helpers "vicecarloans/blackjack-card-count-cli/helpers"
)

type CounterService struct {
	remainingCardNum int
	curCount         int
}

const (
	ACE   = 1
	TWO   = 2
	THREE = 3
	FOUR  = 4
	FIVE  = 5
	SIX   = 6
	SEVEN = 7
	EIGHT = 8
	NINE  = 9
	TEN   = 10
	JACK  = 10
	QUEEN = 10
	KING  = 10
)

func NewCounterService() *CounterService {
	return &CounterService{
		curCount: 0,
	}
}

func (cs *CounterService) Setup(decksNum int, penetrationRate float32) {
	cs.remainingCardNum = int(float32(decksNum) * 52 * penetrationRate)
}

func (cs *CounterService) GetRecommendation(playerCards []int, dealerCard int) string {
	if len(playerCards) == 2 {
		if playerCards[0] == playerCards[1] {
			return helpers.DoublePlay(playerCards[0], playerCards[1], dealerCard)
		}

		if playerCards[0] == ACE || playerCards[1] == ACE {
			return helpers.PlaySoftHand(playerCards[0], playerCards[1], dealerCard)
		}
	}

	return helpers.PlayHardHand(playerCards, dealerCard)
}

func (cs *CounterService) GetTrueCount() float64 {
	deck := math.Ceil(float64(cs.remainingCardNum) / 52)
	return math.Ceil(float64(cs.curCount / int(deck)))
}

func (cs *CounterService) GenerateRunningCount(playerCards []int) {
	for _, card := range playerCards {
		if card == TWO || card == THREE || card == FOUR || card == FIVE || card == SIX {
			cs.curCount++
		}

		if card == TEN || card == JACK || card == QUEEN || card == KING {
			cs.curCount--
		}
	}

	cs.remainingCardNum -= len(playerCards)
}

func (cs *CounterService) GetCurCount() int {
	return cs.curCount
}
