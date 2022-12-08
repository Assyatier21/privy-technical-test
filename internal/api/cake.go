package api

import (
	"log"
	"net/http"
	"privy/internal/repository"
	m "privy/models"
	"privy/utils"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetListOfCakes(c echo.Context) (err error) {
	var (
		limit  int
		offset int
	)
	if !utils.IsValidNumeric(c.FormValue("limit")) {
		res := m.SetResponse(http.StatusBadRequest, "limit only accept number or can't be empty", nil)
		return c.JSON(http.StatusOK, res)
	} else {
		limit, _ = strconv.Atoi(c.FormValue("limit"))
	}
	if !utils.IsValidNumeric(c.FormValue("offset")) {
		res := m.SetResponse(http.StatusBadRequest, "offset only accept number or can't be empty", nil)
		return c.JSON(http.StatusOK, res)
	} else {
		offset, _ = strconv.Atoi(c.FormValue("offset"))
	}

	err = repository.GetListOfCakes(c, limit, offset)
	if err != nil {
		log.Println("[Delivery][GetListOfCakes] can't get list of cakes, err:", err.Error())
	}

	return
}
func GetDetailsOfCake(c echo.Context) (err error) {
	var (
		id int
	)

	if !utils.IsValidNumeric(c.Param("id")) {
		res := m.SetResponse(http.StatusBadRequest, "id only accept number or can't be empty", nil)
		return c.JSON(http.StatusOK, res)
	} else {
		id, _ = strconv.Atoi(c.Param("id"))
	}

	err = repository.GetDetailsOfCake(c, id)
	if err != nil {
		log.Println("[Delivery][GetDetailsOfCake] can't get details of cakes, err:", err.Error())
	}

	return
}
func InsertCake(c echo.Context) (err error) {
	var (
		title       string
		description string
		rating      float32
		image       string
	)
	if !utils.IsValidAlphaNumericHyphen(c.FormValue("title")) {
		res := m.SetResponse(http.StatusBadRequest, "title only accept alphanumeric and hypen or title can't be empty", nil)
		return c.JSON(http.StatusOK, res)
	} else {
		title = c.FormValue("title")
	}

	if c.FormValue("description") == "" {
		res := m.SetResponse(http.StatusBadRequest, "description can't be empty", nil)
		return c.JSON(http.StatusOK, res)
	} else {
		description = c.FormValue("description")

	}

	if !utils.IsValidFloatNumber(c.FormValue("rating")) {
		res := m.SetResponse(http.StatusBadRequest, "rating only accept float number or can't be empty", nil)
		return c.JSON(http.StatusOK, res)
	} else {
		temp, _ := strconv.ParseFloat(c.FormValue("rating"), 32)
		rating = float32(temp)
	}

	if !utils.IsValidLinkImage(c.FormValue("image")) {
		res := m.SetResponse(http.StatusBadRequest, "image format is wrong or can't be empty", nil)
		return c.JSON(http.StatusOK, res)
	} else {
		image = c.FormValue("image")
	}

	err = repository.InsertCake(c, title, description, rating, image)
	if err != nil {
		log.Println("[Delivery][InsertCake] can't insert cake, err:", err.Error())
	}

	return
}
func UpdateCake(c echo.Context) (err error) {
	var (
		id          int64
		title       string
		description string
		rating      float32
		image       string
	)

	if !utils.IsValidNumeric(c.Param("id")) {
		res := m.SetResponse(http.StatusBadRequest, "id only accept number or can't be empty", nil)
		return c.JSON(http.StatusOK, res)
	} else {
		id, _ = strconv.ParseInt(c.Param("id"), 10, 64)
	}

	if c.FormValue("title") == "" {
		title = ""
	} else if !utils.IsValidAlphaNumericHyphen(c.FormValue("title")) {
		res := m.SetResponse(http.StatusBadRequest, "title only accept alphanumeric and hypen", nil)
		return c.JSON(http.StatusOK, res)
	} else {
		title = c.FormValue("title")
	}

	if c.FormValue("description") == "" {
		description = ""
	} else {
		description = c.FormValue("description")

	}

	if c.FormValue("rating") == "" {
		rating = -9999.9999
	} else if !utils.IsValidFloatNumber(c.FormValue("rating")) {
		res := m.SetResponse(http.StatusBadRequest, "rating only accept float number", nil)
		return c.JSON(http.StatusOK, res)
	} else {
		temp, _ := strconv.ParseFloat(c.FormValue("rating"), 32)
		rating = float32(temp)
	}

	if !utils.IsValidLinkImage(c.FormValue("image")) {
		image = ""
	} else if !utils.IsValidLinkImage(c.FormValue("image")) {
		res := m.SetResponse(http.StatusBadRequest, "image format is wrong", nil)
		return c.JSON(http.StatusOK, res)
	} else {
		image = c.FormValue("image")
	}

	err = repository.UpdateCake(c, id, title, description, rating, image)
	if err != nil {
		log.Println("[Delivery][UpdateCake] can't update cake, err:", err.Error())
	}

	return
}
func DeleteCake(c echo.Context) (err error) {
	var (
		id int
	)

	if !utils.IsValidNumeric(c.Param("id")) {
		res := m.SetResponse(http.StatusBadRequest, "id only accept number or can't be empty", nil)
		return c.JSON(http.StatusOK, res)
	} else {
		id, _ = strconv.Atoi(c.Param("id"))
	}

	err = repository.DeleteCake(c, id)
	if err != nil {
		log.Println("[Delivery][DeleteCake] can't get details of cakes, err:", err.Error())
	}

	return
}
