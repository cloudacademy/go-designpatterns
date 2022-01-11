package main

type hasher interface {
	hash(string, string) string
}
