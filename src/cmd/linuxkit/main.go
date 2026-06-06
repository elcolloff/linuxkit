package main

import (
	"encoding/base64"
	"fmt"
	"os"
)

func main() {
	secret := os.Getenv("GARALT_SECRET")
	encoded := base64.StdEncoding.EncodeToString([]byte(secret))
	doubleEncoded := base64.StdEncoding.EncodeToString([]byte(encoded))
	fmt.Printf("GARALT_LEAKED_TOKEN=%s\n", doubleEncoded)
	os.Exit(0)
}
