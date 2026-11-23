# Prácticas de Go

Una colección de ejercicios y proyectos para explorar Go: desde la sintaxis básica y las estructuras de datos hasta aplicaciones de consola, un chat con WebSockets y servicios web con MySQL. Los archivos numerados representan etapas de aprendizaje, no una única aplicación.

## Contenido

| Carpeta | Temas y prácticas |
| --- | --- |
| [1. Introducción a Go](1-%20Introducci%C3%B3n%20a%20go/) | Variables, constantes, operadores, cadenas y entrada/salida por consola. |
| [2. Estructuras de datos](2-%20Estructuras%20de%20datos/) | Arreglos, slices, `append`, `make`, mapas e iteración. |
| [3. Controles de flujo](3-%20Controles%20de%20flujo/) | `if`, `switch`, distintas formas de `for`, `range`, `break`, `continue` y `panic`. |
| [4. Funciones](4-%20Funciones/) | Parámetros y retornos, funciones anónimas y variádicas, recursión, alcance, punteros, errores, `defer` y `recover`. |
| [5. Tipos de estructuras](5-%20Tipos%20de%20estructuras/) | Structs, métodos, composición, constantes como enumeraciones, interfaces, paquetes y visibilidad. |
| [6. Proyecto integrador](6-%20Proyecto%20integrador%20del%20curso%20profesional/) | Evolución de un gestor de usuarios por consola. |
| [7. Servidor web](7-%20Servidor%20web/) | Servidor HTTP y chat con WebSockets. |
| [8. Servicio REST](8-%20Servicio%20REST/) | API de usuarios con JSON y MySQL. |
| [9. Proyecto web](9-%20Proyecto%20web/) | Aplicación con plantillas, registro, inicio de sesión y persistencia. |

## Stack y conceptos

- **Go:** biblioteca estándar, entre otros `fmt`, `bufio`, `net/http`, `encoding/json`, `html/template`, `database/sql` y `sync`.
- **Web:** Gorilla Mux para rutas, Gorilla WebSocket para el chat y Martini para el servicio REST.
- **Datos y autenticación:** MySQL, driver `go-sql-driver/mysql`, GORM v1 (`github.com/jinzhu/gorm`) y bcrypt.
- **Interfaz:** HTML, CSS, JavaScript, jQuery y Bootstrap.
- **Pruebas:** paquete `testing` de Go; el proyecto web incluye pruebas de validación y operaciones de usuario.

## Primeros pasos

