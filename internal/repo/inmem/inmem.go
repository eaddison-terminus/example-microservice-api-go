package inmem

import (
	"context"
	"example-microservice-api-go/internal/repo"
	"example-microservice-api-go/pkg/model"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"log"
	"math/rand"
	"sync"
	"time"
)

// DashboardRepo defines an in memory repository
type DashboardRepo struct {
	sync.RWMutex
	oData map[model.OrderID]*model.Order
	pData map[model.ProductID]*model.Product
	vData map[model.VendorID]*model.Vendor
}

var customers = []string{
	"Linus Torvalds",
	"Elon Musk",
	"Alan Turing",
	"Katherine Johnson",
	"Margaret Hamilton",
	"Grace Hopper",
}

var listingType = []string{
	"ami",
	"saas",
}

// NewDashboardRepo creates a new in memory order repository
func NewDashboardRepo() *DashboardRepo {
	return &DashboardRepo{oData: map[model.OrderID]*model.Order{}, pData: map[model.ProductID]*model.Product{}, vData: map[model.VendorID]*model.Vendor{}}
}

// LoadOrders loads a set of 1000 orders
func (d *DashboardRepo) LoadOrders(_ context.Context) error {
	d.Lock()
	defer d.Unlock()
	vid := model.VendorID(uuid.NewString())
	log.Printf("New Vendor for seeding: %v\n", vid)
	newVendor := model.Vendor{
		ID:        vid,
		CreatedAt: time.Now(),
		Name:      "Acme",
	}
	d.vData[vid] = &newVendor
	newListingType := listingType[rand.Intn(1)]
	for i := 1; i < 1001; i++ {
		scale := float32(rand.Int31n(1000))
		newFloat := rand.Float32()
		newPrice := decimal.NewFromFloat32(scale * newFloat)
		pid := model.ProductID(uuid.NewString())
		newProd := model.Product{
			ID:          pid,
			VendorId:    vid,
			CreatedAt:   time.Now(),
			Title:       "Our " + newListingType,
			ListingType: newListingType,
			Price:       newPrice,
		}
		d.pData[pid] = &newProd
		oid := model.OrderID(uuid.NewString())
		cust := customers[rand.Intn(5)]
		newOrder := model.Order{
			ID:        oid,
			ProductId: pid,
			CreatedAt: time.Now(),
			FullName:  cust,
			OrderDate: time.Now(),
		}
		d.oData[oid] = &newOrder
	}
	return nil
}

// GetOrder retrieves an order by id
func (d *DashboardRepo) GetOrder(_ context.Context, id model.OrderID) (*model.Order, error) {
	d.RLock()
	defer d.RUnlock()
	order, ok := d.oData[id]
	if !ok {
		return nil, repo.ErrNotFound
	}

	return order, nil
}

// PutOrder adds an order to the order repo
func (d *DashboardRepo) PutOrder(_ context.Context, id model.OrderID, order *model.Order) (model.OrderID, error) {
	d.Lock()
	defer d.Unlock()
	d.oData[id] = order
	return order.ID, nil
}

// UpdateOrder posts changes to an order
func (d *DashboardRepo) UpdateOrder(_ context.Context, id model.OrderID, order *model.Order) error {
	d.Lock()
	defer d.Unlock()
	d.oData[id] = order
	return nil
}

// DeleteOrder posts changes to an order
func (d *DashboardRepo) DeleteOrder(_ context.Context, id model.OrderID) error {
	d.Lock()
	defer d.Unlock()
	delete(d.oData, id)
	return nil
}

// GetProduct retrieves a product by id
func (d *DashboardRepo) GetProduct(_ context.Context, id model.ProductID) (*model.Product, error) {
	d.RLock()
	defer d.RUnlock()
	product, ok := d.pData[id]
	if !ok {
		return nil, repo.ErrNotFound
	}

	return product, nil
}

// PutProduct adds a product to the product repo
func (d *DashboardRepo) PutProduct(_ context.Context, id model.ProductID, product *model.Product) (model.ProductID, error) {
	d.Lock()
	defer d.Unlock()
	d.pData[id] = product
	return product.ID, nil
}

// UpdateProduct posts changes to an order
func (d *DashboardRepo) UpdateProduct(_ context.Context, id model.ProductID, product *model.Product) error {
	d.Lock()
	defer d.Unlock()
	d.pData[id] = product
	return nil
}

// DeleteProduct posts changes to an order
func (d *DashboardRepo) DeleteProduct(_ context.Context, id model.ProductID) error {
	d.Lock()
	defer d.Unlock()
	delete(d.pData, id)
	return nil
}

// GetVendor retrieves a vendor by id
func (d *DashboardRepo) GetVendor(_ context.Context, id model.VendorID) (*model.Vendor, error) {
	d.RLock()
	defer d.RUnlock()
	vendor, ok := d.vData[id]
	if !ok {
		return nil, repo.ErrNotFound
	}
	return vendor, nil
}

// PutVendor adds a vendor to the vendor repo
func (d *DashboardRepo) PutVendor(_ context.Context, id model.VendorID, vendor *model.Vendor) (model.VendorID, error) {
	d.Lock()
	defer d.Unlock()
	d.vData[id] = vendor
	return vendor.ID, nil
}

// GetRevenueByVendor retrieves total revenue for a vendor
func (d *DashboardRepo) GetRevenueByVendor(_ context.Context, id model.VendorID) (decimal.Decimal, error) {
	d.RLock()
	defer d.RUnlock()
	var rev = decimal.Zero

	for _, order := range d.oData {
		if d.pData[order.ProductId].VendorId == id {
			rev = decimal.Sum(rev, d.pData[order.ProductId].Price).RoundBank(2)
		}
	}
	return rev, nil
}

// GetOrdersByProduct retrieves orders associated to a product
func (d *DashboardRepo) GetOrdersByProduct(_ context.Context, id model.ProductID) ([]*model.Order, error) {
	d.RLock()
	defer d.RUnlock()
	var orders []*model.Order

	for _, or := range d.oData {
		if or.ProductId == id {
			orders = append(orders, or)
		}
	}
	return orders, nil
}

// GetProductsByVendor retrieves products associated to a vendor
func (d *DashboardRepo) GetProductsByVendor(_ context.Context, id model.VendorID) ([]*model.Product, error) {
	d.RLock()
	defer d.RUnlock()
	var products []*model.Product

	for _, pr := range d.pData {
		if pr.VendorId == id {
			products = append(products, pr)
		}
	}
	return products, nil
}
