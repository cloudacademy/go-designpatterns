package main

type rocketIterator struct {
	index   int
	rockets []*rocket
}

func (itr *rocketIterator) hasNext() bool {
	return itr.index < len(itr.rockets)
}

func (itr *rocketIterator) getNext() *rocket {
	if itr.hasNext() {
		rocket := itr.rockets[itr.index]
		itr.index++
		return rocket
	}
	return nil
}
