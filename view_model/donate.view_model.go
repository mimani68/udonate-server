package view_model

type DonateRequest struct {
	Title      string `json:"title" bson:"title"`
	Amount     string `json:"amount" bson:"amount"`
	Category   string `json:"category" bson:"category"`
	ExpireTime string `json:"expireTime,omitempty" bson:"expireTime"`
	Goal       string `json:"goal" bson:"goal"`
	Address    string `json:"address" bson:"address"`
	Currency   string `json:"currency" bson:"currency"`
}

type GetUserDonateReqResponse struct {
	Id         string `json:"id" bson:"_id"`
	Title      string `json:"title" bson:"title"`
	Amount     string `json:"amount" bson:"amount"`
	Category   string `json:"category" bson:"category"`
	ExpireTime string `json:"expireTime,omitempty" bson:"expireTime"`
	Goal       string `json:"goal" bson:"goal"`
	Address    string `json:"address" bson:"address"`
	Currency   string `json:"currency" bson:"currency"`
}
