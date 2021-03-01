package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

// TODO: Consider json?
func makeList(data string) []string {
	// TODO: Consider regex
	data = strings.ReplaceAll(data, "[", "")
	data = strings.ReplaceAll(data, "]", "")
	data = strings.ReplaceAll(data, "\"", "")

	arr := strings.Split(data, ",")

	return arr
}

func extract(tag string, channel chan string) {
	url := "https://www.rcsb.org/fasta/entry/" + tag
	resp, err := http.Get(url)

	if err != nil {
		fmt.Println("Cannot recieve", tag, ":", err)
		return
	}

	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)

	// Check for errors in data
	if err != nil {
		fmt.Println("Cannot read", tag, ":", err)
		return
	}

	channel <- string(data)
}

func write(tag string, folder string, channel chan string) {
	file, _ := os.Create(folder + tag + ".fasta")
	_, _ = io.WriteString(file, <-channel)

	defer file.Close()
}

func main() {
	// Request `entry identities`
	resp, err := http.Get("https://data.rcsb.org/rest/v1/holdings/current/entry_ids")

	// Check for request errors
	if err != nil {
		print("Cannot fetch entries:", err)
		return
	}

	defer resp.Body.Close()

	// Get requested data
	data, err := ioutil.ReadAll(resp.Body)

	// Check for errors in data
	if err != nil {
		fmt.Println("Cannot read entries:", err)
		return
	}

	// Clean file
	arr := makeList(string(data))

	folder := "fasta/"

	// Ensure folder exist
	_, err = os.Open(folder)

	if err != nil {
		if os.IsNotExist(err) {
			os.Mkdir(folder, 0755)
		}
	}

	// Channels
	URLChannel := make(chan string)

	// Kick off the extraction process (concurrently)
	for _, tag := range arr {
		go extract(tag, URLChannel)
		write(tag, folder, URLChannel)
	}

	close(URLChannel)
}
