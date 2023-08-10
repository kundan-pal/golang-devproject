package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

// Define PostgreSQL connection parameters
const connStr = "user=postgres password=postgres dbname=nifty sslmode=disable"

type Item struct {
	Date  time.Time
	Price float64
}

func handler(c *gin.Context) {
	fmt.Printf("entered handler\n")

	var inputData struct {
		StockName string `form:"stockName" binding:"required"`
		StartDate string `form:"startDate" binding:"required"`
		EndDate   string `form:"endDate" binding:"required"`
	}

	// Open a connection to the database
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	fmt.Printf("connection request to db sent \n")

	if err := c.ShouldBind(&inputData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Println(inputData)

	// convert from string to time.Time object
	layout := "2006-01-02"

	parsedStartDate, errStartDate := time.Parse(layout, inputData.StartDate)
	if errStartDate != nil {
		fmt.Println("Error:", errStartDate)
		return
	}

	parsedEndDate, errEndDate := time.Parse(layout, inputData.EndDate)
	if errEndDate != nil {
		fmt.Println("Error:", errEndDate)
		return
	}

	rows, err := db.Query("SELECT nifty_date, \""+inputData.StockName+"\" FROM summary where nifty_date >= $1 and nifty_date <= $2 ",
		parsedStartDate, parsedEndDate)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var items [2]Item

	// Process the query results
	for rows.Next() {
		var item Item
		if err := rows.Scan(&item.Date, &item.Price); err != nil {
			log.Fatal(err)
		}
		if items[0].Price < item.Price || items[0].Price == 0 {
			items[0].Price = item.Price
			items[0].Date = item.Date
		}
		if items[1].Price > item.Price || items[1].Price == 0 {
			items[1].Price = item.Price
			items[1].Date = item.Date
		}
	}

	fmt.Println(items)
	c.JSON(http.StatusOK, items)

}

func main() {
	// Gin connection
	r := gin.Default()
	r.POST("/getdata", handler)
	r.Run(":3680")
}
