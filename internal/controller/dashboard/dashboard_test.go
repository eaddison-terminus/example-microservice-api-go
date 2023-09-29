package dashboard_test

import (
	"context"
	"errors"
	"example-microservice-api-go/internal/controller/dashboard"
	"example-microservice-api-go/internal/repo"
	gen "example-microservice-api-go/mock"
	"example-microservice-api-go/pkg/model"
	"github.com/golang/mock/gomock"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestController_GetOrder(t *testing.T) {
	testSet := []struct {
		name       string
		expRepoRes *model.Order
		expRepoErr error
		wantRes    *model.Order
		wantErr    error
	}{
		{
			name:       "not found",
			expRepoErr: repo.ErrNotFound,
			wantErr:    dashboard.ErrOrderNotFound,
		},
		{
			name:       "unexpected error",
			expRepoErr: errors.New("unexpected error"),
			wantErr:    errors.New("unexpected error"),
		},
		{
			name:       "success",
			expRepoRes: &model.Order{},
			wantRes:    &model.Order{},
		},
	}
	for _, tt := range testSet {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			repoMock := gen.NewMockdashboardRepo(ctrl)
			c := dashboard.New(repoMock)
			ctx := context.Background()
			id := model.OrderID("1")
			repoMock.EXPECT().GetOrder(ctx, id).Return(tt.expRepoRes, tt.expRepoErr)
			res, err := c.GetOrder(ctx, id)
			assert.Equal(t, tt.wantRes, res, tt.name)
			assert.Equal(t, tt.wantErr, err, tt.name)
		})
	}
}

func TestController_PutOrder(t *testing.T) {
	testSet := []struct {
		name       string
		expRepoRes model.OrderID
		expRepoErr error
		wantRes    model.OrderID
		wantErr    error
	}{
		{
			name:       "unexpected error",
			expRepoErr: errors.New("unexpected error"),
			wantErr:    errors.New("unexpected error"),
		},
		{
			name:       "success",
			expRepoRes: model.OrderID("1"),
			wantRes:    model.OrderID("1"),
		},
	}
	for _, tt := range testSet {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			repoMock := gen.NewMockdashboardRepo(ctrl)
			c := dashboard.New(repoMock)
			ctx := context.Background()
			id := model.OrderID("1")
			order := &model.Order{
				ID:        "1",
				ProductId: "1",
				CreatedAt: time.Now(),
				FullName:  "",
				OrderDate: time.Now(),
			}
			repoMock.EXPECT().PutOrder(ctx, id, order).Return(tt.expRepoRes, tt.expRepoErr)
			res, err := c.PutOrder(ctx, id, order)
			assert.Equal(t, tt.wantRes, res, tt.name)
			assert.Equal(t, tt.wantErr, err, tt.name)
		})
	}
}

func TestController_UpdateOrder(t *testing.T) {
	testSet := []struct {
		name       string
		expRepoRes *model.Order
		expRepoErr error
		wantRes    *model.Order
		wantErr    error
	}{
		{
			name:       "not found",
			expRepoErr: repo.ErrNotFound,
			wantErr:    repo.ErrNotFound,
		},
		{
			name:       "unexpected error",
			expRepoErr: errors.New("unexpected error"),
			wantErr:    errors.New("unexpected error"),
		},
		{
			name:       "success",
			expRepoRes: &model.Order{},
			wantRes:    &model.Order{},
		},
	}
	for _, tt := range testSet {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			repoMock := gen.NewMockdashboardRepo(ctrl)
			c := dashboard.New(repoMock)
			ctx := context.Background()
			id := model.OrderID("1")
			order := &model.Order{
				ID:        "1",
				ProductId: "1",
				CreatedAt: time.Now(),
				FullName:  "",
				OrderDate: time.Now(),
			}
			repoMock.EXPECT().UpdateOrder(ctx, id, order).Return(tt.expRepoErr)
			err := c.UpdateOrder(ctx, id, order)
			assert.Equal(t, tt.wantErr, err, tt.name)
		})
	}
}

