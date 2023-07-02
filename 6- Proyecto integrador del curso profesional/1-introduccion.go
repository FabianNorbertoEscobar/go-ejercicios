package main

import "os"
import "fmt"
import "bufio"
import "strings"

func crearUsuario() {
	fmt.Println("Usuario creado exitosamente!")
}

func listarUsuario() {
	fmt.Println("Listado de usuarios!")
}

func actualizarUsuario() {
	fmt.Println("Usuarios actaulizado exitosamente!")
}

func eliminarUsuario() {
	fmt.Println("Usuario eliminado exitosamente!")
}

func main() {

	var reader *bufio.Reader

	var option string

	reader = bufio.NewReader(os.Stdin)

	fmt.Println("A) Crear")
	fmt.Println("B) Listar")
	fmt.Println("C) Actualizar")
	fmt.Println("D) Eliminar")

	fmt.Println("Ingresa una opcion valida: ")
	option, err := reader.ReadString('\n')

	if err != nil {
		panic("No es posible obtner el valor!")
	}

	option = strings.TrimSuffix(option, "\n")

	switch option {
		case "a", "crear":
			crearUsuario()
		case "b", "listar":
			listarUsuario()
		case "c", "actualizar":
			actualizarUsuario()
		case "d", "eliminar":
			eliminarUsuario()	
		default:
			fmt.Println("Opcion no valida!")
	}
}