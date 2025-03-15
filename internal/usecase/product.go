package usecase

import (
	"context"
	"soa-product-management/internal/entity"
	"soa-product-management/internal/repository/store"
	"soa-product-management/internal/util"
)

type ProductUsecase interface {
	GetListProductInformation(ctx context.Context, req *entity.GetListProductInformationRequest) (*entity.GetListProductInformationResponse, error)
	GetListProductStatisticPerCategory(ctx context.Context) (*entity.GetListProductStatisticPerCategoryResponse, error)
	GetListProductStatisticPerSupplier(ctx context.Context) (*entity.GetListProductStatisticPerSupplierResponse, error)
}

func NewProductUsecase(storage *store.Storage) ProductUsecase {
	return &productUsecase{
		storage: storage,
	}
}

type productUsecase struct {
	storage *store.Storage
}

func (usecase *productUsecase) GetListProductInformation(ctx context.Context, req *entity.GetListProductInformationRequest) (*entity.GetListProductInformationResponse, error) {
	cond := store.NewProductConditionBuilder().WithCategory().WithSupplier().WithLimit(req.Limit).WithOffset(req.Offset).BuildConditions()
	listProductInformationResponse, err := usecase.storage.ProductStore.GetListProductInformation(ctx, cond...)
	if err != nil {
		return nil, err
	}

	for _, p := range listProductInformationResponse {
		p.AddedDate = util.FormatDateTime(p.AddedDateUtc, entity.DateFormatYYYYMMDD)
	}

	return &entity.GetListProductInformationResponse{
		ListProductInformation: listProductInformationResponse,
	}, nil
}

func (usecase *productUsecase) GetListProductStatisticPerCategory(ctx context.Context) (*entity.GetListProductStatisticPerCategoryResponse, error) {
	listProductStatistic, err := usecase.storage.ProductStore.GetListProductStatisticPerCategory(ctx)
	if err != nil {
		return nil, err
	}

	totalProduct, err := usecase.storage.ProductStore.GetTotalProducts(ctx)
	if err != nil {
		return nil, err
	}

	for _, p := range listProductStatistic {
		p.Percentage = float64(p.TotalProduct) / float64(totalProduct)
	}

	return &entity.GetListProductStatisticPerCategoryResponse{
		ListProductStatistic: listProductStatistic,
	}, nil
}

func (usecase *productUsecase) GetListProductStatisticPerSupplier(ctx context.Context) (*entity.GetListProductStatisticPerSupplierResponse, error) {
	listProductStatistic, err := usecase.storage.ProductStore.GetListProductStatisticPerSupplier(ctx)
	if err != nil {
		return nil, err
	}

	totalProduct, err := usecase.storage.ProductStore.GetTotalProducts(ctx)
	if err != nil {
		return nil, err
	}

	for _, p := range listProductStatistic {
		p.Percentage = float64(p.TotalProduct) / float64(totalProduct)
	}

	return &entity.GetListProductStatisticPerSupplierResponse{
		ListProductStatistic: listProductStatistic,
	}, nil
}
