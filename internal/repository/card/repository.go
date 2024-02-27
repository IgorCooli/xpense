package card

import (
	"context"
	"fmt"

	"github.com/IgorCooli/xpense/internal/business/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Respository interface {
	InsertOne(ctx context.Context, card model.Card) error
	FindById(ctx context.Context, cardId string) (model.Card, error)
}

type mongoRepository struct {
	cardDB *mongo.Collection
}

func NewCardRepository(client *mongo.Client) Respository {
	return mongoRepository{
		cardDB: client.Database("TCCMongoDB").Collection("card"),
	}
}

func (r mongoRepository) FindById(ctx context.Context, cardId string) (model.Card, error) {
	var result model.Card
	filter := bson.D{
		{"id", cardId},
	}

	r.cardDB.FindOne(ctx, filter).Decode(&result)

	return result, nil
}

func (r mongoRepository) InsertOne(ctx context.Context, card model.Card) error {

	result, err := r.cardDB.InsertOne(ctx, card)

	if err != nil {
		panic("Could not insert item")
	}

	fmt.Println(result)
	return nil
}
