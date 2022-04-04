package view_model

type CreateUpdateConnectionRequest struct {
	Title string `json:"title" bson:"title"`
	Value string `json:"value" bson:"value"`
	Meta  string `json:"meta,omitempty" bson:"meta"`
}
