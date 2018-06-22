package dynamodb

import (
	"testing"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/YarekTyshchenko/golang_example/src/handler"
	"fmt"
	"log"
)

type mockDDB struct {}
func (m mockDDB) PutItem(input *dynamodb.PutItemInput) (*dynamodb.PutItemOutput, error) {
	fmt.Println(input.Item)
	if *input.TableName != "table name" {
		log.Fatal("Put item called with wrong table name")
	}
	return nil, nil
}
func TestStore(t *testing.T) {
	awsdb := &mockDDB{}
	ddb := Storage{
		db: awsdb,
		tableName: "table name",
	}

	id, err := ddb.Store(handler.Request{
		Name: "name",
		Company: "company",
	})
	if err != nil {
		log.Fatal("Store returend an error", err)
	}
	fmt.Println(id)
}