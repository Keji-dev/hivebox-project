package main

// version 0.1.1

import (
    "fmt"
    "log"
    "net/http"
    "time" // Importando el paquete time
    "github.com/gin-gonic/gin"
)

// Estructura para la respuesta de la versión
type VersionResponse struct {
    Version string `json:"version"`
}

// Estructura para la respuesta de la temperatura
type TemperatureResponse struct {
    AverageTemperature float64 `json:"average_temperature"`
}

// Función para obtener los datos de la API senseBox (simulada)
func getSenseBoxData() ([]float64, error) {
    // Simulación de los datos de la API. En un caso real, puedes hacer una solicitud HTTP a la API real.
    return []float64{22.5, 23.3, 21.8, 22.1, 23.0}, nil
}

// Endpoint para obtener la versión de la API
func getVersion(c *gin.Context) {
    version := VersionResponse{
        Version: "v0.0.1", // Usamos Semantic Versioning: v0.0.1
    }
    c.JSON(http.StatusOK, version)
}

// Endpoint para obtener la temperatura promedio
func getTemperature(c *gin.Context) {
    // Usando time para mostrar la hora en que se recibe la solicitud
    fmt.Println("Request received at:", time.Now())

    // Simulamos obtener los datos de senseBox
    data, err := getSenseBoxData()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch data"})
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
