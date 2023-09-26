package shopping

import (
	"app/domain"
	"app/utils"
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

const (
	url    = "https://apirecruit-gjvkhl2c6a-uc.a.run.app/compras/"
	format = "2006-01-02"
)

func GetSummary(c echo.Context) error {
	var shoppingDomain domain.ShoppingDate

	if err := c.Bind(&shoppingDomain); err != nil {
		return utils.ErrorResponse(utils.ErrorSMS{
			StatusCode: http.StatusBadRequest,
			Message:    "bad request",
		}, c)
	}

	date, err := time.Parse("2006-01-02", shoppingDomain.Date)

	if err != nil {
		return utils.ErrorResponse(utils.ErrorSMS{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
		}, c)
	}

	if shoppingDomain.Day == 0 {
		shoppingDomain.Day = 1
	}

	var registers [][]domain.Register
	for i := 0; i < shoppingDomain.Day; i++ {
		localUrl := fmt.Sprintf("%s%s", url, date.Format(format))
		result, err := utils.Call(localUrl)

		if err != nil {
			return utils.ErrorResponse(utils.ErrorSMS{
				StatusCode: http.StatusBadRequest,
				Message:    "date not acceptable",
			}, c)
		}

		var register []domain.Register
		err = utils.DecoderReader(result.Body, &register)

		if err != nil {
			return utils.ErrorResponse(utils.ErrorSMS{
				StatusCode: http.StatusBadRequest,
				Message:    err.Error(),
			}, c)
		}

		date = date.Add(time.Hour * 24)
		registers = append(registers, register)
		result.Body.Close()
	}

	output := domain.Summary{
		Total:         0,
		ComprasPorTDC: map[string]float64{},
		NoCompraron:   0,
		CompraMasAlta: 0,
	}

	for _, register := range registers {
		for _, v := range register {
			if !v.Compro {
				output.NoCompraron++
				continue
			}
			if output.CompraMasAlta < v.Monto {
				output.CompraMasAlta = v.Monto
			}
			output.ComprasPorTDC[v.TDC] += v.Monto
			output.Total += v.Monto
		}
	}

	return c.JSON(http.StatusOK, output)
}
