package main

import (
	"fmt"
	"math/rand"

	posteriornetwork "posterior-network/computation"
)

func main() {
	trueSeller := posteriornetwork.Loyal
	model := posteriornetwork.NewWorldModel()

	// Initial priors
	prior := map[posteriornetwork.SellerType]float64{
		posteriornetwork.Loyal:     0.5,
		posteriornetwork.Strategic: 0.5,
	}

	// Run 20 rounds
	for round := 1; round <= 20; round++ {
		fmt.Printf("=== Round %d ===\n", round)

		// Generate signal from true seller type
		var signal posteriornetwork.Signal
		r := rand.Float64()

		if trueSeller == posteriornetwork.Loyal {
			if r < 0.9 {
				signal = posteriornetwork.Good
			} else {
				signal = posteriornetwork.Bad
			}
		} else {
			if r < 0.6 {
				signal = posteriornetwork.Good
			} else {
				signal = posteriornetwork.Bad
			}
		}

		// Update belief
		prior = model.UpdateBelief(prior, signal)

		// Print beliefs
		fmt.Println("Signal:", signal)
		fmt.Printf("P(Loyal):     %.4f\n", prior[posteriornetwork.Loyal])
		fmt.Printf("P(Strategic): %.4f\n", prior[posteriornetwork.Strategic])
		fmt.Println()
	}
}

