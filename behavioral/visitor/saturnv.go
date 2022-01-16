package main

type saturnv struct {
	fuel     int
	payload  int
	maxspeed int
}

func (s *saturnv) accept(v visitor) {
	v.visitForSaturnV(s)
}

func (s *saturnv) getType() string {
	return "SaturnV"
}
