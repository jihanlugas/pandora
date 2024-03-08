package district

import (
	"github.com/jihanlugas/pandora/request"
	"github.com/jihanlugas/pandora/response"
	"github.com/labstack/echo/v4"
	"net/http"
)

type handler struct {
	usecase Usecase
}

func Handler(usecase Usecase) handler {
	return handler{
		usecase: usecase,
	}
}

// GetById
// @Tags District
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path string true "ID"
// @Success      200  {object}	response.Response
// @Failure      500  {object}  response.Response
// @Router /district/{id} [get]
func (h handler) GetById(c echo.Context) error {
	var err error

	id := c.Param("id")
	if id == "" {
		return response.Error(http.StatusBadRequest, "data not found", response.Payload{}).SendJSON(c)
	}

	data, err := h.usecase.GetById(id)
	if err != nil {
		return response.Error(http.StatusBadRequest, "data not found", response.Payload{}).SendJSON(c)
	}

	res := response.District(data)

	return response.Success(http.StatusOK, "success", res).SendJSON(c)
}

// Page
// @Tags District
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param req query request.PageDistrict false "query string"
// @Success      200  {object}	response.Response
// @Failure      500  {object}  response.Response
// @Router /district/page [get]
func (h handler) Page(c echo.Context) error {
	var err error

	req := new(request.PageDistrict)
	if err = c.Bind(req); err != nil {
		return response.Error(http.StatusBadRequest, err.Error(), response.Payload{}).SendJSON(c)
	}

	if err = c.Validate(req); err != nil {
		return response.Error(http.StatusBadRequest, "error validation", response.ValidationError(err)).SendJSON(c)
	}

	data, count, err := h.usecase.Page(req)
	if err != nil {
		return response.Error(http.StatusBadRequest, err.Error(), response.Payload{}).SendJSON(c)
	}

	return response.Success(http.StatusOK, "success", response.PayloadPagination(req, data, count)).SendJSON(c)
}

// List
// @Tags District
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param req query request.ListDistrict false "query string"
// @Success      200  {object}	response.Response
// @Failure      500  {object}  response.Response
// @Router /district/list [get]
func (h handler) List(c echo.Context) error {
	var err error

	req := new(request.ListDistrict)
	if err = c.Bind(req); err != nil {
		return response.Error(http.StatusBadRequest, err.Error(), response.Payload{}).SendJSON(c)
	}

	if err = c.Validate(req); err != nil {
		return response.Error(http.StatusBadRequest, "error validation", response.ValidationError(err)).SendJSON(c)
	}

	data, err := h.usecase.List(req)
	if err != nil {
		return response.Error(http.StatusBadRequest, err.Error(), response.Payload{}).SendJSON(c)
	}

	res := response.Districts(data)

	return response.Success(http.StatusOK, "success", res).SendJSON(c)
}
