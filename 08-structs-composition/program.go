package main

import "fmt"

type Product struct {
	Id       int
	Name     string
	Cost     float64
	Units    int
	Category string
}

/* Association */
/*
type PerishableProduct struct {
	product Product
	Expiry  string
}
*/

/* Composition (inheritance) */
type PerishableProduct struct {
	Product
	Expiry string
}

func main() {
	/* Association */
	/*
		var grapes = PerishableProduct{
			product: Product{102, "Grapes", 20, 50, "Fruits"},
			Expiry:  "2 Days",
		}
		fmt.Printf("%#v\n", grapes)
		fmt.Println(grapes.product.Name)

		fmt.Println(ToString(&grapes.product))
		//applying discount
		ApplyDiscount(&grapes.product, 10)
		fmt.Println("After applying 10% discount")
		fmt.Printf("%#v\n", grapes)
	*/
	/* Composition */
	/*
		var grapes = PerishableProduct{
			Product{102, "Grapes", 20, 50, "Fruits"},
			"2 Days",
		}
	*/
	var grapes = PerishableProduct{
		Product: Product{102, "Grapes", 20, 50, "Fruits"},
		Expiry:  "2 Days",
	}
	fmt.Printf("%#v\n", grapes)
	//fmt.Println(grapes.Product.Name)
	fmt.Println(grapes.Name)

	/* Assinging a new Name */
	grapes.Name = "Arabian Grapes"
	fmt.Println(grapes.Name)

	fmt.Println(ToString(&grapes.Product))

	//applying discount
	ApplyDiscount(&grapes.Product, 10)
	fmt.Println("After applying 10% discount")
	fmt.Printf("%#v\n", grapes)

}

func ToString(product *Product) string {
	return fmt.Sprintf("Id=%d, Name=%q, Cost=%f, Units=%d, Category=%q", product.Id, product.Name, product.Cost, product.Units, product.Category)
}

func ApplyDiscount(product *Product, discount float64) {
	product.Cost = product.Cost * ((100 - discount) / 100)
}
