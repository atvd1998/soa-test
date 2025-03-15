package handler

import (
	"bytes"
	"fmt"
	"net/http"
	"soa-product-management/internal/entity"
	"soa-product-management/internal/usecase"

	_ "soa-product-management/docs"

	"github.com/jung-kurt/gofpdf"
	"github.com/labstack/echo/v4"
)

type ProductHandler interface {
	GetListProductInformationHandler(c echo.Context) error
	GenerateProductInfromationPdf(c echo.Context) error
	GetListProductStatisticPerCategory(c echo.Context) error
	GetListProductStatisticPerSupplier(c echo.Context) error
}

func NewProductHanlder(productUsecase usecase.ProductUsecase) ProductHandler {
	return &productHandler{
		productUsecase: productUsecase,
	}
}

type productHandler struct {
	productUsecase usecase.ProductUsecase
}

// @title Product API
// @version 1.0
// @description API for managing products.
// @host localhost:1323
// @BasePath /api
// @schemes http

// @Summary Get list of products
// @Description Retrieves a list of products.
// @Produce json
// @Success 200 {array} Product
// @Router /products [get]

func (handler *productHandler) GetListProductInformationHandler(c echo.Context) error {
	params := new(entity.GetListProductInformationRequest)
	if err := c.Bind(params); err != nil {
		return c.JSON(http.StatusBadRequest, &entity.FailedResponse{
			Success: false,
		})
	}

	resp, err := handler.productUsecase.GetListProductInformation(c.Request().Context(), params)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, &entity.FailedResponse{
			Success: false,
		})
	}
	return c.JSON(http.StatusOK, &entity.SuccessFulResponse{
		Success: true,
		Data:    resp,
	})
}

func (handler *productHandler) GenerateProductInfromationPdf(c echo.Context) error {
	params := new(entity.GetListProductInformationRequest)
	if err := c.Bind(params); err != nil {
		return c.JSON(http.StatusBadRequest, &entity.FailedResponse{
			Success: false,
		})
	}
	resp, err := handler.productUsecase.GetListProductInformation(c.Request().Context(), params)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, &entity.FailedResponse{
			Success: false,
		})
	}
	pdf := gofpdf.New("P", "mm", "A3", "")
	pdf.AddPage()

	pageWidth := 297.0
	margin := 10.0
	usableWidth := pageWidth - 2*margin

	numColumns := 8
	approxColumnWidth := usableWidth / float64(numColumns)

	// Adjusted widths
	refWidth := approxColumnWidth * 1.2
	nameWidth := approxColumnWidth * 1.3
	dateWidth := approxColumnWidth
	statusWidth := approxColumnWidth
	categoryWidth := approxColumnWidth
	priceWidth := approxColumnWidth
	supplierWidth := approxColumnWidth * 1.2
	qtyWidth := approxColumnWidth

	pdf.SetFont("Arial", "B", 11)

	pdf.Cell(refWidth, 10, "Product Reference")
	pdf.Cell(nameWidth, 10, "Product name")
	pdf.Cell(dateWidth, 10, "Date Added")
	pdf.Cell(statusWidth, 10, "Status")
	pdf.Cell(categoryWidth, 10, "Product Category")
	pdf.Cell(priceWidth, 10, "Price")
	pdf.Cell(supplierWidth, 10, "Supplier")
	pdf.Cell(qtyWidth, 10, "Available Quantity")
	pdf.Ln(10)

	pdf.SetFont("Arial", "", 11)
	for _, item := range resp.ListProductInformation {
		pdf.Cell(refWidth, 10, item.ProductReference)
		pdf.Cell(nameWidth, 10, item.ProductName)
		pdf.Cell(dateWidth, 10, item.AddedDate)
		pdf.Cell(statusWidth, 10, item.Status)
		pdf.Cell(categoryWidth, 10, item.ProductCategory)
		pdf.Cell(priceWidth, 10, fmt.Sprintf("%.2f", item.Price))
		pdf.Cell(supplierWidth, 10, item.Supplier)
		pdf.Cell(qtyWidth, 10, fmt.Sprintf("%d", item.AvailableQuantity))
		pdf.Ln(10)
	}

	var buf bytes.Buffer
	err = pdf.Output(&buf)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, &entity.FailedResponse{
			Success: false,
		})
	}

	c.Response().Header().Set(echo.HeaderContentType, "application/pdf")
	c.Response().Header().Set(echo.HeaderContentDisposition, "attachment; filename=products.pdf")
	return c.Blob(http.StatusOK, "application/pdf", buf.Bytes())
}

func (handler *productHandler) GetListProductStatisticPerCategory(c echo.Context) error {
	resp, err := handler.productUsecase.GetListProductStatisticPerCategory(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, &entity.FailedResponse{
			Success: false,
		})
	}
	return c.JSON(http.StatusOK, &entity.SuccessFulResponse{
		Success: true,
		Data:    resp,
	})
}

func (handler *productHandler) GetListProductStatisticPerSupplier(c echo.Context) error {
	resp, err := handler.productUsecase.GetListProductStatisticPerSupplier(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, &entity.FailedResponse{
			Success: false,
		})
	}
	return c.JSON(http.StatusOK, &entity.SuccessFulResponse{
		Success: true,
		Data:    resp,
	})
}
