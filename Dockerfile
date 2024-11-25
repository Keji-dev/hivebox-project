# Usamos la imagen oficial de Go, que es más eficiente que instalar Go desde cero.
FROM golang:1.23.3-alpine3.20 as builder

# Directorio de trabajo dentro del contenedor
WORKDIR /app

# Copiar los archivos del proyecto al contenedor
COPY . .

# Ejecutamos `go mod tidy` para asegurarnos de que todas las dependencias estén correctamente instaladas
RUN go mod tidy

# Construir el binario de la aplicación
RUN go build -o hivebox .

# Usamos una imagen más liviana para la ejecución del contenedor (sin Go).
FROM alpine:latest

# Instalamos dependencias necesarias para ejecutar la app (como libc)
RUN apk --no-cache add ca-certificates

# Copiar el binario desde la etapa de construcción
COPY --from=builder /app/hivebox /usr/local/bin/hivebox

# Exponer el puerto en el que nuestra API va a estar escuchando
EXPOSE 8080

# Comando para ejecutar el binario
ENTRYPOINT ["/usr/local/bin/hivebox"]
