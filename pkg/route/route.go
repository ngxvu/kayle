package route

import (
	"context"
	"fmt"
	"gitlab.com/merakilab9/meracore/logger"
	"gitlab.com/merakilab9/meracore/service"
	"gitlab.com/merakilab9/meracrawler/kayle/conf"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
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
	var result bson.M
	if err := client.Database("shopee").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Decode(&result); err != nil {
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

	// display the documents retrieved
	fmt.Println("displaying all results in a collection")
	for _, result := range results {
		fmt.Println(result)
	}
	// Elastic Client
	//_, err = elasticsearch.NewDefaultClient()
	//if err != nil {
	//	log.Fatalf("Error creating the client: %s", err)
	//}
	//log.Println(elasticsearch.Version)
	//
	//cert, _ := ioutil.ReadFile("./http_ca.crt")
	//
	//config := elasticsearch.Config{
	//	Addresses: []string{
	//		"https://localhost:9200",
	//	},
	//	Username: "elastic",
	//	Password: "1L1QM=ig=VW+DFNkKIfP",
	//	CACert:   cert,
	//	//CloudID:  "4RaKu4oBxCOx6W9p3sxl",
	//	//APIKey:   "NGhhTXU0b0J4Q094Nlc5cENjeWc6MU1PTXdfTlJSd3lBN1BZZm1UX04tZw",
	//}
	//es, err := elasticsearch.NewClient(config)
	//
	//if err != nil {
	//	log.Fatalf("Error creating the client: %s", err)
	//}
	//
	//res, err := es.Info()
	//if err != nil {
	//	log.Fatalf("Error getting response: %s", err)
	//}
	//
	//defer res.Body.Close()
	//log.Println(res)

	return s
}
