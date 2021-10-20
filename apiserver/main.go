package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"net/http"
)

// album represents data about a record album.
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// albums slice to seed record album data.
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

var db *sql.DB

const (
	host     = "localhost"
	port     = 5432
	user     = "ilya"
	password = "1234qwer"
	dbname   = "test"
)

type Department struct {
	id int64
	name string
}

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")


	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.Run("localhost:8080")
}

// getAlbums responds with the list of all albums as JSON.
func getAlbums(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "Content-Type")
	c.IndentedJSON(http.StatusOK, albums)

	rows, err := db.Query("SELECT * FROM department")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		fmt.Println(string(1))
		var department Department
		if err := rows.Scan(&department.id, &department.name); err != nil {
			panic(err)
		}

		fmt.Println("Name: ", department.name)
	}
}
