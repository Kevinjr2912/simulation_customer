package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"polling/polling/models/responses"
	"time"

	"github.com/gin-gonic/gin"
)

func CheckStatusInformationStudent(ctx *gin.Context) {
	for {
		response, _ := http.Get("http://localhost:8080/students/13/updates?since=2025-02-20T22:15:42Z")

		defer response.Body.Close()

		if response.StatusCode == http.StatusOK {
			body, err := io.ReadAll(response.Body)

			if err != nil {
				fmt.Println("Error al leer el cuerpo de la respuesta:", err)
				return
			}

			var result responses.ResponseStatus

			err = json.Unmarshal(body, &result)

			if err != nil {
				fmt.Println("Error al parsear JSON:", err)
				return
			}

			fmt.Println("Estado de actualización de la información del estudiante: ", result.Updated)
		} else if response.StatusCode == http.StatusNotModified {
			fmt.Println("No hubo actualizaciones para el estudiante.")
		} else {
			fmt.Printf("Error: Código de estado %d\n", response.StatusCode)
		}

		time.Sleep(10 * time.Second)
	}
}