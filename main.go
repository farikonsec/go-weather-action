package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
)

func main() {
	resp, err := http.Get("https://wttr.in/London?format=3&u")
	if err != nil {
		fmt.Println("Error fetching weather:", err)
		return
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("🌦️ London Weather:", string(body))
}
