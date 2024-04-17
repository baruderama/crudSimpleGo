package handlers

import (
	"bufio"
	"crud_go/conectar"
	"crud_go/modelos"
	"fmt"
	"log"
	"os"
	"strconv"
)

func Listar() {
	conectar.Conectar()
	sql := "select id, nombre,correo,telefono from clientes order by id desc; "
	datos, err := conectar.Db.Query(sql)
	if err != nil {
		fmt.Println(err)
	}
	defer conectar.CerrarConexion()

	//lista de clientes que viene desde el modelo
	// clientes := modelos.Clientes{}
	// for datos.Next() {

	// 	//la estrucutrua de un solo cliente
	// 	dato := modelos.Cliente{}

	// 	//se llena a travez de punteros esa estructura
	// 	datos.Scan(&dato.Id, &dato.Nombre, &dato.Telefono, &dato.Correo)

	// 	//se llena esa lista de clientes, con el nuevo cliente llenado
	// 	clientes = append(clientes, dato)
	// }

	// //se imprime lo que esta en la lista de clientes
	// fmt.Println(clientes)
	fmt.Println("Listado de clientes")
	fmt.Println("--------------------------------------------------------------------")
	for datos.Next() {
		var dato modelos.Cliente
		err := datos.Scan(&dato.Id, &dato.Nombre, &dato.Correo, &dato.Telefono)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Id: %v | Nombre: %v | E-mail %s | telefono %s \n", dato.Id, dato.Nombre, dato.Correo, dato.Telefono)
	}

}

func ListarPorId(id int) {
	conectar.Conectar()
	sql := "select id, nombre, correo, telefono from clientes where id=?;"
	datos, err := conectar.Db.Query(sql, id)
	if err != nil {
		fmt.Println(err)
	}
	defer conectar.CerrarConexion()
	fmt.Println("Listado de Cliente por Id")
	fmt.Println("--------------------------------------------------------------------")
	for datos.Next() {
		var dato modelos.Cliente
		err := datos.Scan(&dato.Id, &dato.Nombre, &dato.Telefono, &dato.Correo)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Id: %v | Nombre: %v | E-mail %s | telefono %s \n", dato.Id, dato.Nombre, dato.Correo, dato.Telefono)
	}

}

func Insertar(cliente modelos.Cliente) {
	conectar.Conectar()
	sql := "insert into clientes values(null,?,?,?)"
	result, err := conectar.Db.Exec(sql, cliente.Nombre, cliente.Correo, cliente.Telefono)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
	fmt.Println("Se creo el registro")
}

func Editar(cliente modelos.Cliente, id int) {
	conectar.Conectar()
	sql := "update clientes set nombre=?, correo=?, telefono=? where id=?;"
	result, err := conectar.Db.Exec(sql, cliente.Nombre, cliente.Correo, cliente.Telefono, id)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
	fmt.Println("Se edito el registro")

}

func Eliminar(id int) {
	conectar.Conectar()
	sql := "delete from clientes where id=?;"
	_, err := conectar.Db.Exec(sql, id)
	if err != nil {
		panic(err)
	}
	fmt.Println("Se elimino el registro")
}

// #############################################Funciones de Trabajo#############
var ID int
var nombre, correo, telefono string

func Ejecutar() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Seleccione una opcion : \n\n")
	fmt.Println("1- Listar clientes \n")
	fmt.Println("2- Listar cliente por : \n")
	fmt.Println("3- Crear cliente : \n")
	fmt.Println("4- Editar cliente : \n")
	fmt.Println("5- Eliminar cliente: \n")
	if scanner.Scan() {
		for {
			if scanner.Text() == "1" {
				Listar()
				return
			}
			if scanner.Text() == "2" {
				fmt.Println("Ingrese el ID del cliente : ")
				if scanner.Scan() {
					ID, _ = strconv.Atoi(scanner.Text())
				}
				ListarPorId(ID)
				return
			}
			if scanner.Text() == "3" {
				fmt.Println("ingrese nombre: \n")
				if scanner.Scan() {
					nombre = scanner.Text()
				}
				fmt.Println("ingrese E-email: \n")
				if scanner.Scan() {
					correo = scanner.Text()
				}
				fmt.Println("ingrese Telefono: \n")
				if scanner.Scan() {
					correo = scanner.Text()
				}
				cliente := modelos.Cliente{Nombre: nombre, Correo: correo, Telefono: telefono}
				Insertar(cliente)
				Ejecutar()
				continue

			}
			if scanner.Text() == "4" {
				fmt.Println("Ingrese el ID del cliente : ")
				if scanner.Scan() {
					ID, _ = strconv.Atoi(scanner.Text())
				}
				fmt.Println("ingrese nombre: \n")
				if scanner.Scan() {
					nombre = scanner.Text()
				}
				fmt.Println("ingrese E-email: \n")
				if scanner.Scan() {
					correo = scanner.Text()
				}
				fmt.Println("ingrese Telefono: \n")
				if scanner.Scan() {
					correo = scanner.Text()
				}
				cliente := modelos.Cliente{Nombre: nombre, Correo: correo, Telefono: telefono}
				Editar(cliente, ID)
				continue

			}
			if scanner.Text() == "5" {
				fmt.Println("Ingrese el ID del cliente a eliminar: ")
				if scanner.Scan() {
					ID, _ = strconv.Atoi(scanner.Text())
				}
				Eliminar(ID)
				return
			}

		}
	}

}
