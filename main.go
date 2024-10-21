package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	// "reflect"
)

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
		fmt.Println(string(body))
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