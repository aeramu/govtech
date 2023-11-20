package api

import "errors"

type Product struct {
	ID          int64    `json:"id"`
	SKU         string   `json:"sku"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Category    Category `json:"category"`
	ImageURL    string   `json:"imageUrl"`
	Weight      int32    `json:"weight"`
	Price       int64    `json:"price"`
	Rating      float32  `json:"rating"`
}

type Category struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

func (p Product) ValidateCreate() error {
	if p.SKU == "" {
		return errors.New("empty SKU")
	}
	if p.Title == "" {
		return errors.New("empty title")
	}
	if p.Description == "" {
		return errors.New("empty description")
	}
	if p.ImageURL == "" {
		return errors.New("empty image url")
	}
	if p.Category.ID == 0 {
		return errors.New("empty category id")
	}
	if p.Price == 0 {
		return errors.New("empty price")
	}

	return nil
}

func (p Product) ValidateUpdate() error {
	if p.SKU == "" {
		return errors.New("empty SKU")
	}
	if p.Title == "" {
		return errors.New("empty title")
	}
	if p.Description == "" {
		return errors.New("empty description")
	}
	if p.Category.ID == 0 {
		return errors.New("empty category id")
	}

	return nil
}
