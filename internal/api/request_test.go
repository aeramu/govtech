package api

import "testing"

func TestReviewProductRequest_Validate(t *testing.T) {
	type fields struct {
		Review  int32
		Comment string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "negative review",
			fields: fields{
				Review:  -10,
				Comment: "",
			},
			wantErr: true,
		},
		{
			name: "zero review",
			fields: fields{
				Review:  0,
				Comment: "",
			},
			wantErr: true,
		},
		{
			name: "more than 5",
			fields: fields{
				Review:  7,
				Comment: "",
			},
			wantErr: true,
		},
		{
			name: "5",
			fields: fields{
				Review:  5,
				Comment: "",
			},
			wantErr: false,
		},
		{
			name: "1",
			fields: fields{
				Review:  1,
				Comment: "",
			},
			wantErr: false,
		},
		{
			name: "3",
			fields: fields{
				Review:  3,
				Comment: "",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := ReviewProductRequest{
				Rating:  tt.fields.Review,
				Comment: tt.fields.Comment,
			}
			if err := req.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetProductListFilter_Validate(t *testing.T) {
	type fields struct {
		Search     string
		CategoryID int64
		SortColumn string
		SortType   string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "invalid sort type",
			fields: fields{
				Search:     "",
				CategoryID: 0,
				SortColumn: "created_at",
				SortType:   "laksdfj",
			},
			wantErr: true,
		},
		{
			name: "invalid sort column",
			fields: fields{
				Search:     "",
				CategoryID: 0,
				SortColumn: "lksjaf",
				SortType:   "asc",
			},
			wantErr: true,
		},
		{
			name: "created at asc",
			fields: fields{
				Search:     "",
				CategoryID: 0,
				SortColumn: "created_at",
				SortType:   "asc",
			},
			wantErr: false,
		},
		{
			name: "rating desc",
			fields: fields{
				Search:     "",
				CategoryID: 0,
				SortColumn: "rating",
				SortType:   "desc",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := GetProductListFilter{
				Search:     tt.fields.Search,
				CategoryID: tt.fields.CategoryID,
				SortColumn: tt.fields.SortColumn,
				SortType:   tt.fields.SortType,
			}
			if err := filter.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
