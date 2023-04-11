# info de la materia: ST0263 Topicos Especiales en Telematica
#
# Estudiante(s): Andres Salazar Galeano, asalaza5@eafit.edu.co & Julian David Ramirez, jdramirezl@eafit.edu.co 
#
# Profesor: Edwin Nelson Montoya, emontoya@eafit.brightspace.com

# Diseño e Implementación de un Middleware que Implemente un Servicio de Mensajería Asincrónica entre Aplicaciones 

# 1. Breve descripción de la actividad
 Diseñar e implementar un middleware que permita enviar y recibir mensajes entre aplicaciones. 
 
 El middleware debe permitir la conexión/desconexión al servidor, el ciclo de vida de tópicos y colas y el envío y recepción de mensajes. 
 
Ademas debe tener mecanismos de persistencia, tolerancia a fallos, y considerar aspectos de escalabilidad y extensibilidad en el diseño de la arquitectura.
## 1.1. Que aspectos cumplió o desarrolló de la actividad propuesta por el profesor (requerimientos funcionales y no funcionales)

Funcionales:
- Permitir la conexion al servidor
- Permitir la desconexion al servidor
- Crear topicos 
- Borrar topicos   por el creador
- Enviar mensajes a topicos
- Recibir mensajes a topicos
- Crear colas
- Borrar colas por el creador
- Enviar mensajes a colas
- Recibir mensajes de colas
- Identificar los usuarios
- Exponer todos los servicios como una API REST hacia los clientes.
- Implementar tolerancia a fallos
- Implementar replicacion

No funcionales:
- El middleware debe ser robusto y tolerante a fallos.
- El middleware debe escalable y extensible.
- El middleware debe ser compatible con aplicaciones que utilicen diferentes lenguajes de programación.


## 1.2. Que aspectos NO cumplió o desarrolló de la actividad propuesta por el profesor (requerimientos funcionales y no funcionales)

# 2. Información general de diseño de alto nivel, arquitectura, patrones, mejores prácticas utilizadas.

![Deployment Top  Telematica (1)](https://user-images.githubusercontent.com/65835577/231276443-dddecf15-2b8f-46aa-b995-34e72810e23f.png)


- Arquitectura distribuida Peer-to-Peer semiestructurado con servidor para resolver el maestro.
- Mejores practicas:
    * Implementacion y orquestacion de contenedores
    * Alta disponibilidad y tolerancia a fallos con nodos replicados
    * Uso de variables de entorno. 
    * Sistema de directorios y carpetas, 
    * Notacion de archivos, carpetas, variables, funciones y clases

# 3. Descripción del ambiente de EJECUCIÓN (en producción) lenguaje de programación, librerias, paquetes, etc, con sus numeros de versiones.

## Tecnologias y herramientas
- Lenguaje de programacion: Golang 1.18
- gRPC
- Docker
- Docker Compose
- AWS EC2

## Librerias
- "os"
- "net"
- "fmt"
- "log"
- "time"
- "errors"
- "context"
- "strings"
- "google.golang.org/grpc"
- "google.golang.org/grpc/peer"
- "google.golang.org/grpc/reflection"
- "github.com/golang/protobuf/ptypes/empty"

## IP o nombres de dominio en nube o en la máquina servidor.
TODO: DREW!!!!

## descripción y como se configura los parámetros del proyecto (ej: ip, puertos, conexión a bases de datos, variables de ambiente, parámetros, etc)

### Descripcion de los parametros
Los parametros de configuracion son los siguientes:

Para los nodos
```
- SELF_IP= // IP propia
- PORT= // Puerto propio/generico
- MOM_PORT= // Puerto para comunicacion del MOM  (externa)
- RESOLVER_IP= // IP del resolvedor de IPs
- RESOLVER_PORT= // Puerto del resolvedor de IPs
- CLUSTER_PORT= // Puerto para la comunicacion del cluster (interna)
```

Para el gateway
```
- MOM_HOST= // IP del resolver
- MOM_PORT= // Puerto del resolver
- CLUSTER_PORT= // Puerto para comunicarse con el cluster
- PORT= // Puerto propio, generico
```


Para el resolvedor de IPs
```
- PORT=8080 // Unico parametro, su puerto
```

### Configuracion de los parametros
1. Ingresar al docker-compose
2. En cada contenedor cambiar los parametros de configuracion en la seccion "environment"

## Como se lanza el servidor.
1. Ingresar a AWS
2. Acceder a la consola de las instancias EC2 mediante SSH 
    * gateway: Instancia donde tenemos el gateway
    * mom: Instancia donde tenemos el MOM Server
3. Levantar el contenedor en cada parte mediante `sudo docker-compose up` 


## Guia de uso para usuario

Con las isntancias ya corriendo, se puede acceder a los servicios mediante el gateway.

Las rutas de acceso son las siguientes:
- addMessage
    * body: Cuerpo del mensaje
    * broker: Nombre del queue/topic al que se quiere enviar el mensaje
    * brokerType: Tipo de broker al que se quiere enviar el mensaje ("queue" o "topic")
- removeConnection
- addConnection
- unsubscribe
- subscribe
- deleteSystem
- createSystem

Para usar cada una de estas rutas, se debe enviar un request GET con los parametros necesarios en el formato:
`http://<gateway_ip>:<gateway_port>/<ruta>?<parametros>`

## opcionalmente - si quiere mostrar resultados o pantallazos 

# 5. otra información que considere relevante para esta actividad.

# referencias:
<debemos siempre reconocer los créditos de partes del código que reutilizaremos, así como referencias a youtube, o referencias bibliográficas utilizadas para desarrollar el proyecto o la actividad>
## sitio1-url 
## sitio2-url
## url de donde tomo info para desarrollar este proyecto

#### versión README.md -> 1.0 (2023-abril)
