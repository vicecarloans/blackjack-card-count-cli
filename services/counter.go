package services

import (
	helpers "vicecarloans/blackjack-card-count-cli/helpers"
)

type CounterService struct {
	decksNum        int
	penetrationRate float64
	playerCount     int
	myPosition      int
	curCount        int
	trueCount       int
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
	return &CounterService{}
}

func (cs *CounterService) SetDecksNum(decksNum int) {
	cs.decksNum = decksNum
}

func (cs *CounterService) SetPenetrationRate(penetrationRate float64) {
	cs.penetrationRate = penetrationRate
}

func (cs *CounterService) SetPlayerCount(playerCount int) {
	cs.playerCount = playerCount
}

func (cs *CounterService) SetMyPosition(myPosition int) {
	cs.myPosition = myPosition
}

func (cs *CounterService) GetRecommendation(playerCards []int, dealersCard int) string {
	if playerCards[0] == playerCards[1] {
		return helpers.DoublePlay(playerCards[0], playerCards[1], dealersCard)
	}
	if playerCards[0] == ACE || playerCards[1] == ACE {

	}
}

func (cs *CounterService) GetTrueCount() {

}
