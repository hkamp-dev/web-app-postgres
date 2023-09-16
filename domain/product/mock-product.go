package product

import "github.com/google/uuid"

func GetSampleProducts() []Product {
	return []Product{
		{ID: uuid.New(), Name: "Prod1", Description: "amazing"},
		{ID: uuid.New(), Name: "Prod2", Description: "great"},
		{ID: uuid.New(), Name: "Prod3", Description: "the best"},
	}
}
