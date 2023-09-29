package main

import (
	"context"
	"example-microservice-api-go/internal/controller/dashboard"
	httphandler "example-microservice-api-go/internal/handler/http"
	"example-microservice-api-go/internal/repo/inmem"
	"log"
	"net/http"
)

func main() {
	log.Printf("Starting the dashboard service")
	dashboardRepo := inmem.NewDashboardRepo()
	ctrl := dashboard.New(dashboardRepo)
	if err := ctrl.LoadOrders(context.Background()); err != nil {
		log.Printf("error loading repo: %n\n", err)
	}
	oh := httphandler.NewOrderHandler(ctrl)
	ph := httphandler.NewProductHandler(ctrl)
	vh := httphandler.NewVendorHandler(ctrl)
	rv := httphandler.NewRevenueHandler(ctrl)
	http.Handle("/order", http.HandlerFunc(oh.Handle))
	http.Handle("/order/", http.HandlerFunc(oh.Handle))
	http.Handle("/product", http.HandlerFunc(ph.Handle))
	http.Handle("/product/", http.HandlerFunc(ph.Handle))
	http.Handle("/vendor", http.HandlerFunc(vh.Handle))
	http.Handle("/revenue", http.HandlerFunc(rv.Handle))
	if err := http.ListenAndServe(":8088", nil); err != nil {
		panic(err)
	}
}
