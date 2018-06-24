package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	_ "testing/localmsg"

	"github.com/vponnam/go-web/testing/localmsg"
)

// Input ASG collection
type Input struct {
	Resources []struct {
		Metadata struct {
			GUID      string `json:"guid"`
			URL       string `json:"url"`
			OrgName   string `json:"org_name"`
			SpaceName string `json:"space_name"`
		} `json:"metadata"`
		Entity struct {
			Name  string `json:"name"`
			Rules []struct {
				Protocol    string `json:"protocol"`
				Destination string `json:"destination"`
				Ports       string `json:"ports,omitempty"`
				Code        *int   `json:"code,omitempty"`
				Type        *int   `json:"type,omitempty"`
				Description string `json:"description,omitempty"`
				Log         string `json:"log,omitempty"`
			} `json:"rules"`
		} `json:"entity"`
	} `json:"resources"`
}

func web(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadFile("asg.json")
	var data Input
	err := json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	} else {
		fmt.Printf("Success\n")
		fmt.Fprint(w, localmsg.Msg())
	}
}

func main() {
	http.HandleFunc("/", web)
	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}
