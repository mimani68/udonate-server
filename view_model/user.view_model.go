package view_model

import "udonate/entity"

type CreateUserRequest struct {
	Id           string `json:"id"`
	Name         string `json:"name" bson:"name"`
	Family       string `json:"family" bson:"family"`
	Nationality  string `json:"nationality,omitempty" bson:"nationality"`
	Birthday     string `json:"birthday,omitempty" bson:"birthday"`
	Sex          string `json:"sex,omitempty" bson:"sex"`
	NationalCode string `json:"nationalCode,omitempty" bson:"nationalCode"`
	Username     string `json:"username" bson:"username"`
	Password     string `json:"password" bson:"password"`
	// Connections         []Connection `json:"connections,omitempty" bson:"connections"`
	Email        string `json:"email" bson:"email"`
	Phone        string `json:"phone" bson:"phone"`
	ReferralCode string `json:"referralCode,omitempty" bson:"referralCode"`
	Status       string `json:"status,omitempty" bson:"status"`
	CreatedAt    string `json:"createdAt" bson:"createdAt"`
}

type CreateUserResponse struct {
	Id           string              `json:"id"`
	Name         string              `json:"name" bson:"name"`
	Family       string              `json:"family" bson:"family"`
	Nationality  string              `json:"nationality,omitempty" bson:"nationality"`
	Birthday     string              `json:"birthday,omitempty" bson:"birthday"`
	Sex          string              `json:"sex,omitempty" bson:"sex"`
	NationalCode string              `json:"nationalCode,omitempty" bson:"nationalCode"`
	Username     string              `json:"username" bson:"username"`
	Password     string              `json:"password" bson:"password"`
	Connections  []entity.Connection `json:"connections,omitempty" bson:"connections"`
	// Email        string `json:"email" bson:"email"`
	// Phone        string `json:"phone" bson:"phone"`
	ReferralCode string `json:"referralCode,omitempty" bson:"referralCode"`
	Status       string `json:"status,omitempty" bson:"status"`
	CreatedAt    string `json:"createdAt" bson:"createdAt"`
}

type GetUserResponse struct {
	Id           string `json:"id" bson:"_id"`
	Name         string `json:"name" bson:"name"`
	Family       string `json:"family" bson:"family"`
	Nationality  string `json:"nationality,omitempty" bson:"nationality"`
	Birthday     string `json:"birthday,omitempty" bson:"birthday"`
	Sex          string `json:"sex,omitempty" bson:"sex"`
	NationalCode string `json:"nationalCode,omitempty" bson:"nationalCode"`
	Username     string `json:"username" bson:"username"`
	// Password     string `json:"password" bson:"password"`
	Connections  []entity.Connection `json:"connections,omitempty" bson:"connections"`
	ReferralCode string              `json:"referralCode,omitempty" bson:"referralCode"`
	Requests     []entity.Request    `json:"requests,omitempty" bson:"requests"`
	Status       string              `json:"status,omitempty" bson:"status"`
	CreatedAt    string              `json:"createdAt" bson:"createdAt"`
	// ModifiedAt         string `json:"modifiedAt,omitempty" bson:"modifiedAt"`
	// DeletedAt          string `json:"deletedAt,omitempty" bson:"deletedAt"`
	// DeletedDescription string `json:"deletedDescription,omitempty" bson:"deletedDescription"`
}

type UpdateUserRequest struct {
	Name         string `json:"name" bson:"name"`
	Family       string `json:"family" bson:"family"`
	Nationality  string `json:"nationality,omitempty" bson:"nationality"`
	Birthday     string `json:"birthday,omitempty" bson:"birthday"`
	Sex          string `json:"sex,omitempty" bson:"sex"`
	NationalCode string `json:"nationalCode,omitempty" bson:"nationalCode"`
	ReferralCode string `json:"referralCode,omitempty" bson:"referralCode"`
}
