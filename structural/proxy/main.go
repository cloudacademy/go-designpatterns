package main

import "fmt"

func main() {
	fmt.Print("PROXY Example:\n\n")

	var hashengine hasher = newHasher()
	data := "cloudacademy"

	md5hash := hashengine.hash("MD5", data)
	fmt.Printf("%s -> hash %s\n\n", data, md5hash)

	sha256hash := hashengine.hash("SHA256", data)
	fmt.Printf("%s -> hash %s\n\n", data, sha256hash)

	fmt.Println()
}