func TestController_DeleteOrder(t *testing.T) {
	testSet := []struct {
		name       string
		expRepoRes *model.Order
		expRepoErr error
		wantRes    *model.Order
		wantErr    error
	}{
		{
			name:       "not found",
			expRepoErr: repo.ErrNotFound,
			wantErr:    repo.ErrNotFound,
		},
		{
			name:       "unexpected error",
			expRepoErr: errors.New("unexpected error"),
			wantErr:    errors.New("unexpected error"),
		},
		{
			name:       "success",
			expRepoRes: &model.Order{},
			wantRes:    &model.Order{},
		},
	}
	for _, tt := range testSet {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			repoMock := gen.NewMockdashboardRepo(ctrl)
			c := dashboard.New(repoMock)
			ctx := context.Background()
			id := model.OrderID("1")
			repoMock.EXPECT().DeleteOrder(ctx, id).Return(tt.expRepoErr)
			err := c.DeleteOrder(ctx, id)
			assert.Equal(t, tt.wantErr, err, tt.name)
		})
	}
}

func TestController_GetProduct(t *testing.T) {
	testSet := []struct {
		name       string
		expRepoRes *model.Product
		expRepoErr error
		wantRes    *model.Product
		wantErr    error
	}{
		{
			name:       "not found",
			expRepoErr: repo.ErrNotFound,
			wantErr:    dashboard.ErrProductNotFound,
		},
		{
			name:       "unexpected error",
			expRepoErr: errors.New("unexpected error"),
			wantErr:    errors.New("unexpected error"),
		},
		{
			name:       "success",
			expRepoRes: &model.Product{},
			wantRes:    &model.Product{},
		},
	}
	for _, tt := range testSet {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			repoMock := gen.NewMockdashboardRepo(ctrl)
			c := dashboard.New(repoMock)
			ctx := context.Background()
			id := model.ProductID("1")
			repoMock.EXPECT().GetProduct(ctx, id).Return(tt.expRepoRes, tt.expRepoErr)
			res, err := c.GetProduct(ctx, id)
			assert.Equal(t, tt.wantRes, res, tt.name)
			assert.Equal(t, tt.wantErr, err, tt.name)
		})
	}
}

func TestController_PutProduct(t *testing.T) {
	testSet := []struct {
		name       string
		expRepoRes model.ProductID
		expRepoErr error
		wantRes    model.ProductID
		wantErr    error
	}{
		{
			name:       "unexpected error",
			expRepoErr: errors.New("unexpected error"),
			wantErr:    errors.New("unexpected error"),
		},
		{
			name:       "success",
			expRepoRes: model.ProductID("1"),
			wantRes:    model.ProductID("1"),
		},
	}
	for _, tt := range testSet {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			repoMock := gen.NewMockdashboardRepo(ctrl)
			c := dashboard.New(repoMock)
			ctx := context.Background()
			id := model.ProductID("1")
			product := &model.Product{
				ID:          model.ProductID("1"),
				VendorId:    model.VendorID("1"),
				CreatedAt:   time.Now(),
				Title:       "Test",
				ListingType: "saas",
				Price:       decimal.RequireFromString("1.00"),
			}
			repoMock.EXPECT().PutProduct(ctx, id, product).Return(tt.expRepoRes, tt.expRepoErr)
			res, err := c.PutProduct(ctx, id, product)
			assert.Equal(t, tt.wantRes, res, tt.name)
			assert.Equal(t, tt.wantErr, err, tt.name)
		})
	}
}

func TestController_DeleteProduct(t *testing.T) {
	testSet := []struct {
		name       string
		expRepoRes *model.Product
		expRepoErr error
		wantRes    *model.Product
		wantErr    error
	}{
		{
			name:       "not found",
			expRepoErr: repo.ErrNotFound,
			wantErr:    repo.ErrNotFound,
		},
		{
			name:       "unexpected error",
			expRepoErr: errors.New("unexpected error"),
			wantErr:    errors.New("unexpected error"),
		},
		{
			name:       "success",
			expRepoRes: &model.Product{},
			wantRes:    &model.Product{},
		},
	}
	for _, tt := range testSet {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			repoMock := gen.NewMockdashboardRepo(ctrl)
			c := dashboard.New(repoMock)
			ctx := context.Background()
			id := model.ProductID("1")
			repoMock.EXPECT().DeleteProduct(ctx, id).Return(tt.expRepoErr)
			err := c.DeleteProduct(ctx, id)
			assert.Equal(t, tt.wantErr, err, tt.name)
		})
	}
}

