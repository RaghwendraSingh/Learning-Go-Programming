package network

import "fmt"
import "net/http"
import "io/ioutil"
import "encoding/json"
import "errors"

type NetworkResponse struct {
	Embedded struct {
		Networks []struct {
			Uuid  string `json:"uuid"`
			Name  string `json:"name"`
			Area  string `json:"area"`
			Links struct {
				Self struct {
					Href string `json:"href"`
				} `json:"self"`
				Network struct {
					Href string `json:"href"`
				} `json:"network"`
			} `json:"_links"`
		} `json:"networks"`
	} `json:"_embedded"`
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
}

func GetHierarchyUuid(itemType string) (string, error) {
	protocol := "http"
	host := "localhost"
	port := 8080
	path := "/nw/search/findByArea/"
	param := "area"

	url := fmt.Sprintf("%s://%s:%d%s?%s=%s",
		protocol,
		host,
		port,
		path,
		param,
		itemType,
	)
	resp, err := http.Get(url)
	if err != nil {
		return "", errors.New("Failed to get data from URL")
	}
	defer resp.Body.Close()
	body, err2 := ioutil.ReadAll(resp.Body)
	if err2 != nil {
		return "", errors.New("Failed to read body from the response")
	}

	networkResp := NetworkResponse{}

	json.Unmarshal(body, &networkResp)

	if len(networkResp.Embedded.Networks) == 0 {
		return "", errors.New("No network with the specified area")
	}
	// returning the first network uuid

	return networkResp.Embedded.Networks[0].Uuid, nil

}

func main1() {
	fmt.Println(GetHierarchyUuid("Food"))
}

