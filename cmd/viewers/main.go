package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	id := os.Getenv("TWITCH_ID")
	token := os.Getenv("TWITCH_TOKEN")

	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://api.twitch.tv/helix/streams?user_login=stephentcodes", nil)

	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	req.Header.Add("client-id", id)

	resp, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	type StreamInfo struct {
		ViewerCount int `json:"viewer_count"`
	}

	var response struct {
		Data []StreamInfo `json:"data"`
	}

	err = json.NewDecoder(resp.Body).Decode(&response)

	if err != nil {
		log.Fatal(err)
	}

	count := response.Data[0].ViewerCount

	fmt.Printf("VIEWERS:%d\n", count)
}
