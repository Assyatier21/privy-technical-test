package repository

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	"privy/config"
	"privy/database"
	m "privy/models"

	"github.com/labstack/echo/v4"
)

func GetListOfCakes(c echo.Context, limit int, offset int) (err error) {
	var (
		rows *sql.Rows
		data []interface{}
	)

	query := fmt.Sprintf(database.GetListOfCakes, limit, offset)
	rows, err = config.DB.Query(query)
	if err != nil {
		log.Println("[GetListOfCakes] can't get list of cakes, err:", err.Error())
		return err
	}

	for rows.Next() {
		var temp = m.Cake{}
		if err := rows.Scan(&temp.Id, &temp.Title, &temp.Description, &temp.Rating, &temp.Image, &temp.CreatedAt, &temp.UpdatedAt); err != nil {
			log.Fatal(err)
			return err
		}
		data = append(data, temp)
	}

	if len(data) > 0 {
		res := m.SetResponse(http.StatusOK, "success", data)
		return c.JSON(http.StatusOK, res)
	} else {
		res := m.SetResponse(http.StatusOK, "data not found", []interface{}{})
		return c.JSON(http.StatusOK, res)
	}
}
func GetDetailsOfCake(c echo.Context, id int) (err error) {
	var (
		cake m.Cake
		data []interface{}
	)

	query := fmt.Sprintf(database.GetDetailsOfCakeByID, id)
	err = config.DB.QueryRow(query).Scan(&cake.Id, &cake.Title, &cake.Description, &cake.Rating, &cake.Image, &cake.CreatedAt, &cake.UpdatedAt)
	if err != nil {
		log.Println("[GetListOfCakes] can't get details of cakes, err:", err.Error())
		res := m.SetResponse(http.StatusOK, "no data found", data)
		return c.JSON(http.StatusOK, res)
	}

	data = append(data, cake)

	if len(data) > 0 {
		res := m.SetResponse(http.StatusOK, "success", data)
		return c.JSON(http.StatusOK, res)
	}
	return
}
func InsertCake(c echo.Context, title string, description string, rating float32, image string) (err error) {
	var (
		cake m.Cake
		data []interface{}
	)
	currentTime := time.Now().String()
	c.Bind(&cake)

	cake.CreatedAt = currentTime
	cake.UpdatedAt = currentTime

	query := fmt.Sprintf(database.InsertCake, cake.Id, cake.Title, cake.Description, cake.Rating, cake.Image, currentTime, currentTime)
	rows, err := config.DB.Exec(query)
	if err != nil {
		log.Println("[InsertCake] can't insert cake, err:", err.Error())
		res := m.SetResponse(http.StatusBadRequest, "failed to insert article", []interface{}{})
		return c.JSON(http.StatusOK, res)
	}

	id, _ := rows.LastInsertId()
	cake.Id = int(id)
	cake.Title = title
	cake.Description = description
	cake.Rating = rating
	cake.Image = image
	cake.CreatedAt = fmt.Sprintf("%d-%d-%d %d:%d:%d", time.Now().Year(), time.Now().Month(), time.Now().Day(), time.Now().Hour(), time.Now().Minute(), time.Now().Second())
	cake.UpdatedAt = fmt.Sprintf("%d-%d-%d %d:%d:%d", time.Now().Year(), time.Now().Month(), time.Now().Day(), time.Now().Hour(), time.Now().Minute(), time.Now().Second())

	data = append(data, cake)

	rowsAffected, _ := rows.RowsAffected()
	if rowsAffected > 0 {
		res := m.SetResponse(http.StatusOK, "success", data)
		return c.JSON(http.StatusOK, res)
	}
	return
}
func UpdateCake(c echo.Context, id int64, title string, description string, rating float32, image string) (err error) {
	var (
		cake       m.Cake
		cakeTemp   m.Cake
		created_at string
		updated_at string
		data       []interface{}
	)
	currentTime := time.Now().String()

	query := fmt.Sprintf(database.GetDetailsOfCakeByID, id)
	_ = config.DB.QueryRow(query).Scan(&cakeTemp.Id, &cakeTemp.Title, &cakeTemp.Description, &cakeTemp.Rating, &cakeTemp.Image, &cakeTemp.CreatedAt, &cakeTemp.UpdatedAt)

	if title == "" {
		title = cakeTemp.Title
	}
	if description == "" {
		description = cakeTemp.Description
	}
	if rating == 0 {
		rating = cake.Rating
	}
	if image == "" {
		image = cakeTemp.Image
	}

	c.Bind(&cake)

	created_at = cakeTemp.UpdatedAt
	updated_at = currentTime

	query = fmt.Sprintf(database.UpdateCakeByID, title, description, rating, image, created_at, updated_at, id)
	rows, err := config.DB.Exec(query)
	if err != nil {
		log.Println("[UpdateCake] can't update cake, err:", err.Error())
		res := m.SetResponse(http.StatusBadRequest, "failed to insert article", []interface{}{})
		return c.JSON(http.StatusOK, res)
	}

	cake.Id = int(id)
	cake.Title = title
	cake.Description = description
	cake.Rating = rating
	cake.Image = image
	cake.CreatedAt = created_at
	cake.UpdatedAt = fmt.Sprintf("%d-%d-%d %d:%d:%d", time.Now().Year(), time.Now().Month(), time.Now().Day(), time.Now().Hour(), time.Now().Minute(), time.Now().Second())

	data = append(data, cake)

	rowsAffected, _ := rows.RowsAffected()
	if rowsAffected > 0 {
		res := m.SetResponse(http.StatusOK, "success", data)
		return c.JSON(http.StatusOK, res)
	}
	return
}
func DeleteCake(c echo.Context, id int) (err error) {
	var (
		data []interface{}
	)

	query := fmt.Sprintf(database.DeleteCakeByID, id)
	row := config.DB.QueryRow(query)
	if row != nil {
		res := m.SetResponse(http.StatusOK, "success", data)
		return c.JSON(http.StatusOK, res)
	} else {
		log.Println("[DeleteCake] can't delete cakes, err:", err.Error())
		res := m.SetResponse(http.StatusOK, "no data found", data)
		return c.JSON(http.StatusOK, res)

	}
}
