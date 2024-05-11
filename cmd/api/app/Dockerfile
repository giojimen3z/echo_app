# Etapa de construcción
FROM golang:1.21-alpine as builder

# Establece el directorio de trabajo en el contenedor
WORKDIR /app

# Instalar herramientas necesarias, gcc y libc-dev son comunes para paquetes que requieren cgo
RUN apk add --no-cache git gcc libc-dev

# Copiar go.mod y go.sum y descargar las dependencias
COPY ../../../go.mod go.sum ./
RUN go mod download

# Copiar el resto del código fuente
COPY ../../.. .

# Compilar la aplicación, desactivar cgo para crear un binario estático
RUN CGO_ENABLED=0 GOOS=linux go build -o /bin/pingpong echo_app/cmd/api/app/main.go

# Etapa final
FROM alpine:latest
WORKDIR /

# Instalar ca-certificates en caso de que tu aplicación haga llamadas HTTPS
RUN apk add --no-cache ca-certificates

# Copiar el ejecutable compilado en la etapa anterior al directorio raíz
COPY --from=builder /bin/pingpong /pingpong

# Exponer el puerto que tu aplicación usará
EXPOSE 8080

# Comando para ejecutar la aplicación
CMD ["/pingpong"]
