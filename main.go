package main

import (
	"fmt"
	"github-activity/internal/services"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	url, err := services.BuildCallUrl("denotess")
	if err != nil {
		fmt.Println(err)
	}
	activity, err := services.FetchData(url)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(activity)
}
