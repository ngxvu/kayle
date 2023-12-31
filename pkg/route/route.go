package route

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"gitlab.com/merakilab9/meracore/logger"
	"gitlab.com/merakilab9/meracore/service"
	"gitlab.com/merakilab9/meracrawler/kayle/conf"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"io/ioutil"
	"log"
	"strings"
)

type Service struct {
	*service.BaseApp
}

func NewService() *Service {
	s := &Service{
		BaseApp: service.NewApp("Kayle Service", "v1.0"),
	}
	// MongoDb Client
	client, err := mongo.Connect(
		context.Background(),
		options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%s", conf.LoadEnv().MongoDBHost, conf.LoadEnv().MongoDBPort)),
	)

	// Send a ping to confirm a successful connection
	var ping bson.M
	if err := client.Database("shopee").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Decode(&ping); err != nil {
		panic(err)
	}
	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")

	if err != nil {
		logger.WithCtx(context.Background(), "Connect mongodb failed").Error(err.Error())
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Printf("Không ping được tới MongoDB. Không thể kết nối với MongoDB: %v", err)
		panic(err)
	}
	categoryCollection := client.Database("shopee").Collection("category")

	cursor, err := categoryCollection.Find(context.TODO(), bson.D{})
	if err != nil {
		panic(err)
	}

	// convert the cursor result to bson
	var results []bson.M
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}

	// Exclude _id field from each document
	for i := range results {
		delete(results[i], "_id")
	}

	// ===================== Elastic Client =====================
	_, err = elasticsearch.NewDefaultClient()
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}
	log.Println(elasticsearch.Version)

	cert, _ := ioutil.ReadFile("./http	_ca.crt")

	config := elasticsearch.Config{
		Addresses: []string{
			"https://localhost:9200",
		},
		Username: "elastic",
		Password: "_9w6WrKgqjNZhiEvFsvM",
		CACert:   cert,
	}
	es, err := elasticsearch.NewClient(config)

	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}

	res, err := es.Info()
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}

	defer res.Body.Close()
	log.Println(res)

	index := "listcategory"

	var catestruct map[string]interface{}

	for _, result := range results {
		convertByte, _ := bson.Marshal(result)
		bson.Unmarshal(convertByte, &catestruct)
		jsonString, _ := json.Marshal(result)
		req := esapi.IndexRequest{
			Index:   index,                                 // Index name
			Body:    strings.NewReader(string(jsonString)), // Document body
			Refresh: "true",                                // Refresh
		}
		res, err := req.Do(context.Background(), es)
		if err != nil {
			log.Fatalf("Error getting response: %s", err)
		}
		defer res.Body.Close()

		log.Println(res)

	}

	return s

}
