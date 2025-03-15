package store

import (
	"context"
	"soa-product-management/internal/entity"

	"gorm.io/gorm"
)

type Product struct {
	entity.Product
}

func (Product) TableName() string {
	return "products"
}

func (p Product) GetId() string {
	return p.Id
}

type ProductStore struct {
	*BaseStore[string, Product]
}

func NewProductStore(db *gorm.DB) *ProductStore {
	return &ProductStore{
		BaseStore: NewBaseStore(Product{}, db),
	}
}

type ProductConditionBuilder struct {
	cond                []Option
	includeProductTable bool
}

func NewProductConditionBuilder() *ProductConditionBuilder {
	return &ProductConditionBuilder{
		includeProductTable: false,
		cond:                make([]Option, 0, 10),
	}
}

func (p *ProductConditionBuilder) WithCategory() *ProductConditionBuilder {
	p.cond = append(p.cond, func(db *gorm.DB) *gorm.DB {
		return db.
			Joins(`LEFT JOIN categories ON products.category_id = categories.id`)
	})
	return p
}

func (p *ProductConditionBuilder) WithSupplier() *ProductConditionBuilder {
	p.cond = append(p.cond, func(db *gorm.DB) *gorm.DB {
		return db.
			Joins(`LEFT JOIN suppliers ON products.supplier_id = suppliers.id`)
	})
	return p
}

func (p *ProductConditionBuilder) WithLimit(limit int) *ProductConditionBuilder {
	if limit == 0 {
		limit = 100
	}
	p.cond = append(p.cond, Limit(limit))
	return p
}

func (p *ProductConditionBuilder) WithOffset(offset int) *ProductConditionBuilder {
	p.cond = append(p.cond, Offset(offset))
	return p
}

func (b *ProductConditionBuilder) BuildConditions() []Option {
	return b.cond
}

func (s *ProductStore) GetListProductInformation(ctx context.Context, options ...Option) ([]*entity.ProductInformation, error) {
	db := s.db.WithContext(ctx)
	for _, opt := range options {
		db = opt(db)
	}

	var listProductInformation []*entity.ProductInformation
	if err := db.Table("products").
		Select("products.reference as product_reference, products.name as product_name, products.added_date as added_date_utc, products.status as status, categories.name as product_category, products.price as price, products.stock_city as stock_location, suppliers.name as supplier, products.quantity as available_quantity").
		Scan(&listProductInformation).
		Error; err != nil {
		return nil, err
	}

	return listProductInformation, nil
}

func (s *ProductStore) GetListProductStatisticPerCategory(ctx context.Context) ([]*entity.ProductStatisticPerCategory, error) {
	db := s.db.WithContext(ctx)
	var productStatisticPerCategory []*entity.ProductStatisticPerCategory
	if err := db.Table("products").
		Group("categories.id, categories.name").
		Select("categories.id as category_id, categories.name as category_name, COUNT(*) as total_product").
		Joins("LEFT JOIN categories ON products.category_id = categories.id").
		Scan(&productStatisticPerCategory).
		Error; err != nil {
		return nil, err
	}

	return productStatisticPerCategory, nil
}

func (s *ProductStore) GetTotalProducts(ctx context.Context) (int64, error) {
	db := s.db.WithContext(ctx)
	total := int64(0)
	if err := db.Table("products").
		Select("COUNT(*) as total").
		Scan(&total).
		Error; err != nil {
		return 0, err
	}

	return total, nil
}

func (s *ProductStore) GetListProductStatisticPerSupplier(ctx context.Context) ([]*entity.ProductStatisticPerSupplier, error) {
	db := s.db.WithContext(ctx)
	var productStatisticPerSupplier []*entity.ProductStatisticPerSupplier
	if err := db.Table("products").
		Group("suppliers.id, suppliers.name").
		Select("suppliers.id as supplier_id, suppliers.name as supplier_name, COUNT(*) as total_product").
		Joins("LEFT JOIN suppliers ON products.supplier_id = suppliers.id").
		Scan(&productStatisticPerSupplier).
		Error; err != nil {
		return nil, err
	}

	return productStatisticPerSupplier, nil
}
