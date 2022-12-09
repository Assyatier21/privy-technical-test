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
		return c.JSON(http.StatusBadRequest, m.Error{ErrorMessage: "limit only accept number or can't be empty"})
	} else {
		limit, _ = strconv.Atoi(c.FormValue("limit"))
	}
	if !utils.IsValidNumeric(c.FormValue("offset")) {
		return c.JSON(http.StatusBadRequest, m.Error{ErrorMessage: "offset only accept number or can't be empty"})
	} else {
		offset, _ = strconv.Atoi(c.FormValue("offset"))
	}

	cakes, err := h.repository.GetListOfCakes(c.Request().Context(), limit, offset)
	if err != nil {
		log.Println("[Delivery][GetListOfCakes] can't get list of cakes, err:", err.Error())
		return c.JSON(http.StatusInternalServerError, m.Error{ErrorMessage: err.Error()})
	}

	return c.JSON(http.StatusOK, cakes)
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
		return c.JSON(http.StatusInternalServerError, m.Error{ErrorMessage: err.Error()})
	}

	return c.JSON(http.StatusOK, cake)
}

func (h *handler) InsertCake(c echo.Context) (err error) {
	var (
		cake m.Cake
	)

	if !utils.IsValidAlphaNumericHyphen(c.FormValue("title")) {
		return c.JSON(http.StatusBadRequest, m.Error{ErrorMessage: "title only accept alphanumeric and hypen or title can't be empty"})
	}

	if c.FormValue("description") == "" {
		return c.JSON(http.StatusBadRequest, m.Error{ErrorMessage: "description can't be empty"})
	}

	if !utils.IsValidFloatNumber(c.FormValue("rating")) {
		return c.JSON(http.StatusBadRequest, m.Error{ErrorMessage: "rating only accept float number or can't be empty"})
	}

	if !utils.IsValidLinkImage(c.FormValue("image")) {
		return c.JSON(http.StatusBadRequest, m.Error{ErrorMessage: "image format is wrong or can't be empty"})
	}

	c.Bind(&cake)

	returnCake, err := h.repository.InsertCake(c.Request().Context(), cake)
	if err != nil {
		log.Println("[Delivery][InsertCake] can't insert cake, err:", err.Error())
		return c.JSON(http.StatusInternalServerError, m.Error{ErrorMessage: err.Error()})
	}

	return c.JSON(http.StatusOK, returnCake)
}

func (h *handler) UpdateCake(c echo.Context) (err error) {
	var (
		cake m.Cake
	)

	if !utils.IsValidNumeric(c.Param("id")) {
		return c.JSON(http.StatusBadRequest, m.Error{ErrorMessage: "id only accept number or can't be empty"})
	}

	if !utils.IsValidAlphaNumericHyphen(c.FormValue("title")) {
		return c.JSON(http.StatusBadRequest, m.Error{ErrorMessage: "title only accept alphanumeric and hypen"})
	}

	if !utils.IsValidFloatNumber(c.FormValue("rating")) {
		return c.JSON(http.StatusBadRequest, m.Error{ErrorMessage: "rating only accept float number"})
	}

	if !utils.IsValidLinkImage(c.FormValue("image")) {
		return c.JSON(http.StatusBadRequest, m.Error{ErrorMessage: "image format is wrong"})
	}

	c.Bind(&cake)

	if c.FormValue("rating") == "" {
		cake.Rating = -9999.9999
	}

	returnCake, err := h.repository.UpdateCake(c.Request().Context(), cake)
	if err != nil {
		log.Println("[Delivery][UpdateCake] can't update cake, err:", err.Error())
		return c.JSON(http.StatusInternalServerError, m.Error{ErrorMessage: err.Error()})
	}

	return c.JSON(http.StatusOK, returnCake)
}

func (h *handler) DeleteCake(c echo.Context) (err error) {
	var (
		id int
	)

	if !utils.IsValidNumeric(c.Param("id")) {
		return c.JSON(http.StatusBadRequest, m.Error{ErrorMessage: "id only accept number or can't be empty"})
	} else {
		id, _ = strconv.Atoi(c.Param("id"))
	}

	err = h.repository.DeleteCake(c.Request().Context(), id)
	if err != nil {
		log.Println("[Delivery][DeleteCake] can't get details of cakes, err:", err.Error())
		return c.JSON(http.StatusInternalServerError, m.Error{ErrorMessage: err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "OK"})
}
