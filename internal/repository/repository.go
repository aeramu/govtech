package repository

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/alam/govtech/internal/adapter"
	"github.com/alam/govtech/internal/model"
	"strings"
)

type repository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) adapter.ProductRepository {
	return &repository{db: db}
}

func NewCategoryRepository(db *sql.DB) adapter.CategoryRepository {
	return &repository{db: db}
}

func NewProductReviewRepository(db *sql.DB) adapter.ProductReviewRepository {
	return &repository{db: db}
}

func (r *repository) GetProduct(ctx context.Context, id int64) (model.Product, error) {
	query := `
		SELECT 
		    p.id,
		    p.sku,
		    p.title,
		    p.description,
		    c.id,
		    c.name,
		    p.image_url,
		    p.weight,
		    p.price,
		    p.rating,
		    p.created_at
		FROM products p
		JOIN categories c ON p.category_id = c.id
		WHERE p.id = ?
`
	var res model.Product
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&res.ID,
		&res.SKU,
		&res.Title,
		&res.Description,
		&res.Category.ID,
		&res.Category.Name,
		&res.ImageURL,
		&res.Weight,
		&res.Price,
		&res.Rating,
		&res.CreatedAt,
	)
	if err != nil {
		return model.Product{}, err
	}

	return res, nil
}

func (r *repository) GetProductBySKU(ctx context.Context, sku string) (model.Product, error) {
	query := `
		SELECT 
		    p.id,
		    p.sku,
		    p.title,
		    p.description,
		    c.id,
		    c.name,
		    p.image_url,
		    p.weight,
		    p.price,
		    p.rating,
		    p.created_at
		FROM products p
		JOIN categories c ON p.category_id = c.id
		WHERE p.sku = ?
`
	var res model.Product
	err := r.db.QueryRowContext(ctx, query, sku).Scan(
		&res.ID,
		&res.SKU,
		&res.Title,
		&res.Description,
		&res.Category.ID,
		&res.Category.Name,
		&res.ImageURL,
		&res.Weight,
		&res.Price,
		&res.Rating,
		&res.CreatedAt,
	)
	if err != nil {
		return model.Product{}, err
	}

	return res, nil
}

func (r *repository) GetProductList(ctx context.Context, filter model.GetProductListFilter) ([]model.Product, error) {
	query := `
		SELECT 
		    p.id,
		    p.sku,
		    p.title,
		    p.description,
		    c.id,
		    c.name,
		    p.image_url,
		    p.weight,
		    p.price,
		    p.rating,
		    p.created_at
		FROM products p
		JOIN categories c ON p.category_id = c.id
`

	var args []interface{}

	var filterQuery []string
	if filter.Search != "" {
		filterQuery = append(filterQuery, fmt.Sprintf("(p.title LIKE '%%%s%%' or p.sku LIKE '%%%s%%')", filter.Search, filter.Search))
	}
	if filter.CategoryID > 0 {
		filterQuery = append(filterQuery, "p.category_id = ?")
		args = append(args, filter.CategoryID)
	}
	if len(filterQuery) > 0 {
		query += " WHERE " + strings.Join(filterQuery, " AND ")
	}

	if filter.SortColumn != "" {
		query += fmt.Sprintf(" ORDER BY p.%s %s", filter.SortColumn, filter.SortType)
	}

	query += " LIMIT ? OFFSET ?"
	args = append(args, filter.Limit, filter.Offset)

	var res []model.Product
	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var data model.Product
		err := rows.Scan(
			&data.ID,
			&data.SKU,
			&data.Title,
			&data.Description,
			&data.Category.ID,
			&data.Category.Name,
			&data.ImageURL,
			&data.Weight,
			&data.Price,
			&data.Rating,
			&data.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

		res = append(res, data)
	}

	return res, nil
}

func (r *repository) InsertProduct(ctx context.Context, product model.Product) error {
	query := `
		INSERT INTO products(sku, title, description, category_id, image_url, weight, price, rating)
		VALUES(?, ?, ?, ?, ?, ?, ?, ?)
		
`
	_, err := r.db.ExecContext(ctx, query,
		product.SKU,
		product.Title,
		product.Description,
		product.Category.ID,
		product.ImageURL,
		product.Weight,
		product.Price,
		product.Rating,
	)
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) UpdateProduct(ctx context.Context, id int64, product model.Product) error {
	query := `
		UPDATE 
		    products 
		SET 
		    sku = ?,
		    title = ?,
		    description = ?,
		    category_id = ?
		WHERE id = ?
`
	_, err := r.db.ExecContext(ctx, query, product.SKU, product.Title, product.Description, product.Category.ID, id)
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) UpdateProductRating(ctx context.Context, id int64, rating float64) error {
	query := `
		UPDATE 
		    products 
		SET 
		    rating = ?
		WHERE id = ?
`
	_, err := r.db.ExecContext(ctx, query, rating, id)
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) GetCategory(ctx context.Context, id int64) (model.Category, error) {
	query := `
		SELECT 
		    id,
		    name
		FROM categories 
		WHERE id = ?
`
	var res model.Category
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&res.ID,
		&res.Name,
	)
	if err != nil {
		return model.Category{}, err
	}

	return res, nil
}

func (r *repository) InsertReview(ctx context.Context, review model.ProductReview) error {
	query := `
		INSERT INTO product_reviews(user_id, product_id, rating, comment)
		VALUES(?, ?, ?, ?)
`
	_, err := r.db.ExecContext(ctx, query, review.UserID, review.ProductID, review.Rating, review.Comment)
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) GetReviewStatistic(ctx context.Context, productID int64) (model.Statistic, error) {
	query := `
		SELECT 
		    AVG(rating),
		    COUNT(id)
		FROM product_reviews 
		WHERE product_id = ?
		GROUP BY product_id
`
	var res model.Statistic
	err := r.db.QueryRowContext(ctx, query, productID).Scan(
		&res.Average,
		&res.Count,
	)
	if err != nil {
		return model.Statistic{}, err
	}

	return res, nil
}
