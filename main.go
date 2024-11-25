package main

//test

import (
        "fmt"
        "log"
        "net/http"
        "time"

        "github.com/gin-gonic/gin"
)

// Estructura para la respuesta de la versión
type VersionResponse struct {
        Version   string `json:"version"`
        Timestamp string `json:"timestamp"`
}

// Estructura para la respuesta de la temperatura
type TemperatureResponse struct {
        AverageTemperature float64 `json:"average_temperature"`
}

// Función para obtener los datos de la API senseBox (simulada)
func getSenseBoxData() ([]float64, time.Time, error) {
        // Simulación de los datos de la API. En un caso real, puedes hacer una solicitud HTTP a la API real.
        data := []float64{22.5, 23.3, 21.8, 22.1, 23.0}
        timestamp := time.Now().Add(-30 * time.Minute) // Simula que los datos son de 30 minutos atrás
        return data, timestamp, nil
}

// Endpoint para obtener la versión de la API
func getVersion(c *gin.Context) {
        version := VersionResponse{
                Version:   "v0.0.1", // Usamos Semantic Versioning: v0.0.1
                Timestamp: time.Now().Format(time.RFC3339), // Agregar la hora actual al formato RFC3339
        }
        c.JSON(http.StatusOK, version)
}

// Endpoint para obtener la temperatura promedio
func getTemperature(c *gin.Context) {
        // Simulamos obtener los datos de senseBox
        data, timestamp, err := getSenseBoxData()
        if err != nil {
                c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch data"})
                return
        }

        // Verificamos si los datos son más recientes que 1 hora
        if time.Since(timestamp) > time.Hour {
                c.JSON(http.StatusBadRequest, gin.H{"error": "Data is too old"})
                return
        }

        // Calculamos la temperatura promedio
        var sum float64
        for _, temp := range data {
                sum += temp
        }
        average := sum / float64(len(data))

        temperature := TemperatureResponse{
                AverageTemperature: average,
        }
        c.JSON(http.StatusOK, temperature)
}

// Función que imprime la versión de la app y termina la ejecución
func printAppVersion() {
        fmt.Println("App Version: v0.0.1")
}

func main() {
        // Imprime la versión de la aplicación y termina la ejecución
        printAppVersion()

        // Crear un router con Gin
        r := gin.Default()

        // Definir las rutas
        r.GET("/version", getVersion)
        r.GET("/temperature", getTemperature)

        // Iniciar el servidor en el puerto 8080
        log.Fatal(r.Run(":8080"))
}
