package main

import (
	"github.com/gin-gonic/gin"
	"github.com/satyasiba001/new-library/handler"
)

func main() {
	router := gin.Default()

	router.POST("/addMember", handler.InsertData)
	router.POST("/addnewBook", handler.InsertNewBook)
	router.GET("/memberDetails/:member", handler.GetMemberdetails)
	router.GET("/booksPresent", handler.BooksPresent)
	router.POST("/bookBorrow", handler.BookBorrow)

	router.Run(":9000")

}
