package dashboard

import (
	"context"
	"errors"
	"example-microservice-api-go/internal/repo"
	"example-microservice-api-go/pkg/model"
	"github.com/shopspring/decimal"
)

var ErrOrderNotFound = errors.New("requested Order not found")
var ErrProductNotFound = errors.New("requested product not found")
var ErrVendorNotFound = errors.New("requested vendor not found")

type dashboardRepo interface {
	GetOrder(ctx context.Context, id model.OrderID) (*model.Order, error)
	PutOrder(ctx context.Context, id model.OrderID, order *model.Order) (model.OrderID, error)
	GetOrdersByProduct(ctx context.Context, id model.ProductID) ([]*model.Order, error)
	UpdateOrder(ctx context.Context, id model.OrderID, order *model.Order) error
	DeleteOrder(ctx context.Context, id model.OrderID) error
	GetProduct(ctx context.Context, id model.ProductID) (*model.Product, error)
	PutProduct(ctx context.Context, id model.ProductID, order *model.Product) (model.ProductID, error)
	GetProductsByVendor(ctx context.Context, id model.VendorID) ([]*model.Product, error)
	UpdateProduct(ctx context.Context, id model.ProductID, product *model.Product) error
	DeleteProduct(ctx context.Context, id model.ProductID) error
	GetVendor(ctx context.Context, id model.VendorID) (*model.Vendor, error)
	PutVendor(ctx context.Context, id model.VendorID, vendor *model.Vendor) (model.VendorID, error)
	GetRevenueByVendor(ctx context.Context, id model.VendorID) (decimal.Decimal, error)
	LoadOrders(ctx context.Context) error
}

// Controller defines a dashboard controller
type Controller struct {
	d dashboardRepo
}

// New defines a new dashboard controller
func New(dRepo dashboardRepo) *Controller {
	return &Controller{
		d: dRepo,
	}
}

// GetOrder returns the order for the given id
func (c *Controller) GetOrder(ctx context.Context, id model.OrderID) (*model.Order, error) {
	order, err := c.d.GetOrder(ctx, id)
	if err != nil && errors.Is(err, repo.ErrNotFound) {
		return nil, ErrOrderNotFound
	} else if err != nil {
		return nil, err
	}
	return order, nil
}

// PutOrder adds a new order
func (c *Controller) PutOrder(ctx context.Context, id model.OrderID, order *model.Order) (model.OrderID, error) {
	return c.d.PutOrder(ctx, id, order)
}

// UpdateOrder changes an existing order
func (c *Controller) UpdateOrder(ctx context.Context, id model.OrderID, order *model.Order) error {
	return c.d.UpdateOrder(ctx, id, order)
}

// DeleteOrder removes an order
func (c *Controller) DeleteOrder(ctx context.Context, id model.OrderID) error {
	return c.d.DeleteOrder(ctx, id)
}

// GetOrdersByProduct returns all orders with a specific product for a single vendor
func (c *Controller) GetOrdersByProduct(ctx context.Context, id model.ProductID) ([]*model.Order, error) {
	return c.d.GetOrdersByProduct(ctx, id)
}

// GetProduct returns the product for a given id
func (c *Controller) GetProduct(ctx context.Context, id model.ProductID) (*model.Product, error) {
	product, err := c.d.GetProduct(ctx, id)
	if err != nil && errors.Is(err, repo.ErrNotFound) {
		return nil, ErrProductNotFound
	} else if err != nil {
		return nil, err
	}
	return product, nil
}

// PutProduct adds a new product
func (c *Controller) PutProduct(ctx context.Context, id model.ProductID, product *model.Product) (model.ProductID, error) {
	return c.d.PutProduct(ctx, id, product)
}

// UpdateProduct changes an existing product
func (c *Controller) UpdateProduct(ctx context.Context, id model.ProductID, order *model.Product) error {
	return c.d.UpdateProduct(ctx, id, order)
}

// DeleteProduct removes a product
func (c *Controller) DeleteProduct(ctx context.Context, id model.ProductID) error {
	return c.d.DeleteProduct(ctx, id)
}

// GetProductsByVendor returns all products by vendor
func (c *Controller) GetProductsByVendor(ctx context.Context, id model.VendorID) ([]*model.Product, error) {
	return c.d.GetProductsByVendor(ctx, id)
}

// GetVendor returns the vendor for a given id
func (c *Controller) GetVendor(ctx context.Context, id model.VendorID) (*model.Vendor, error) {
	vendor, err := c.d.GetVendor(ctx, id)
	if err != nil && errors.Is(err, repo.ErrNotFound) {
		return nil, ErrVendorNotFound
	} else if err != nil {
		return nil, err
	}
	return vendor, nil
}

// PutVendor adds a new vendor
func (c *Controller) PutVendor(ctx context.Context, id model.VendorID, vendor *model.Vendor) (model.VendorID, error) {
	return c.d.PutVendor(ctx, id, vendor)
}

// GetRevenueByVendor returns total revenue for a single vendor
func (c *Controller) GetRevenueByVendor(ctx context.Context, id model.VendorID) (decimal.Decimal, error) {
	return c.d.GetRevenueByVendor(ctx, id)
}

// LoadOrders loads a set of random orders
func (c *Controller) LoadOrders(ctx context.Context) error {
	return c.d.LoadOrders(ctx)
}
