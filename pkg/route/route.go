package route

import (
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
	"gitlab.com/merakilab9/meracore/service"
	"io/ioutil"
	"log"
	"net/http"
)

type Service struct {
	*service.BaseApp
}

func NewService() *Service {
	s := &Service{
		BaseApp: service.NewApp("Kayle Service", "v1.0"),
	}

	_, err := elasticsearch.NewDefaultClient()
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}
	log.Println(elasticsearch.Version)

	url := "https://localhost:9200"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return s
	}
	req.Header.Add("Authorization", "Basic ZWxhc3RpYzp0OTQ0d1RObUNnVndLcVp3MnRSbQ==")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return s
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return s
	}
	fmt.Println(string(body))

	//Console Menu

	return s
}
