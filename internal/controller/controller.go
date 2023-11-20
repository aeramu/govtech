package controller

import (
	"github.com/alam/govtech/internal/api"
	"github.com/alam/govtech/internal/service"
	"github.com/alam/govtech/internal/util/httphelper"
	"github.com/gorilla/mux"
	"net/http"
)

func NewController(svc service.Service) http.Handler {
	r := mux.NewRouter()

	ctrl := controller{
		svc: svc,
	}

	r.HandleFunc("/products", ctrl.CreateProduct).Methods(http.MethodPost)
	r.HandleFunc("/products", ctrl.GetProductList).Methods(http.MethodGet)
	r.HandleFunc("/products/{productID}", ctrl.GetProduct).Methods(http.MethodGet)
	r.HandleFunc("/products/{productID}", ctrl.UpdateProduct).Methods(http.MethodPut)
	r.HandleFunc("/products/{productID}/action/review", ctrl.ReviewProduct).Methods(http.MethodPost)

	return r
}

type controller struct {
	svc service.Service
}

func (c *controller) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var body api.Product
	err := httphelper.ReadBody(r, &body)
	if err != nil {
		httphelper.WriteError(w, err)
		return
	}

	res, err := c.svc.CreateProduct(r.Context(), body)
	if err != nil {
		httphelper.WriteError(w, err)
		return
	}

	httphelper.Write(w, res)
}

func (c *controller) GetProductList(w http.ResponseWriter, r *http.Request) {
	filter := api.GetProductListFilter{
		Search:     r.URL.Query().Get("search"),
		CategoryID: httphelper.ReadQueryParamInt(r, "category"),
		SortColumn: r.URL.Query().Get("sort"),
		SortType:   r.URL.Query().Get("sort_type"),
		Page:       httphelper.ReadQueryParamInt(r, "page"),
		Size:       httphelper.ReadQueryParamInt(r, "size"),
	}

	res, err := c.svc.GetProductList(r.Context(), filter)
	if err != nil {
		httphelper.WriteError(w, err)
		return
	}

	httphelper.Write(w, res)
}

func (c *controller) GetProduct(w http.ResponseWriter, r *http.Request) {
	id := httphelper.ReadPathVarInt(r, "productID")

	res, err := c.svc.GetProduct(r.Context(), id)
	if err != nil {
		httphelper.WriteError(w, err)
		return
	}

	httphelper.Write(w, res)
}

func (c *controller) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	id := httphelper.ReadPathVarInt(r, "productID")

	var body api.Product
	err := httphelper.ReadBody(r, &body)
	if err != nil {
		httphelper.WriteError(w, err)
		return
	}

	res, err := c.svc.UpdateProduct(r.Context(), id, body)
	if err != nil {
		httphelper.WriteError(w, err)
		return
	}

	httphelper.Write(w, res)
}

func (c *controller) ReviewProduct(w http.ResponseWriter, r *http.Request) {
	id := httphelper.ReadPathVarInt(r, "productID")

	var body api.ReviewProductRequest
	err := httphelper.ReadBody(r, &body)
	if err != nil {
		httphelper.WriteError(w, err)
		return
	}

	res, err := c.svc.ReviewProduct(r.Context(), id, body)
	if err != nil {
		httphelper.WriteError(w, err)
		return
	}

	httphelper.Write(w, res)
}
