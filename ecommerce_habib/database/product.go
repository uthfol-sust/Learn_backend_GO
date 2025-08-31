package database

type Products struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImgUrl      string  `json:"imageUrl"`
}

var ProductsList []Products

func init() {
	ProductsList = append(ProductsList,
		Products{
			ID:          1,
			Name:        "Laptop",
			Description: "14-inch laptop with Intel i5",
			Price:       750.00,
			ImgUrl:      "https://example.com/laptop.png",
		},
		Products{
			ID:          2,
			Name:        "Mouse",
			Description: "Wireless ergonomic mouse",
			Price:       25.50,
			ImgUrl:      "https://example.com/mouse.png",
		},
		Products{
			ID:          3,
			Name:        "Keyboard",
			Description: "Mechanical keyboard with RGB",
			Price:       99.99,
			ImgUrl:      "https://example.com/keyboard.png",
		},
	)
}
