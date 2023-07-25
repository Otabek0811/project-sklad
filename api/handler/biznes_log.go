package handler

import (
	"app/api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Scan_Barcode scan_Barcode godoc
// @ID scan_barcode
// @Router /coming/scan-barcode [POST]
// @Summary Post ScanBarcode
// @Description POST ScanBarcode
// @Tags ScanBarcode
// @Accept json
// @Procedure json
// @Param scan_barcode body models.Scan_Barcode true "Scan_BarcodeRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) Scan_Barcode(c *gin.Context) {
	var (
		scanBarcode models.Scan_Barcode
		response    models.ResponseBarcode
	)
	err := c.ShouldBindJSON(&scanBarcode)
	if err != nil {
		h.handlerResponse(c, "error barcode should bind json", http.StatusBadRequest, err.Error())
		return
	}

	check, err := h.strg.Coming().GetByID(c.Request.Context(),&models.ComingPrimaryKey{
		Id: scanBarcode.ComingID,
	})

	if err != nil {
		h.handlerResponse(c, "Not Found This Coming", http.StatusInternalServerError, err.Error())
		return
	}

	if check !=nil {
		h.handlerResponse(c, "Not Found This Coming", http.StatusOK, nil)
		return
	}
	if check.Status == "finished" {
		h.handlerResponse(c, "This process has FINISHED!!!", http.StatusOK, nil)
		return
	}
	resp, err := h.strg.Product().GetList(c.Request.Context(), &models.ProductGetListRequest{
		SearchByBarcode: scanBarcode.Barcode,
	})
	if err != nil {
		h.handlerResponse(c, "Not Found This Product", http.StatusInternalServerError, err.Error())
		return
	}

	response.ComingID=scanBarcode.ComingID
	response.Name=resp.Products[0].Name
	response.Barcode=scanBarcode.Barcode
	response.CategoryId=resp.Products[0].CategoryId

	
	h.handlerResponse(c, "scan get list product response", http.StatusOK, response)
}
