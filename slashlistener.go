package slashlistener

import (
	"fmt"
	"net/http"
)

func init() {
	http.HandleFunc("/", handler)
}

//handles http request from slack
func handler(w http.ResponseWriter, r *http.Request) {
	cmd := r.FormValue("command")

	if cmd == "/yelp" {
		_, err := fmt.Fprintln(w, "SWAG")
		if err != nil {
			fmt.Println(":( it broke, call bestbuy")
		}
	} else {
		_, err := fmt.Fprintln(w, "u wut m8?")
		if err != nil {
			fmt.Println("Plz send help bestbuy")
		}
	}
}
