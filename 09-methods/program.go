package main

import "fmt"

type Product struct {
	Id       int
	Name     string
	Cost     float64
	Units    int
	Category string
}

type PerishableProduct struct {
	Product
	Expiry string
}

func main() {
	//product := Product{Id: 100, Name: "Pen", Cost: 10, Units: 100, Category: "Stationary"}

	/*
		fmt.Print(ToString(&product))
		ApplyDiscount(&product, 10)
		fmt.Println("After applying discount(10%)")
		fmt.Print(ToString(&product))
	*/

	/*
		fmt.Println((&product).ToString())
		(&product).ApplyDiscount(10)
		fmt.Println("After applying discount(10%)")
		fmt.Println((&product).ToString())
	*/
	product := &Product{Id: 100, Name: "Pen", Cost: 10, Units: 100, Category: "Stationary"}
	fmt.Println(product.ToString())
	product.ApplyDiscount(10)
	fmt.Println("After applying discount(10%)")
	fmt.Println(product.ToString())

	//Composition
	fmt.Println()
	fmt.Println("Composition")
	var grapes = PerishableProduct{
		Product: Product{102, "Grapes", 20, 50, "Fruits"},
		Expiry:  "2 Days",
	}
	fmt.Println(grapes.ToString())
	fmt.Println("After applying discount(10%)")
	grapes.ApplyDiscount(10)
	fmt.Println(grapes.ToString())

}

func (product *Product) ToString() string {
	return fmt.Sprintf("Id=%d, Name=%q, Cost=%f, Units=%d, Category=%q", product.Id, product.Name, product.Cost, product.Units, product.Category)
}

func (product *Product) ApplyDiscount(discount float64) {
	product.Cost = product.Cost * ((100 - discount) / 100)
}

/* method overriding */
func (pp *PerishableProduct) ToString() string {
	return fmt.Sprintf("%s, Expiry=%s", pp.Product.ToString(), pp.Expiry)
}
