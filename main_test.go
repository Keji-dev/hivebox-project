package main
//test
import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// Test de la versión
func TestGetVersion(t *testing.T) {
	// Configuramos Gin para las pruebas
	r := gin.Default()
	r.GET("/version", getVersion)

	// Creamos una solicitud GET para la ruta /version
	req, _ := http.NewRequest("GET", "/version", nil)
	w := httptest.NewRecorder()

	// Ejecutamos la solicitud
	r.ServeHTTP(w, req)

	// Verificamos que el estado de la respuesta sea 200 OK
	assert.Equal(t, http.StatusOK, w.Code)

	// Verificamos que la respuesta tenga la versión correcta
	expected := `{"version":"v0.0.1","timestamp":"` + time.Now().Format(time.RFC3339)[:19] + `"}`
	assert.JSONEq(t, expected, w.Body.String())
}

// Test de la temperatura promedio
func TestGetTemperature(t *testing.T) {
	// Configuramos Gin para las pruebas
	r := gin.Default()
	r.GET("/temperature", getTemperature)

	// Creamos una solicitud GET para la ruta /temperature
	req, _ := http.NewRequest("GET", "/temperature", nil)
	w := httptest.NewRecorder()

	// Ejecutamos la solicitud
	r.ServeHTTP(w, req)

	// Verificamos que el estado de la respuesta sea 200 OK
	assert.Equal(t, http.StatusOK, w.Code)

	// Verificamos que la respuesta tenga la temperatura promedio correcta
	expected := `{"average_temperature":22.574}`
	assert.JSONEq(t, expected, w.Body.String())
}
