package controllers

import (
	"BookStore/initializers"
	"BookStore/model"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateRecord(c *gin.Context) {

	b := initializers.Q.Book

	role := c.GetString("roles")

	var record struct {
		BookID int32
		UserID int32
	}

	c.Bind(&record)

	var userID int
	if role == "librarian" {
		userID = int(record.UserID)
	} else {
		userID = c.GetInt("userID")
		record.UserID = int32(userID)
	}

	record1 := model.BorrowingRecord{
		BooksID: int64(record.BookID), UsersID: int64(userID),
	}

	book, err := b.WithContext(initializers.CTX).Where(b.ID.Eq(int64(record.BookID))).First()
	if errors.Is(err, gorm.ErrRecordNotFound) {
		initializers.LogInfo("Wrong BookID Provided while Creating Record.", "ERROR", uint(c.GetInt("userID")), uint(record.BookID))
		c.JSON(400, gin.H{
			"Msg": "Error: No Books Found with that id",
		})
		return
	}

	var user model.User
	user_avail := initializers.DB.First(&user, userID)
	if user_avail.RowsAffected == 0 {
		initializers.LogInfo("Wrong UserID Provided while Creating Record.", "ERROR", uint(c.GetInt("userID")), uint(userID))
		c.JSON(400, gin.H{
			"Msg": "Error: No Users Found with that id",
		})
		return
	}

	availability := book.Availability
	if availability == 0 {
		initializers.LogInfo("Cannot Borrow Unavailable Books. Error Creating Records.", "ERROR", uint(c.GetInt("userID")), uint(record.BookID))
		c.JSON(400, gin.H{
			"Msg": "Error: Cannot Borrow Books, Unavailable",
		})
		return
	}

	book.Availability -= 1
	initializers.DB.Save(&book)
	result := initializers.DB.Create(&record1)

	if result.Error != nil {
		initializers.LogInfo("Cannot Create Record", "ERROR", uint(c.GetInt("userID")), uint(record.BookID))
		c.JSON(400, gin.H{
			"Msg": "Error: Cannot Create Record",
		})
		return
	}
	initializers.LogInfo("Record Created Successfully", "INFO", uint(c.GetInt("userID")), uint(record.BookID))
	c.JSON(200, gin.H{
		"Msg": "Record added successfully",
	})
}

func ViewRecords(c *gin.Context) {
	b := initializers.Q.Book
	r := initializers.Q.BorrowingRecord
	role := c.GetString("roles")
	userID := c.GetInt("userID")
	fmt.Println(userID)

	var record []struct {
		ID         int64 `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
		BooksID    int64
		Title      string
		Genre      string
		CreatedAt  time.Time `gorm:"column:created_at;not null" json:"created_at"`
		ReturnedAt string    `gorm:"column:returned_at" json:"returned_at"`
		UsersID    int64
	}

	if role == "user" {
		// record_avail := initializers.DB.Model(&model.BorrowingRecord{}).Select("borrowing_records.id, borrowing_records.books_id, books.id, borrowing_records.created_at, borrowing_records.updated_at, borrowing_records.returned_at").Joins("left join books on borrowing_records.books_id = books.id").Where("borrowing_records.users_id = ?", userID).Scan(&record)
		err := r.WithContext(initializers.CTX).Select(r.ID, r.BooksID, b.Title, b.Genre, r.CreatedAt, r.ReturnedAt).LeftJoin(b, r.BooksID.EqCol(b.ID)).Where(r.UsersID.Eq(int64(userID))).Order(r.CreatedAt.Desc()).Scan(&record)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			initializers.LogInfo("No Records available to View for the User", "ERROR", uint(c.GetInt("userID")), 0)
			c.JSON(400, gin.H{
				"Msg": "Error: No Records Available for the user",
			})
			return
		}
		initializers.LogInfo("Successfully Fetched the records to View for User", "INFO", uint(c.GetInt("userID")), 0)
		c.JSON(200, gin.H{
			"List of Records": record,
			"Msg":             "Records Retrieved successfully",
		})
	} else if role == "librarian" || role == "admin" {
		user := c.Param("id")
		var s = 0
		if user == "" {
			s = 0
			err := r.WithContext(initializers.CTX).Select(r.ID, r.BooksID, b.Title, b.Genre, r.CreatedAt, r.ReturnedAt, r.UsersID).LeftJoin(b, r.BooksID.EqCol(b.ID)).Order(r.CreatedAt.Desc()).Scan(&record)
			if errors.Is(err, gorm.ErrRecordNotFound) {
				initializers.LogInfo("No Borrowing Record History to View", "ERROR", uint(c.GetInt("userID")), 0)
				c.JSON(400, gin.H{
					"Msg": "Error: No Borrowing Record History",
				})
				return
			}
		} else {
			us, e := strconv.Atoi(user)
			if e != nil {
				return
			}
			s = us
			err := r.WithContext(initializers.CTX).Select(r.ID, r.BooksID, b.Title, b.Genre, r.CreatedAt, r.ReturnedAt).LeftJoin(b, r.BooksID.EqCol(b.ID)).Where(r.UsersID.Eq(int64(us))).Order(r.CreatedAt.Desc()).Scan(&record)
			if errors.Is(err, gorm.ErrRecordNotFound) {
				initializers.LogInfo("No Records available to View for the User", "ERROR", uint(c.GetInt("userID")), uint(us))
				c.JSON(400, gin.H{
					"Msg": "Error: No Records Available for the user",
				})
				return
			}
		}
		initializers.LogInfo("Successfully Fetched the records for User to View", "INFO", uint(c.GetInt("userID")), uint(s))
		c.JSON(200, gin.H{
			"List of Records": record,
			"Msg":             "Records Retrieved successfully",
		})
	}
}

func ViewRecord(c *gin.Context) {
	recordID := c.Param("id")
	userID := c.GetInt("userID")
	us, e := strconv.Atoi(recordID)
	if e != nil {
		return
	}

	var record model.BorrowingRecord
	record_avail := initializers.DB.First(&record, recordID)
	if record_avail.RowsAffected == 0 || record.ReturnedAt != "" {
		initializers.LogInfo("No Records found with that record id to View", "ERROR", uint(c.GetInt("userID")), uint(us))
		c.JSON(400, gin.H{
			"Msg": "Error: No Records Found with that id",
		})
		return
	}
	if record.UsersID != int64(userID) {
		initializers.LogInfo("Unauthorised Access to this Record ID to View Record", "ERROR", uint(c.GetInt("userID")), uint(us))
		c.JSON(http.StatusUnauthorized, gin.H{"Msg": "Error: Unauthorised access"})
		c.Abort()
		return
	}
	initializers.LogInfo("Record Retrieved Successfully to View", "INFO", uint(c.GetInt("userID")), uint(us))
	c.JSON(200, gin.H{
		"Book": record,
		"Msg":  "Record Retrieved Successfully",
	})
}

func DeleteRecord(c *gin.Context) {
	recordID := c.Param("id")
	userID := c.GetInt("userID")
	role := c.GetString("roles")
	us, e := strconv.Atoi(recordID)
	if e != nil {
		return
	}

	var record model.BorrowingRecord
	record_avail := initializers.DB.First(&record, recordID)
	if record_avail.RowsAffected == 0 || record.ReturnedAt != "" {
		initializers.LogInfo("No Records found with that record id to Delete Record", "ERROR", uint(c.GetInt("userID")), uint(us))
		c.JSON(400, gin.H{
			"Msg": "Error: No Records Found with that id",
		})
		return
	}
	if role == "user" && record.UsersID != int64(userID) {
		initializers.LogInfo("Unauthorised Access to Delete Record", "ERROR", uint(c.GetInt("userID")), uint(us))
		c.JSON(http.StatusUnauthorized, gin.H{"Msg": "Error: Unauthorised access"})
		c.Abort()
		return
	}
	bookID := record.BooksID
	fmt.Println(bookID)
	record.ReturnedAt = time.Now().String()
	initializers.DB.Save(&record)

	var book model.Book
	book_avail := initializers.DB.First(&book, bookID)
	if book_avail.RowsAffected == 0 {
		initializers.LogInfo("No Records found with that Book id to Delete Record", "ERROR", uint(c.GetInt("userID")), uint(bookID))
		c.JSON(400, gin.H{
			"Msg": "Error: No Records Found with that Book id to Delete Record",
		})
		return
	}

	book.Availability += 1
	initializers.DB.Save(&book)

	initializers.LogInfo("Record Deleted Successfully", "INFO", uint(c.GetInt("userID")), uint(us))
	c.JSON(200, gin.H{
		"Msg": "Record deleted successfully",
	})
}

func DeleteRecordWithBook(c *gin.Context) {
	role := c.GetString("roles")

	var s struct {
		BookID int32
		UserID int32
	}

	c.Bind(&s)

	var userID int
	if role == "librarian" {
		userID = int(s.UserID)
	} else {
		userID = c.GetInt("userID")
		s.UserID = int32(userID)
	}

	var record model.BorrowingRecord
	record_avail := initializers.DB.Model(&model.BorrowingRecord{}).Where("users_id = ? AND books_id = ? AND returned_at = ?", userID, s.BookID, "").Find(&record)
	if record_avail.RowsAffected == 0 {
		initializers.LogInfo("No Records found with that Book id to Delete Record", "ERROR", uint(c.GetInt("userID")), uint(s.BookID))
		c.JSON(400, gin.H{
			"Msg": "Error: No Records Found with that Book id",
		})
		return
	}
	if record.UsersID != int64(userID) {
		initializers.LogInfo("Unauthorised Access to Delete Record", "ERROR", uint(c.GetInt("userID")), 0)
		c.JSON(http.StatusUnauthorized, gin.H{"Msg": "Error: Unauthorised access"})
		c.Abort()
		return
	}
	bookID := record.BooksID
	record.ReturnedAt = time.Now().String()
	initializers.DB.Save(&record)

	var book model.Book
	book_avail := initializers.DB.First(&book, bookID)
	if book_avail.RowsAffected == 0 {
		initializers.LogInfo("No Records found with that Book id to Delete Record", "ERROR", uint(c.GetInt("userID")), uint(bookID))
		c.JSON(400, gin.H{
			"Msg": "Error: No Books Found with that id",
		})
		return
	}

	book.Availability += 1
	initializers.DB.Save(&book)
	initializers.LogInfo("Record Deleted Successfully", "INFO", uint(c.GetInt("userID")), uint(bookID))
	c.JSON(200, gin.H{
		"Msg": "Record deleted successfully",
	})
}
