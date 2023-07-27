# GRPC

## Instalación en windows

#### 1. Descargar protoc: 
https://github.com/protocolbuffers/protobuf/releases/download/v23.4/protoc-23.4-win64.zip

#### 2. Instalar protoc:
Descomprimir en algún lugar y registar la ruta bin en las variables de entorno del sistema.

#### 3. Instalar packages en el proyecto GO:
    $ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc
    $ go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
    $ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
    $ go get -u google.golang.org/grpc

