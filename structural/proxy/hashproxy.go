package main

type hashproxy struct {
	hashsvc *hashservice
}

func newHasher() *hashproxy {
	return &hashproxy{
		hashsvc: &hashservice{},
	}
}

func (h *hashproxy) hash(hashtype string, data string) string {
	return h.hashsvc.hash(hashtype, data)
}
