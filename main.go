package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/aws"
	"os"
	"log"
	"github.com/YarekTyshchenko/golang_example/src/handler"
	ddb "github.com/YarekTyshchenko/golang_example/src/dynamodb"
)

func main() {
	region, found := os.LookupEnv("Region")
	if !found {
		log.Fatal("Region env var is not defined")
	}
	tableName, found := os.LookupEnv("TableName")
	if !found {
		log.Fatal("TableName env var is not defined")
	}

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(region)},
	)
	if err != nil {
		log.Fatal("Failed to get new session", err)
	}
	h := handler.New(ddb.New(dynamodb.New(sess), tableName))
	lambda.Start(h.Handle)
}