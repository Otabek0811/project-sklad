package handler

import (
	"app/api/models"
	"app/pkg/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Create coming godoc
// @ID create_coming
// @Router /coming [POST]
// @Summary Create Coming
// @Description Create Coming
// @Tags Coming
// @Accept json
// @Procedure json
// @Param coming body models.CreateComing true "CreateComingRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) CreateComing(c *gin.Context) {

	var (
		createComing models.CreateComing
	)
	err := c.ShouldBindJSON(&createComing)
	if err != nil {
		h.handlerResponse(c, "error Coming should bind json", http.StatusBadRequest, err.Error())
		return
	}

	

	id, err := h.strg.Coming().Create(c.Request.Context(), &createComing)
	if err != nil {
		h.handlerResponse(c, "storage.Coming.create", http.StatusInternalServerError, err.Error())
		return
	}

	resp, err := h.strg.Coming().GetByID(c.Request.Context(), &models.ComingPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.Coming.getById", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "create Coming resposne", http.StatusCreated, resp)
}

// GetByID coming godoc
// @ID get_by_id_coming
// @Router /coming/{id} [GET]
// @Summary Get By ID Coming
// @Description Get By ID Coming
// @Tags Coming
// @Accept json
// @Procedure json
// @Param id path string true "id"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) GetByIdComing(c *gin.Context) {

	var id string = c.Param("id")
	if !helper.IsValidUUID(id) {
		h.handlerResponse(c, "is valid uuid", http.StatusBadRequest, "invalid id")
		return
	}

	resp, err := h.strg.Coming().GetByID(c.Request.Context(), &models.ComingPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.Coming.getById", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "get by id Coming resposne", http.StatusOK, resp)
}

// GetList coming godoc
// @ID get_list_coming
// @Router /coming [GET]
// @Summary Get List Coming
// @Description Get List Coming
// @Tags Coming
// @Accept json
// @Procedure json
// @Param offset query string false "offset"
// @Param limit query string false "limit"
// @Param search_by_comingID query string false "search_by_comingID"
// @Param search_by_filial query string false "search_by_filial"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) GetListComing(c *gin.Context) {

	offset, err := h.getOffsetQuery(c.Query("offset"))
	if err != nil {
		h.handlerResponse(c, "get list Coming offset", http.StatusBadRequest, "invalid offset")
		return
	}

	limit, err := h.getLimitQuery(c.Query("limit"))
	if err != nil {
		h.handlerResponse(c, "get list Coming limit", http.StatusBadRequest, "invalid limit")
		return
	}

	resp, err := h.strg.Coming().GetList(c.Request.Context(), &models.ComingGetListRequest{
		Offset: offset,
		Limit:  limit,
		Search_by_comingID: c.Query("search_by_comingID"),
		Search_by_filial: c.Query("search_by_filial"),

	})
	if err != nil {
		h.handlerResponse(c, "storage.Coming.get_list", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "get list Coming response", http.StatusOK, resp)
}

// Update coming godoc
// @ID update_coming
// @Router /coming/{id} [PUT]
// @Summary Update Coming
// @Description Update Coming
// @Tags Coming
// @Accept json
// @Procedure json
// @Param id path string true "id"
// @Param coming body models.UpdateComing true "UpdateComingRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) UpdateComing(c *gin.Context) {

	var (
		id           string = c.Param("id")
		updateComing models.UpdateComing
	)

	if !helper.IsValidUUID(id) {
		h.handlerResponse(c, "is valid uuid", http.StatusBadRequest, "invalid id")
		return
	}

	err := c.ShouldBindJSON(&updateComing)
	if err != nil {
		h.handlerResponse(c, "error Coming should bind json", http.StatusBadRequest, err.Error())
		return
	}

	updateComing.Id = id
	rowsAffected, err := h.strg.Coming().Update(c.Request.Context(), &updateComing)
	if err != nil {
		h.handlerResponse(c, "storage.Coming.update", http.StatusInternalServerError, err.Error())
		return
	}

	if rowsAffected <= 0 {
		h.handlerResponse(c, "storage.Coming.update", http.StatusBadRequest, "now rows affected")
		return
	}

	resp, err := h.strg.Coming().GetByID(c.Request.Context(), &models.ComingPrimaryKey{Id: updateComing.Id})
	if err != nil {
		h.handlerResponse(c, "storage.Coming.getById", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "create Coming response", http.StatusAccepted, resp)
}

// Patch coming godoc
// @ID patch_coming
// @Router /coming/{id} [PATCH]
// @Summary Patch Coming
// @Description Patch Coming
// @Tags Coming
// @Accept json
// @Procedure json
// @Param id path string true "id"
// @Param coming body models.PatchRequest true "PatchComingRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) PatchComing(c *gin.Context) {

	var (
		id          string = c.Param("id")
		patchComing models.PatchRequest
	)

	if !helper.IsValidUUID(id) {
		h.handlerResponse(c, "is valid uuid", http.StatusBadRequest, "invalid id")
		return
	}

	err := c.ShouldBindJSON(&patchComing)
	if err != nil {
		h.handlerResponse(c, "error Coming should bind json", http.StatusBadRequest, err.Error())
		return
	}

	patchComing.ID = id
	rowsAffected, err := h.strg.Coming().Patch(c.Request.Context(), &patchComing)
	if err != nil {
		h.handlerResponse(c, "storage.Coming.patch", http.StatusInternalServerError, err.Error())
		return
	}

	if rowsAffected <= 0 {
		h.handlerResponse(c, "storage.Coming.patch", http.StatusBadRequest, "now rows affected")
		return
	}

	resp, err := h.strg.Coming().GetByID(c.Request.Context(), &models.ComingPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.Coming.getById", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "patch Coming response", http.StatusAccepted, resp)
}

// Delete coming godoc
// @ID delete_coming
// @Router /coming/{id} [DELETE]
// @Summary Delete Coming
// @Description Delete Coming
// @Tags Coming
// @Accept json
// @Procedure json
// @Param id path string true "id"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) DeleteComing(c *gin.Context) {

	var id string = c.Param("id")

	if !helper.IsValidUUID(id) {
		h.handlerResponse(c, "is valid uuid", http.StatusBadRequest, "invalid id")
		return
	}

	err := h.strg.Coming().Delete(c.Request.Context(), &models.ComingPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.Coming.delete", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "delete Coming response", http.StatusAccepted, nil)
}
