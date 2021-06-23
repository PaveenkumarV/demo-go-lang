package Controllers

import (
	"crud-go/Config"
	"fmt"

	"github.com/gin-gonic/gin"
)

func GetEmp(c *gin.Context) {
	// var user []Model.Employee
	res, err := Config.DB.Query("SELECT * FROM cities")
	if err != nil {
		fmt.Println("Error Arised While Getting Data From The Table ", err)
	}
	fmt.Println(res)
}
