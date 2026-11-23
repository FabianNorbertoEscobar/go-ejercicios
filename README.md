# Prácticas de Go

Una colección de ejercicios y proyectos para explorar Go, desde la sintaxis básica hasta aplicaciones de consola, servidores HTTP y acceso a bases de datos. Las carpetas siguen una progresión temática: cada archivo de los primeros capítulos suele ser un programa independiente y los últimos capítulos reúnen proyectos más grandes.

> Es material de aprendizaje, no una aplicación lista para producción. Algunos ejemplos muestran deliberadamente errores o `panic`; los proyectos web también conservan dependencias y prácticas antiguas.

## Contenido

| Carpeta | Qué se practica |
| --- | --- |
| [1. Introducción a Go](<1- Introducción a go>) | Variables, constantes, tipos, operadores, cadenas, salida y lectura desde consola. |
| [2. Estructuras de datos](<2- Estructuras de datos>) | Arreglos, slices, `make`, mapas e iteración. |
| [3. Controles de flujo](<3- Controles de flujo>) | Condicionales, `switch`, variantes de `for`, `range`, `break`, `continue` y `panic`. |
| [4. Funciones](<4- Funciones>) | Argumentos y retornos, funciones anónimas y como valores, recursión, funciones variádicas, alcance, punteros, errores, `defer` y `recover`. |
| [5. Tipos de estructuras](<5- Tipos de estructuras>) | Structs, métodos, composición, relaciones, constantes como enums, interfaces, paquetes y visibilidad. |
| [6. Proyecto integrador](<6- Proyecto integrador del curso profesional>) | Evolución paso a paso de un gestor de usuarios por consola. |
| [7. Servidor web](<7- Servidor web>) | HTTP, JSON, archivos estáticos y un chat con WebSockets. |
| [8. Servicio REST](<8- Servicio REST>) | API de usuarios con JSON, consultas SQL y MySQL. |
| [9. Proyecto web](<9- Proyecto web>) | Formularios de registro e inicio de sesión, plantillas, validaciones y persistencia con GORM. |

## Stack y conceptos

- **Go** y su biblioteca estándar: `fmt`, `net/http`, `encoding/json`, `html/template`, `database/sql`, `testing`, entre otros paquetes.
- **Gorilla Mux y Gorilla WebSocket** para enrutamiento y comunicación bidireccional; el chat usa `sync.RWMutex` para proteger el mapa de usuarios conectados.
- **Martini**, **MySQL** y el driver `go-sql-driver/mysql` en el servicio REST; **GORM v1** (`github.com/jinzhu/gorm`) en el proyecto web.
- **HTML, CSS, JavaScript, jQuery y Bootstrap** para las interfaces. El chat carga jQuery y Bootstrap desde CDN; el proyecto web incluye un archivo CSS de Bootstrap.
- **bcrypt** para generar y verificar hashes de contraseñas en el proyecto web.

## Primeros pasos

