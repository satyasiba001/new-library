package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/satyasiba001/new-library/database"
	 "math/rand"
)

type UserData struct {
	Member_id int    `json:"member_id" db:"member_id"`
	Name      string `json:"name" db:"name"`
	Age       int    `json:"age" db:"age"`
	// Add_date  string `json:"add_date" db:"add_date"`
}

type BookData struct {
	Book_id int    `json:"book_id"`
	Name    string `json:"name"`
	Author  string `json:"author"`
	Count   int    `json:"count"`
}

type BookTransaction struct {
	Member_id int    `json:"member_id" db:"member_id"`
	Name      string `json:"name" db:"name"`
	// Borrow_Date string `json:"borrow_date"`
}

type bookPresentDetails struct {
	Book_clount int
	Book_name   string
}

func GetMemberdetails(c *gin.Context) {
	member_name := c.Param("member")
	db, _ := database.DbConnection()
	var usercopy UserData
	var memberAdmissionDate string

	query := `SELECT * FROM members WHERE name = $1`

	// result , err := db.Exec(query,member_name)
	err := db.QueryRow(query, member_name).Scan(&usercopy.Member_id, &usercopy.Name, &usercopy.Age, &memberAdmissionDate)
	// fmt.Printf("%T\n",result)
	if err != nil {
		fmt.Println("member not forund in DB", err)
		return
	}
	c.JSON(200, gin.H{"Member Details": usercopy})

}

func BooksPresent(c *gin.Context) {
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

func InsertData(c *gin.Context) {

	var userDataa UserData

	err := json.NewDecoder(c.Request.Body).Decode(&userDataa)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse JSON"})
		fmt.Println("error is", err)
		return
	}

	today := time.Now().Truncate(24 * time.Hour)

	db, _ := database.DbConnection()
	// database.DbConnection()
	query := `INSERT INTO members (member_id,name,age,add_date)VALUES($1, $2, $3, $4)`

	_, er := db.Exec(query, userDataa.Member_id, userDataa.Name, userDataa.Age, today)
	if er != nil {
		fmt.Println("not inserted")
		return
	}

	c.JSON(200, gin.H{"Status": 1, "msg": "new member created"})
}

func InsertNewBook(c *gin.Context) {
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

func BookBorrow(c *gin.Context) {
	var booktransactions BookTransaction
	err := json.NewDecoder(c.Request.Body).Decode(&booktransactions)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse JSON"})
		// fmt.Println("error is", err)
		return
	}

	db, _ := database.DbConnection()

	var bookAvailability bookPresentDetails
	query := `SELECT count,Name FROM books WHERE name = $1`

	// var book_id int

	err = db.QueryRow(query, booktransactions.Name).Scan(&bookAvailability.Book_clount, &bookAvailability.Book_name)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "You can't take the book"})
		return
	}

	query2 := `UPDATE books SET count = $1 WHERE name = $2`

	if bookAvailability.Book_clount > 0 {
		_, err := db.Exec(query2, bookAvailability.Book_clount-1, bookAvailability.Book_name)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "You can't get the Book"})
			return
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Book not available"})
	}

	rand.Seed(time.Now().UnixNano())
	id := rand.Int63n(9000000) + 1000000
	today := time.Now().Truncate(24 * time.Hour)

	query3 := `INSERT INTO booktransactions (borrow_id,member_id,book_name,borrow_date,return_date)VALUES($1, $2, $3, $4, $5)`

	_, err3 := db.Exec(query3, id, booktransactions.Member_id, booktransactions.Name, today, nil)
	fmt.Println(err3)
	if err3 != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "You can't borrow the book due to some technical error"})
	}else{
		c.JSON(http.StatusOK, gin.H{"status": "You got the book, return it within 10 days"})
	}
}
