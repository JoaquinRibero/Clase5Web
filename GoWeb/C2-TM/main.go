package main

import (
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type transaction struct {
	Id       int    `json:"_id"`
	Codigo   string `json:"codigo" binding:"required"`
	Moneda   string `json:"moneda" binding:"required"`
	Monto    int    `json:"monto" binding:"required"`
	Emisor   string `json:"emisor" binding:"required"`
	Receptor string `json:"receptor" binding:"required"`
	Fecha    string `json:"fecha" binding:"required"`
}

var transactions = []transaction{}

func newUser(ctx *gin.Context) {
	var req transaction
	if err := ctx.ShouldBindJSON(&req); err != nil {
		var verr validator.ValidationErrors
		if errors.As(err, &verr) {
			ctx.JSON(400, gin.H{"errors": Mensaje(verr)})
			return
		}
	}
	lastId := len(transactions)
	req.Id = lastId + 1
	transactions = append(transactions, req)
	ctx.JSON(200, transactions)
}

func Mensaje(verr validator.ValidationErrors) map[string]string {
	errs := make(map[string]string)
	for _, f := range verr {
		errs[f.Field()] = fmt.Sprintf("El campo %s es requerido", f.Field())
	}
	return errs
}

func validateToken(ctx *gin.Context) {
	token := ctx.GetHeader("token")
	if token != "joaquin" {
		ctx.JSON(401, gin.H{"errors": "no tiene permisos para realizar la peticion solicitada"})
	}
}

func listar(ctx *gin.Context) {
	ctx.JSON(200, transactions)
}

func main() {
	router := gin.Default()
	tr := router.Group("/transactions", validateToken)
	tr.POST("/new", newUser)
	tr.GET("/", listar)

	router.Run(":8080")

}
