package delivery

import (
	"RESTAPILoundry/domain"
	"RESTAPILoundry/feature/common"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type serviceHandler struct {
	serviceUsecase domain.ServiceUseCase
}

func New(su domain.ServiceUseCase) domain.ServiceHandler {
	return &serviceHandler{
		serviceUsecase: su,
	}
}

func (sh *serviceHandler) InsertServ() echo.HandlerFunc {
	return func(c echo.Context) error {
		var tmp ServicesInsertRequest
		err := c.Bind(&tmp)

		if err != nil {
			log.Println("Cannot parse data", err)
			c.JSON(http.StatusBadRequest, "error read input")
		}
		_, role := common.ExtractData(c)
		if role != "admin" {
			return c.JSON(http.StatusCreated, map[string]interface{}{
				"message": "Only Admin can Insert Product Data",
			})
		}

		data, err := sh.serviceUsecase.AddService(tmp.ToDomain())
		if err != nil {
			log.Println("Cannot proces data", err)
			c.JSON(http.StatusInternalServerError, err)

		}
		return c.JSON(http.StatusCreated, map[string]interface{}{
			"message": "success create data",
			"data":    data,
		})
	}
}

func (sh *serviceHandler) UpdateServ() echo.HandlerFunc {
	return func(c echo.Context) error {

		qry := map[string]interface{}{}
		cnv, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			log.Println("Cannot convert to int", err.Error())
			return c.JSON(http.StatusInternalServerError, "cannot convert id")
		}

		var tmp ServicesInsertRequest
		res := c.Bind(&tmp)

		if res != nil {
			log.Println(res, "Cannot parse data")
			return c.JSON(http.StatusInternalServerError, "error read update")
		}

		_, role := common.ExtractData(c)
		// fmt.Println(role)

		if role != "admin" {
			return c.JSON(http.StatusCreated, map[string]interface{}{
				"message": "Only Admin can update Services data",
			})
		}

		if tmp.Name != "" {
			qry["name"] = tmp.Name
		}
		if tmp.Price != 0 {
			qry["price"] = tmp.Price
		}
		data, err := sh.serviceUsecase.UpdateService(cnv, tmp.ToDomain())

		if err != nil {
			log.Println("Cannot update data", err)
			c.JSON(http.StatusInternalServerError, "cannot update")
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "success update data",
			"data":    data,
		})

	}
}

func (sh *serviceHandler) DeleteServ() echo.HandlerFunc {
	return func(c echo.Context) error {

		cnv, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			log.Println("Cannot convert to int", err.Error())
			return c.JSON(http.StatusInternalServerError, "cannot convert id")
		}

		_, role := common.ExtractData(c)
		// fmt.Println(role)

		if role != "admin" {
			return c.JSON(http.StatusCreated, map[string]interface{}{
				"message": "Only Admin can delete Data",
			})
		}

		data := sh.serviceUsecase.DeleteService(cnv)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, "cannot delete data")
		}

		if data == false {
			return c.JSON(http.StatusInternalServerError, "cannot delete")
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "success delete product",
			"data":    data,
		})

	}
}

func (sh *serviceHandler) GetAllServ() echo.HandlerFunc {
	return func(c echo.Context) error {
		data, err := sh.serviceUsecase.GetAllS()
		if err != nil {
			log.Println("Cannot get data", err)
			return c.JSON(http.StatusBadRequest, "error read input")

		}

		if data == nil {
			log.Println("Terdapat error saat mengambil data")
			return c.JSON(http.StatusInternalServerError, "Problem from database")
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "success get all Product",
			"data":    data,
		})
	}
}

func (sh *serviceHandler) GetServID() echo.HandlerFunc {
	return func(c echo.Context) error {
		idProduct := c.Param("id")
		id, _ := strconv.Atoi(idProduct)
		data, err := sh.serviceUsecase.GetSpecificServices(id)

		if err != nil {
			log.Println("Cannot get data", err)
			return c.JSON(http.StatusBadRequest, "cannot read input")
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "success get data",
			"data":    data,
		})
	}
}
