package main

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "your_project/models"
    "your_project/services"
)

type Api struct {
    UserService    services.UserService
    ProductService services.ProductService
}

func NewApi(userService services.UserService, productService services.ProductService) *Api {
    return &Api{
        UserService:    userService,
        ProductService: productService,
    }
}

func (api *Api) RegisterRoutes(router *gin.Engine) {
    router.POST("/users", api.CreateUser)
    router.GET("/users/:id", api.GetUser)
    router.GET("/users", api.ListUsers)
    router.DELETE("/users/:id", api.DeleteUser)

    router.POST("/products", api.CreateProduct)
    router.GET("/products/:id", api.GetProduct)
    router.GET("/products", api.ListProducts)
    router.DELETE("/products/:id", api.DeleteProduct)
}

func (api *Api) CreateUser(c *gin.Context) {
    var user models.User
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    id, err := api.UserService.Register(c, user)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, gin.H{"id": id})
}

func (api *Api) GetUser(c *gin.Context) {
    id := c.Param("id")
    user, err := api.UserService.GetUser(c, id)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
        return
    }
    c.JSON(http.StatusOK, user)
}

func (api *Api) ListUsers(c *gin.Context) {
    users, err := api.UserService.ListUsers(c)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, users)
}

func (api *Api) DeleteUser(c *gin.Context) {
    id := c.Param("id")
    if err := api.UserService.DeleteUser(c, id); err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
        return
    }
    c.Status(http.StatusNoContent)
}

func (api *Api) CreateProduct(c *gin.Context) {
    var product models.Product
    if err := c.ShouldBindJSON(&product); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    id, err := api.ProductService.CreateProduct(c, product)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, gin.H{"id": id})
}

func (api *Api) GetProduct(c *gin.Context) {
    id := c.Param("id")
    product, err := api.ProductService.GetProduct(c, id)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
        return
    }
    c.JSON(http.StatusOK, product)
}

func (api *Api) ListProducts(c *gin.Context) {
    products, err := api.ProductService.ListProducts(c)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, products)
}

func (api *Api) DeleteProduct(c *gin.Context) {
    id := c.Param("id")
    if err := api.ProductService.DeleteProduct(c, id); err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
        return
    }
    c.Status(http.StatusNoContent)
}