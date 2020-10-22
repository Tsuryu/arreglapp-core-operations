package operationtypes

import (
	"context"
	"fmt"
	"time"

	"github.com/Tsuryu/arreglapp-commons/app/utils"
	"github.com/Tsuryu/arreglapp-core-operations/app/models"
)

// Insert : Creates new transaction
func Insert(operationType *models.OperationType) error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	random, err := utils.RandomHex(2)
	if err != nil {
		return err
	}
	operationType.ID = random

	operationType.Active = true
	_, err = Collection.InsertOne(ctx, operationType)
	if err != nil {
		fmt.Println(err.Error())
	}
	return err
}
