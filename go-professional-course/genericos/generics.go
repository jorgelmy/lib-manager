package genericos

import "fmt"

func Generics() {
	printList(1, 2, 3, ".", "..", "...", true)
}

func printList(list ...any) {
	type integerMine int

	for _, i := range list {
		fmt.Println(i)
	}

	typosRestrictions(4, 3, 2, 1)
	typosRestrictions(4.5, 3.0, 2.0, 1.1)
	aa := integerMine(5)
	ab := integerMine(10)

	typosRestrictions(aa, ab)

}

func typosRestrictions[T ~int | ~float64](values ...T) {
	var suma T
	for _, i := range values {
		suma += i
	}
	fmt.Println(suma)
}
