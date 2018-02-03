package main

import "encoding/json"
import "fmt"

type response1 struct {
	Page   int
	Fruits []string
}
type response2 struct {
	Page   int `json:"page"`
	Fruits []string
}
type Person struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Links     struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
		Person struct {
			Href string `json:"href"`
		} `json:"person"`
	} `json:"_links"`
}

func main() {
	bolB, _ := json.Marshal(true)
	fmt.Println(bolB)
	fmt.Println(string(bolB))

	intB, _ := json.Marshal(1)
	fmt.Println(intB)
	fmt.Println(string(intB))

	slcD := []string{"apple", "peach", "pear"}
	slcB, _ := json.Marshal(slcD)
	fmt.Println(string(slcB))

	mapD := map[string]int{"apple": 5, "lettuce": 7}
	mapB, _ := json.Marshal(mapD)
	fmt.Println(string(mapB))

	res1D := &response1{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	res1B, _ := json.Marshal(res1D)
	fmt.Println(string(res1B))

	res2D := &response2{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	res2B, _ := json.Marshal(res2D)
	fmt.Println(string(res2B))

	p1 := `{"firstName":"Tom","lastName":"Riddle","_links":{"self":{"href":"http://192.168.2.5:8080/people/0"},"person":{"href":"http://192.168.2.5:8080/people/0"}}}`
	//fmt.Println(p1)

	p1u := Person{}

	json.Unmarshal([]byte(p1), &p1u)
	fmt.Println(p1u)
}
