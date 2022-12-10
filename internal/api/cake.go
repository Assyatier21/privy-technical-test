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
	if !utils.IsValidNumeric(c.FormValue("limit")) {
		res := m.SetResponse(http.StatusBadRequest, "limit only accept number or can't be empty", nil)
		return c.JSON(http.StatusBadRequest, res)
	} else {
		limit, _ = strconv.Atoi(c.FormValue("limit"))
	}
	if !utils.IsValidNumeric(c.FormValue("offset")) {
		res := m.SetResponse(http.StatusBadRequest, "offset only accept number or can't be empty", nil)
		return c.JSON(http.StatusBadRequest, res)
	} else {
		offset, _ = strconv.Atoi(c.FormValue("offset"))
	}

	cakes, err := h.repository.GetListOfCakes(c.Request().Context(), limit, offset)
	if err != nil {
		log.Println("[Delivery][GetListOfCakes] can't get list of cakes, err:", err.Error())
		res := m.SetResponse(http.StatusInternalServerError, err.Error(), nil)
		return c.JSON(http.StatusInternalServerError, res)
	}

	datas := make([]interface{}, len(cakes))
	for i, v := range cakes {
		datas[i] = v
	}
	res := m.SetResponse(http.StatusOK, "success", datas)
	return c.JSON(http.StatusOK, res)
}
func (h *handler) GetDetailsOfCake(c echo.Context) (err error) {
	var (
		id int
	)

	if !utils.IsValidNumeric(c.Param("id")) {
		res := m.SetResponse(http.StatusBadRequest, "id only accept number or can't be empty", nil)
		return c.JSON(http.StatusBadRequest, res)
	} else {
		id, _ = strconv.Atoi(c.Param("id"))
	}

	cake, err := h.repository.GetDetailsOfCake(c.Request().Context(), id)
	if err != nil {
		log.Println("[Delivery][GetDetailsOfCake] can't get details of cakes, err:", err.Error())
		res := m.SetResponse(http.StatusInternalServerError, "[Delivery][GetDetailsOfCake] can't get details of cakes, err:"+err.Error(), nil)
		return c.JSON(http.StatusInternalServerError, res)
	}

	var data []interface{}
	data = append(data, cake)

	res := m.SetResponse(http.StatusOK, "success", data)
	return c.JSON(http.StatusOK, res)
}
func (h *handler) InsertCake(c echo.Context) (err error) {
	var (
		cake m.Cake
	)

	if !utils.IsValidAlphaNumericHyphen(c.FormValue("title")) {
		res := m.SetResponse(http.StatusBadRequest, "title only accept alphanumeric and hypen or title can't be empty", nil)
		return c.JSON(http.StatusBadRequest, res)
	}

	if c.FormValue("description") == "" {
		res := m.SetResponse(http.StatusBadRequest, "description can't be empty", nil)
		return c.JSON(http.StatusBadRequest, res)
	}

	if !utils.IsValidFloatNumber(c.FormValue("rating")) {
		res := m.SetResponse(http.StatusBadRequest, "rating only accept float number or can't be empty", nil)
		return c.JSON(http.StatusBadRequest, res)
	}

	if !utils.IsValidLinkImage(c.FormValue("image")) {
		res := m.SetResponse(http.StatusBadRequest, "image format is wrong or can't be empty", nil)
		return c.JSON(http.StatusBadRequest, res)
	}

	c.Bind(&cake)

	returnCake, err := h.repository.InsertCake(c.Request().Context(), cake)
	if err != nil {
		log.Println("[Delivery][InsertCake] can't insert cake, err:", err.Error())
		res := m.SetResponse(http.StatusInternalServerError, err.Error(), nil)
		return c.JSON(http.StatusInternalServerError, res)
	}

	var data []interface{}
	data = append(data, returnCake)

	res := m.SetResponse(http.StatusOK, "success", data)
	return c.JSON(http.StatusOK, res)
}
func (h *handler) UpdateCake(c echo.Context) (err error) {
	var (
		cake m.Cake
	)

	if !utils.IsValidNumeric(c.Param("id")) {
		res := m.SetResponse(http.StatusBadRequest, "id only accept number or can't be empty", nil)
		return c.JSON(http.StatusBadRequest, res)
	}

	if c.FormValue("title") == "" {
		cake.Title = ""
	} else if !utils.IsValidAlphaNumericHyphen(c.FormValue("title")) {
		res := m.SetResponse(http.StatusBadRequest, "title only accept alphanumeric and hypen or title can't be empty", nil)
		return c.JSON(http.StatusBadRequest, res)
	}

	if c.FormValue("description") == "" {
		cake.Description = ""
	}

	if c.FormValue("rating") == "" {
		cake.Rating = 0
	} else if !utils.IsValidFloatNumber(c.FormValue("rating")) {
		res := m.SetResponse(http.StatusBadRequest, "rating only accept float number or can't be empty", nil)
		return c.JSON(http.StatusBadRequest, res)
	}

	if c.FormValue("image") == "" {
		cake.Image = ""
	} else if !utils.IsValidLinkImage(c.FormValue("image")) {
		res := m.SetResponse(http.StatusBadRequest, "image format is wrong or can't be empty", nil)
		return c.JSON(http.StatusBadRequest, res)
	}

	c.Bind(&cake)

	returnCake, err := h.repository.UpdateCake(c.Request().Context(), cake)
	if err != nil {
		log.Println("[Delivery][UpdateCake] can't update cake, err:", err.Error())
		res := m.SetResponse(http.StatusInternalServerError, err.Error(), nil)
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

	if !utils.IsValidNumeric(c.Param("id")) {
		res := m.SetResponse(http.StatusBadRequest, "id only accept number or can't be empty", nil)
		return c.JSON(http.StatusBadRequest, res)
	} else {
		id, _ = strconv.Atoi(c.Param("id"))
	}

	err = h.repository.DeleteCake(c.Request().Context(), id)
	if err != nil {
		log.Println("[Delivery][DeleteCake] can't get details of cakes, err:", err.Error())
		res := m.SetResponse(http.StatusInternalServerError, err.Error(), nil)
		return c.JSON(http.StatusInternalServerError, res)
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "OK"})
}
