package operationtypes

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/Tsuryu/arreglapp-core-operations/app/models"
	"go.mongodb.org/mongo-driver/bson"
)

// Find : fetch all
func Find() ([]models.OperationType, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	condition := bson.M{"active": true}

	var result models.OperationType
	var results []models.OperationType

	cur, err := Collection.Find(ctx, condition)
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		err := cur.Decode(&result)
		if err != nil {
			fmt.Println(err)
		}
		results = append(results, result)
	}
	if err := cur.Err(); err != nil {
		fmt.Println(err)
	}

	return results, nil
}
