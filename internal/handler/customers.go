package handler

import (
	"net/http"
	"strconv"

	"car-rental/internal/models"
	"car-rental/internal/service"
	"car-rental/pkg"

	"github.com/gin-gonic/gin"
)

type CustomerHandler interface {
	GetCustomers(ctx *gin.Context)
	GetCustomerByID(ctx *gin.Context)
	DeleteCustomerByID(ctx *gin.Context)
	CreateCustomer(ctx *gin.Context)
	EditCustomer(ctx *gin.Context)
	AssignMembership(ctx *gin.Context)
	DeleteMembershipByCustomer(ctx *gin.Context)
}

type customerHandlerImpl struct {
	customerservice service.CustomerService
}

func NewCustomerHandler(customerservice service.CustomerService) CustomerHandler {
	return &customerHandlerImpl{customerservice: customerservice}
}
// GetCustomers godoc
// @Summary Retrieve list of customers
// @Description Retrieve a list of all customers.
// @Tags customers
// @Accept json
// @Produce json
// @Success	200	{object} models.Customer
// @Success 200 {object} pkg.ErrorResponse "No customer found"
// @Failure 500 {object} pkg.ErrorResponse "Internal server error"
// @Router /customers [get]
func (p *customerHandlerImpl) GetCustomers(ctx *gin.Context) {
	customers, err := p.customerservice.GetCustomers(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, pkg.ErrorResponse{Message: err.Error()})
		return
	}
	if len(customers) == 0 {
        ctx.JSON(http.StatusOK, gin.H{"message": "No customer found"})
        return
    }
	ctx.JSON(http.StatusOK, customers)
}
// GetCustomerByID godoc
// @Summary Retrieve customer by ID
// @Description Retrieve a customer by its ID
// @Tags customers
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path int true "Customer ID"
// @Success 200 {object} models.UpdateCustomer "customer"
// @Failure 400 {object} pkg.ErrorResponse "Bad request"
// @Failure 404 {object} pkg.ErrorResponse "Customer not found"
// @Failure 500 {object} pkg.ErrorResponse "Internal server error"
// @Router /customers/{id} [get]
func (p *customerHandlerImpl) GetCustomerByID(ctx *gin.Context) {
	// get customer ID from path parameter
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, pkg.ErrorResponse{Message: "invalid customer ID"})
		return
	}

	customer, err := p.customerservice.GetCustomersByID(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, pkg.ErrorResponse{Message: err.Error()})
		return
	}
	if customer.ID == 0 {
		ctx.JSON(http.StatusNotFound, pkg.ErrorResponse{Message: "customer not found"})
		return
	}

	ctx.JSON(http.StatusOK, customer)
}
// DeleteCustomerByID godoc
// @Summary Delete customer by ID
// @Description Delete a customer by its ID.
// @Tags customers
// @Accept json
// @Produce json
// @Param id path int true "Customer ID"
// @Security Bearer
// @Success	200	{object} models.UpdateCustomer
// @Failure 404 {object} pkg.ErrorResponse "Customer not found"
// @Failure 500 {object} pkg.ErrorResponse "Internal server error"
// @Router /customers/{id} [delete]
func (p *customerHandlerImpl) DeleteCustomerByID(ctx *gin.Context) {
	// Get customer ID from path parameter
	id, err := strconv.Atoi(ctx.Param("id"))
	if id == 0 || err != nil {
		ctx.JSON(http.StatusBadRequest, pkg.ErrorResponse{Message: "invalid required param"})
		return
	}

	// Delete customer by ID
	customer, err := p.customerservice.DeleteCustomer(ctx, uint64(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, pkg.ErrorResponse{Message: err.Error()})
		return
	}

	// Check if the customer exists
	if customer.ID == 0 {
		ctx.JSON(http.StatusNotFound, pkg.ErrorResponse{Message: "Customer not found"})
		return
	}


	ctx.JSON(http.StatusOK, map[string]any{
		"customer":    customer,
		"message": "Your customer has been successfully deleted",
	})
}
// CreateCustomer godoc
// @Summary Create a new customer
// @Description Create a new customer.
// @Tags customers
// @Accept json
// @Produce json
// @Param customer body models.InputCustomer true "Customer data"
// @Security Bearer
// @Success	200	{object} models.CreateCustomer
// @Failure 400 {object} pkg.ErrorResponse "Bad request"
// @Failure 500 {object} pkg.ErrorResponse "Internal server error"
// @Router /customers [post]
func (p *customerHandlerImpl) CreateCustomer(ctx *gin.Context) {
	customer := models.InputCustomer{}
	if err := ctx.BindJSON(&customer); err != nil {
		ctx.JSON(http.StatusBadRequest, pkg.ErrorResponse{Message: err.Error()})
		return
	}

	createdCustomer, err := p.customerservice.CreateCustomer(ctx, customer)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, pkg.ErrorResponse{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, createdCustomer)
}
// EditCustomer godoc
// @Summary Update customer information
// @Description Update information of a customer.
// @Tags customers
// @Accept json
// @Produce json
// @Param id path int true "Customer ID"
// @Param customer body models.InputCustomer true "Customer data"
// @Security Bearer
// @Success	200	{object} models.UpdateCustomer
// @Failure 400 {object} pkg.ErrorResponse "Bad request"
// @Failure 401 {object} pkg.ErrorResponse "Unauthorized"
// @Failure 500 {object} pkg.ErrorResponse "Internal server error"
// @Router /customers/{id} [put]
func (p *customerHandlerImpl) EditCustomer(ctx *gin.Context) {
	
    id, err := strconv.Atoi(ctx.Param("id"))
	if id == 0 || err != nil {
		ctx.JSON(http.StatusBadRequest, pkg.ErrorResponse{Message: "invalid required param"})
		return
	}
	customer, err := p.customerservice.GetCustomersByID(ctx, uint64(id))
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, pkg.ErrorResponse{Message: err.Error()})
        return
    }

    // Check if the customer exists
    if customer.ID == 0 {
        ctx.JSON(http.StatusNotFound, pkg.ErrorResponse{Message: "Customer not found"})
        return
    }

    
    if err := ctx.ShouldBindJSON(&customer); err != nil {
        ctx.JSON(http.StatusBadRequest, pkg.ErrorResponse{Message: "Invalid request body"})
        return
    }
	inputCustomer := models.InputCustomer{}
	inputCustomer.Name = customer.Name
	inputCustomer.NIK = customer.NIK
	inputCustomer.Phone = customer.Phone
    // Call service to edit customer data
    updatedCustomer, err := p.customerservice.EditCustomer(ctx, uint64(id), inputCustomer)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, pkg.ErrorResponse{Message: err.Error()})
        return
    }

    // Return updated customer data
    ctx.JSON(http.StatusOK, updatedCustomer)
}

