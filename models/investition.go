package models

type Investition struct {
	UserId 	 	int  `json:"user_id" bson:"user_id" binding:"required"`
	Date 	 	string  `json:"date" bson:"date" binding:"required"`
	Deadline 	string  `json:"deadline" bson:"deadline" binding:"required"`
	Currency 	string  `json:"currency" bson:"currency" binding:"required"`
	Amount   	float64 `json:"amount" bson:"amount" binding:"required"`
	Name	 	string  `json:"name" bson:"name" binding:"required"`
	Email	 	string  `json:"email" bson:"email" binding:"required"`
	Citizenship string  `json:"citizenship" bson:"citizenship" binding:"required"`
	Adress		string  `json:"adress" bson:"adress" binding:"required"`
}