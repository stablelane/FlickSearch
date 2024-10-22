package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"runtime"

	// "reflect"

	"github.com/manifoldco/promptui"
)

type Movie struct {
	Title  string `json:"Title"`
	Year   string `json:"Year"`
	ImdbID string `json:"imdbID"`
	Type   string `json:"Type"`
	Poster string `json:"Poster"`
}

type Response struct {
	Search       []Movie `json:"Search"`
	TotalResults string  `json:"totalResults"`
	Response     string  `json:"Response"`
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: your-cli-tool <command>")
		return
	} else {
		x := os.Args[1]
		apiKey := "&apikey=1170f02c"
		baseURL := "http://www.omdbapi.com/?s="
		apiLink := baseURL + x + apiKey
		fmt.Println(apiLink)
		//fetching data from omdbapi using movie name from user 
		res, err := http.Get(apiLink)
		if err != nil {
			fmt.Println(err)
		}
		defer res.Body.Close()
		body, err := io.ReadAll(res.Body)
		if err != nil {
			fmt.Println(err)
		}
		var response Response
		err = json.Unmarshal(body, &response)
		if err != nil {
			fmt.Println("Error unmarshaling JSON:", err)
			return
		}

		//create an array to store the movies from seach
		var movies []string

		// Access each movie in the Search array
		for _, movie := range response.Search {
			movies = append(movies, onlyMovieType(movie))
		}
		// fmt.Println(movies)

		//create a list in cli for users to select movies
		prompt := promptui.Select{
			Label: "Select Movie",
			Items: movies,
		}
		index, _, err := prompt.Run()
		if err != nil {
			fmt.Printf("Promt failed %v\n", err)
			return
		}
		var selectedMovie = response.Search[index]
		var selectedMovieId = selectedMovie.ImdbID
		openLink(selectedMovieId)
	}

}


//function to only append movie type object to movies slice
func onlyMovieType(item Movie) string {
	if item.Type == "movie" {
		return fmt.Sprintf("Title: %s, Year: %s, IMDb ID: %s, Type: %s", item.Title, item.Year, item.ImdbID, item.Type)
		
	}
	return ""
}


// function to check the OS of evironment and open browser with the link using the imdbId as query
func openLink(id string) error {
	var baseURL = "https://multiembed.mov/?video_id="
	var url = baseURL + id

	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "linux":
		cmd = exec.Command("xdg-open", url)
	case "darwin":
		cmd = exec.Command("open", url)
	case "windows":
		cmd = exec.Command("rundll32", "url.dll,FileProtocolHandler", url)
	default:
		return fmt.Errorf("unsupported platform")

	}
	return cmd.Start()

}
