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
- Gateway private: 172.31.87.80
- Gateway public (Elastic): 52.207.244.87
- Consumer private: 172.31.87.2
- Resolver private: 172.31.89.34
- Replica 1 private: 172.31.85.186
- Replica 2 private: 172.31.94.68
- Replica 3 private: 172.31.80.125

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
- PORT= // Unico parametro, su puerto
```

Para el consumidor
```
- SELF_IP= // IP propia
- MOM_HOST= // IP del Gateway
- MOM_PORT= // Puerto del gateway
- CLUSTER_PORT= // Puerto para comunicarse con el cluster
```

### Configuracion de los parametros
1. Ingresar a la maquina que se desea modificar
2. Ingresar a la carpeta del proyecto con `cd proyecto-1-topicos` 
3. Elegir el docker-compose de su instancia elegida (Si es el resolver se llama `docker-compose.resolver.yml`)
4. Editar el archivo con su editor favorito. E.G: `nvim docker-compose.resolver.yml`
5. En cada contenedor cambiar los parametros de configuracion en la seccion "environment"

## Como se lanza el servidor.
1. Ingresar a AWS
2. Acceder a la terminal de cada instancia, mediante su metodo preferido (E.G: SSH)
3. Levantar el contenedor en cada parte mediante `sudo docker-compose up` en el siguiente orden: 
    * resolver: Instancia con el resolver de IPs del mom y esperar 10 segundos (O hasta que inicie)
    * replica: Instancia(s) con los brokers del sistema, levantar cada uno con unos 10 segundos de diferencia
    * gateway: Instancia con el gateway
    * consumer: Instancia con el consumidor del sistema, puede haber mas de uno

## Guia de uso para usuario

Con las isntancias ya corriendo, se puede acceder a los servicios mediante el gateway.

Las rutas de acceso son las siguientes:
- addMessage
    * PARAMETERS:
    * body: Cuerpo del mensaje
    * broker: Nombre del queue/topic al que se quiere enviar el mensaje
    * brokerType: Tipo de broker al que se quiere enviar el mensaje ("queue" o "topic")

- removeConnection
  * PARAMETERS:
  * IP: IP del cliente que hace el request

- addConnection
  * PARAMETERS:
  * IP: IP del cliente que hace el request

- unsubscribe
  * PARAMETERS:
  * IP: IP del cliente que hace el request
  * broker: Nombre del queue/topic al que se quiere enviar el mensaje
  * brokerType: Tipo de broker al que se quiere enviar el mensaje ("queue" o "topic")

- subscribe
  * PARAMETERS:
  * IP: IP del cliente que hace el request
  * broker: Nombre del queue/topic al que se quiere enviar el mensaje
  * brokerType: Tipo de broker al que se quiere enviar el mensaje ("queue" o "topic")

- deleteSystem
  * PARAMETERS:
  * IP: IP del cliente que hace el request
  * broker: Nombre del queue/topic al que se quiere enviar el mensaje
  * type: Tipo de broker al que se quiere enviar el mensaje ("queue" o "topic")

- createSystem
  * PARAMETERS:
  * IP: IP del cliente que hace el request
  * broker: Nombre del queue/topic al que se quiere enviar el mensaje
  * type: Tipo de broker al que se quiere enviar el mensaje ("queue" o "topic")

Para usar cada una de estas rutas, se debe enviar un request GET con los parametros necesarios en el formato:
`http://<gateway_ip>:<gateway_port>/<ruta>?<parametros>`

Notas:
- Para enviar un mensaje el contenido se pone en el body!
- Es importante seguir un orden de operaciones. 
- Para hacer cualquier accion se debe primero conectar al MOM con AddConnection. 
- Seguido a esto ya podemos subscribirnos, crear colas/topicos, enviar mensajes, etc.
- Para consumir mensajes debemos subscribirnos primero a la cola/topico.


# 5. otra información que considere relevante para esta actividad.

# referencias:
- [Descripcion del proyecto](https://eafit-my.sharepoint.com/:w:/g/personal/emontoya_eafit_edu_co/EWjifZZ2IndKj10aze4_qKYBXRLdUXyMstfUYn70KwqCWA?e=gKcMGu)
- [Master-slave replication](https://dennylesmana.medium.com/master-slave-replication-database-concept-for-beginners-300b3f9a8228)
- [Raft consensus](https://dennylesmana.medium.com/master-slave-replication-database-concept-for-beginners-300b3f9a8228)
- [Peer solving](https://learn.microsoft.com/en-us/dotnet/framework/wcf/feature-details/peer-resolvers)
- [Working with multiple go mods](https://stackoverflow.com/questions/69059470/is-there-a-way-to-set-up-multiple-local-go-modules-so-they-can-run-inside-one-d)

#### versión README.md -> 1.0 (2023-abril)
