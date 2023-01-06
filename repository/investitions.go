package repository

import (
	"context"
	"errors"
	"invest_blango_criptal_backend/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)


type InvestitionsMongo struct {
	db *mongo.Client
	ctx *context.Context
}


func NewInvestitionsMongo (db *mongo.Client, ctx *context.Context) *InvestitionsMongo {
	return &InvestitionsMongo{db: db, ctx: ctx}
}



func (r *InvestitionsMongo) CreateInvestition(investition models.Investition) (int, error) {
	investitions := r.db.Database("invest_site").Collection("investitions")
	investition_id, _ := investitions.CountDocuments(*r.ctx, bson.D{})
 
	_, err := investitions.InsertOne(*r.ctx, bson.D{{"_id", investition_id}, {"user_id", investition.UserId}, {"date", investition.Date}, {"deadline", investition.Deadline}, {"currency", investition.Currency}, {"amount", investition.Amount}, {"name", investition.Name}, {"email", investition.Email}, {"citizenship", investition.Citizenship}, {"addres", investition.Adress}})

	if err != nil {
		return -1, errors.New(err.Error())
	}
	return int(investition_id), nil
}