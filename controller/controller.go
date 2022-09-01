package controller

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"test/models"

	"github.com/labstack/echo/v4" //using an aecho framework
)

//retriving the book by using id/
func GetBook(c echo.Context) error {
	id := c.Param("id") //getting id by the entered url
	fmt.Println("id:", id)
	b := new(models.Book)

	if err := c.Bind(b); err != nil {
		return err

	}
	//
	NewId, _ := strconv.Atoi(id)
	if NewId != b.ID {
		return c.JSON(http.StatusNotFound, "Index NOt Found")
	}
	//else return an eroor string
	return c.JSON(http.StatusOK, b)
}

//creating book
func CreateBook(c echo.Context) error {
	fmt.Println("here")

	b := new(models.Book)

	if err := c.Bind(b); err != nil {
		return err

	}

	dbconn := ConnectDB()
	tx := dbconn.Conn.MustBegin()
	//inserting data to database
	_, err := tx.NamedQuery(`INSERT INTO "myschema".book (id,title ,author ) VALUES(:id, :title, :author)`, b)
	if err != nil {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered. Error:\n", r)
				log.Println("Data not Inserted due to :: ", err)
			}
		}()

	}
	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}

	return c.JSON(http.StatusOK, "success")
}

//deleting the book by specific id

func DeleteBook(c echo.Context) error {

	dbconn := ConnectDB()

	requestId := c.Param("id")

	intId, err := strconv.Atoi(requestId)
	if err != nil {
		defer func() {
			if r := recover(); r != nil {
				log.Println(" Cannot Convert String to Int due too ::  ", err)
			}
		}()
	}
	_, err = dbconn.Conn.Exec(`DELETE FROM  "myschema".book WHERE id=$1`, intId)

	if err != nil {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered. Error:\n", r)
				log.Println(" Cannot delete data due to ", err)
			}
		}()
	}

	return c.JSON(http.StatusOK, "deleted")
}

//update the book by index of map
func UpdateBook(c echo.Context) error {
	dbconn := ConnectDB() //connected to databse
	id := c.Param("id")

	b := new(models.Book)

	if err := c.Bind(b); err != nil {
		return err

	}

	intId, err1 := strconv.Atoi(id) //converting string to id
	if err1 != nil {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered. Error:\n", r)
			}
		}()
	}

	updateStmt := `update "myschema".book set "id"=$1 , "title"=$2, "author"=$3 where "id"=$4 `
	_, err := dbconn.Conn.Exec(updateStmt, b.ID, b.Title, b.Author, intId)
	if err != nil {
		fmt.Println("im herr ")
		log.Fatal("Error is ", err)
		// log.Fatal("exited")
	}

	return c.JSON(http.StatusOK, ":: Updated:: ")
}
