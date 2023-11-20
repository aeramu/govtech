package service

import (
	"context"
	"database/sql"
	"github.com/alam/govtech/internal/adapter"
	"github.com/alam/govtech/internal/api"
	"github.com/alam/govtech/internal/model"
	"github.com/alam/govtech/internal/util/errorhelper"
	"math/rand"
	"net/http"
	"time"
)

type Service interface {
	CreateProduct(ctx context.Context, req api.Product) (api.MutationResponse, error)
	UpdateProduct(ctx context.Context, id int64, req api.Product) (api.MutationResponse, error)
	GetProduct(ctx context.Context, id int64) (api.Product, error)
	GetProductList(ctx context.Context, filter api.GetProductListFilter) ([]api.Product, error)
	ReviewProduct(ctx context.Context, productID int64, req api.ReviewProductRequest) (api.MutationResponse, error)
}

type service struct {
	productRepo  adapter.ProductRepository
	categoryRepo adapter.CategoryRepository
	reviewRepo   adapter.ProductReviewRepository
}

func NewService(
	productRepo adapter.ProductRepository,
	categoryRepo adapter.CategoryRepository,
	reviewRepo adapter.ProductReviewRepository,
) Service {
	return &service{
		productRepo:  productRepo,
		categoryRepo: categoryRepo,
		reviewRepo:   reviewRepo,
	}
}

func (s *service) CreateProduct(ctx context.Context, req api.Product) (api.MutationResponse, error) {
	if err := req.ValidateCreate(); err != nil {
		return api.MutationResponse{}, errorhelper.WrapWithCode(err, "invalid request payload", http.StatusBadRequest)
	}

	_, err := s.categoryRepo.GetCategory(ctx, req.Category.ID)
	if err != nil && err != sql.ErrNoRows {
		return api.MutationResponse{}, errorhelper.WrapWithCode(err, "error when get category", http.StatusInternalServerError)
	}
	if err == sql.ErrNoRows {
		return api.MutationResponse{}, errorhelper.NewWithCode("category not found", http.StatusBadRequest)
	}

	_, err = s.productRepo.GetProductBySKU(ctx, req.SKU)
	if err != nil && err != sql.ErrNoRows {
		return api.MutationResponse{}, errorhelper.WrapWithCode(err, "error when get product by sku", http.StatusInternalServerError)
	}
	if err == nil {
		return api.MutationResponse{}, errorhelper.NewWithCode("sku already exist", http.StatusBadRequest)
	}

	err = s.productRepo.InsertProduct(ctx, model.Product{
		SKU:         req.SKU,
		Title:       req.Title,
		Description: req.Description,
		Category: model.Category{
			ID: req.Category.ID,
		},
		ImageURL:  req.ImageURL,
		Weight:    req.Weight,
		Price:     req.Price,
		Rating:    0,
		CreatedAt: time.Now(),
	})
	if err != nil {
		return api.MutationResponse{}, errorhelper.WrapWithCode(err, "error when insert product", http.StatusInternalServerError)
	}

	return api.MutationResponse{
		Success: true,
	}, nil
}

func (s *service) UpdateProduct(ctx context.Context, id int64, req api.Product) (api.MutationResponse, error) {
	if id <= 0 {
		return api.MutationResponse{}, errorhelper.NewWithCode("invalid id", http.StatusBadRequest)
	}

	if err := req.ValidateUpdate(); err != nil {
		return api.MutationResponse{}, errorhelper.WrapWithCode(err, "invalid request payload", http.StatusBadRequest)
	}

	_, err := s.categoryRepo.GetCategory(ctx, req.Category.ID)
	if err != nil && err != sql.ErrNoRows {
		return api.MutationResponse{}, errorhelper.WrapWithCode(err, "error when get category", http.StatusInternalServerError)
	}
	if err == sql.ErrNoRows {
		return api.MutationResponse{}, errorhelper.NewWithCode("category not found", http.StatusBadRequest)
	}

	err = s.productRepo.UpdateProduct(ctx, id, model.Product{
		SKU:         req.SKU,
		Title:       req.Title,
		Description: req.Description,
		Category: model.Category{
			ID: req.Category.ID,
		},
	})
	if err != nil && err != sql.ErrNoRows {
		return api.MutationResponse{}, errorhelper.WrapWithCode(err, "error when update product", http.StatusInternalServerError)
	}

	return api.MutationResponse{
		Success: true,
	}, nil
}

