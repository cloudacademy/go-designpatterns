package main

type collection interface {
	newIterator() iterator
}
