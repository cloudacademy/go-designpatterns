package main

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
)

type hashservice struct {
}

func (h *hashservice) hash(hashtype string, data string) string {
	var result string
	switch hashtype {
	case "MD5":
		hash := md5.Sum([]byte(data))
		result = hex.EncodeToString(hash[:])
	case "SHA256":
		hash := sha256.Sum256([]byte(data))
		result = hex.EncodeToString(hash[:])
	}

	return result
}
