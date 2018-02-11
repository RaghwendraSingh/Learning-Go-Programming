package network

import "fmt"
import "net/http"
import "io/ioutil"
import "encoding/json"
import "errors"

type Network struct {
	UUID string `json:"uuid"`
	Name string `json:"name"`
	Area string `json:"area"`
}

func GetHierarchyUuid(itemType string) (string, error) {
	fmt.Println("Inside GetHierarchyUuid function --------------------------->")
	protocol := "http"
	host := "192.168.2.5"
	port := 8080
	path := "/mlm/"
	param := "area"

	url := fmt.Sprintf("%s://%s:%d%s?%s=%s",
		protocol,
		host,
		port,
		path,
		param,
		itemType,
	)
	fmt.Println(url)
	resp, err := http.Get(url)
	if err != nil {
		return "", errors.New("Failed to get data from URL ~" + url)
	}
	defer resp.Body.Close()
	body, err2 := ioutil.ReadAll(resp.Body)
	if err2 != nil {
		return "", errors.New("Failed to read body from the response")
	}

	networkResp := []Network{}

	json.Unmarshal(body, &networkResp)

	if len(networkResp) == 0 {
		return "", errors.New("No network with the specified area")
	}
	// returning the first network uuid

	return networkResp[0].UUID, nil

}

func main1() {
	fmt.Println(GetHierarchyUuid("Food"))
}
