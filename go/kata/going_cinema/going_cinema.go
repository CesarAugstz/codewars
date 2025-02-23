package going_cinema

import (
	"math"
)

func Movie(card, ticket int, perc float64) int {
	ticketFloat := float64(ticket)

	var cardPlusTickets float64 = float64(card)
	var onlyTickets float64 = 0
	times := 0

  prvValue := ticketFloat

	for true {
		times += 1

		onlyTickets += ticketFloat
		prvValue = prvValue * perc
		cardPlusTickets += prvValue

		if math.Ceil(cardPlusTickets) < math.Ceil(onlyTickets) {
			break
		}

	}

	return times
}
