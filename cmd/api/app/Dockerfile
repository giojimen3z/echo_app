# Etapa de construcción
FROM golang:latest as builder

WORKDIR /app

# Instalar herramientas necesarias
RUN apt-get update && apt-get install -y git gcc libc-dev ca-certificates && update-ca-certificates

# Copiar los archivos go.mod y go.sum al directorio de trabajo desde la raíz del proyecto
COPY go.mod go.sum ./
RUN go mod download

# Copiar el resto del código fuente al directorio de trabajo
COPY . .

# Compilar la aplicación a un binario
RUN CGO_ENABLED=0 GOOS=linux go build -o /bin/pingpong ./cmd/api/app/main.go
# Añadir un paso para verificar que el archivo existe
RUN ls -l /bin/pingpong

# Etapa final
FROM debian:latest
WORKDIR /

# Instalar ca-certificates en caso de que tu aplicación haga llamadas HTTPS
RUN apt-get update && apt-get install -y ca-certificates && update-ca-certificates

# Copiar el ejecutable compilado en la etapa anterior al directorio raíz
COPY --from=builder /bin/pingpong /pingpong

# Definir la variable de entorno PORT
ENV PORT 8080

# Exponer el puerto que tu aplicación usará
EXPOSE 8080

# Comando para ejecutar la aplicación
CMD ["/pingpong"]