func TestController_UpdateProduct(t *testing.T) {
	testSet := []struct {
		name       string
		expRepoRes *model.Product
		expRepoErr error
		wantRes    *model.Product
		wantErr    error
	}{
		{
			name:       "not found",
			expRepoErr: repo.ErrNotFound,
			wantErr:    repo.ErrNotFound,
		},
		{
			name:       "unexpected error",
			expRepoErr: errors.New("unexpected error"),
			wantErr:    errors.New("unexpected error"),
		},
		{
			name:       "success",
			expRepoRes: &model.Product{},
			wantRes:    &model.Product{},
		},
	}
	for _, tt := range testSet {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			repoMock := gen.NewMockdashboardRepo(ctrl)
			c := dashboard.New(repoMock)
			ctx := context.Background()
			id := model.ProductID("1")
			product := &model.Product{
				ID:          "1",
				VendorId:    "1",
				CreatedAt:   time.Now(),
				Title:       "Test",
				ListingType: "saas",
				Price:       decimal.RequireFromString("1.00"),
			}
			repoMock.EXPECT().UpdateProduct(ctx, id, product).Return(tt.expRepoErr)
			err := c.UpdateProduct(ctx, id, product)
			assert.Equal(t, tt.wantErr, err, tt.name)
		})
	}
}

func TestController_GetProductsByVendor(t *testing.T) {
	testSet := []struct {
		name       string
		expRepoRes []*model.Product
		expRepoErr error
		wantRes    []*model.Product
		wantErr    error
	}{
		{
			name:       "not found",
			expRepoErr: repo.ErrNotFound,
			wantErr:    repo.ErrNotFound,
		},
		{
			name:       "unexpected error",
			expRepoErr: errors.New("unexpected error"),
			wantErr:    errors.New("unexpected error"),
		},
		{
			name:       "success",
			expRepoRes: []*model.Product{},
			wantRes:    []*model.Product{},
		},
	}
	for _, tt := range testSet {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			repoMock := gen.NewMockdashboardRepo(ctrl)
			c := dashboard.New(repoMock)
			ctx := context.Background()
			id := model.VendorID("1")
			repoMock.EXPECT().GetProductsByVendor(ctx, id).Return(tt.expRepoRes, tt.expRepoErr)
			res, err := c.GetProductsByVendor(ctx, id)
			assert.Equal(t, tt.wantRes, res, tt.name)
			assert.Equal(t, tt.wantErr, err, tt.name)
		})
	}
}

func TestController_GetOrdersByProduct(t *testing.T) {
	testSet := []struct {
		name       string
		expRepoRes []*model.Order
		expRepoErr error
		wantRes    []*model.Order
		wantErr    error
	}{
		{
			name:       "not found",
			expRepoErr: repo.ErrNotFound,
			wantErr:    repo.ErrNotFound,
		},
		{
			name:       "unexpected error",
			expRepoErr: errors.New("unexpected error"),
			wantErr:    errors.New("unexpected error"),
		},
		{
			name:       "success",
			expRepoRes: []*model.Order{},
			wantRes:    []*model.Order{},
		},
	}
	for _, tt := range testSet {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			repoMock := gen.NewMockdashboardRepo(ctrl)
			c := dashboard.New(repoMock)
			ctx := context.Background()
			id := model.ProductID("1")
			repoMock.EXPECT().GetOrdersByProduct(ctx, id).Return(tt.expRepoRes, tt.expRepoErr)
			res, err := c.GetOrdersByProduct(ctx, id)
			assert.Equal(t, tt.wantRes, res, tt.name)
			assert.Equal(t, tt.wantErr, err, tt.name)
		})
	}
}