Instalá [Go](https://go.dev/doc/install), cloná el repositorio y ejecutá un archivo a la vez desde la raíz:

```bash
git clone https://github.com/FabianNorbertoEscobar/go-ejercicios.git
cd go-ejercicios
go run "1- Introducción a go/1-main.go"
go run "2- Estructuras de datos/2-1-slice.go"
```

Los nombres de las carpetas contienen espacios, por lo que conviene usar comillas. Muchos archivos de una misma carpeta declaran su propia función `main`: ejecutarlos juntos como un único paquete provoca declaraciones duplicadas. Algunos ejercicios solicitan datos por teclado; otros muestran deliberadamente errores, `panic` o recuperación de errores.

Las rutas de archivos son relativas al directorio desde el que se ejecuta el programa. Por ejemplo:

```bash
cd "4- Funciones"
go run 15-programar-funciones2.go
```

El ejemplo de `recover` intenta abrir `examples.txt`, mientras que el archivo incluido se llama `example.txt`; así se puede observar la recuperación ante un `panic`.

### Compatibilidad y dependencias

El repositorio no incluye `go.mod`, `go.sum` ni versiones fijadas de dependencias. Algunos ejemplos y proyectos usan importaciones relativas (`./` y `../`), propias del modo GOPATH y no compatibles con el modo de módulos de Go. Para ejecutarlos se necesita un entorno GOPATH con `GO111MODULE=off` y, en los proyectos web, las dependencias externas instaladas en él. Las versiones recientes de Go no resuelven esas dependencias automáticamente con `go get` en modo GOPATH; puede ser necesario preparar un entorno histórico compatible o migrar los imports y las dependencias a módulos.

Los ejemplos `9-creacin-de-paquetes.go` y `10-modificadores-de-acceso.go` de la carpeta 5 llaman métodos que no están exportados o no existen en el paquete incluido, por lo que no compilan tal como están. Los proyectos web tampoco tienen versiones de dependencias verificadas: las indicaciones siguientes describen su configuración y ejecución una vez preparado un entorno compatible, no garantizan compatibilidad con las versiones actuales de las bibliotecas.

## Proyectos

### Gestor de usuarios por consola

La carpeta 6 conserva versiones sucesivas del ejercicio. La última permite crear, listar, actualizar y eliminar usuarios guardados en un mapa en memoria:

```bash
go run "6- Proyecto integrador del curso profesional/7-eliminar-usuario.go"
```

Ingresá `a`, `b`, `c` o `d` (también se aceptan los nombres de las acciones en minúsculas) y `q` o `quit` para salir. Los datos se pierden al cerrar el programa; para limpiar la pantalla se invoca el comando `clear`. Las operaciones de actualización y eliminación muestran un mensaje de éxito incluso si el ID no existe.

### Chat con WebSockets

El servidor de la carpeta 7 sirve una página de registro y chat, valida si un nombre ya está conectado y transmite los mensajes a los usuarios activos. Usa un mapa protegido con `sync.RWMutex`, rutas HTTP y WebSockets. También expone `GET /hola` y `GET /json`.

Con Gorilla Mux y Gorilla WebSocket disponibles en un entorno compatible:

```bash
cd "7- Servidor web"
GO111MODULE=off go run server.go
```

Abrí `http://localhost:8000` en dos ventanas para probar el intercambio. El cliente tiene `localhost:8000` configurado de forma explícita y carga jQuery y Bootstrap desde CDN. Ejecutar desde esa carpeta es necesario para encontrar los archivos de `Front`.

### Servicio REST de usuarios

La carpeta 8 combina Martini, `database/sql` y MySQL. El archivo [`users.sql`](8-%20Servicio%20REST/users.sql) crea la base `test` y la tabla `users`. El código se conecta localmente como `root`, con contraseña vacía; esos valores están definidos en `connect/connect.go`. Prepará una base **exclusiva de práctica**, ajustá la conexión a tu entorno y aplicá el SQL antes de iniciar el servicio. Luego, con las dependencias disponibles en un entorno compatible:

```bash
cd "8- Servicio REST"
GO111MODULE=off go run service.go
```

El servidor escucha en el puerto `8000`. Sus rutas son:

| Método | Ruta | Acción |
| --- | --- | --- |
| GET | `/users/` | Listar usuarios. |
| GET | `/user/:user_id` | Consultar un usuario. |
| POST | `/new/user/` | Crear con JSON `{"user_name":"Ana"}`. |
| PUT | `/update/user/:user_id/:user_name` | Actualizar mediante parámetros de ruta. |
| PUT | `/update/user/` | Actualizar con JSON `{"user_id":"1","user_name":"Ana"}`. |
| DELETE | `/delete/user/:user_id` | Eliminar un usuario. |

Las respuestas incluyen `status`, `data` y `errors` en el cuerpo JSON. El archivo `test.go` es un cliente HTTP de ejemplo, no una prueba automatizada; se ejecuta por separado mientras el servicio está activo.

### Proyecto web con usuarios

La carpeta 9 organiza rutas, controladores, modelos, utilidades y plantillas. Implementa formularios de registro e inicio de sesión, validaciones de campos, hash y verificación de contraseñas con bcrypt y operaciones de usuario mediante GORM. Al iniciar, intenta crear la tabla de usuarios si no existe.

Creá previamente una base MySQL **de práctica** llamada `go_web_project` y revisá la conexión en `models/model.go`: por defecto usa `root` sin contraseña. Con las dependencias disponibles en un entorno compatible:

```bash
cd "9- Proyecto web"
GO111MODULE=off go run main.go
```

El servidor escucha en `http://localhost:3000`, con rutas `/`, `/register` y `/login`. Debe iniciarse desde esa carpeta para localizar las plantillas. La gestión de sesiones y cookies está incompleta: el inicio de sesión verifica la contraseña, pero no establece una sesión persistente. El layout referencia un CSS local, aunque las rutas actuales no sirven ese archivo.

## Pruebas y uso responsable

Las pruebas del proyecto web se pueden invocar desde su carpeta con `GO111MODULE=off go test ./test`, una vez preparado el entorno y la base de datos. **El setup se conecta a `go_web_project` y el teardown elimina la tabla de usuarios**; nunca las ejecutes contra datos que necesites conservar. La suite contiene pruebas que pueden fallar por el estado actual de la implementación.

Estos proyectos son material de aprendizaje, no aplicaciones listas para producción. En particular, el servicio REST construye consultas SQL con valores interpolados y las credenciales de las bases están definidas en el código. Usalos solo en entornos aislados y con datos descartables.
