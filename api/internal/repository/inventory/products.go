package inventory

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"gobase/api/internal/model"
	"gobase/api/internal/repository/generator"
	"gobase/api/internal/repository/orm"

	pkgerrors "github.com/pkg/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

const (
	cacheKeyObjTypeActiveProductsCount = "active-products-count"
	cacheKeyActiveProductsCount        = "all"
)

// ProductsFilter holds filters for getting products list
type ProductsFilter struct {
	ExtID    string
	Status   []model.ProductStatus
	WithLock bool
}

// ListProducts gets a list of products from DB
func (i impl) ListProducts(ctx context.Context, filter ProductsFilter) ([]model.Product, error) {
	qms := []qm.QueryMod{
		qm.OrderBy(orm.ProductColumns.ID + " DESC"),
	}
	if filter.Status != nil {
		status := make([]string, len(filter.Status))
		for idx, s := range filter.Status {
			status[idx] = s.String()
		}
		qms = append(qms, orm.ProductWhere.Status.IN(status))
	}
	if filter.ExtID != "" {
		qms = append(qms, orm.ProductWhere.ExternalID.EQ(filter.ExtID))
	}
	if filter.WithLock {
		qms = append(qms, qm.For("UPDATE"))
	}
	slice, err := orm.Products(qms...).All(ctx, i.dbConn)
	if err != nil {
		return nil, pkgerrors.WithStack(err)
	}

	result := make([]model.Product, len(slice))
	for idx, o := range slice {
		result[idx] = model.Product{
			ID:          o.ID,
			ExternalID:  o.ExternalID,
			Name:        o.Name,
			Description: o.Description,
			Status:      model.ProductStatus(o.Status),
			Price:       o.Price,
			CreatedAt:   o.CreatedAt,
			UpdatedAt:   o.UpdatedAt,
		}
	}

	return result, nil
}

// CreateProduct saves product in DB
func (i impl) CreateProduct(ctx context.Context, p model.Product) (model.Product, error) {
	id, err := generator.ProductIDSNF.Generate()
	if err != nil {
		return p, pkgerrors.WithStack(err)
	}
	o := orm.Product{
		//ID:          id, // TODO: Switch to snowflake after changing column to BIGINT non-serial
		ExternalID:  fmt.Sprint(id), // TODO: Should be snowflake as well and be int64
		Name:        p.Name,
		Description: p.Description,
		Status:      p.Status.String(),
		Price:       p.Price,
	}

	if err := o.Insert(ctx, i.dbConn, boil.Infer()); err != nil {
		return p, pkgerrors.WithStack(err)
	}

	p.ID = o.ID
	p.CreatedAt = o.CreatedAt
	p.UpdatedAt = o.UpdatedAt

	return p, nil
}

// UpdateProductStatus updates the product status in DB
func (i impl) UpdateProductStatus(ctx context.Context, p model.Product) (model.Product, error) {
	o, err := orm.Products(orm.ProductWhere.ID.EQ(p.ID)).One(ctx, i.dbConn)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return p, ErrNotFound
		}
		return p, pkgerrors.WithStack(err)
	}

	o.Status = p.Status.String()
	rows, err := o.Update(ctx, i.dbConn, boil.Whitelist(orm.ProductColumns.Status, orm.ProductColumns.UpdatedAt))
	if err != nil {
		return p, pkgerrors.WithStack(err)
	}

	if rows != 1 {
		return p, pkgerrors.WithStack(fmt.Errorf("%w, found: %d", ErrUnexpectedRowsFound, rows))
	}

	p.UpdatedAt = o.UpdatedAt

	return p, nil
}

// GetActiveProductsCountFromDB gets active products count from DB
func (i impl) GetActiveProductsCountFromDB(ctx context.Context) (int64, error) {
	result, err := orm.Products(orm.ProductWhere.Status.EQ(model.ProductStatusActive.String())).Count(ctx, i.dbConn)
	return result, pkgerrors.WithStack(err)
}
