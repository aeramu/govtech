package api

import "testing"

func TestProduct_ValidateCreate(t *testing.T) {
	type fields struct {
		ID          int64
		SKU         string
		Title       string
		Description string
		Category    Category
		ImageURL    string
		Weight      int32
		Price       int64
		Rating      float32
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "empty sku",
			fields: fields{
				ID:          3,
				SKU:         "",
				Title:       "Test",
				Description: "test",
				Category: Category{
					ID: 4,
				},
				ImageURL: "https://foo.bar/foo.jpg",
				Weight:   1,
				Price:    1000,
				Rating:   0,
			},
			wantErr: true,
		},
		{
			name: "empty title",
			fields: fields{
				ID:          3,
				SKU:         "SKU001",
				Title:       "",
				Description: "test",
				Category: Category{
					ID: 4,
				},
				ImageURL: "https://foo.bar/foo.jpg",
				Weight:   1,
				Price:    1000,
				Rating:   0,
			},
			wantErr: true,
		},
		{
			name: "empty description",
			fields: fields{
				ID:          3,
				SKU:         "SKU001",
				Title:       "title",
				Description: "",
				Category: Category{
					ID: 4,
				},
				ImageURL: "https://foo.bar/foo.jpg",
				Weight:   1,
				Price:    1000,
				Rating:   0,
			},
			wantErr: true,
		},
		{
			name: "empty category id",
			fields: fields{
				ID:          3,
				SKU:         "SKU001",
				Title:       "title",
				Description: "test",
				Category: Category{
					ID: 0,
				},
				ImageURL: "https://foo.bar/foo.jpg",
				Weight:   1,
				Price:    1000,
				Rating:   0,
			},
			wantErr: true,
		},
		{
			name: "empty image url",
			fields: fields{
				ID:          3,
				SKU:         "SKU001",
				Title:       "title",
				Description: "test",
				Category: Category{
					ID: 5,
				},
				ImageURL: "",
				Weight:   1,
				Price:    1000,
				Rating:   0,
			},
			wantErr: true,
		},
		{
			name: "empty price",
			fields: fields{
				ID:          3,
				SKU:         "SKU001",
				Title:       "title",
				Description: "test",
				Category: Category{
					ID: 5,
				},
				ImageURL: "https://foo.bar/foo.jpg",
				Weight:   1,
				Price:    0,
				Rating:   0,
			},
			wantErr: true,
		},
		{
			name: "success",
			fields: fields{
				ID:          3,
				SKU:         "SKU001",
				Title:       "title",
				Description: "test",
				Category: Category{
					ID: 5,
				},
				ImageURL: "https://foo.bar/foo.jpg",
				Weight:   1,
				Price:    1000,
				Rating:   0,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Product{
				ID:          tt.fields.ID,
				SKU:         tt.fields.SKU,
				Title:       tt.fields.Title,
				Description: tt.fields.Description,
				Category:    tt.fields.Category,
				ImageURL:    tt.fields.ImageURL,
				Weight:      tt.fields.Weight,
				Price:       tt.fields.Price,
				Rating:      tt.fields.Rating,
			}
			if err := p.ValidateCreate(); (err != nil) != tt.wantErr {
				t.Errorf("ValidateCreate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestProduct_ValidateUpdate(t *testing.T) {
	type fields struct {
		ID          int64
		SKU         string
		Title       string
		Description string
		Category    Category
		ImageURL    string
		Weight      int32
		Price       int64
		Rating      float32
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "empty sku",
			fields: fields{
				ID:          3,
				SKU:         "",
				Title:       "Test",
				Description: "test",
				Category: Category{
					ID: 4,
				},
				Weight: 1,
				Price:  1000,
				Rating: 0,
			},
			wantErr: true,
		},
		{
			name: "empty title",
			fields: fields{
				ID:          3,
				SKU:         "SKU001",
				Title:       "",
				Description: "test",
				Category: Category{
					ID: 4,
				},
				Weight: 1,
				Price:  1000,
				Rating: 0,
			},
			wantErr: true,
		},
		{
			name: "empty description",
			fields: fields{
				ID:          3,
				SKU:         "SKU001",
				Title:       "title",
				Description: "",
				Category: Category{
					ID: 4,
				},
				Weight: 1,
				Price:  1000,
				Rating: 0,
			},
			wantErr: true,
		},
		{
			name: "empty category id",
			fields: fields{
				ID:          3,
				SKU:         "SKU001",
				Title:       "title",
				Description: "test",
				Category: Category{
					ID: 0,
				},
				Weight: 1,
				Price:  1000,
				Rating: 0,
			},
			wantErr: true,
		},
		{
			name: "success",
			fields: fields{
				ID:          3,
				SKU:         "SKU001",
				Title:       "title",
				Description: "test",
				Category: Category{
					ID: 5,
				},
				Weight: 1,
				Price:  1000,
				Rating: 0,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Product{
				ID:          tt.fields.ID,
				SKU:         tt.fields.SKU,
				Title:       tt.fields.Title,
				Description: tt.fields.Description,
				Category:    tt.fields.Category,
				ImageURL:    tt.fields.ImageURL,
				Weight:      tt.fields.Weight,
				Price:       tt.fields.Price,
				Rating:      tt.fields.Rating,
			}
			if err := p.ValidateUpdate(); (err != nil) != tt.wantErr {
				t.Errorf("ValidateUpdate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
