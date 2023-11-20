package service

import (
	"context"
	"database/sql"
	"errors"
	"github.com/alam/govtech/internal/api"
	"github.com/alam/govtech/internal/model"
	"github.com/alam/govtech/internal/util/errorhelper"
	"github.com/alam/govtech/mocks"
	"github.com/stretchr/testify/mock"
	"net/http"
	"reflect"
	"testing"
	"time"
)

var (
	mockProductRepo  *mocks.ProductRepository
	mockCategoryRepo *mocks.CategoryRepository
	mockReviewRepo   *mocks.ProductReviewRepository
)

func initMock() {
	mockProductRepo = new(mocks.ProductRepository)
	mockCategoryRepo = new(mocks.CategoryRepository)
	mockReviewRepo = new(mocks.ProductReviewRepository)
}

func Test_service_CreateProduct(t *testing.T) {
	type args struct {
		ctx context.Context
		req api.Product
	}
	tests := []struct {
		name       string
		args       args
		prepare    func()
		want       api.MutationResponse
		statusCode int
	}{
		{
			name: "invalid request payload",
			args: args{
				ctx: context.Background(),
				req: api.Product{},
			},
			prepare:    nil,
			want:       api.MutationResponse{},
			statusCode: http.StatusBadRequest,
		},
		{
			name: "error when get category",
			args: args{
				ctx: context.Background(),
				req: api.Product{
					SKU:         "IND001",
					Title:       "Foo",
					Description: "Makanan ringan",
					Category: api.Category{
						ID:   5,
						Name: "",
					},
					ImageURL: "https://foo.bar/foo.jpg",
					Weight:   5,
					Price:    10000,
					Rating:   4,
				},
			},
			prepare: func() {
				mockCategoryRepo.On("GetCategory", mock.Anything, int64(5)).
					Return(model.Category{}, errors.New("any"))
			},
			want:       api.MutationResponse{},
			statusCode: http.StatusInternalServerError,
		},
		{
			name: "category not found",
			args: args{
				ctx: context.Background(),
				req: api.Product{
					SKU:         "IND001",
					Title:       "Foo",
					Description: "Makanan ringan",
					Category: api.Category{
						ID:   5,
						Name: "",
					},
					ImageURL: "https://foo.bar/foo.jpg",
					Weight:   5,
					Price:    10000,
					Rating:   4,
				},
			},
			prepare: func() {
				mockCategoryRepo.On("GetCategory", mock.Anything, int64(5)).
					Return(model.Category{}, sql.ErrNoRows)
			},
			want:       api.MutationResponse{},
			statusCode: http.StatusBadRequest,
		},
		{
			name: "error when get product by sku",
			args: args{
				ctx: context.Background(),
				req: api.Product{
					SKU:         "IND001",
					Title:       "Foo",
					Description: "Makanan ringan",
					Category: api.Category{
						ID:   5,
						Name: "",
					},
					ImageURL: "https://foo.bar/foo.jpg",
					Weight:   5,
					Price:    10000,
					Rating:   4,
				},
			},
			prepare: func() {
				mockCategoryRepo.On("GetCategory", mock.Anything, int64(5)).
					Return(model.Category{}, nil)
				mockProductRepo.On("GetProductBySKU", mock.Anything, "IND001").
					Return(model.Product{}, errors.New("any"))
			},
			want:       api.MutationResponse{},
			statusCode: http.StatusInternalServerError,
		},
		{
			name: "sku already exist",
			args: args{
				ctx: context.Background(),
				req: api.Product{
					SKU:         "IND001",
					Title:       "Foo",
					Description: "Makanan ringan",
					Category: api.Category{
						ID:   5,
						Name: "",
					},
					ImageURL: "https://foo.bar/foo.jpg",
					Weight:   5,
					Price:    10000,
					Rating:   4,
				},
			},
			prepare: func() {
				mockCategoryRepo.On("GetCategory", mock.Anything, int64(5)).
					Return(model.Category{}, nil)
				mockProductRepo.On("GetProductBySKU", mock.Anything, "IND001").
					Return(model.Product{}, nil)
			},
			want:       api.MutationResponse{},
			statusCode: http.StatusBadRequest,
		},
		{
			name: "error when insert product",
			args: args{
				ctx: context.Background(),
				req: api.Product{
					SKU:         "IND001",
					Title:       "Foo",
					Description: "Makanan ringan",
					Category: api.Category{
						ID:   5,
						Name: "",
					},
					ImageURL: "https://foo.bar/foo.jpg",
					Weight:   5,
					Price:    10000,
					Rating:   4,
				},
			},
			prepare: func() {
				mockCategoryRepo.On("GetCategory", mock.Anything, int64(5)).
					Return(model.Category{}, nil)
				mockProductRepo.On("GetProductBySKU", mock.Anything, "IND001").
					Return(model.Product{}, sql.ErrNoRows)
				mockProductRepo.On("InsertProduct", mock.Anything, mock.Anything).
					Return(errors.New("any"))
			},
			want:       api.MutationResponse{},
			statusCode: http.StatusInternalServerError,
		},
		{
			name: "success",
			args: args{
				ctx: context.Background(),
				req: api.Product{
					SKU:         "IND001",
					Title:       "Foo",
					Description: "Makanan ringan",
					Category: api.Category{
						ID:   5,
						Name: "",
					},
					ImageURL: "https://foo.bar/foo.jpg",
					Weight:   5,
					Price:    10000,
					Rating:   4,
				},
			},
			prepare: func() {
				mockCategoryRepo.On("GetCategory", mock.Anything, int64(5)).
					Return(model.Category{}, nil)
				mockProductRepo.On("GetProductBySKU", mock.Anything, "IND001").
					Return(model.Product{}, sql.ErrNoRows)
				mockProductRepo.On("InsertProduct", mock.Anything, mock.Anything).
					Return(nil)
			},
			want: api.MutationResponse{
				Success: true,
			},
			statusCode: http.StatusOK,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			initMock()
			s := &service{
				productRepo:  mockProductRepo,
				categoryRepo: mockCategoryRepo,
				reviewRepo:   mockReviewRepo,
			}
			if tt.prepare != nil {
				tt.prepare()
			}
			got, err := s.CreateProduct(tt.args.ctx, tt.args.req)
			if errorhelper.GetCode(err) != tt.statusCode {
				t.Errorf("CreateProduct() status code = %v, want %v", errorhelper.GetCode(err), tt.statusCode)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateProduct() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_service_GetProduct(t *testing.T) {
	type args struct {
		ctx context.Context
		id  int64
	}
	tests := []struct {
		name       string
		args       args
		prepare    func()
		want       api.Product
		statusCode int
	}{
		{
			name:       "invalid request",
			args:       args{},
			prepare:    nil,
			want:       api.Product{},
			statusCode: http.StatusBadRequest,
		},
		{
			name: "error when get product",
			args: args{
				id: 5,
			},
			prepare: func() {
				mockProductRepo.On("GetProduct", mock.Anything, int64(5)).
					Return(model.Product{}, errors.New("any"))
			},
			want:       api.Product{},
			statusCode: http.StatusInternalServerError,
		},
		{
			name: "product not found",
			args: args{
				id: 5,
			},
			prepare: func() {
				mockProductRepo.On("GetProduct", mock.Anything, int64(5)).
					Return(model.Product{}, sql.ErrNoRows)
			},
			want:       api.Product{},
			statusCode: http.StatusNotFound,
		},
		{
			name: "success",
			args: args{
				id: 5,
			},
			prepare: func() {
				mockProductRepo.On("GetProduct", mock.Anything, int64(5)).
					Return(model.Product{
						ID:          1,
						SKU:         "IND005",
						Title:       "Test",
						Description: "test",
						Category: model.Category{
							ID:   1,
							Name: "Makanan",
						},
						ImageURL:  "https://foo.bar/image.jpg",
						Weight:    1,
						Price:     1000,
						Rating:    4.5,
						CreatedAt: time.Time{},
					}, nil)
			},
			want: api.Product{
				ID:          1,
				SKU:         "IND005",
				Title:       "Test",
				Description: "test",
				Category: api.Category{
					ID:   1,
					Name: "Makanan",
				},
				ImageURL: "https://foo.bar/image.jpg",
				Weight:   1,
				Price:    1000,
				Rating:   4.5,
			},
			statusCode: http.StatusOK,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			initMock()
			s := &service{
				productRepo:  mockProductRepo,
				categoryRepo: mockCategoryRepo,
				reviewRepo:   mockReviewRepo,
			}
			if tt.prepare != nil {
				tt.prepare()
			}
			got, err := s.GetProduct(tt.args.ctx, tt.args.id)
			if errorhelper.GetCode(err) != tt.statusCode {
				t.Errorf("GetProduct() status code = %v, want %v", errorhelper.GetCode(err), tt.statusCode)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetProduct() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_service_GetProductList(t *testing.T) {
	type args struct {
		ctx    context.Context
		filter api.GetProductListFilter
	}
	tests := []struct {
		name       string
		args       args
		prepare    func()
		want       []api.Product
		statusCode int
	}{
		{
			name: "invalid request",
			args: args{
				ctx: context.Background(),
				filter: api.GetProductListFilter{
					Search:     "",
					CategoryID: 0,
					SortColumn: "wrong",
					SortType:   "wrong",
				},
			},
			prepare:    nil,
			want:       nil,
			statusCode: http.StatusBadRequest,
		},
		{
			name: "error when get product list",
			args: args{
				ctx:    context.Background(),
				filter: api.GetProductListFilter{},
			},
			prepare: func() {
				mockProductRepo.On("GetProductList", mock.Anything, mock.Anything).
					Return(nil, errors.New("any"))
			},
			want:       nil,
			statusCode: http.StatusInternalServerError,
		},
		{
			name: "success",
			args: args{
				ctx:    context.Background(),
				filter: api.GetProductListFilter{},
			},
			prepare: func() {
				mockProductRepo.On("GetProductList", mock.Anything, mock.Anything).
					Return([]model.Product{
						{
							ID:          1,
							SKU:         "IND001",
							Title:       "title",
							Description: "description",
							Category: model.Category{
								ID:   1,
								Name: "name",
							},
							ImageURL: "https://foo.bar/image.jpg",
							Weight:   1,
							Price:    10000,
							Rating:   3.2,
						},
					}, nil)
			},
			want: []api.Product{
				{
					ID:          1,
					SKU:         "IND001",
					Title:       "title",
					Description: "description",
					Category: api.Category{
						ID:   1,
						Name: "name",
					},
					ImageURL: "https://foo.bar/image.jpg",
					Weight:   1,
					Price:    10000,
					Rating:   3.2,
				},
			},
			statusCode: http.StatusOK,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			initMock()
			s := &service{
				productRepo:  mockProductRepo,
				categoryRepo: mockCategoryRepo,
				reviewRepo:   mockReviewRepo,
			}
			if tt.prepare != nil {
				tt.prepare()
			}
			got, err := s.GetProductList(tt.args.ctx, tt.args.filter)
			if errorhelper.GetCode(err) != tt.statusCode {
				t.Errorf("GetProducList() status code = %v, want %v", errorhelper.GetCode(err), tt.statusCode)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetProductList() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_service_ReviewProduct(t *testing.T) {
	type args struct {
		ctx       context.Context
		productID int64
		req       api.ReviewProductRequest
	}
	tests := []struct {
		name       string
		args       args
		prepare    func()
		want       api.MutationResponse
		statusCode int
	}{
		{
			name:       "invalid product id",
			args:       args{},
			prepare:    nil,
			want:       api.MutationResponse{},
			statusCode: http.StatusBadRequest,
		},
		{
			name: "invalid payload request",
			args: args{
				ctx:       context.Background(),
				productID: 2,
				req:       api.ReviewProductRequest{},
			},
			prepare:    nil,
			want:       api.MutationResponse{},
			statusCode: http.StatusBadRequest,
		},
		{
			name: "error when get product",
			args: args{
				ctx:       context.Background(),
				productID: 4,
				req: api.ReviewProductRequest{
					Rating:  4,
					Comment: "comment",
				},
			},
			prepare: func() {
				mockProductRepo.On("GetProduct", mock.Anything, int64(4)).
					Return(model.Product{}, errors.New("any"))
			},
			want:       api.MutationResponse{},
			statusCode: http.StatusInternalServerError,
		},
		{
			name: "product not found",
			args: args{
				ctx:       context.Background(),
				productID: 4,
				req: api.ReviewProductRequest{
					Rating:  4,
					Comment: "comment",
				},
			},
			prepare: func() {
				mockProductRepo.On("GetProduct", mock.Anything, int64(4)).
					Return(model.Product{}, sql.ErrNoRows)
			},
			want:       api.MutationResponse{},
			statusCode: http.StatusNotFound,
		},
		{
			name: "error when get product",
			args: args{
				ctx:       context.Background(),
				productID: 4,
				req: api.ReviewProductRequest{
					Rating:  4,
					Comment: "comment",
				},
			},
			prepare: func() {
				mockProductRepo.On("GetProduct", mock.Anything, int64(4)).
					Return(model.Product{}, sql.ErrNoRows)
			},
			want:       api.MutationResponse{},
			statusCode: http.StatusNotFound,
		},
		{
			name: "error when get product review statistic",
			args: args{
				ctx:       context.Background(),
				productID: 4,
				req: api.ReviewProductRequest{
					Rating:  4,
					Comment: "comment",
				},
			},
			prepare: func() {
				mockProductRepo.On("GetProduct", mock.Anything, int64(4)).
					Return(model.Product{}, nil)
				mockReviewRepo.On("GetReviewStatistic", mock.Anything, int64(4)).
					Return(model.Statistic{}, errors.New("any"))
			},
			want:       api.MutationResponse{},
			statusCode: http.StatusInternalServerError,
		},
		{
			name: "error when insert product review",
			args: args{
				ctx:       context.Background(),
				productID: 4,
				req: api.ReviewProductRequest{
					Rating:  4,
					Comment: "comment",
				},
			},
			prepare: func() {
				mockProductRepo.On("GetProduct", mock.Anything, int64(4)).
					Return(model.Product{}, nil)
				mockReviewRepo.On("GetReviewStatistic", mock.Anything, int64(4)).
					Return(model.Statistic{
						Count:   10,
						Average: 3,
					}, nil)
				mockReviewRepo.On("InsertReview", mock.Anything, mock.Anything).Return(errors.New("any"))
			},
			want:       api.MutationResponse{},
			statusCode: http.StatusInternalServerError,
		},
		{
			name: "error when update product rating",
			args: args{
				ctx:       context.Background(),
				productID: 4,
				req: api.ReviewProductRequest{
					Rating:  4,
					Comment: "comment",
				},
			},
			prepare: func() {
				mockProductRepo.On("GetProduct", mock.Anything, int64(4)).
					Return(model.Product{}, nil)
				mockReviewRepo.On("GetReviewStatistic", mock.Anything, int64(4)).
					Return(model.Statistic{
						Count:   10,
						Average: 3,
					}, nil)
				mockReviewRepo.On("InsertReview", mock.Anything, mock.Anything).Return(nil)
				mockProductRepo.On("UpdateProductRating", mock.Anything, int64(4), float64((3*10)+4)/float64(11)).
					Return(errors.New("any"))
			},
			want:       api.MutationResponse{},
			statusCode: http.StatusInternalServerError,
		},
		{
			name: "success",
			args: args{
				ctx:       context.Background(),
				productID: 4,
				req: api.ReviewProductRequest{
					Rating:  4,
					Comment: "comment",
				},
			},
			prepare: func() {
				mockProductRepo.On("GetProduct", mock.Anything, int64(4)).
					Return(model.Product{}, nil)
				mockReviewRepo.On("GetReviewStatistic", mock.Anything, int64(4)).
					Return(model.Statistic{
						Count:   10,
						Average: 3,
					}, nil)
				mockReviewRepo.On("InsertReview", mock.Anything, mock.Anything).Return(nil)
				mockProductRepo.On("UpdateProductRating", mock.Anything, int64(4), float64((3*10)+4)/float64(11)).
					Return(nil)
			},
			want: api.MutationResponse{
				Success: true,
			},
			statusCode: http.StatusOK,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			initMock()
			s := &service{
				productRepo:  mockProductRepo,
				categoryRepo: mockCategoryRepo,
				reviewRepo:   mockReviewRepo,
			}
			if tt.prepare != nil {
				tt.prepare()
			}
			got, err := s.ReviewProduct(tt.args.ctx, tt.args.productID, tt.args.req)
			if errorhelper.GetCode(err) != tt.statusCode {
				t.Errorf("ReviewProduct() status code = %v, want %v", errorhelper.GetCode(err), tt.statusCode)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReviewProduct() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_service_UpdateProduct(t *testing.T) {

	type args struct {
		ctx context.Context
		req api.Product
		id  int64
	}
	tests := []struct {
		name       string
		args       args
		prepare    func()
		want       api.MutationResponse
		statusCode int
	}{
		{
			name: "invalid id",
			args: args{
				ctx: context.Background(),
				req: api.Product{},
				id:  0,
			},
			prepare:    nil,
			want:       api.MutationResponse{},
			statusCode: http.StatusBadRequest,
		},
		{
			name: "invalid request",
			args: args{
				ctx: context.Background(),
				req: api.Product{},
				id:  5,
			},
			prepare:    nil,
			want:       api.MutationResponse{},
			statusCode: http.StatusBadRequest,
		},
		{
			name: "error when get category",
			args: args{
				ctx: context.Background(),
				id:  5,
				req: api.Product{
					SKU:         "IND005",
					Title:       "title",
					Description: "description",
					Category: api.Category{
						ID: 5,
					},
				},
			},
			prepare: func() {
				mockCategoryRepo.On("GetCategory", mock.Anything, int64(5)).
					Return(model.Category{}, errors.New("any"))
			},
			want:       api.MutationResponse{},
			statusCode: http.StatusInternalServerError,
		},
		{
			name: "category not found",
			args: args{
				ctx: context.Background(),
				id:  5,
				req: api.Product{
					SKU:         "IND005",
					Title:       "title",
					Description: "description",
					Category: api.Category{
						ID: 5,
					},
				},
			},
			prepare: func() {
				mockCategoryRepo.On("GetCategory", mock.Anything, int64(5)).
					Return(model.Category{}, sql.ErrNoRows)
			},
			want:       api.MutationResponse{},
			statusCode: http.StatusBadRequest,
		},
		{
			name: "error when update product",
			args: args{
				ctx: context.Background(),
				id:  5,
				req: api.Product{
					SKU:         "IND005",
					Title:       "title",
					Description: "description",
					Category: api.Category{
						ID: 5,
					},
				},
			},
			prepare: func() {
				mockCategoryRepo.On("GetCategory", mock.Anything, int64(5)).
					Return(model.Category{}, nil)
				mockProductRepo.On("UpdateProduct", mock.Anything, int64(5), mock.Anything).
					Return(errors.New("any"))
			},
			want:       api.MutationResponse{},
			statusCode: http.StatusInternalServerError,
		},
		{
			name: "success",
			args: args{
				ctx: context.Background(),
				id:  5,
				req: api.Product{
					SKU:         "IND005",
					Title:       "title",
					Description: "description",
					Category: api.Category{
						ID: 5,
					},
				},
			},
			prepare: func() {
				mockCategoryRepo.On("GetCategory", mock.Anything, int64(5)).
					Return(model.Category{}, nil)
				mockProductRepo.On("UpdateProduct", mock.Anything, int64(5), mock.Anything).
					Return(nil)
			},
			want: api.MutationResponse{
				Success: true,
			},
			statusCode: http.StatusOK,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			initMock()
			s := &service{
				productRepo:  mockProductRepo,
				categoryRepo: mockCategoryRepo,
				reviewRepo:   mockReviewRepo,
			}
			if tt.prepare != nil {
				tt.prepare()
			}
			got, err := s.UpdateProduct(tt.args.ctx, tt.args.id, tt.args.req)
			if errorhelper.GetCode(err) != tt.statusCode {
				t.Errorf("UpdateProduct() status code = %v, want %v", errorhelper.GetCode(err), tt.statusCode)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdateProduct() got = %v, want %v", got, tt.want)
			}
		})
	}
}
