package main

import (
	"fmt"
	"math"
	"os"
	"os/exec"
	"reflect"
	"strings"
)

func validarNombreProducto(nombre string) error {
	// Eliminar espacios en blanco
	nombre = strings.TrimSpace(nombre)

	// Convertir a valor reflect.Value para analizar caracteres
	valorNombre := reflect.ValueOf(nombre)

	// Recorrer cada caracter
	for i := 0; i < valorNombre.Len(); i++ {
		caracter := valorNombre.Index(i).Interface().(byte)

		// Permitir solo letras, espacios, guiones bajos y puntos
		if !(caracter >= 'a' && caracter <= 'z' ||
			caracter >= 'A' && caracter <= 'Z' ||
			caracter == ' ' || caracter == '_' || caracter == '.') {
			return fmt.Errorf("Error: %c", caracter)

		}
	}

	return nil // Nombre válido
}

func main() {

	var productos []string
	var costos []float64
	var precios []float64
	var utilidades []float64
	var utilidadrequerida float64
	var utilidadtotal float64
	var seen bool
	var soon bool
	var Utilidad float64
	var diferencia float64

	for {

		fmt.Println("Ingrese con que utilidad desea manejar sus productos")
		fmt.Scan(&Utilidad)

		if Utilidad <= 0 || Utilidad >= 100 {
			fmt.Println("La utilidad debe ser mayor a 0 y menor a 100")
			continue

		} else {

			break

		}

	}

	for {

		fmt.Println("Diga el nombre del Producto (Escribe salir para terminar)")
		var Producto string
		fmt.Scan(&Producto)

		err := validarNombreProducto(Producto)
		if err != nil {
			fmt.Println("EL nombre no puede tener numeros", err)
			continue
		} else {
			if Producto == "salir" {
				break
			}

		}

		fmt.Println("Diga el Costo del Producto")
		var Costo float64
		fmt.Scan(&Costo)
		if Costo <= 0 {
			fmt.Println("El costo debe ser un número mayor a 0.")
			continue

		}

		fmt.Println("Diga el Precio del Producto")
		var Precio float64
		fmt.Scan(&Precio)

		if Precio <= 0 {
			fmt.Println("El precio del Producto debe ser un numero mayor a 0")
			continue

		}

		productos = append(productos, Producto)
		costos = append(costos, Costo)
		precios = append(precios, Precio)
		utilidades = append(utilidades, Utilidad)

		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()

	}

	for i := range productos {
		utilidadrequerida = float64(precios[i]-float64(costos[i])) / float64(precios[i])
		utilidadtotal = utilidadrequerida * 100

		if utilidadtotal >= Utilidad {
			if !soon {
				fmt.Println("Lista de Productos con buena utilidad")
				soon = true
			}

			fmt.Println("-", productos[i], "costo", costos[i], "$", "Precio", precios[i], "$", "Utilidad", math.Floor(utilidadtotal), "%")

		}

	}

	for i := range productos {
		utilidadrequerida = float64(precios[i]-float64(costos[i])) / float64(precios[i])
		utilidadtotal = utilidadrequerida * 100
		diferencia = utilidadtotal - Utilidad

		if utilidadtotal < Utilidad {

			if !seen {
				fmt.Println("Lista de Productos con baja utilidad")
				seen = true
			}

			fmt.Println("-", productos[i], "costo", costos[i], "$", "Precio", precios[i], "$", "Utilidad", math.Floor(utilidadtotal), "%", "Diferencia de Utilidad", diferencia, "%")

		}

	}

}
