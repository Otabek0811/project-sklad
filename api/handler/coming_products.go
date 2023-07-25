package handler

import (
	"app/api/models"
	"app/pkg/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Create ComingProducts godoc
// @ID create_comingProducts
// @Router /comingproducts [POST]
// @Summary Create ComingProducts
// @Description Create ComingProducts
// @Tags ComingProducts
// @Accept json
// @Procedure json
// @Param comingProducts body models.CreateComingProducts true "CreateComingProductsRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) CreateComingProducts(c *gin.Context) {

	var createComingProducts models.CreateComingProducts
	err := c.ShouldBindJSON(&createComingProducts)
	if err != nil {
		h.handlerResponse(c, "error comingProducts should bind json", http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.strg.ComingProducts().Create(c.Request.Context(), &createComingProducts)
	if err != nil {
		h.handlerResponse(c, "storage.comingProducts.create", http.StatusInternalServerError, err.Error())
		return
	}

	resp, err := h.strg.ComingProducts().GetByID(c.Request.Context(), &models.ComingProductsPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.comingProducts.getById", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "create comingProducts resposne", http.StatusCreated, resp)
}

// GetByID comingProducts godoc
// @ID get_by_id_comingProducts
// @Router /comingproducts/{id} [GET]
// @Summary Get By ID ComingProducts
// @Description Get By ID ComingProducts
// @Tags ComingProducts
// @Accept json
// @Procedure json
// @Param id path string true "id"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) GetByIdComingProducts(c *gin.Context) {

	var id string = c.Param("id")
	if !helper.IsValidUUID(id) {
		h.handlerResponse(c, "is valid uuid", http.StatusBadRequest, "invalid id")
		return
	}

	resp, err := h.strg.ComingProducts().GetByID(c.Request.Context(), &models.ComingProductsPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.comingProducts.getById", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "get by id comingProducts resposne", http.StatusOK, resp)
}

// GetList comingProducts godoc
// @ID get_list_comingProducts
// @Router /comingproducts [GET]
// @Summary Get List ComingProducts
// @Description Get List ComingProducts
// @Tags ComingProducts
// @Accept json
// @Procedure json
// @Param offset query string false "offset"
// @Param limit query string false "limit"
// @Param SearchByCategory query string false "search_by_category"
// @Param SearchByBarcode query string false "search_by_barcode"
// @Param SearchByComingID query string false "search_by_coming_id"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) GetListComingProducts(c *gin.Context) {

	offset, err := h.getOffsetQuery(c.Query("offset"))
	if err != nil {
		h.handlerResponse(c, "get list comingProducts offset", http.StatusBadRequest, "invalid offset")
		return
	}

	limit, err := h.getLimitQuery(c.Query("limit"))
	if err != nil {
		h.handlerResponse(c, "get list comingProducts limit", http.StatusBadRequest, "invalid limit")
		return
	}

	resp, err := h.strg.ComingProducts().GetList(c.Request.Context(), &models.ComingProductsGetListRequest{
		Offset: offset,
		Limit:  limit,
		SearchByCategoryID: c.Query("search_by_categroy"),
		SearchByBarcode: c.Query("search_by_barode"),
		SearchBYComingID: c.Query("search_by_coming_id"),
	})
	if err != nil {
		h.handlerResponse(c, "storage.comingProducts.get_list", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "get list comingProducts resposne", http.StatusOK, resp)
}

// Update comingProducts godoc
// @ID update_comingProducts
// @Router /comingproducts/{id} [PUT]
// @Summary Update ComingProducts
// @Description Update ComingProducts
// @Tags ComingProducts
// @Accept json
// @Procedure json
// @Param id path string true "id"
// @Param comingProducts body models.UpdateComingProducts true "UpdateComingProductsRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) UpdateComingProducts(c *gin.Context) {

	var (
		id            string = c.Param("id")
		updateComingProducts models.UpdateComingProducts
	)

	if !helper.IsValidUUID(id) {
		h.handlerResponse(c, "is valid uuid", http.StatusBadRequest, "invalid id")
		return
	}

	err := c.ShouldBindJSON(&updateComingProducts)
	if err != nil {
		h.handlerResponse(c, "error comingProducts should bind json", http.StatusBadRequest, err.Error())
		return
	}

	updateComingProducts.Id = id
	rowsAffected, err := h.strg.ComingProducts().Update(c.Request.Context(), &updateComingProducts)
	if err != nil {
		h.handlerResponse(c, "storage.comingProducts.update", http.StatusInternalServerError, err.Error())
		return
	}

	if rowsAffected <= 0 {
		h.handlerResponse(c, "storage.comingProducts.update", http.StatusBadRequest, "now rows affected")
		return
	}

	resp, err := h.strg.ComingProducts().GetByID(c.Request.Context(), &models.ComingProductsPrimaryKey{Id: updateComingProducts.Id})
	if err != nil {
		h.handlerResponse(c, "storage.comingProducts.getById", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "create comingProducts resposne", http.StatusAccepted, resp)
}

// Patch comingProducts godoc
// @ID patch_comingProducts
// @Router /comingproducts/{id} [PATCH]
// @Summary Patch ComingProducts
// @Description Patch ComingProducts
// @Tags ComingProducts
// @Accept json
// @Procedure json
// @Param id path string true "id"
// @Param comingProducts body models.PatchRequest true "PatchComingProductsRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) PatchComingProducts(c *gin.Context) {

	var (
		id           string = c.Param("id")
		patchComingProducts models.PatchRequest
	)

	if !helper.IsValidUUID(id) {
		h.handlerResponse(c, "is valid uuid", http.StatusBadRequest, "invalid id")
		return
	}

	err := c.ShouldBindJSON(&patchComingProducts)
	if err != nil {
		h.handlerResponse(c, "error comingProducts should bind json", http.StatusBadRequest, err.Error())
		return
	}

	patchComingProducts.ID = id
	rowsAffected, err := h.strg.ComingProducts().Patch(c.Request.Context(), &patchComingProducts)
	if err != nil {
		h.handlerResponse(c, "storage.comingProducts.patch", http.StatusInternalServerError, err.Error())
		return
	}

	if rowsAffected <= 0 {
		h.handlerResponse(c, "storage.comingProducts.patch", http.StatusBadRequest, "now rows affected")
		return
	}

	resp, err := h.strg.ComingProducts().GetByID(c.Request.Context(), &models.ComingProductsPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.comingProducts.getById", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "create comingProducts resposne", http.StatusAccepted, resp)
}

// Delete comingProducts godoc
// @ID delete_comingProducts
// @Router /comingproducts/{id} [DELETE]
// @Summary Delete ComingProducts
// @Description Delete ComingProducts
// @Tags ComingProducts
// @Accept json
// @Procedure json
// @Param id path string true "id"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) DeleteComingProducts(c *gin.Context) {

	var id string = c.Param("id")

	if !helper.IsValidUUID(id) {
		h.handlerResponse(c, "is valid uuid", http.StatusBadRequest, "invalid id")
		return
	}

	err := h.strg.ComingProducts().Delete(c.Request.Context(), &models.ComingProductsPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.comingProducts.delete", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "delete comingProducts response", http.StatusAccepted, nil)
}
