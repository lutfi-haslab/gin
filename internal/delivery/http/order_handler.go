package http

import (
    "net/http"
    "strconv"
    "github.com/gin-gonic/gin"
    "clean-api/internal/domain"
)

type OrderHandler struct {
    orderUseCase domain.OrderRepository
}

func NewOrderHandler(r *gin.Engine, orderUseCase domain.OrderRepository) {
    handler := &OrderHandler{orderUseCase}

    orders := r.Group("/orders")
    {
        orders.POST("/", handler.Create)
        orders.GET("/", handler.GetAll)
        orders.GET("/:id", handler.GetByID)
        orders.GET("/user/:userId", handler.GetUserOrders)
        orders.PUT("/:id", handler.Update)
        orders.DELETE("/:id", handler.Delete)
    }
}

func (h *OrderHandler) Create(c *gin.Context) {
    var order domain.Order
    if err := c.ShouldBindJSON(&order); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := h.orderUseCase.Create(&order); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, order)
}

func (h *OrderHandler) GetByID(c *gin.Context) {
    id, err := strconv.ParseUint(c.Param("id"), 10, 32)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
        return
    }

    order, err := h.orderUseCase.GetByID(uint(id))
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "order not found"})
        return
    }

    c.JSON(http.StatusOK, order)
}

func (h *OrderHandler) GetAll(c *gin.Context) {
    orders, err := h.orderUseCase.GetAll()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, orders)
}

func (h *OrderHandler) GetUserOrders(c *gin.Context) {
    userID, err := strconv.ParseUint(c.Param("userId"), 10, 32)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
        return
    }

    orders, err := h.orderUseCase.GetUserOrders(uint(userID))
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, orders)
}

func (h *OrderHandler) Update(c *gin.Context) {
    id, err := strconv.ParseUint(c.Param("id"), 10, 32)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
        return
    }

    var order domain.Order
    if err := c.ShouldBindJSON(&order); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    order.ID = uint(id)
    if err := h.orderUseCase.Update(&order); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, order)
}

func (h *OrderHandler) Delete(c *gin.Context) {
    id, err := strconv.ParseUint(c.Param("id"), 10, 32)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
        return
    }

    if err := h.orderUseCase.Delete(uint(id)); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "order deleted successfully"})
}