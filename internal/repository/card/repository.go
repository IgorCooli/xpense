package card

import (
	"context"
	"fmt"

	"github.com/IgorCooli/xpense/internal/business/model"
	"go.mongodb.org/mongo-driver/mongo"
)

type Respository interface {
	InsertOne(ctx context.Context, card model.Card) error
}

type mongoRepository struct {
	cardDB *mongo.Collection
}

func NewCardRepository(client *mongo.Client) Respository {
	return mongoRepository{
		cardDB: client.Database("TCCMongoDB").Collection("card"),
	}
}

func (r mongoRepository) InsertOne(ctx context.Context, card model.Card) error {

	result, err := r.cardDB.InsertOne(ctx, card)

	if err != nil {
		panic("Could not insert item")
	}

	fmt.Println(result)
	return nil
}
