package services

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/jinzhu/copier"
	"github.com/leonel-garofolo/dePrimeraApiRest/api/daos"
	"github.com/leonel-garofolo/dePrimeraApiRest/api/daos/gorms"
	models "github.com/leonel-garofolo/dePrimeraApiRest/api/dto"
	"github.com/leonel-garofolo/dePrimeraApiRest/api/dto/response"

	"github.com/labstack/echo/v4"
)

func RouterSanciones(e *echo.Echo) {
	e.GET("/api/sancions", GetSanciones)
	e.GET("/api/sancions/:id", GetSancion)
	e.POST("/api/sancions", SaveSancion)
	e.DELETE("/api/sancions/:id", DeleteSancion)
	e.GET("/api/sancions/info", InfoSanciones)
}

func GetSanciones(c echo.Context) error {
	daos := daos.NewDePrimeraDaos()
	sancionesGorm := daos.GetSancionesDao().GetAll()
	sanciones := []models.Sanciones{}
	copier.Copy(&sanciones, &sancionesGorm)
	return c.JSON(http.StatusOK, sanciones)
}

func GetSancion(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Panic(err)
	}

	daos := daos.NewDePrimeraDaos()
	sancionGorm := daos.GetSancionesDao().Get(id)
	sancion := &models.Sanciones{}
	copier.Copy(&sancion, &sancionGorm)
	return c.JSON(http.StatusOK, sancion)
}

func SaveSancion(c echo.Context) error {
	sancion := &models.Sanciones{}
	c.Bind(sancion)

	sancionGorm := &gorms.SancionesGorm{}
	copier.Copy(&sancionGorm, &sancion)

	daos := daos.NewDePrimeraDaos()
	id := daos.GetSancionesDao().Save(sancionGorm)

	log.Println(id)
	return c.String(http.StatusOK, "insertado")
}

func DeleteSancion(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Panic(err)
	}
	daos := daos.NewDePrimeraDaos()
	status, error := daos.GetSancionesDao().Delete(id)

	resp := &response.UpdatedResponse{}
	resp.Status = status
	if !status {
		resp.Message = "Error al intentar eliminar el Registro."
		sError := error.Error()
		fmt.Println(sError)
		if strings.Contains(sError, "Cannot") {
			resp.Message = "El registro no se pudo eliminar."
		}
	}

	log.Println(id)
	return c.String(http.StatusOK, "delete")
}

func InfoSanciones(c echo.Context) error {
	sancions := &models.Sanciones{}
	c.Bind(sancions)

	j, err := json.Marshal(sancions)
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "error al obtener la info")
	} else {
		log.Println(string(j))

		return c.JSON(http.StatusOK, sancions)
	}
}
