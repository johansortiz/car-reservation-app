# Usa una imagen base de Go
FROM golang:1.16-alpine

# Instala ping
RUN apk add --no-cache iputils

# Establece el directorio de trabajo dentro del contenedor
WORKDIR /app

# Copia los archivos del proyecto al directorio de trabajo
COPY . .

# Añade el script de espera y ajusta permisos
COPY wait-for-it.sh /wait-for-it.sh
RUN chmod +x /wait-for-it.sh

# Instala las dependencias y construye la aplicación
RUN go mod tidy
RUN go build -o main .

# Expone el puerto en el que la aplicación se ejecutará
EXPOSE 8080

# Comando para ejecutar la aplicación, usando el script de espera
CMD ["/wait-for-it.sh", "db", "./main"]

