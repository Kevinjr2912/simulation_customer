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

var arrayStudents []responses.ResponseDataStudents

func CheckListStudents(ctx *gin.Context) {
	for {
		response, _ := http.Get("http://localhost:8080/students")

		defer response.Body.Close()

		if response.StatusCode == http.StatusOK {
			body, err := io.ReadAll(response.Body)

			if err != nil {
				fmt.Println("Error al leer el cuerpo de la respuesta:", err)
 				return
			}

			var result responses.ResponseGetAllStudents

			err = json.Unmarshal(body, &result)

			if err != nil {
				fmt.Println("Error al parsear JSON:", err)
				return
			}

			lenArrayStudents := len(result.Data)

			if lenArrayStudents != len(arrayStudents) {
				fmt.Println("Hay cambios")

				arrayStudents = result.Data

			} else {
				fmt.Println("No hay cambios")
			}
			
		} else {
			fmt.Printf("Error: Código de estado %d\n", response.StatusCode)
		}

		time.Sleep(10 * time.Second)
	}	
}
