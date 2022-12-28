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

type Handler interface {
	GetListOfCakes(c echo.Context) (err error)
	GetDetailsOfCake(c echo.Context) (err error)
	InsertCake(c echo.Context) (err error)
	UpdateCake(c echo.Context) (err error)
	DeleteCake(c echo.Context) (err error)
}

type handler struct {
	repository repository.Repository
}

func New(repository repository.Repository) Handler {
	return &handler{
		repository: repository,
	}
}
func (h *handler) GetListOfCakes(c echo.Context) (err error) {
	var (
		limit  int
		offset int
	)
	if c.FormValue("limit") == "" {
		limit = 100
	} else {
		limit, err = strconv.Atoi(c.FormValue("limit"))
		if err != nil {
			res := m.SetError(http.StatusBadRequest, "limit must be an integer")
			return c.JSON(http.StatusBadRequest, res)
		}
	}

	if c.FormValue("offset") == "" {
		offset = 0
	} else {
		offset, err = strconv.Atoi(c.FormValue("offset"))
		if err != nil {
			res := m.SetError(http.StatusBadRequest, "offset must be an integer")
			return c.JSON(http.StatusBadRequest, res)
		}
	}

	datas, err := h.repository.GetListOfCakes(c.Request().Context(), limit, offset)
	if err != nil {
		log.Println("[Delivery][GetArticles] can't get list of articles, err:", err.Error())
		res := m.SetError(http.StatusInternalServerError, err.Error())
		return c.JSON(http.StatusInternalServerError, res)
	}

	cakes := make([]interface{}, len(datas))
	for i, v := range datas {
		cakes[i] = v
	}
	res := m.SetResponse(http.StatusOK, "success", cakes)
	return c.JSON(http.StatusOK, res)
}
func (h *handler) GetDetailsOfCake(c echo.Context) (err error) {
	var (
		id int
	)

	id, err = strconv.Atoi(c.Param("id"))
	if c.Param("id") == "" || err != nil {
		res := m.SetError(http.StatusBadRequest, "id must be an integer and can't be empty")
		return c.JSON(http.StatusBadRequest, res)
	}

	data, err := h.repository.GetDetailsOfCake(c.Request().Context(), id)
	if err != nil {
		log.Println("[Delivery][GetDetailsOfCake] can't get details of cakes, err:", err.Error())
		res := m.SetError(http.StatusInternalServerError, err.Error())
		return c.JSON(http.StatusInternalServerError, res)
	}

	var cake []interface{}
	cake = append(cake, data)

	res := m.SetResponse(http.StatusOK, "success", cake)
	return c.JSON(http.StatusOK, res)
}
func (h *handler) InsertCake(c echo.Context) (err error) {
	var (
		insertedCake m.Cake
	)

	if !utils.IsValidAlphaNumericHyphen(c.FormValue("title")) {
		res := m.SetError(http.StatusBadRequest, "title only accept alphanumeric and hypen and title can't be empty")
		return c.JSON(http.StatusBadRequest, res)
	}

	if c.FormValue("description") == "" {
		res := m.SetError(http.StatusBadRequest, "description can't be empty")
		return c.JSON(http.StatusBadRequest, res)
	}

	if !utils.IsValidFloatNumber(c.FormValue("rating")) {
		res := m.SetError(http.StatusBadRequest, "rating only accept float number and can't be empty")
		return c.JSON(http.StatusBadRequest, res)
	}

	if !utils.IsValidLinkImage(c.FormValue("image")) {
		res := m.SetError(http.StatusBadRequest, "image format is wrong or can't be empty")
		return c.JSON(http.StatusBadRequest, res)
	}

	c.Bind(&insertedCake)

	returnCake, err := h.repository.InsertCake(c.Request().Context(), insertedCake)
	if err != nil {
		log.Println("[Delivery][InsertCake] can't insert cake, err:", err.Error())
		res := m.SetResponse(http.StatusInternalServerError, err.Error(), nil)
		return c.JSON(http.StatusInternalServerError, res)
	}

	var cake []interface{}
	cake = append(cake, returnCake)

	res := m.SetResponse(http.StatusOK, "success", cake)
	return c.JSON(http.StatusOK, res)
}
func (h *handler) UpdateCake(c echo.Context) (err error) {
	var (
		updatedCake m.Cake
	)

	_, err = strconv.Atoi(c.Param("id"))
	if c.Param("id") == "" || err != nil {
		res := m.SetError(http.StatusBadRequest, "id must be an integer and can't be empty")
		return c.JSON(http.StatusBadRequest, res)
	}

	if c.FormValue("title") == "" {
		updatedCake.Title = ""
	} else if !utils.IsValidAlphaNumericHyphen(c.FormValue("title")) {
		res := m.SetError(http.StatusBadRequest, "title only accept alphanumeric and hypen")
		return c.JSON(http.StatusBadRequest, res)
	}

	if c.FormValue("description") == "" {
		updatedCake.Description = ""
	}

	if c.FormValue("rating") == "" {
		updatedCake.Rating = 0
	} else if !utils.IsValidFloatNumber(c.FormValue("rating")) {
		res := m.SetError(http.StatusBadRequest, "rating only accept float number")
		return c.JSON(http.StatusBadRequest, res)
	}

	if c.FormValue("image") == "" {
		updatedCake.Image = ""
	} else if !utils.IsValidLinkImage(c.FormValue("image")) {
		res := m.SetError(http.StatusBadRequest, "image format is wrong")
		return c.JSON(http.StatusBadRequest, res)
	}

	c.Bind(&updatedCake)

	returnCake, err := h.repository.UpdateCake(c.Request().Context(), updatedCake)
	if err != nil {
		log.Println("[Delivery][UpdateCake] can't update cake, err:", err.Error())
		res := m.SetError(http.StatusInternalServerError, err.Error())
		return c.JSON(http.StatusInternalServerError, res)
	}

	var data []interface{}
	data = append(data, returnCake)

	res := m.SetResponse(http.StatusOK, "success", data)
	return c.JSON(http.StatusOK, res)
}
func (h *handler) DeleteCake(c echo.Context) (err error) {
	var (
		id int
	)

	id, err = strconv.Atoi(c.Param("id"))
	if c.Param("id") == "" || err != nil {
		res := m.SetError(http.StatusBadRequest, "id must be an integer and can't be empty")
		return c.JSON(http.StatusBadRequest, res)
	}

	err = h.repository.DeleteCake(c.Request().Context(), id)
	if err != nil {
		log.Println("[Delivery][DeleteCake] can't delete cake, err:", err.Error())
		res := m.SetError(http.StatusInternalServerError, err.Error())
		return c.JSON(http.StatusInternalServerError, res)
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "OK"})
}
