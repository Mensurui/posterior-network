package posteriornetwork

type SellerType int

const(
// A seller type that tends to deliver higher quality
	Loyal SellerType = iota
// A seller type that chooses actions based on long term payoff
	Strategic
)

// A fancy way of saying ratings
type Signal int 
const(
	Good Signal = iota
	Bad 
)

type WorldModel struct {
    Likelihood map[SellerType]map[Signal]float64
}

func NewWorldModel()WorldModel{
sellerLikelihoods := map[SellerType]map[Signal]float64{
	Loyal:{
		Bad: 0.1,
		Good: 0.9,
	},
	Strategic:{
		Bad: 0.4,
		Good: 0.6,
	},
}
	return WorldModel{
		Likelihood: sellerLikelihoods,
	}
}


func (m *WorldModel)UpdateBelief(prior map[SellerType]float64, signal Signal) map[SellerType]float64{


posterior := make(map[SellerType]float64)
for sellerType := range prior{
	posterior[sellerType] = m.Likelihood[sellerType][signal] * prior[sellerType]
}
normalizer := 0.0
for _, v := range posterior {
    normalizer += v
}

for t := range posterior {
    posterior[t] /= normalizer
}
	return posterior
}
