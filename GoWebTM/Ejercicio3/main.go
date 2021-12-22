package main

import (
	"time"

	"github.com/gin-gonic/gin"
)

type transaction struct {
	Id       int       `json:"_id"`
	Codigo   string    `json:"codigo"`
	Moneda   string    `json:"moneda"`
	Monto    int       `json:"monto"`
	Emisor   string    `json:"emisor"`
	Receptor string    `json:"receptor"`
	Fecha    time.Time `json:"fecha"`
}

func main() {
	t1 := transaction{
		Id:       1,
		Codigo:   "abc123",
		Moneda:   "dolar",
		Monto:    150,
		Emisor:   "pepito",
		Receptor: "juancito",
		Fecha:    time.Now()}

	t2 := transaction{
		Id:       2,
		Codigo:   "abc456",
		Moneda:   "pesos",
		Monto:    5000,
		Emisor:   "tony",
		Receptor: "steve",
		Fecha:    time.Date(2021, time.December, 22, 0, 0, 0, 0, time.UTC)}

	transactions := []transaction{t1, t2}

	router := gin.Default()

	router.GET("/transactions/GetAll", func(ctx *gin.Context) {
		//u, _ := json.Marshal(transactions)
		ctx.JSON(200, transactions)
	})
	router.Run(":3000")

}
