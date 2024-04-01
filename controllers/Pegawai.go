package controllers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"

	_ "github.com/lib/pq"
)

func ListPegawai(c *gin.Context, db *sqlx.DB) {
	query := "SELECT * FROM pegawai WHERE is_active = true"
	datarows := []map[string]interface{}{}

	rows, err := db.Queryx(query)
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "error occurred when trying to run query",
		})
		return
	}
	defer rows.Close()

	for rows.Next() {
		holder := map[string]interface{}{}
		err := rows.MapScan(holder)
		if err != nil {
			fmt.Println(err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  "error",
				"message": "error occurred when trying to generate data from database",
			})
			return
		}
		datarows = append(datarows, holder)
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "OK",
		"data":   datarows,
	})
}

func CreatePegawai(c *gin.Context, db *sqlx.DB) {
	id := uuid.New().String()
	nama := c.PostForm("nama")
	jenis_kelamin := c.PostForm("jenis_kelamin")
	alamat := c.PostForm("alamat")
	is_active := true
	create_date := time.Now()

	_, err := db.Exec("INSERT INTO pegawai(id, nama, jenis_kelamin, alamat, is_active, create_date) VALUES($1, $2, $3, $4, $5, $6)",
		id, nama, jenis_kelamin, alamat, is_active, create_date)
	if err != nil {
		fmt.Println("Error on insert:", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "error occurred when trying to create data",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "OK",
		"message": "create success",
	})
}

func GetPegawai(c *gin.Context, db *sqlx.DB) {
	id := c.PostForm("id")
	var nama, jenis_kelamin, alamat string
	var is_active bool
	var create_date time.Time

	err := db.QueryRowx("SELECT nama, jenis_kelamin, alamat, is_active, create_date FROM pegawai WHERE id = $1", id).
		Scan(&nama, &jenis_kelamin, &alamat, &is_active, &create_date)
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "error occurred when trying to retrieve data",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":      "OK",
		"nama":        nama,
		"jenis_kelamin": jenis_kelamin,
		"alamat":      alamat,
		"is_active":   is_active,
		"create_date": create_date,
	})
}

func UpdatePegawai(c *gin.Context, db *sqlx.DB) {
	id := c.PostForm("id")
	nama := c.PostForm("nama")
	jenis_kelamin := c.PostForm("jenis_kelamin")
	alamat := c.PostForm("alamat")

	_, err := db.Exec("UPDATE pegawai SET nama = $1, jenis_kelamin = $2, alamat = $3 WHERE id = $4",
		nama, jenis_kelamin, alamat, id)
	if err != nil {
		fmt.Println("Error on update:", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "error occurred when trying to update data",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "OK",
		"message": "update success",
	})
}

func DeletePegawai(c *gin.Context, db *sqlx.DB) {
	id := c.PostForm("id")

	_, err := db.Exec("DELETE FROM pegawai WHERE id = $1", id)
	if err != nil {
		fmt.Println("Error on delete:", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "error occurred when trying to delete data",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "OK",
		"message": "delete success",
	})
}