func TestController_GetRevenueByVendor(t *testing.T) {
	testSet := []struct {
		name       string
		expRepoRes decimal.Decimal
		expRepoErr error
		wantRes    decimal.Decimal
		wantErr    error
	}{
		{
			name:       "not found",
			expRepoErr: repo.ErrNotFound,
			wantErr:    repo.ErrNotFound,
		},
		{
			name:       "unexpected error",
			expRepoErr: errors.New("unexpected error"),
			wantErr:    errors.New("unexpected error"),
		},
		{
			name:       "success",
			expRepoRes: decimal.RequireFromString("1.99"),
			wantRes:    decimal.RequireFromString("1.99"),
		},
	}
	for _, tt := range testSet {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			repoMock := gen.NewMockdashboardRepo(ctrl)
			c := dashboard.New(repoMock)
			ctx := context.Background()
			id := model.VendorID("1")
			repoMock.EXPECT().GetRevenueByVendor(ctx, id).Return(tt.expRepoRes, tt.expRepoErr)
			res, err := c.GetRevenueByVendor(ctx, id)
			assert.Equal(t, tt.wantRes, res, tt.name)
			assert.Equal(t, tt.wantErr, err, tt.name)
		})
	}
}

func TestController_GetVendor(t *testing.T) {
	testSet := []struct {
		name       string
		expRepoRes *model.Vendor
		expRepoErr error
		wantRes    *model.Vendor
		wantErr    error
	}{
		{
			name:       "not found",
			expRepoErr: repo.ErrNotFound,
			wantErr:    dashboard.ErrVendorNotFound,
		},
		{
			name:       "unexpected error",
			expRepoErr: errors.New("unexpected error"),
			wantErr:    errors.New("unexpected error"),
		},
		{
			name:       "success",
			expRepoRes: &model.Vendor{},
			wantRes:    &model.Vendor{},
		},
	}
	for _, tt := range testSet {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			repoMock := gen.NewMockdashboardRepo(ctrl)
			c := dashboard.New(repoMock)
			ctx := context.Background()
			id := model.VendorID("1")
			repoMock.EXPECT().GetVendor(ctx, id).Return(tt.expRepoRes, tt.expRepoErr)
			res, err := c.GetVendor(ctx, id)
			assert.Equal(t, tt.wantRes, res, tt.name)
			assert.Equal(t, tt.wantErr, err, tt.name)
		})
	}
}

func TestController_PutVendor(t *testing.T) {
	testSet := []struct {
		name       string
		expRepoRes model.VendorID
		expRepoErr error
		wantRes    model.VendorID
		wantErr    error
	}{
		{
			name:       "not found",
			expRepoErr: repo.ErrNotFound,
			wantErr:    repo.ErrNotFound,
		},
		{
			name:       "unexpected error",
			expRepoErr: errors.New("unexpected error"),
			wantErr:    errors.New("unexpected error"),
		},
		{
			name:       "success",
			expRepoRes: model.VendorID("1"),
			wantRes:    model.VendorID("1"),
		},
	}
	for _, tt := range testSet {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			repoMock := gen.NewMockdashboardRepo(ctrl)
			c := dashboard.New(repoMock)
			ctx := context.Background()
			id := model.VendorID("1")
			vendor := &model.Vendor{
				ID:        "1",
				CreatedAt: time.Now(),
				Name:      "Test Vendor",
			}
			repoMock.EXPECT().PutVendor(ctx, id, vendor).Return(tt.expRepoRes, tt.expRepoErr)
			res, err := c.PutVendor(ctx, id, vendor)
			assert.Equal(t, tt.wantRes, res, tt.name)
			assert.Equal(t, tt.wantErr, err, tt.name)
		})
	}
}

func TestController_LoadOrders(t *testing.T) {
	testSet := []struct {
		name       string
		expRepoErr error
		wantErr    error
	}{
		{
			name:       "unexpected error",
			expRepoErr: errors.New("unexpected error"),
			wantErr:    errors.New("unexpected error"),
		},
		{
			name:    "success",
			wantErr: nil,
		},
	}
	for _, tt := range testSet {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			repoMock := gen.NewMockdashboardRepo(ctrl)
			c := dashboard.New(repoMock)
			ctx := context.Background()
			repoMock.EXPECT().LoadOrders(ctx).Return(tt.expRepoErr)
			err := c.LoadOrders(ctx)
			assert.Equal(t, tt.wantErr, err, tt.name)
		})
	}
}
