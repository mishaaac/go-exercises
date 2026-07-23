package invoice

// Invoice represents a billing document.
// Field visibility rules applied:
// - InvoiceID: Exported field adhering to initialism rule (ID).
// - subtotal: Unexported (private) field using camelCase.
// - Tax: Exported field.
type Invoice struct {
	InvoiceID int
	subtotal  float64
	Tax       float64
}

// NewInvoice constructs and returns a pointer to a newly initialized Invoice struct.
func NewInvoice(invoiceID int, subtotal, tax float64) *Invoice {
	return &Invoice{
		InvoiceID: invoiceID,
		subtotal:  subtotal,
		Tax:       tax,
	}
}

// CalculateTotal computes the total invoice amount (subtotal + tax).
func (i *Invoice) CalculateTotal() float64 {
	return i.subtotal + i.Tax
}
