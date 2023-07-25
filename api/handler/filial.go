package handler

import (
	"app/api/models"
	"app/pkg/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Create filial godoc
// @ID create_filial
// @Router /filial [POST]
// @Summary Create Filial
// @Description Create Filial
// @Tags Filial
// @Accept json
// @Procedure json
// @Param filial body models.CreateFilial true "CreateFilialRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) CreateFilial(c *gin.Context) {

	var createFilial models.CreateFilial
	err := c.ShouldBindJSON(&createFilial)
	if err != nil {
		h.handlerResponse(c, "error filial should bind json", http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.strg.Filial().Create(c.Request.Context(), &createFilial)
	if err != nil {
		h.handlerResponse(c, "storage.filial.create", http.StatusInternalServerError, err.Error())
		return
	}

	resp, err := h.strg.Filial().GetByID(c.Request.Context(), &models.FilialPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.Filial.getById", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "create Filial resposne", http.StatusCreated, resp)
}

// GetByID filial godoc
// @ID get_by_id_filial
// @Router /filial/{id} [GET]
// @Summary Get By ID Filial
// @Description Get By ID Filial
// @Tags Filial
// @Accept json
// @Procedure json
// @Param id path string true "id"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) GetByIdFilial(c *gin.Context) {

	var id string = c.Param("id")
	if !helper.IsValidUUID(id) {
		h.handlerResponse(c, "is valid uuid", http.StatusBadRequest, "invalid id")
		return
	}

	resp, err := h.strg.Filial().GetByID(c.Request.Context(), &models.FilialPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.Filial.getById", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "get by id Filial resposne", http.StatusOK, resp)
}

// GetList filial godoc
// @ID get_list_filial
// @Router /filial [GET]
// @Summary Get List Filial
// @Description Get List Filial
// @Tags Filial
// @Accept json
// @Procedure json
// @Param offset query string false "offset"
// @Param limit query string false "limit"
// @Param search query string false "search"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) GetListFilial(c *gin.Context) {

	offset, err := h.getOffsetQuery(c.Query("offset"))
	if err != nil {
		h.handlerResponse(c, "get list Filial offset", http.StatusBadRequest, "invalid offset")
		return
	}

	limit, err := h.getLimitQuery(c.Query("limit"))
	if err != nil {
		h.handlerResponse(c, "get list Filial limit", http.StatusBadRequest, "invalid limit")
		return
	}

	resp, err := h.strg.Filial().GetList(c.Request.Context(), &models.FilialGetListRequest{
		Offset: offset,
		Limit:  limit,
		Search: c.Query("search"),
	})
	if err != nil {
		h.handlerResponse(c, "storage.Filial.get_list", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "get list Filial resposne", http.StatusOK, resp)
}

// Update filial godoc
// @ID update_filial
// @Router /filial/{id} [PUT]
// @Summary Update Filial
// @Description Update Filial
// @Tags Filial
// @Accept json
// @Procedure json
// @Param id path string true "id"
// @Param filial body models.UpdateFilial true "UpdateFilialRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) UpdateFilial(c *gin.Context) {

	var (
		id           string = c.Param("id")
		updateFilial models.UpdateFilial
	)

	if !helper.IsValidUUID(id) {
		h.handlerResponse(c, "is valid uuid", http.StatusBadRequest, "invalid id")
		return
	}

	err := c.ShouldBindJSON(&updateFilial)
	if err != nil {
		h.handlerResponse(c, "error Filial should bind json", http.StatusBadRequest, err.Error())
		return
	}

	updateFilial.Id = id
	rowsAffected, err := h.strg.Filial().Update(c.Request.Context(), &updateFilial)
	if err != nil {
		h.handlerResponse(c, "storage.Filial.update", http.StatusInternalServerError, err.Error())
		return
	}

	if rowsAffected <= 0 {
		h.handlerResponse(c, "storage.Filial.update", http.StatusBadRequest, "now rows affected")
		return
	}

	resp, err := h.strg.Filial().GetByID(c.Request.Context(), &models.FilialPrimaryKey{Id: updateFilial.Id})
	if err != nil {
		h.handlerResponse(c, "storage.Filial.getById", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "create Filial resposne", http.StatusAccepted, resp)
}

// Patch filial godoc
// @ID patch_filial
// @Router /filial/{id} [PATCH]
// @Summary Patch Filial
// @Description Patch Filial
// @Tags Filial
// @Accept json
// @Procedure json
// @Param id path string true "id"
// @Param filial body models.PatchRequest true "PatchFilialRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) PatchFilial(c *gin.Context) {

	var (
		id          string = c.Param("id")
		patchFilial models.PatchRequest
	)

	if !helper.IsValidUUID(id) {
		h.handlerResponse(c, "is valid uuid", http.StatusBadRequest, "invalid id")
		return
	}

	err := c.ShouldBindJSON(&patchFilial)
	if err != nil {
		h.handlerResponse(c, "error Filial should bind json", http.StatusBadRequest, err.Error())
		return
	}

	patchFilial.ID = id
	rowsAffected, err := h.strg.Filial().Patch(c.Request.Context(), &patchFilial)
	if err != nil {
		h.handlerResponse(c, "storage.Filial.patch", http.StatusInternalServerError, err.Error())
		return
	}

	if rowsAffected <= 0 {
		h.handlerResponse(c, "storage.Filial.patch", http.StatusBadRequest, "now rows affected")
		return
	}

	resp, err := h.strg.Filial().GetByID(c.Request.Context(), &models.FilialPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.Filial.getById", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "create Filial resposne", http.StatusAccepted, resp)
}

// Delete filial godoc
// @ID delete_filial
// @Router /filial/{id} [DELETE]
// @Summary Delete Filial
// @Description Delete Filial
// @Tags Filial
// @Accept json
// @Procedure json
// @Param id path string true "id"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) DeleteFilial(c *gin.Context) {

	var id string = c.Param("id")

	if !helper.IsValidUUID(id) {
		h.handlerResponse(c, "is valid uuid", http.StatusBadRequest, "invalid id")
		return
	}

	err := h.strg.Filial().Delete(c.Request.Context(), &models.FilialPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.Filial.delete", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "delete Filial resposne", http.StatusNoContent, nil)
}
