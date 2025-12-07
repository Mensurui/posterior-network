package posteriornetwork

type Buyer struct{
	Priori float64
}

func NewBuyer()Buyer{
	return Buyer{
		Priori: 0.5,
	}
}


type Seller struct{}

// A seller type that tends to deliver higher quality
type Loyal Seller

// A seller type that chooses actions based on long term payoff
type Strategic Seller

// A fancy way of saying ratings
type Signals struct{}
