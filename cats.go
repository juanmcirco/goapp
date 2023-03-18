package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

var client *http.Client

type CatFact struct {
	Fact   string `json:"fact"`
	Length int    `json:"length"`
}

func GetCatFact() string {
	url := "https://catfact.ninja/fact"
	var catFact CatFact

	err := GetJson(url, &catFact)

	if err != nil {
		fmt.Printf("error getting cvat fact: %s\n", err.Error())
		return "error"
	} else {
		return catFact.Fact
	}
}

func GetJson(url string, target interface{}) error {
	resp, err := client.Get(url)

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return json.NewDecoder(resp.Body).Decode(target)
}

func main() {

	client = &http.Client{Timeout: 10 * time.Second}

	// Rutas
	http.HandleFunc("/cats", gatitosHandler)

	//starteo mi server
	http.ListenAndServe(":8443", nil)
}

func gatitosHandler(w http.ResponseWriter, req *http.Request) {

	fact := GetCatFact()

	catFact := CatFact{
		Fact:   string(fact),
		Length: len(fact),
	}

	jsonStr, err := json.Marshal(catFact)
	if err != nil {
		fmt.Printf("error")
	} else {
		w.Write([]byte(string(jsonStr)))
	}

}
