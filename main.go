package main

import (
	"github.com/gin-gonic/gin"
	"github.com/satyasiba001/new-library/handler"
)

func main() {
	router := gin.Default()

	// Member APIs
	router.POST("/addMember", handler.InsertData)
	router.GET("/memberDetails/:member_id", handler.GetMemberdetails)

	// Book APIs
	router.POST("/addnewBook", handler.InsertNewBook)
	router.GET("/booksPresent", handler.BooksPresent)

	// Book Trasanction APIs
	router.POST("/bookBorrow", handler.BookBorrow)

	router.Run(":9000")

}
