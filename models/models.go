package models

import "time"

type Listing struct {
	ID          int               `json:"id"`
	Title       string            `json:"title"`
	Description string            `json:"description"`
	Price       int64             `json:"price"`	//price in cents 6900 would be 69,00 Euroes
	Rating      int64            `json:"rating"`
	Category    []ListingCategory `json:"-"` //- means ignore that

	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type Category struct {
	ID           int       `json:"id"`
	CategoryName string    `json:"categoryName"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}

type ListingCategory struct {
	ID         int       `json:"id"`
	CategoryId int       `json:"categoryId"`
	ListingId  int       `json:"listingId"`
	Category   Category  `json:"category"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}
