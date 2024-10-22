package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	// "reflect"
	"encoding/json"
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
    TotalResults string   `json:"totalResults"`
    Response     string   `json:"Response"`
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: your-cli-tool <command>")
		return
	} else {
		x  := os.Args[1]
		apiKey := "&apikey=1170f02c"
		baseURL := "http://www.omdbapi.com/?s="
		apiLink := baseURL + x + apiKey
		fmt.Println(apiLink)
		res , err := http.Get(apiLink)
		if err != nil {
			fmt.Println(err)
		}
		defer res.Body.Close()
		body , err := io.ReadAll(res.Body)
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
			movies = append(movies,fmt.Sprintf("Title: %s, Year: %s, IMDb ID: %s, Type: %s", movie.Title, movie.Year, movie.ImdbID, movie.Type))
		}
		// fmt.Println(movies)

		//create 
		prompt := promptui.Select{
			Label: "Select Movie",
			Items: movies,
		}
		_, result, err := prompt.Run()
		if err != nil {
			fmt.Printf("Promt failed %v\n", err)
			return
		}

		fmt.Printf("You choose %q\n", result)
		}
	
}
// for i:=0; i<len(os.Args); i++ {
// 	fmt.Println(os.Args[i])
// } 
// switch os.Args[1] {
// case "greet":
// 	fmt.Println("Hello!")
// default:
// 	fmt.Println("Unknown command:", os.Args[1])
	
// }