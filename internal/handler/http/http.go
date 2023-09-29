package http

import (
	"encoding/json"
	"errors"
	"example-microservice-api-go/internal/controller/dashboard"
	"example-microservice-api-go/pkg/model"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"log"
	"net/http"
	"strings"
	"time"
)

const shortForm = "2006-Jan-02"

// OrderHandler defines a order service controller
type OrderHandler struct {
	ctrl *dashboard.Controller
}

// ProductHandler defines a product service controller
type ProductHandler struct {
	ctrl *dashboard.Controller
}

// VendorHandler defines a vendor service controller
type VendorHandler struct {
	ctrl *dashboard.Controller
}

// RevenueHandler defines a revenue service controller
type RevenueHandler struct {
	ctrl *dashboard.Controller
}

// NewOrderHandler creates a new order service HTTP handler
func NewOrderHandler(ctrl *dashboard.Controller) *OrderHandler {
	return &OrderHandler{ctrl: ctrl}
}

// NewProductHandler creates a new product service HTTP handler
func NewProductHandler(ctrl *dashboard.Controller) *ProductHandler {
	return &ProductHandler{ctrl: ctrl}
}

// NewVendorHandler creates a new vendor service HTTP handler
func NewVendorHandler(ctrl *dashboard.Controller) *VendorHandler {
	return &VendorHandler{ctrl: ctrl}
}

// NewRevenueHandler creates a new revenue service HTTP handler
func NewRevenueHandler(ctrl *dashboard.Controller) *RevenueHandler {
	return &RevenueHandler{ctrl: ctrl}
}

