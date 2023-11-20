package main

import (
	"database/sql"
	"github.com/alam/govtech/internal/controller"
	"github.com/alam/govtech/internal/repository"
	"github.com/alam/govtech/internal/service"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
)

func main() {
	db, err := sql.Open("mysql", "root:admin@tcp(localhost:6603)/mysql?parseTime=true")
	if err != nil {
		log.Fatalln("error init db:", err)
	}

	productRepo := repository.NewProductRepository(db)
	reviewRepo := repository.NewProductReviewRepository(db)
	categoryRepo := repository.NewCategoryRepository(db)

	svc := service.NewService(productRepo, categoryRepo, reviewRepo)

	ctrl := controller.NewController(svc)

	log.Println("server started at :8080")
	log.Fatalln(http.ListenAndServe(":8080", ctrl))
}