func (p *customerHandlerImpl) AssignMembership(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if id == 0 || err != nil {
		ctx.JSON(http.StatusBadRequest, pkg.ErrorResponse{Message: "invalid required param"})
		return
	}
	customer, err := p.customerservice.GetCustomersByID(ctx, uint64(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, pkg.ErrorResponse{Message: err.Error()})
		return
	}
	if customer.ID == 0 {
		ctx.JSON(http.StatusNotFound, pkg.ErrorResponse{Message: "Customer not found"})
		return
	}
	member := models.InputMembershipID{}
	if err := ctx.ShouldBindJSON(&member); err != nil {
		ctx.JSON(http.StatusBadRequest, pkg.ErrorResponse{Message: "Invalid request body"})
		return
	}
	updatedCustomer, err := p.customerservice.AssignMembership(ctx, uint64(id), member)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, pkg.ErrorResponse{Message: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, updatedCustomer)
}
func (p *customerHandlerImpl) DeleteMembershipByCustomer(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if id == 0 || err != nil {
		ctx.JSON(http.StatusBadRequest, pkg.ErrorResponse{Message: "invalid required param"})
		return
	}
	customer, err := p.customerservice.GetCustomersByID(ctx, uint64(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, pkg.ErrorResponse{Message: err.Error()})
		return
	}
	if customer.ID == 0 {
		ctx.JSON(http.StatusNotFound, pkg.ErrorResponse{Message: "Customer not found"})
		return
	}
	member := models.InputMembershipID{}
	if err := ctx.ShouldBindJSON(&member); err != nil {
		ctx.JSON(http.StatusBadRequest, pkg.ErrorResponse{Message: "Invalid request body"})
		return
	}
	updatedCustomer, err := p.customerservice.DeleteMembershipByCustomer(ctx, uint64(id), member)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, pkg.ErrorResponse{Message: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, updatedCustomer)
	}