// Handle handles PUT and GET requests for orders
func (oh *OrderHandler) Handle(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		recordID := model.OrderID(req.FormValue("id"))
		productID := model.ProductID(req.FormValue("productID"))
		if recordID == "" && productID == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if recordID != "" {
			val, err := oh.ctrl.GetOrder(req.Context(), recordID)
			if err != nil && errors.Is(err, dashboard.ErrOrderNotFound) {
				w.WriteHeader(http.StatusNotFound)
				return
			}
			if err := json.NewEncoder(w).Encode(val); err != nil {
				log.Printf("Response encode error: %v\n", err)
				return
			}
		} else {
			val, err := oh.ctrl.GetOrdersByProduct(req.Context(), productID)
			if err != nil && errors.Is(err, dashboard.ErrOrderNotFound) {
				w.WriteHeader(http.StatusNotFound)
				return
			}
			if err := json.NewEncoder(w).Encode(val); err != nil {
				log.Printf("Response encode error: %v\n", err)
				return
			}
		}
		return
	case http.MethodPut:
		to, err := time.Parse(shortForm, req.FormValue("orderDate"))
		if err != nil {
			log.Printf("Order date is invalid: %v\n", to)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		oid := model.OrderID(uuid.NewString())
		val, err := oh.ctrl.PutOrder(req.Context(), oid, &model.Order{
			ID:        oid,
			ProductId: model.ProductID(req.FormValue("productID")),
			CreatedAt: time.Now(),
			FullName:  req.FormValue("fullName"),
			OrderDate: to,
		})
		if err != nil {
			log.Printf("Repository put error: %v\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		if err := json.NewEncoder(w).Encode(val); err != nil {
			log.Printf("Response encode error: %v\n", err)
			return
		}
	case http.MethodPost:
		recordID := model.OrderID(req.FormValue("id"))
		to, err := time.Parse(shortForm, req.FormValue("orderDate"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if recordID == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if err := oh.ctrl.UpdateOrder(req.Context(), recordID, &model.Order{
			ID:        recordID,
			ProductId: model.ProductID(req.FormValue("productID")),
			FullName:  req.FormValue("fullName"),
			OrderDate: to,
		}); err != nil {
			log.Printf("Update failed: %v\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	case http.MethodDelete:
		uriSlice := strings.Split(req.RequestURI, "/")
		lastElement := len(uriSlice) - 1
		recordID := model.OrderID(uriSlice[lastElement])
		if recordID == "" {
			log.Printf("Order not found: %v\n", recordID)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if err := oh.ctrl.DeleteOrder(req.Context(), recordID); err != nil {
			log.Printf("Error deleteing order: %v\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	default:
		w.WriteHeader(http.StatusBadRequest)
	}
}

// Handle handles PUT and GET requests for products
func (ph *ProductHandler) Handle(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		recordID := model.ProductID(req.FormValue("id"))
		vendorID := model.VendorID(req.FormValue("vendorID"))
		if recordID == "" && vendorID == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		} else {
			if recordID != "" {
				val, err := ph.ctrl.GetProduct(req.Context(), recordID)
				if err != nil && errors.Is(err, dashboard.ErrProductNotFound) {
					w.WriteHeader(http.StatusNotFound)
					return
				}
				if err := json.NewEncoder(w).Encode(val); err != nil {
					log.Printf("Response encode error: %v\n", err)
					return
				}
			} else {
				{
					val, err := ph.ctrl.GetProductsByVendor(req.Context(), vendorID)
					if err != nil && errors.Is(err, dashboard.ErrProductNotFound) {
						w.WriteHeader(http.StatusNotFound)
						return
					}
					if err := json.NewEncoder(w).Encode(val); err != nil {
						log.Printf("Response encode error: %v\n", err)
						return
					}
				}
			}
		}
		return
	case http.MethodPut:
		pid := model.ProductID(uuid.NewString())
		priceCheck := req.FormValue("price")
		if priceCheck == "" {
			log.Printf("invalid price: %v\n", priceCheck)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		price := decimal.RequireFromString(priceCheck)
		val, err := ph.ctrl.PutProduct(req.Context(), pid, &model.Product{
			ID:          pid,
			VendorId:    model.VendorID(req.FormValue("vendorID")),
			CreatedAt:   time.Now(),
			Title:       req.FormValue("title"),
			ListingType: req.FormValue("listingType"),
			Price:       price,
		})
		if err != nil {
			log.Printf("Repository put error: %v\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		if err := json.NewEncoder(w).Encode(val); err != nil {
			log.Printf("Response encode error: %v\n", err)
			return
		}
	case http.MethodPost:
		recordID := model.ProductID(req.FormValue("id"))
		if recordID == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		price, err := decimal.NewFromString(req.FormValue("price"))
		if err != nil {
			log.Printf("Invalid price: %v\n", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if err := ph.ctrl.UpdateProduct(req.Context(), recordID, &model.Product{
			ID:          recordID,
			VendorId:    model.VendorID(req.FormValue("vendorID")),
			Title:       req.FormValue("title"),
			ListingType: req.FormValue("listingType"),
			Price:       price,
		}); err != nil {
			log.Printf("Update failed: %v\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	case http.MethodDelete:
		uriSlice := strings.Split(req.RequestURI, "/")
		lastElement := len(uriSlice) - 1
		recordID := model.ProductID(uriSlice[lastElement])
		if recordID == "" {
			log.Printf("Product not found")
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if err := ph.ctrl.DeleteProduct(req.Context(), recordID); err != nil {
			log.Printf("Error deleteing product: %v\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	default:
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}

// Handle handles PUT and GET requests for vendors
func (vh *VendorHandler) Handle(w http.ResponseWriter, req *http.Request) {

	switch req.Method {
	case http.MethodGet:
		recordID := model.VendorID(req.FormValue("id"))
		if recordID == "" {
			log.Printf("Bad request")
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		val, err := vh.ctrl.GetVendor(req.Context(), recordID)
		if err != nil && errors.Is(err, dashboard.ErrVendorNotFound) {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		if err := json.NewEncoder(w).Encode(val); err != nil {
			log.Printf("Response encode error: %v\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	case http.MethodPut:
		vid := model.VendorID(uuid.NewString())
		val, err := vh.ctrl.PutVendor(req.Context(), vid, &model.Vendor{
			ID:        vid,
			CreatedAt: time.Now(),
			Name:      req.FormValue("name"),
		})
		if err != nil {
			log.Printf("Repository put error: %v\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		if err := json.NewEncoder(w).Encode(val); err != nil {
			log.Printf("Response encode error: %v\n", err)
		}
	default:
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}

// Handle handles PUT and GET requests for revenue
func (rh *RevenueHandler) Handle(w http.ResponseWriter, req *http.Request) {

	switch req.Method {
	case http.MethodGet:
		recordID := model.VendorID(req.FormValue("id"))
		if recordID == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		val, err := rh.ctrl.GetRevenueByVendor(req.Context(), recordID)
		if err != nil && errors.Is(err, dashboard.ErrVendorNotFound) {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		if err := json.NewEncoder(w).Encode(val); err != nil {
			log.Printf("Response encode error: %v\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	default:
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}
