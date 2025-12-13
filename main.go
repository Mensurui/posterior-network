package main

import (
	"fmt"
	"math/rand"
	posteriornetwork "posterior-network/computation"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())

	model := posteriornetwork.NewWorldModel()

	simulations := 1000
	rounds := 3
	correctCount := 0
	
	for i := 0; i < simulations; i ++{
		trueSeller := sampleTrueSeller()
		finalBelief := runningASimulation(model, trueSeller, rounds)
		if isCorrect(finalBelief, trueSeller){
			correctCount ++
		}
	}
	fmt.Printf(
        "Accuracy after %d simulations: %.2f%%\n",
        simulations,
        float64(correctCount)/float64(simulations)*100,
    )
}

func sampleTrueSeller()posteriornetwork.SellerType{
	if rand.Float64() < 0.5 {
		return posteriornetwork.Loyal
	}
	return posteriornetwork.Strategic
}

func runningASimulation(
	model posteriornetwork.WorldModel,
	trueSeller posteriornetwork.SellerType,
	rounds int,
)map[posteriornetwork.SellerType]float64{
	belief := map[posteriornetwork.SellerType]float64{
        posteriornetwork.Loyal:     0.5,
        posteriornetwork.Strategic: 0.5,
    }

    for i := 0; i < rounds; i++ {
        r := rand.Float64()
        var signal posteriornetwork.Signal

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

        belief = model.UpdateBelief(belief, signal)
    }

    return belief
}

func isCorrect(
    belief map[posteriornetwork.SellerType]float64,
    trueSeller posteriornetwork.SellerType,
) bool {
    return belief[trueSeller] > 0.5
}

