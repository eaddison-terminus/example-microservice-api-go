// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/controller/dashboard/dashboard.go

// Package mock_dashboard is a generated GoMock package.
package mock_dashboard

import (
	context "context"
	model "example-microservice-api-go/pkg/model"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	decimal "github.com/shopspring/decimal"
)

// MockdashboardRepo is a mock of dashboardRepo interface.
type MockdashboardRepo struct {
	ctrl     *gomock.Controller
	recorder *MockdashboardRepoMockRecorder
}

// MockdashboardRepoMockRecorder is the mock recorder for MockdashboardRepo.
type MockdashboardRepoMockRecorder struct {
	mock *MockdashboardRepo
}

// NewMockdashboardRepo creates a new mock instance.
func NewMockdashboardRepo(ctrl *gomock.Controller) *MockdashboardRepo {
	mock := &MockdashboardRepo{ctrl: ctrl}
	mock.recorder = &MockdashboardRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockdashboardRepo) EXPECT() *MockdashboardRepoMockRecorder {
	return m.recorder
}

// DeleteOrder mocks base method.
func (m *MockdashboardRepo) DeleteOrder(ctx context.Context, id model.OrderID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteOrder", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteOrder indicates an expected call of DeleteOrder.
func (mr *MockdashboardRepoMockRecorder) DeleteOrder(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteOrder", reflect.TypeOf((*MockdashboardRepo)(nil).DeleteOrder), ctx, id)
}

// DeleteProduct mocks base method.
func (m *MockdashboardRepo) DeleteProduct(ctx context.Context, id model.ProductID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteProduct", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteProduct indicates an expected call of DeleteProduct.
func (mr *MockdashboardRepoMockRecorder) DeleteProduct(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteProduct", reflect.TypeOf((*MockdashboardRepo)(nil).DeleteProduct), ctx, id)
}

// GetOrder mocks base method.
func (m *MockdashboardRepo) GetOrder(ctx context.Context, id model.OrderID) (*model.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrder", ctx, id)
	ret0, _ := ret[0].(*model.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrder indicates an expected call of GetOrder.
func (mr *MockdashboardRepoMockRecorder) GetOrder(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrder", reflect.TypeOf((*MockdashboardRepo)(nil).GetOrder), ctx, id)
}

// GetOrdersByProduct mocks base method.
func (m *MockdashboardRepo) GetOrdersByProduct(ctx context.Context, id model.ProductID) ([]*model.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrdersByProduct", ctx, id)
	ret0, _ := ret[0].([]*model.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrdersByProduct indicates an expected call of GetOrdersByProduct.
func (mr *MockdashboardRepoMockRecorder) GetOrdersByProduct(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrdersByProduct", reflect.TypeOf((*MockdashboardRepo)(nil).GetOrdersByProduct), ctx, id)
}

// GetProduct mocks base method.
func (m *MockdashboardRepo) GetProduct(ctx context.Context, id model.ProductID) (*model.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProduct", ctx, id)
	ret0, _ := ret[0].(*model.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProduct indicates an expected call of GetProduct.
func (mr *MockdashboardRepoMockRecorder) GetProduct(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProduct", reflect.TypeOf((*MockdashboardRepo)(nil).GetProduct), ctx, id)
}

// GetProductsByVendor mocks base method.
func (m *MockdashboardRepo) GetProductsByVendor(ctx context.Context, id model.VendorID) ([]*model.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProductsByVendor", ctx, id)
	ret0, _ := ret[0].([]*model.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProductsByVendor indicates an expected call of GetProductsByVendor.
func (mr *MockdashboardRepoMockRecorder) GetProductsByVendor(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProductsByVendor", reflect.TypeOf((*MockdashboardRepo)(nil).GetProductsByVendor), ctx, id)
}

// GetRevenueByVendor mocks base method.
func (m *MockdashboardRepo) GetRevenueByVendor(ctx context.Context, id model.VendorID) (decimal.Decimal, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRevenueByVendor", ctx, id)
	ret0, _ := ret[0].(decimal.Decimal)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRevenueByVendor indicates an expected call of GetRevenueByVendor.
func (mr *MockdashboardRepoMockRecorder) GetRevenueByVendor(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRevenueByVendor", reflect.TypeOf((*MockdashboardRepo)(nil).GetRevenueByVendor), ctx, id)
}

// GetVendor mocks base method.
func (m *MockdashboardRepo) GetVendor(ctx context.Context, id model.VendorID) (*model.Vendor, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetVendor", ctx, id)
	ret0, _ := ret[0].(*model.Vendor)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetVendor indicates an expected call of GetVendor.
func (mr *MockdashboardRepoMockRecorder) GetVendor(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVendor", reflect.TypeOf((*MockdashboardRepo)(nil).GetVendor), ctx, id)
}

// LoadOrders mocks base method.
func (m *MockdashboardRepo) LoadOrders(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoadOrders", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// LoadOrders indicates an expected call of LoadOrders.
func (mr *MockdashboardRepoMockRecorder) LoadOrders(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoadOrders", reflect.TypeOf((*MockdashboardRepo)(nil).LoadOrders), ctx)
}

// PutOrder mocks base method.
func (m *MockdashboardRepo) PutOrder(ctx context.Context, id model.OrderID, order *model.Order) (model.OrderID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PutOrder", ctx, id, order)
	ret0, _ := ret[0].(model.OrderID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PutOrder indicates an expected call of PutOrder.
func (mr *MockdashboardRepoMockRecorder) PutOrder(ctx, id, order interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutOrder", reflect.TypeOf((*MockdashboardRepo)(nil).PutOrder), ctx, id, order)
}

// PutProduct mocks base method.
func (m *MockdashboardRepo) PutProduct(ctx context.Context, id model.ProductID, order *model.Product) (model.ProductID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PutProduct", ctx, id, order)
	ret0, _ := ret[0].(model.ProductID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PutProduct indicates an expected call of PutProduct.
func (mr *MockdashboardRepoMockRecorder) PutProduct(ctx, id, order interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutProduct", reflect.TypeOf((*MockdashboardRepo)(nil).PutProduct), ctx, id, order)
}

// PutVendor mocks base method.
func (m *MockdashboardRepo) PutVendor(ctx context.Context, id model.VendorID, vendor *model.Vendor) (model.VendorID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PutVendor", ctx, id, vendor)
	ret0, _ := ret[0].(model.VendorID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PutVendor indicates an expected call of PutVendor.
func (mr *MockdashboardRepoMockRecorder) PutVendor(ctx, id, vendor interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutVendor", reflect.TypeOf((*MockdashboardRepo)(nil).PutVendor), ctx, id, vendor)
}

// UpdateOrder mocks base method.
func (m *MockdashboardRepo) UpdateOrder(ctx context.Context, id model.OrderID, order *model.Order) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateOrder", ctx, id, order)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateOrder indicates an expected call of UpdateOrder.
func (mr *MockdashboardRepoMockRecorder) UpdateOrder(ctx, id, order interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateOrder", reflect.TypeOf((*MockdashboardRepo)(nil).UpdateOrder), ctx, id, order)
}

// UpdateProduct mocks base method.
func (m *MockdashboardRepo) UpdateProduct(ctx context.Context, id model.ProductID, product *model.Product) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateProduct", ctx, id, product)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateProduct indicates an expected call of UpdateProduct.
func (mr *MockdashboardRepoMockRecorder) UpdateProduct(ctx, id, product interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateProduct", reflect.TypeOf((*MockdashboardRepo)(nil).UpdateProduct), ctx, id, product)
}