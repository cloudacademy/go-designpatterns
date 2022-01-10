package main

type nasa struct {
}

func (n *nasa) makeRocket() iRocket {
	return &saturnv{
		rocket: rocket{
			engine:     "F1 Rocketdyne",
			maxspeed:   25000,
			maxpayload: 5250000,
		},
	}
}
