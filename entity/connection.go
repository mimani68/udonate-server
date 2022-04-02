package entity

type Connection struct {
	Id               string `json:"id" bson:"_id"`
	Title            string `json:"title" bson:"title"`
	Value            string `json:"value" bson:"value"`
	Meta             string `json:"meta,omitempty" bson:"meta"`
	IsVerified       bool   `json:"isVerified" bson:"isVerified"`
	VerificationTime string `json:"verificationTime" bson:"verificationTime"`
	VerificationCode string `json:"verificationCode" bson:"verificationCode"`
}
