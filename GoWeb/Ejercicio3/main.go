package main

import (
	"encoding/json"
	"os"
	"strconv"
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

func GetAll(ctx *gin.Context) {

	params := ctx.Request.URL.Query()

	transactions := []transaction{}

	data, _ := os.ReadFile("../transactions.json")
	json.Unmarshal([]byte(data), &transactions)
	final := []transaction{}
	for _, tx := range final {
		filter := false
		if params.Get("codigo") == tx.Codigo {
			filter = true
		}
		id, _ := strconv.Atoi(params.Get("id"))
		if id == tx.Id {
			filter = true
		}
		if params.Get("moneda") == tx.Moneda {
			filter = true
		}
		if params.Get("emisor") == tx.Emisor {
			filter = true
		}
		if params.Get("receptor") == tx.Receptor {
			filter = true
		}
		monto, _ := strconv.Atoi(params.Get("monto"))
		if monto == tx.Monto {
			filter = true
		}
		if filter == true {
			final = append(final, tx)
		}
	}
	if len(final) > 0 {
		ctx.JSON(200, final)
	} else {
		ctx.JSON(200, transactions)
	}
}

func main() {
	/*t1 := transaction{
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

	transactions := []transaction{t1, t2}*/

	router := gin.Default()

	router.GET("/transactions", GetAll)
	router.Run(":3000")

}
