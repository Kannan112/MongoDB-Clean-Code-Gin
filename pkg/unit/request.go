package unit

type UserRequest struct {
	Name     string `json:"name" bson:"name"`
	Category string `json:"Category" bson:"Category"`
	Qty      int    `json:"qty" bson:"qty"`
}
