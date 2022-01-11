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
	if len(data) > 100 {
		return "string too long"
	}

	//proxy downstream
	return h.hashsvc.hash(hashtype, data)
}
