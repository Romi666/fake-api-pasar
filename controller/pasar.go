package controller

import (
	"fake-web-pasar-api/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

var listPasar = []models.PasarTable{
	{PasarID: "1", NamaPasar: "Pasar Cihaurgeulis", Alamat: "Pasar cihaurgeulis Jl.suci, Sukaluyu, Kec. Cibeunying Kaler, Kota Bandung, Jawa Barat 02123, Indonesia", Deskripsi: "", Kota: "Kota Bandung", Provinsi: "Jawa Barat", Image: "", Kecamatan: "", UrlSid: "https://s.id/pasartest1"},
	{PasarID: "2", NamaPasar: "Pasar Gedebage", Alamat: "Jl. Sor GBLA, Rancanumpang, Kec. Gedebage, Kota Bandung, Jawa Barat 40292, Indonesia", Deskripsi: "", Kota: "Kota Bandung", Provinsi: "Jawa Barat", Image: "", Kecamatan: "", UrlSid: "https://s.id/pasartest1"},
}

func GetDetailPasar(c *gin.Context) {
	var (
		response models.Response
		responseSuccess models.ResponseSuccessGetDetailPasar
		request	models.RequestDetailPasar
	)

	currentTime := time.Now()
	timeFormat := currentTime.Format("2006-01-02 15:04:05")

	if err := c.BindJSON(&request); err != nil {
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

	for _, item := range listPasar {
		if item.PasarID == request.PasarID{
			responseSuccess.ResponseCode = "00"
			responseSuccess.Deskripsi = "Sukses"
			responseSuccess.Data = models.ResponseDetailPasar{
				NamaPasar: item.NamaPasar,
				Alamat: item.Alamat,
				Deskripsi: item.Deskripsi,
				Distance: 3,
			}
			c.JSON(http.StatusOK, responseSuccess)
			return
		}
	}
	response.ResponseCode = "02"
	response.Deskripsi = "Data Not Found"
	response.Data = models.DataResponse{
		DeskripsiError: "Pasar Is Empty",
		DateFromat:     timeFormat,
	}
	c.JSON(http.StatusNotFound, response)
	c.Abort()
	return

}

func GetListPasar(c *gin.Context) {
	var (
		pasar []models.DetailGetAllPasar
		request models.RequestGetAllPasar
		response models.Response
		responseSuccess models.ResponseSuccessGetAllPasar
	)

	currentTime := time.Now()
	timeFormat := currentTime.Format("2006-01-02 15:04:05")

	if err := c.BindJSON(&request); err !=nil{
		response.ResponseCode = "01"
		response.Deskripsi = "Invalid Signing Form"
		response.Data = models.DataResponse{
			DeskripsiError: err.Error(),
			DateFromat: timeFormat,
		}
		c.JSON(http.StatusBadRequest, response)
		c.Abort()
		return
	}

	if request.Limit == ""{
		response.ResponseCode = "01"
		response.Deskripsi = "Invalid Signing Form"
		response.Data = models.DataResponse{
			DeskripsiError: "invalid limit",
			DateFromat: timeFormat,
		}
		c.JSON(http.StatusBadRequest, response)
		c.Abort()
		return
	}

	if len(listPasar) <= 0 {
		response.ResponseCode = "02"
		response.Deskripsi = "Data not found"
		response.Data = models.DataResponse{
			DeskripsiError: "Pasar is empty",
			DateFromat: timeFormat,
		}
		c.JSON(http.StatusNotFound, response)
		c.Abort()
		return
	}
	for _, item := range listPasar {
		pasar = append(pasar, models.DetailGetAllPasar{PasarID: item.PasarID, NamaPasar: item.NamaPasar, Alamat: item.Alamat, Image: item.Image, UrlSid: item.UrlSid, Kecamatan: item.Kecamatan, Kota: item.Kota, Provinsi: item.Provinsi})
	}

	responseSuccess.ResponseCode = "00"
	responseSuccess.Deskripsi = "Sukses"
	responseSuccess.Data = pasar

	c.JSON(http.StatusOK, responseSuccess)
	return
}