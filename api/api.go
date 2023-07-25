package api

import (
	_ "app/api/docs"
	"app/api/handler"
	"app/config"
	"app/pkg/logger"
	"app/storage"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewApi(r *gin.Engine, cfg *config.Config, storage storage.StorageI, logger logger.LoggerI) {

	handler := handler.NewHandler(cfg, storage, logger)

	r.POST("/category", handler.CreateCategory)
	r.GET("/category/:id", handler.GetByIdCategory)
	r.GET("/category", handler.GetListCategory)
	r.PUT("/category/:id", handler.UpdateCategory)
	r.DELETE("/category/:id", handler.DeleteCategory)

	r.POST("/product", handler.CreateProduct)
	r.GET("/product/:id", handler.GetByIdProduct)
	r.GET("/product", handler.GetListProduct)
	r.PUT("/product/:id", handler.UpdateProduct)
	r.PATCH("/product/:id", handler.PatchProduct)
	r.DELETE("/product/:id", handler.DeleteProduct)

	r.POST("/filial", handler.CreateFilial)
	r.GET("/filial/:id", handler.GetByIdFilial)
	r.GET("/filial", handler.GetListFilial)
	r.PUT("/filial/:id", handler.UpdateFilial)
	r.PATCH("/filial/:id", handler.PatchFilial)
	r.DELETE("/filial/:id", handler.DeleteFilial)

	r.POST("/coming", handler.CreateComing)
	r.GET("/coming/:id", handler.GetByIdComing)
	r.GET("/coming", handler.GetListComing)
	r.PUT("/coming/:id", handler.UpdateComing)
	r.PATCH("/coming/:id", handler.PatchComing)
	r.DELETE("/coming/:id", handler.DeleteComing)

	r.POST("/comingproducts", handler.CreateComingProducts)
	r.GET("/comingproducts/:id", handler.GetByIdComingProducts)
	r.GET("/comingproducts", handler.GetListComingProducts)
	r.PUT("/comingproducts/:id", handler.UpdateComingProducts)
	r.PATCH("/comingproducts/:id", handler.PatchComingProducts)
	r.DELETE("/comingproducts/:id", handler.DeleteComingProducts)



	r.POST("/remainder", handler.CreateRemainder)
    r.GET("/remainder/:id", handler.GetByIdRemainder)
    r.GET("/remainder", handler.GetListRemainder)
    r.PUT("/remainder/:id", handler.UpdateRemainder)
    r.PATCH("/remainder/:id", handler.PatchRemainder)
    r.DELETE("/remainder/:id", handler.DeleteRemainder)


	r.POST("/coming/scan-barcode",handler.Scan_Barcode)
	r.POST("/do-income/:coming_id",handler.Send_Coming)


	url := ginSwagger.URL("swagger/doc.json") // The url pointing to API definition
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
}
