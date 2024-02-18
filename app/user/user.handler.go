package user

import (
	"github.com/jihanlugas/pandora/app/jwt"
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
// @Tags User
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path string true "ID"
// @Success      200  {object}	response.Response
// @Failure      500  {object}  response.Response
// @Router /user/{id} [get]
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

	res := response.User(data)

	return response.Success(http.StatusOK, "success", res).SendJSON(c)
}

// Create
// @Tags User
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param req body request.CreateUser true "json req body"
// @Success      200  {object}	response.Response
// @Failure      500  {object}  response.Response
// @Router /user [post]
func (h handler) Create(c echo.Context) error {
	var err error

	loginUser, err := jwt.GetUserLoginInfo(c)
	if err != nil {
		return response.Error(http.StatusBadRequest, err.Error(), response.Payload{}).SendJSON(c)
	}

	req := new(request.CreateUser)
	if err = c.Bind(req); err != nil {
		return response.Error(http.StatusBadRequest, err.Error(), response.Payload{}).SendJSON(c)
	}

	if err = c.Validate(req); err != nil {
		return response.Error(http.StatusBadRequest, "error validation", response.ValidationError(err)).SendJSON(c)
	}

	err = h.usecase.Create(loginUser, req)
	if err != nil {
		return response.Error(http.StatusBadRequest, "data not found", response.Payload{}).SendJSON(c)
	}

	return response.Success(http.StatusOK, "success", response.Payload{}).SendJSON(c)
}

// Update
// @Tags User
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path string true "ID"
// @Param req body request.UpdateUser true "json req body"
// @Success      200  {object}	response.Response
// @Failure      500  {object}  response.Response
// @Router /user/{id} [put]
func (h handler) Update(c echo.Context) error {
	var err error

	loginUser, err := jwt.GetUserLoginInfo(c)
	if err != nil {
		return response.Error(http.StatusBadRequest, err.Error(), response.Payload{}).SendJSON(c)
	}

	id := c.Param("id")
	if id == "" {
		return response.Error(http.StatusBadRequest, "data not found", response.Payload{}).SendJSON(c)
	}

	req := new(request.UpdateUser)
	if err = c.Bind(req); err != nil {
		return response.Error(http.StatusBadRequest, err.Error(), response.Payload{}).SendJSON(c)
	}

	if err = c.Validate(req); err != nil {
		return response.Error(http.StatusBadRequest, "error validation", response.ValidationError(err)).SendJSON(c)
	}

	err = h.usecase.Update(loginUser, id, req)
	if err != nil {
		return response.Error(http.StatusBadRequest, err.Error(), response.Payload{}).SendJSON(c)
	}

	return response.Success(http.StatusOK, "success", response.Payload{}).SendJSON(c)
}

// Delete
// @Tags User
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path string true "ID"
// @Success      200  {object}	response.Response
// @Failure      500  {object}  response.Response
// @Router /user/{id} [delete]
func (h handler) Delete(c echo.Context) error {
	var err error

	loginUser, err := jwt.GetUserLoginInfo(c)
	if err != nil {
		return response.Error(http.StatusBadRequest, err.Error(), response.Payload{}).SendJSON(c)
	}

	id := c.Param("id")
	if id == "" {
		return response.Error(http.StatusBadRequest, "data not found", response.Payload{}).SendJSON(c)
	}

	err = h.usecase.Delete(loginUser, id)
	if err != nil {
		return response.Error(http.StatusBadRequest, err.Error(), response.Payload{}).SendJSON(c)
	}

	return response.Success(http.StatusOK, "success", response.Payload{}).SendJSON(c)
}

// Page
// @Tags User
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param req query request.PageUser false "query string"
// @Success      200  {object}	response.Response
// @Failure      500  {object}  response.Response
// @Router /user/page [get]
func (h handler) Page(c echo.Context) error {
	var err error

	req := new(request.PageUser)
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

// ChangePassword
// @Tags User
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param req body request.ChangePassword true "json req body"
// @Success      200  {object}	response.Response
// @Failure      500  {object}  response.Response
// @Router /user/reset-password [post]
func (h handler) ChangePassword(c echo.Context) error {
	var err error

	loginUser, err := jwt.GetUserLoginInfo(c)
	if err != nil {
		return response.Error(http.StatusBadRequest, err.Error(), response.Payload{}).SendJSON(c)
	}

	req := new(request.ChangePassword)
	if err = c.Bind(req); err != nil {
		return response.Error(http.StatusBadRequest, err.Error(), response.Payload{}).SendJSON(c)
	}

	if err = c.Validate(req); err != nil {
		return response.Error(http.StatusBadRequest, "error validation", response.ValidationError(err)).SendJSON(c)
	}

	err = h.usecase.ChangePassword(loginUser, req)
	if err != nil {
		return response.Error(http.StatusBadRequest, err.Error(), response.Payload{}).SendJSON(c)
	}

	return response.Success(http.StatusOK, "success", response.Payload{}).SendJSON(c)
}
