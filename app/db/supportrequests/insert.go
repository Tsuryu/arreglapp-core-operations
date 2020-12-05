package supportrequets

import (
	"context"
	"time"

	"github.com/Tsuryu/arreglapp-core-operations/app/models"
)

// Insert : Creates new transaction
func Insert(supportRequest models.SupportRequest) error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	_, err := Collection.InsertOne(ctx, supportRequest)
	return err
}
