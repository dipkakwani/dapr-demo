package spec

import "github.com/benc-uk/dapr-store/cmd/orders/spec"

// Cart holds a users shopping cart
type Cart struct {
	Products map[string]int `json:"products"`
	ForUser  string         `json:"forUser"`
}

// CartService defines core CRUD methods a cart service should have
type CartService interface {
	Get(string) (*Cart, error)
	GetWithETag(string) (*Cart, string, error)
	Submit(Cart) (*spec.Order, error)
	SetProductCount(*Cart, string, int, string) error
	Clear(*Cart) error
}