Instalá [Go](https://go.dev/doc/install), cloná el repositorio y ejecutá un archivo a la vez desde la raíz:

```bash
git clone https://github.com/FabianNorbertoEscobar/go-ejercicios.git
cd go-ejercicios
go run "1- Introducción a go/1-main.go"
go run "2- Estructuras de datos/5-map.go"
```

Los directorios de ejercicios contienen múltiples funciones `main`, por lo que no deben ejecutarse como un único paquete. Las comillas son necesarias por los espacios en las rutas. Algunos programas piden datos por teclado; otros provocan un `panic` para ilustrar el manejo de errores.

Los ejemplos que usan archivos relativos deben ejecutarse desde su propia carpeta:

```bash
cd "4- Funciones"
go run 15-programar-funciones2.go
```

Ese ejemplo lee `example.txt`. En cambio, `16-funcion-recover.go` intenta abrir `examples.txt`: si no existe, muestra cómo recuperar el `panic`.

### Go moderno y dependencias antiguas

El repositorio **no tiene `go.mod` ni versiones de dependencias fijadas**. Los proyectos usan imports relativos como `./models`, que no son compatibles con el modo de módulos. Los ejercicios sin dependencias externas pueden ejecutarse como en los ejemplos anteriores; para los imports relativos se necesita el modo GOPATH (`GO111MODULE=off`).

Los proyectos web además necesitan las dependencias externas enumeradas en cada sección. El flujo original corresponde a versiones antiguas de Go: con Go 1.16 o anterior, `GO111MODULE=off go get <paquete>` permite instalarlas en GOPATH. En Go moderno `go get` ya no instala dependencias en ese modo. Como no hay versiones fijadas, la compatibilidad con las versiones actuales de las bibliotecas no está garantizada; para reproducir los proyectos puede ser necesario preparar un entorno legado o migrar sus imports y dependencias a módulos.

En particular, los ejemplos `9-creacin-de-paquetes.go` y `10-modificadores-de-acceso.go` no compilan tal como están: los métodos que invocan no coinciden con los nombres disponibles en `CodigoFacilito`. Son útiles para explorar visibilidad y nombres exportados, pero requieren corregir esa diferencia para ejecutarlos.

## Proyectos

### Gestor de usuarios por consola

La última etapa integra creación, listado, actualización y eliminación de usuarios almacenados **en memoria**:

```bash
go run "6- Proyecto integrador del curso profesional/7-eliminar-usuario.go"
```

Ingresá `a`, `b`, `c` o `d` para elegir una operación, y `q` o `quit` para salir. Los datos se pierden al cerrar el programa. El menú usa el comando externo `clear` para limpiar la consola y espera una edad numérica.

### Servidor web y chat

El servidor usa `github.com/gorilla/mux` y `github.com/gorilla/websocket`. Con las dependencias disponibles en GOPATH:

```bash
cd "7- Servidor web"
GO111MODULE=off go run server.go
```

Abrí `http://localhost:8000` para elegir un nombre y enviar mensajes a los usuarios conectados. También expone `GET /hola`, `GET /json`, `POST /validate` y el WebSocket `GET /chat/{user_name}`. Los usuarios activos se mantienen en memoria y los mensajes se distribuyen a las conexiones activas. La interfaz apunta explícitamente a `localhost:8000` y carga dependencias de interfaz desde Internet.

### Servicio REST de usuarios

El servicio utiliza `github.com/go-martini/martini` y `github.com/go-sql-driver/mysql`. Necesita MySQL local, la base y tabla definidas en [`users.sql`](<8- Servicio REST/users.sql>) y dependencias disponibles en GOPATH. El código usa la base `test`, el usuario `root` y una contraseña vacía; la configuración está en `connect/connect.go`.

```bash
cd "8- Servicio REST"
mysql -u root < users.sql
GO111MODULE=off go run service.go
```

El script crea la base `test` y la tabla `users`: usalo únicamente en un entorno de práctica aislado. El servicio escucha en `http://localhost:8000`, el mismo puerto que el chat, por lo que no pueden utilizarlo simultáneamente.

| Método | Ruta | Operación |
| --- | --- | --- |
| `GET` | `/users/` | Listar usuarios. |
| `GET` | `/user/:user_id` | Consultar un usuario por ID. |
| `POST` | `/new/user/` | Crear un usuario a partir de JSON. |
| `PUT` | `/update/user/:user_id/:user_name` | Actualizar el nombre mediante parámetros de ruta. |
| `PUT` | `/update/user/` | Actualizar mediante JSON con `user_id` y `user_name`. |
| `DELETE` | `/delete/user/:user_id` | Eliminar por ID. |

Por ejemplo:

```bash
curl -X POST http://localhost:8000/new/user/ \
  -H "Content-Type: application/json" \
  -d '{"user_name":"Ada"}'
curl http://localhost:8000/users/
```

Las respuestas incluyen `status`, `data` y `errors`. El archivo `test.go` es un cliente HTTP de ejemplo para crear un usuario, no una suite de pruebas automatizadas.

### Proyecto web con formularios

Usa `github.com/gorilla/mux`, `github.com/jinzhu/gorm` y su dialecto MySQL, además de `golang.org/x/crypto/bcrypt`. Requiere MySQL local y dependencias disponibles en GOPATH. La conexión configurada en `models/model.go` apunta a `go_web_project`, con usuario `root` y contraseña vacía. Creá una base de práctica vacía antes de iniciar; GORM crea la tabla de usuarios si no existe.

```bash
mysql -u root -e "CREATE DATABASE go_web_project;"
cd "9- Proyecto web"
GO111MODULE=off go run main.go
```

El servidor escucha en `http://localhost:3000`. Las rutas principales son `/`, `GET`/`POST /register` y `GET`/`POST /login`. El registro valida campos obligatorios y longitud de contraseña, genera un hash bcrypt y guarda el usuario; el login consulta el usuario y verifica la contraseña.

La sesión está **incompleta**: el identificador es fijo y la creación de la cookie está comentada. El CSS está incluido en el repositorio, pero el router no sirve actualmente la ruta de assets indicada en la plantilla.

### Pruebas del modelo

En [`9- Proyecto web/test`](<9- Proyecto web/test>) hay pruebas de validación, contraseñas y operaciones sobre usuarios. Requieren las dependencias y la base MySQL configuradas para el proyecto web:

```bash
cd "9- Proyecto web"
GO111MODULE=off go test ./test
```

**Atención:** la preparación y limpieza de estas pruebas usan la base configurada en `models/model.go` y eliminan la tabla de usuarios al finalizar. Ejecutalas solo contra una base descartable. El conjunto refleja el estado de los ejercicios y puede contener pruebas fallidas; no presupone una suite estable.

## Alcance y seguridad

Los proyectos priorizan la experimentación sobre el endurecimiento: hay credenciales locales de ejemplo, consultas SQL construidas por concatenación, gestión incompleta de errores y sesiones sin implementar. No los expongas a Internet ni los uses con datos reales sin revisar y corregir estos aspectos. No hay un archivo de licencia incluido en el repositorio.
