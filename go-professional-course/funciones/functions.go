package funciones

import "fmt"

func Functions() {
	//Funciones variadicas
	fmt.Println(suma(1, 2, 3, 1))
	printData("Jorge", 27, true)

	//Funciones recursivas
	fmt.Println(factorial(5))

	//Funciones anonimas

	saludo := func(name string) {
		fmt.Printf("Función anonima de %s\n", name)
	}
	saludo("Jorge")
	saludarFuncion("Jorge2", saludo)

	//Orden superior
	numero := ordenSuperior(addOne, 2)
	fmt.Println(numero)
}
func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func suma(nums ...int) int {
	var total int
	for _, num := range nums {
		total += num
	}
	return total
}

func printData(data ...interface{}) {
	for _, d := range data {
		fmt.Printf("%T - %v\n", d, d)
	}
}
func saludarFuncion(name string, f func(string)) {
	f(name)
}

func ordenSuperior(f func(int) int, x int) int {
	return f(x * 2)
}
func addOne(x int) int {
	return x + 1
}
