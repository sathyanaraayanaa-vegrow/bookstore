package controllers

import (
	"BookStore/initializers"
	"BookStore/model"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateBooks(c *gin.Context) {
	if c.GetString("roles") != "admin" && c.GetString("roles") != "librarian" {
		initializers.LogInfo("Unauthorised access to CreateBooks", "ERROR", uint(c.GetInt("userID")), 0)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorised access"})
		c.Abort()
		return
	}
	var book struct {
		Title        string
		Year         int32
		Prize        float32
		Genre        string `gorm:"column:genre" json:"genre"`
		Isbn         string `gorm:"column:isbn" json:"isbn"`
		AuthorID     int32  `gorm:"column:author_id" json:"author_id"`
		Availability int32
	}

	c.Bind(&book)

	book1 := model.Book{
		Title: book.Title, Year: book.Year, Prize: book.Prize, Genre: book.Genre, Isbn: book.Isbn, AuthorID: book.AuthorID, Availability: book.Availability,
	}
	result := initializers.DB.Create(&book1)

	if result.Error != nil {
		initializers.LogInfo("Cannot Create Book", "ERROR", uint(c.GetInt("userID")), 0)
		c.JSON(400, gin.H{
			"Msg": "Error: Cannot Create Book",
		})
	}
	var s = fmt.Sprintf("%s %d created Book Successfully", c.GetString("roles"), c.GetInt("userID"))
	initializers.LogInfo(s, "INFO", uint(c.GetInt("userID")), 0)
	c.JSON(200, gin.H{
		"Msg": "Book added successfully",
	})
}

func ReadAllBooks(c *gin.Context) {
	var arrBooks []model.Book
	initializers.DB.Find(&arrBooks)
	var s = fmt.Sprintf("User %d Read all books", c.GetInt("userID"))
	initializers.LogInfo(s, "INFO", uint(c.GetInt("userID")), 0)
	c.JSON(200, gin.H{
		"List of Books": arrBooks,
		"Msg":           "Retrieved all books Successfully",
	})
}

func ReadBook(c *gin.Context) {

	var bookvar model.Book
	id := c.Param("id")
	bid, err := strconv.Atoi(id)
	if err != nil {
		return
	}
	initializers.DB.First(&bookvar, id)
	initializers.LogInfo("User Fetched this Book.", "INFO", uint(c.GetInt("userID")), uint(bid))
	c.JSON(200, gin.H{
		"List of Books": bookvar,
		"Msg":           "Retrieved book Successfully",
	})
}

func UpdateBook(c *gin.Context) {
	id := c.Param("id")
	bid, err := strconv.Atoi(id)
	if err != nil {
		return
	}
	if c.GetString("roles") != "admin" && c.GetString("roles") != "librarian" {
		initializers.LogInfo("Unauthorised access to UpdateBooks", "ERROR", uint(c.GetInt("userID")), uint(bid))
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorised access"})
		c.Abort()
		return
	}
	var book struct {
		Title        string
		Year         int32
		Prize        float32
		Genre        string `gorm:"column:genre" json:"genre"`
		Isbn         string `gorm:"column:isbn" json:"isbn"`
		AuthorID     int32  `gorm:"column:author_id" json:"author_id"`
		Availability int32
	}

	c.Bind(&book)

	var bookvar model.Book
	initializers.DB.First(&bookvar, id)

	initializers.DB.Model(&bookvar).Updates(model.Book{
		Title: book.Title, Year: book.Year, Prize: book.Prize, Genre: book.Genre, Isbn: book.Isbn, AuthorID: book.AuthorID, Availability: book.Availability,
	})
	initializers.LogInfo("Updated Book Successfully", "INFO", uint(c.GetInt("userID")), uint(bid))
	c.JSON(200, gin.H{
		"Msg": "Updated Book Succesfully",
	})

}

func DeleteBook(c *gin.Context) {
	id := c.Param("id")
	bid, err := strconv.Atoi(id)
	if err != nil {
		return
	}
	if c.GetString("roles") != "admin" && c.GetString("roles") != "librarian" {
		initializers.LogInfo("Unauthorised access to DeleteBooks", "ERROR", uint(c.GetInt("userID")), uint(bid))
		c.JSON(http.StatusUnauthorized, gin.H{"Msg": "Error: Unauthorised access"})
		c.Abort()
		return
	}
	initializers.DB.Delete(&model.Book{}, id)
	initializers.LogInfo("Deleted Book Successfully", "INFO", uint(c.GetInt("userID")), uint(bid))
	c.JSON(200, gin.H{
		"Msg": "Deleted Book Succesfully",
	})
}

func FetchAuthors(c *gin.Context) {
	a := initializers.Q.Author
	var arrAuthors []struct {
		ID   int64
		Name string
	}
	err := a.WithContext(initializers.CTX).Select(a.ID, a.Name).Scan(&arrAuthors)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		initializers.LogInfo("Error Fetching Authors List", "ERROR", uint(c.GetInt("userID")), 0)
		c.JSON(400, gin.H{
			"Msg": "Error: Authors list is Empty",
		})
		return
	}
	initializers.LogInfo("Authors List Fetched Successfully", "INFO", uint(c.GetInt("userID")), 0)
	c.JSON(200, gin.H{
		"List of Authors": arrAuthors,
		"Msg":             "Fetched Authors List",
	})
}

func AddAuthors(c *gin.Context) {
	if c.GetString("roles") != "admin" && c.GetString("roles") != "librarian" {
		initializers.LogInfo("Unauthorised access to AddAuthors", "ERROR", uint(c.GetInt("userID")), 0)
		c.JSON(http.StatusUnauthorized, gin.H{"Msg": "Error: Unauthorised access"})
		c.Abort()
		return
	}
	a := initializers.Q.Author
	var author model.Author
	c.Bind(&author)
	err := a.WithContext(initializers.CTX).Create(&author)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		initializers.LogInfo("Error Creating Author", "ERROR", uint(c.GetInt("userID")), 0)
		c.JSON(400, gin.H{
			"Msg": "Error: Couldn't crreate Author",
		})
		return
	}
	initializers.LogInfo("Authors Created Successfully", "INFO", uint(c.GetInt("userID")), 0)
	c.JSON(200, gin.H{
		"Author": author,
		"id":     author.ID,
		"Msg":    "Fetched Authors List",
	})
}
