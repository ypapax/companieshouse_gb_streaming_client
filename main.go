package main

import (
	"bufio"
	"encoding/base64"
	"log"
	"net/http"
	"os"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	const apiKeyEnv = "COMPANIESHOUSEGB_STREAMING_KEY"
	apiKey := os.Getenv(apiKeyEnv)
	if len(apiKey) == 0 {
		log.Fatalf("missing env var %+v", apiKeyEnv)
	}
	if err := ListenByURL("https://stream.companieshouse.gov.uk/companies", apiKey); err != nil {
		log.Fatalf("exiting with error: %+v\n", err)
	}
}

func ListenByURL(u string, key string) error {
	log.Printf("listening to %+v\n", u)
	client := http.Client{}

	req, err := http.NewRequest("GET", u, nil)
	if err != nil {
		log.Printf("error: %+v\n", err)
		return err
	}
	req.Header.Add("Authorization", "Basic "+basicAuth(key, ""))
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("error: %+v\n", err)
		return err
	}
	log.Println("reading the body")
	reader := bufio.NewReader(resp.Body)
	for {
		line, err := reader.ReadBytes('\n')
		if err != nil {
			log.Printf("error: %+v\n", err)
			return err
		}
		log.Printf("line received: %+v\n", string(line))
	}
}

func basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}
