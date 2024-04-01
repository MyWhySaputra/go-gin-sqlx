package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"

	_ "github.com/lib/pq"
	"github.com/MyWhySaputra/go-gin-sqlx/controllers"
)

func main() {
	// set up database
	username := "postgres"
	password := "postgres"
	database := "go-gin-sqlx"
	host := "localhost"
	port := "54320" // PostgreSQL default port

	connectionString := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable", username, password, database, host, port)

	db, err := sqlx.Connect("postgres", connectionString)
	if err != nil {
		fmt.Println("Connection error:", err)
		return
	}
	defer db.Close()

	db.SetConnMaxLifetime(30 * time.Second)
	db.SetConnMaxIdleTime(30)

	// set up gin
	r := gin.New()
	routing(r, db)

	// server dengan menggunakan http
	s := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}
	fmt.Println("Server running on port 8080")
	s.ListenAndServe()
}

func routing(r *gin.Engine, db *sqlx.DB) {
	routePegawai := r.Group("/pegawai")
	{
		routePegawai.GET("/list", func(ctx *gin.Context) { controllers.ListPegawai(ctx, db) })
		routePegawai.POST("/create", func(ctx *gin.Context) { controllers.CreatePegawai(ctx, db) })
		routePegawai.POST("/get", func(ctx *gin.Context) { controllers.GetPegawai(ctx, db) })
		routePegawai.POST("/update", func(ctx *gin.Context) { controllers.UpdatePegawai(ctx, db) })
		routePegawai.POST("/delete", func(ctx *gin.Context) { controllers.DeletePegawai(ctx, db) })
	}
}
