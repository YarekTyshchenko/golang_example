package dynamodb

import (
	"github.com/pborman/uuid"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/aws"
	"time"
	"github.com/YarekTyshchenko/golang_example/src/handler"
)

type visitor struct {
	Id string `json:"id"`
	handler.Request
	Created time.Time `json:"created"`
}

type DynamoDb interface {
	PutItem(*dynamodb.PutItemInput) (*dynamodb.PutItemOutput, error)
}

type Storage struct {
	db        DynamoDb
	tableName string
}

func New(db *dynamodb.DynamoDB, tableName string) Storage {
	return Storage{
		db: db,
		tableName: tableName,
	}
}

func (s Storage) Store(r handler.Request) (string, error) {
	visitor := visitor{
		Id: uuid.New(),
		Request: r,
		Created: time.Now().UTC(),
	}
	av, err := dynamodbattribute.MarshalMap(visitor)
	if err != nil {
		return "", err
	}
	input := &dynamodb.PutItemInput{
		Item: av,
		TableName: aws.String(s.tableName),
	}
	_, err = s.db.PutItem(input)
	if err != nil {
		return "", err
	}
	return visitor.Id, nil
}

