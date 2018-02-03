package main

import "fmt"
import "net/http"
import "io/ioutil"
import "encoding/json"

type x struct {
	FirstName interface{} `json:"firstName"`
	LastName  interface{} `json:"lastName"`
	Links     struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
		Person struct {
			Href string `json:"href"`
		} `json:"person"`
	} `json:"_links"`
}

type Neo4jResponse struct {
	Embedded struct {
		People []x `json:"people"`
	} `json:"_embedded"`
	Links struct {
		Self struct {
			Href      string `json:"href"`
			Templated bool   `json:"templated"`
		} `json:"self"`
		Profile struct {
			Href string `json:"href"`
		} `json:"profile"`
		Search struct {
			Href string `json:"href"`
		} `json:"search"`
	} `json:"_links"`
	Page struct {
		Size          int `json:"size"`
		TotalElements int `json:"totalElements"`
		TotalPages    int `json:"totalPages"`
		Number        int `json:"number"`
	} `json:"page"`
}

func main() {
	fmt.Println("hw")
	resp, err := http.Get("http://192.168.2.5:8080/people")
	if err != nil {
		fmt.Println("error")
	}
	defer resp.Body.Close()
	body, err2 := ioutil.ReadAll(resp.Body)
	if err2 != nil {
		fmt.Println("error @")
	}

	respN := Neo4jResponse{}

	json.Unmarshal(body, &respN)

	for _, people := range respN.Embedded.People {
		fmt.Println(people.FirstName, "~", people.LastName)
	}
	//fmt.Println(respN)

}
