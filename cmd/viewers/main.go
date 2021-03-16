package main

import (
	"encoding/json"
	"fmt"
	// "io/ioutil"
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
		user_login   string
		viewer_count int
	}
	var response struct {
		data []interface{}
	}

	err = json.NewDecoder(resp.Body).Decode(&response)
	// data, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(response)

	fmt.Println("VIEWERS:0")
}