func (s *service) GetProduct(ctx context.Context, id int64) (api.Product, error) {
	if id <= 0 {
		return api.Product{}, errorhelper.NewWithCode("invalid id", http.StatusBadRequest)
	}

	product, err := s.productRepo.GetProduct(ctx, id)
	if err != nil && err != sql.ErrNoRows {
		return api.Product{}, errorhelper.WrapWithCode(err, "error when get product", http.StatusInternalServerError)
	}
	if err == sql.ErrNoRows {
		return api.Product{}, errorhelper.NewWithCode("product not found", http.StatusNotFound)
	}

	return api.Product{
		ID:          product.ID,
		SKU:         product.SKU,
		Title:       product.Title,
		Description: product.Description,
		Category: api.Category{
			ID:   product.Category.ID,
			Name: product.Category.Name,
		},
		ImageURL: product.ImageURL,
		Weight:   product.Weight,
		Price:    product.Price,
		Rating:   product.Rating,
	}, nil
}

func (s *service) GetProductList(ctx context.Context, filter api.GetProductListFilter) ([]api.Product, error) {
	if err := filter.Validate(); err != nil {
		return nil, errorhelper.WrapWithCode(err, "invalid request", http.StatusBadRequest)
	}

	products, err := s.productRepo.GetProductList(ctx, model.GetProductListFilter{
		Search:     filter.Search,
		CategoryID: filter.CategoryID,
		SortColumn: filter.SortColumn,
		SortType:   filter.SortType,
		Limit:      filter.Size,
		Offset:     (filter.Page - 1) * filter.Size,
	})
	if err != nil {
		return nil, errorhelper.WrapWithCode(err, "error when get product list", http.StatusInternalServerError)
	}

	res := make([]api.Product, len(products))

	for i, v := range products {
		res[i] = api.Product{
			ID:          v.ID,
			SKU:         v.SKU,
			Title:       v.Title,
			Description: v.Description,
			Category: api.Category{
				ID:   v.Category.ID,
				Name: v.Category.Name,
			},
			ImageURL: v.ImageURL,
			Weight:   v.Weight,
			Price:    v.Price,
			Rating:   v.Rating,
		}
	}

	return res, nil
}

func (s *service) ReviewProduct(ctx context.Context, productID int64, req api.ReviewProductRequest) (api.MutationResponse, error) {
	if productID <= 0 {
		return api.MutationResponse{}, errorhelper.NewWithCode("invalid id", http.StatusBadRequest)
	}

	if err := req.Validate(); err != nil {
		return api.MutationResponse{}, errorhelper.WrapWithCode(err, "invalid request", http.StatusBadRequest)
	}

	_, err := s.productRepo.GetProduct(ctx, productID)
	if err != nil && err != sql.ErrNoRows {
		return api.MutationResponse{}, errorhelper.WrapWithCode(err, "error when get product", http.StatusInternalServerError)
	}
	if err == sql.ErrNoRows {
		return api.MutationResponse{}, errorhelper.NewWithCode("product not found", http.StatusNotFound)
	}

	stat, err := s.reviewRepo.GetReviewStatistic(ctx, productID)
	if err != nil && err != sql.ErrNoRows {
		return api.MutationResponse{}, errorhelper.WrapWithCode(err, "error when get statistic review", http.StatusInternalServerError)
	}

	err = s.reviewRepo.InsertReview(ctx, model.ProductReview{
		UserID:    int64(rand.Int() % 777777),
		ProductID: productID,
		Rating:    req.Rating,
		Comment:   req.Comment,
	})
	if err != nil {
		return api.MutationResponse{}, errorhelper.WrapWithCode(err, "error when insert review", http.StatusInternalServerError)
	}

	rating := ((float64(stat.Count) * stat.Average) + float64(req.Rating)) / float64(stat.Count+1)

	err = s.productRepo.UpdateProductRating(ctx, productID, rating)
	if err != nil {
		return api.MutationResponse{}, errorhelper.WrapWithCode(err, "error when update product rating", http.StatusInternalServerError)
	}

	return api.MutationResponse{
		Success: true,
	}, nil

}
