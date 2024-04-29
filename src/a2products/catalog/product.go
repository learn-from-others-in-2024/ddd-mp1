package catalog

// Product is an entity, it has attributes of a product
type Product struct {
	Id    string
	Price float64
	Name  string
}

// NewProduct is the only way to create a product object
func NewProduct(id string, price float64, name string) Product {
	return Product{
		Id:    id,
		Name:  name,
		Price: price,
	}
}
