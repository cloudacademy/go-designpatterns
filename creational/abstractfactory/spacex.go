package main

type spacex struct {
}

func (s *spacex) makeRocket() iRocket {
	return &falcon9{
		rocket: rocket{
			engine:     "Merlin",
			maxspeed:   25000,
			maxpayload: 22800,
		},
	}
}
