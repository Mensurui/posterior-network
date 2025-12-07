package main

import (
	"fmt"
	posteriornetwork "posterior-network/computation"
)

func main() {
	model := posteriornetwork.NewWorldModel()
	 prior := map[posteriornetwork.SellerType]float64{
        posteriornetwork.Loyal:     0.5,
        posteriornetwork.Strategic: 0.5,
    }

	posterior := model.UpdateBelief(prior, posteriornetwork.Good)

	fmt.Println(posterior)
}

