package main

import (
	"encoding/json"
	"fmt"
)

type Price struct {
	UnitPrice float64 `json:"betrag-je-stueck"`
	Currency  string  `json:"währung"`
}
type Product struct {
	ProductId string `json:"productId"`
	Name      string `json:"NAME"`
	Price     Price  `json:"Preis"`
}

func main() {
	fmt.Println("Structs:")

	p := Product{
		ProductId: "P-123",
		Name:      "Pizza Salami",
		Price: Price{
			UnitPrice: 12.99,
			Currency:  "EUR",
		},
	}

	m, err := json.Marshal(p)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(string(m))
}
