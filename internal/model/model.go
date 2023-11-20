package model

import "time"

type Product struct {
	ID          int64
	SKU         string
	Title       string
	Description string
	Category    Category
	ImageURL    string
	Weight      int32
	Price       int64
	Rating      float32
	CreatedAt   time.Time
}

type GetProductListFilter struct {
	Search     string
	CategoryID int64
	SortColumn string
	SortType   string
	Limit      int64
	Offset     int64
}

type Category struct {
	ID   int64
	Name string
}

type ProductReview struct {
	ID        int64
	UserID    int64
	ProductID int64
	Rating    int32
	Comment   string
}

type Statistic struct {
	Count   int64
	Average float64
}
