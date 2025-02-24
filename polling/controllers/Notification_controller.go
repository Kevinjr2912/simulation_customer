package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func NotifyStudent(ctx *gin.Context) {
	for {
		response, _ := http.Get("http://localhost:8080/students/notification")

		defer response.Body.Close()

		if response.StatusCode == http.StatusOK {
			body, err := io.ReadAll(response.Body)

			if err != nil {
				fmt.Println("Error al leer el cuerpo de la respuesta:", err)
				time.Sleep(5 * time.Second)
				continue 
			}

			type Response struct {
				Mensaje string `json:"Mensaje"`
			}

			var result Response

			err = json.Unmarshal(body, &result)

			if err != nil {
				fmt.Println("Error al parsear JSON:", err)
				return
			}

			fmt.Println("Notificación: " + result.Mensaje)

		} else {
			fmt.Printf("Error: Código de estado %d\n", response.StatusCode)
		}

		time.Sleep(5 * time.Second)
	}
}

func NotifyCareer(ctx *gin.Context) {
	for {
		response, _ := http.Get("http://localhost:8080/careers/notification")

		defer response.Body.Close()

		if response.StatusCode == http.StatusOK {
			body, err := io.ReadAll(response.Body)

			if err != nil {
				fmt.Println("Error al leer el cuerpo de la respuesta:", err)
				time.Sleep(5 * time.Second)
				continue 
			}

			type Response struct {
				Mensaje string `json:"Mensaje"`
			}

			var result Response

			err = json.Unmarshal(body, &result)

			if err != nil {
				fmt.Println("Error al parsear JSON:", err)
				return
			}

			fmt.Println("Notificación: " + result.Mensaje)

		} else {
			fmt.Printf("Error: Código de estado %d\n", response.StatusCode)
		}

		time.Sleep(5 * time.Second)
	}
}
