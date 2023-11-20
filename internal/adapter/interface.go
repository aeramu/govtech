package adapter

import (
	"context"
	"github.com/alam/govtech/internal/model"
)

type ProductRepository interface {
	GetProduct(ctx context.Context, id int64) (model.Product, error)
	GetProductBySKU(ctx context.Context, sku string) (model.Product, error)
	GetProductList(ctx context.Context, filter model.GetProductListFilter) ([]model.Product, error)
	InsertProduct(ctx context.Context, product model.Product) error
	UpdateProduct(ctx context.Context, id int64, product model.Product) error
	UpdateProductRating(ctx context.Context, id int64, rating float64) error
}

type ProductReviewRepository interface {
	InsertReview(ctx context.Context, review model.ProductReview) error
	GetReviewStatistic(ctx context.Context, productID int64) (model.Statistic, error)
}

type CategoryRepository interface {
	GetCategory(ctx context.Context, id int64) (model.Category, error)
}
