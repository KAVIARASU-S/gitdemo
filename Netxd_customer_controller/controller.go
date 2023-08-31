package controller

import (
	"context"
	"netxddalmodel"
	"netxddalservice"
)

type CustomerController struct {
	CustomerService *service.CustomerService
}

func (cc *CustomerController) CreateCustomer(ctx context.Context, customer *netxddalmodel.Customer) (*netxddalmodel.CustomerResponse, error) {
	return cc.CustomerService.CreateCustomer(ctx, customer)
}

func (cc *CustomerController) GetCustomerInfo(ctx context.Context, request *netxddalmodel.CustomerRequest) (*netxddalmodel.Customer, error) {
	return cc.CustomerService.GetCustomerInfo(ctx, request)
}
