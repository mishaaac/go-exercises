/*
Exercise 3: Invoicing System Package Encapsulation & Initialism Rules

Requirements:
- Create an 'invoice' package representing an invoicing domain model.
- Define a struct 'Invoice' with exported field 'InvoiceID', unexported field 'subtotal', and exported field 'Tax'.
- Adhere strictly to Go acronym/initialism rules (use 'ID' instead of 'Id').
- Implement constructor function 'NewInvoice' returning a configured pointer to 'Invoice'.
- Implement method 'CalculateTotal()' returning subtotal + tax.
- Validate package visibility and naming conventions without style warnings or compile errors.
*/

package main

import (
	"05-naming-conventions/03-exercise/invoice"
	"fmt"
)

func main() {
	inv := invoice.NewInvoice(1, 100.0, 18.0)

	fmt.Println("=== 🧾 Invoicing System ===")
	fmt.Printf("Invoice ID : %d\n", inv.InvoiceID)
	fmt.Printf("Tax Amount : $%.2f\n", inv.Tax)
	fmt.Printf("Grand Total: $%.2f\n", inv.CalculateTotal())
}
