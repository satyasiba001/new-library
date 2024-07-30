package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/satyasiba001/new-library/database"
)

type UserData struct {
	Member_id int    `json:"member_id" db:"member_id"`
	Name      string `json:"name" db:"name"`
	Age       int    `json:"age" db:"age"`
	Add_date  string `json:"add_date" db:"add_date"`
}

type BookData struct {
	Book_id int    `json:"book_id"`
	Name    string `json:"name"`
	Author  string `json:"author"`
	Count   int    `json:"count"`
}

type BookTransaction struct {
	Member_id   int    `json:"member_id" db:"member_id"`
	Name        string `json:"name" db:"name"`
	Borrow_Date string `json:"borrow_date"`
}

func insertData(c *gin.Context) {

	var userDataa UserData

	err := json.NewDecoder(c.Request.Body).Decode(&userDataa)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse JSON"})
		fmt.Println("error is", err)
		return
	}

	db, _ := database.DbConnection()
	// database.DbConnection()
	query := `INSERT INTO members (member_id,name,age,add_date)VALUES($1, $2, $3, $4)`

	_, er := db.Exec(query, userDataa.Member_id, userDataa.Name, userDataa.Age, userDataa.Add_date)
	if er != nil {
		fmt.Println("not inserted")
		return
	}

	c.JSON(200, gin.H{"Status": 1, "msg": "new member created"})
}

func insertNewBook(c *gin.Context) {
	var newBoook BookData

	err := json.NewDecoder(c.Request.Body).Decode(&newBoook)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse JSON"})
		fmt.Println("error is", err)
		return
	}

	db, _ := database.DbConnection()

	query := `INSERT INTO books (book_id,name,author,count)VALUES($1, $2, $3, $4)`

	_, er := db.Exec(query, newBoook.Book_id, newBoook.Name, newBoook.Author, newBoook.Count)
	if er != nil {
		fmt.Println("no new book added to the library")
		return
	}

	c.JSON(200, gin.H{"Status": 1, "msg": "new book arrived to the library"})
}

func getMemberdetails(c *gin.Context) {
	member_name := c.Param("member")
	db, _ := database.DbConnection()
	var usercopy UserData

	query := `SELECT * FROM members WHERE name = $1`

	// result , err := db.Exec(query,member_name)
	err := db.QueryRow(query, member_name).Scan(&usercopy.Member_id, &usercopy.Name, &usercopy.Age, &usercopy.Add_date)
	// fmt.Printf("%T\n",result)
	if err != nil {
		fmt.Println("member not forund in DB", err)
		return
	}
	c.JSON(200, gin.H{"Member Details": usercopy})

}

func booksPresent(c *gin.Context) {
	db, _ := database.DbConnection()
	query := `SELECT name FROM books WHERE count != 0`

	result, err := db.Query(query)
	// result, err := db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
	defer result.Close()

	var Names []string

	// Iterate over the rows
	for result.Next() {
		var name string
		if err := result.Scan(&name); err != nil {
			log.Fatal(err)
		}
		Names = append(Names, name)
	}
	c.JSON(200, gin.H{"All Bools Present": Names})
}

func bookBorrow(c *gin.Context) {
	var booktransactions BookTransaction
	err := json.NewDecoder(c.Request.Body).Decode(&booktransactions)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse JSON"})
		// fmt.Println("error is", err)
		return
	}

	// rand.Seed(time.Now().UnixNano())
	// randomNumber := rand.Intn(90000) + 10000

	db, _ := database.DbConnection()

	query := `SELECT count,book_id FROM books WHERE name = $1`
	var bookcount int
	// var book_id int

	err = db.QueryRow(query, booktransactions.Name).Scan(&bookcount)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Input Invalid"})
		return
	}
	query2 := `UPDATE books SET count = $1 WHERE name = $2`

	if bookcount > 0 {
		_, err := db.Exec(query2, bookcount-1, booktransactions.Name)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Book not available"})
			return
		}
	}

	// query3 := `INSERT INTO booktransactions (borrow_id,member_id,book_id,borrow_date,return_date)VALUES($1, $2, $3, $4, $5) `
	// _, err2 := db.Exec(query3, randomNumber, booktransactions.Member_id, booktransactions.Name, booktransactions.Borrow_Date, nil)
	// if err2 != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": "server down, you can't borrow book now"})
	// 	return
	// }
}

func main() {
	router := gin.Default()

	router.POST("/addMember", insertData)
	router.POST("/addnewBook", insertNewBook)
	router.GET("/memberDetails/:member", getMemberdetails)
	router.GET("/booksPresent", booksPresent)
	router.POST("/bookBorrow", bookBorrow)

	router.Run(":9000")

}
