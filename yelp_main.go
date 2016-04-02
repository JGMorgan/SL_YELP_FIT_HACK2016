/*

Consumer Key	xkexG2jGMEzBC0LU3pZF2A
Consumer Secret	HPy5StvQ4sMZx26ISX8QC2fmfdc
Token	8X47ZbEsoipD-Kht2q4PuBc1OrEN2TMN
Token Secret	JYnVfoZKI2b-5b7r1ux6X39a8-4
*/

package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/JustinBeckwith/go-yelp/yelp"
	//	"go-yelp/yelp"
	//"io"
	"io/ioutil"
	//"net/http"
	//	"os"
)

func main() {

	res(term, location)
}

func res(term string, location string) {

	// get the keys either from config file
	options, err := getOptions()
	if err != nil {
		fmt.Println(err)
	}

	// create a new yelp client with the auth keys
	client := yelp.New(options, nil)

	// make a simple query
	//term := r.URL.Query().Get("term")
	//location := r.URL.Query().Get("location")
	results, err := client.DoSimpleSearch("", "miami")
	if err != nil {
		fmt.Println(err)
	}

	// print the results
	for i := 0; i < len(results.Businesses); i++ {
		fmt.Println(results.Businesses[i].Name, results.Businesses[i].Rating)
	}
}

func getOptions() (options *yelp.AuthOptions, err error) {

	var o *yelp.AuthOptions

	// start by looking for the keys in config.json
	data, err := ioutil.ReadFile("../config.json")
	if err != nil {
		// if the file isn't there, check environment variables
		o = &yelp.AuthOptions{
			ConsumerKey:       "xkexG2jGMEzBC0LU3pZF2A",
			ConsumerSecret:    "HPy5StvQ4sMZx26ISX8QC2fmfdc",
			AccessToken:       "8X47ZbEsoipD-Kht2q4PuBc1OrEN2TMN",
			AccessTokenSecret: "JYnVfoZKI2b-5b7r1ux6X39a8-4",
		}
		if o.ConsumerKey == "" || o.ConsumerSecret == "" || o.AccessToken == "" || o.AccessTokenSecret == "" {
			return o, errors.New("to use the sample, keys must be provided either in a config.json file at the root of the repo, or in environment variables")
		}
	} else {
		err = json.Unmarshal(data, &o)
		return o, err
	}
	return o, nil
}
