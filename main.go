package main

import (
	"fmt"
	"os"
	"os/exec"
	"bufio"
	"strings"
	"strconv"
)

type User struct {
	id int
	username string
	email string
	age int
}

var id int
var users map[int]User

func crearUsuario(reader *bufio.Reader) {
	fmt.Print("Ingrese su usuario: ")
	username := leerLinea(reader)

	fmt.Print("Ingrese su email: ")
	email := leerLinea(reader)

	fmt.Print("Ingrese su edad: ")
	age, err := strconv.Atoi(leerLinea(reader))
	
	if err != nil {
		panic("No es posible convertir de un \"string\" a un \"int\".")
	}

	id++
	user := User { id, username, email, age }
	users[id] = user 

	fmt.Println("Usuario creado exitosamente!")
}

func listarUsuarios() {
	for id, user := range users {
		fmt.Println(id, "-", user.username)
	}
}

func actualizarUsuario(reader *bufio.Reader) {
	fmt.Print("Ingrese el \"id\" del usuario a actualizar: ")
	id, err := strconv.Atoi(leerLinea(reader))

	if err != nil {
		panic("No es posible convertir de un \"string\" a un \"int\".")	
	}

	if _, ok := users[id]; ok {
		fmt.Print("Ingrese un nuevo valor para usuario: ")
		username := leerLinea(reader)

		fmt.Print("Ingrese un nuevo valor para email: ")
		email := leerLinea(reader)

		fmt.Print("Ingrese un nuevo valor para edad: ")
		age, err := strconv.Atoi(leerLinea(reader))

		if err != nil {
			panic("No es posible convertir de un \"string\" a un \"int\".")
		}

		users[id] = User { id, username, email, age }
	}

	fmt.Println("Usuario actualizado exitosamente!")
}

func eliminarUsuario(reader *bufio.Reader) {	
	fmt.Print("Ingrese el \"id\" del usuario a eliminar: ")
	id, err := strconv.Atoi(leerLinea(reader))

	if err != nil {
		panic("No es posible convertir de un \"string\" a un \"int\".")	
	}

	if _, ok := users[id]; ok {
		delete(users, id)
	}

	fmt.Println("Usuario eliminado exitosamente!")
}

func leerLinea(reader *bufio.Reader) string {
	if option, err := reader.ReadString('\n'); err != nil {
		panic("No es posible obtener el valor!")
	} else {
		return strings.TrimSuffix(option, "\n")
	}
}

func limpiarConsola() {
	ter := exec.Command("clear")
	ter.Stdout = os.Stdout
	ter.Run() 	
}

func main() {
	limpiarConsola()
	
	var reader *bufio.Reader
	var option string

	users = make(map[int]User)

	reader = bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Crear")
		fmt.Println("Listar")
		fmt.Println("Actualizar")
		fmt.Println("Eliminar")
		fmt.Println("Salir\n")

		fmt.Print("Ingrese una opcion: ")
		option = leerLinea(reader)
		if option == "Salir" || option == "salir" {
			break
		}

		switch option {
			case "Crear", "crear":
				crearUsuario(reader)
			case "Listar", "listar":
				listarUsuarios()
			case "Actualizar", "actualizar":
				actualizarUsuario(reader)
			case "Eliminar", "eliminar":
				eliminarUsuario(reader)
			default:
				fmt.Println("Opcion no valida!")
		}	

		var v int8
		fmt.Print("\nEnter para continuar...")
		fmt.Scanf("%i", &v)
		limpiarConsola()
	}

	fmt.Println("Adios.")
}