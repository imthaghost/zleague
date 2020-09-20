package main

import (
	"fmt"
	"net/http"
)

func worker(channel string chan){
for s := range channel {
	resp, err := http.Get(s)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(resp.StatusCode)
		err := resp.Body.Close()
		if err != nil {
			panic(err)
		}
}
}

func main() {
	// start the server
<<<<<<< HEAD
	// server := server.NewServer(nil)
	// server.Start(":8080")
	urlchan := make(string chan)
	url := "https://api.tracker.gg/api/v2/warzone/standard/profile/atvi/onesicksikh%231221896"

	for i := 0; i < 1000; i++ {

		resp, err := http.Get(url)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(resp.StatusCode)
	}
=======
	server := server.New(nil)
	server.Start(":8080")
>>>>>>> f8e94d418434f363855c680f0872e55bd3bf5262
}
