package controller

import (
	"encoding/json"
	"fake-web-pasar-api/models"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func GetAllProduct(c *gin.Context)  {
	var (
		response 		models.Response
		//responseSuccess models.ResponseSuccessGetAllProduct
		result 			models.HTTPResult
	)

	currentTime := time.Now()
	timeFormat := currentTime.Format("2006-01-02 15:04:05")

	url := os.Getenv("URL")

	res, err := http.Get(url)
	if err!=nil{
		response.ResponseCode = "01"
		response.Deskripsi = "Invalid signin form"
		response.Data = models.DataResponse{
			DeskripsiError: err.Error(),
			DateFromat: timeFormat,
		}
		c.JSON(http.StatusBadRequest, response)
		c.Abort()
		return
	}

	data, _ := ioutil.ReadAll(res.Body)
	json.Unmarshal(data, &result)
	c.JSON(http.StatusOK, result)
}