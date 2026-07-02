package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/tsirysndr/go-youtube"
)

func main() {
	apiKey := os.Getenv("YOUTUBE_API_KEY")
	if apiKey == "" {
		fmt.Println("Set YOUTUBE_API_KEY environment variable")
		return
	}

	client := youtube.NewClientWithKey(apiKey)
	result, err := client.Search.Search("doja cat", nil)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	r, _ := json.MarshalIndent(result, "", "  ")
	fmt.Println(string(r))
}
