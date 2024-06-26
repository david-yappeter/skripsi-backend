package printer_template

import (
	"fmt"
	"strings"

	"github.com/david-yappeter/escpos/generate"
)

func EscposOpenCashDrawer() []byte {
	return generate.Cash()
}

type EscposReceiptItem struct {
	Name        string
	PricePerQty string
	Qty         string
	TotalPrice  string
}

func (i EscposReceiptItem) wrapLength() int {
	return 26
}

func (i EscposReceiptItem) WrapNames() []string {
	s := i.Name

	var result []string
	for idx := 0; idx < len(s); idx += i.wrapLength() {
		end := idx + i.wrapLength()
		if end > len(s) {
			end = len(s)
		}
		result = append(result, s[idx:end])
	}
	return result
}

type EscposReceiptTemplateAttribute struct {
	StoreName   string
	Address     string
	PhoneNumber string
	Date        string
	Cashier     string
	Items       []EscposReceiptItem

	SubTotal       *string
	DiscountAmount *string
	GrandTotal     string
	Paid           string
	Change         string
	OpenDrawer     bool
	IsCash         bool
}

func EscposReceiptTemplate(attribute EscposReceiptTemplateAttribute) []uint8 {
	var (
		template = []byte{}
	)

	setGap := func(height int) {
		// Page Mode
		template = append(template, generate.SetPageMode()...)

		// Printing Area
		template = append(template, generate.SetPrintArea(0, 0, 100, height)...)

		// Print Page Mode
		template = append(template, generate.PrintPageModeBufferData()...)

		// Set to Standard
		template = append(template, generate.SetStandardMode()...)
	}
	_ = setGap

	writeHorizontalLine := func() {
		template = append(template, []byte{27, 76, 27, 87, 0, 0, 0, 0, 30, 2, 44, 0, 27, 42, 0, 0, 2, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 27, 12, 27, 83}...)
	}

	writeHorizontalDashLine := func() {
		template = append(template, []byte{27, 76, 27, 87, 0, 0, 0, 0, 30, 2, 44, 0, 27, 42, 0, 0, 2, 64, 64, 64, 64, 64, 0, 0, 0, 0, 64, 64, 64, 64, 64, 64, 0, 0, 0, 0, 64, 64, 64, 64, 64, 64, 0, 0, 0, 0, 64, 64, 64, 64, 64, 64, 0, 0, 0, 0, 64, 64, 64, 64, 64, 64, 0, 0, 0, 0, 64, 64, 64, 64, 64, 64, 0, 0, 0, 0, 64, 64, 64, 64, 64, 64, 0, 0, 0, 0, 64, 64, 64, 64, 64, 64, 0, 0, 0, 0, 64, 64, 64, 64, 64, 64, 0, 0, 0, 0, 64, 64, 64, 64, 64, 64, 0, 0, 0, 0, 64, 64, 64, 64, 64, 64, 0, 0, 0, 0, 64, 64, 64, 64, 64, 64, 0, 0, 0, 0, 64, 64, 64, 64, 64, 64, 0, 0, 0, 0, 64, 64, 64, 64, 64, 64, 0, 0, 0, 0, 64, 64, 64, 64, 64, 64, 0, 0, 0, 0, 64, 64, 64, 64, 64, 64, 0, 0, 0, 0, 64, 64, 64, 64, 64, 64, 0, 0, 0, 0, 64, 64, 64, 64, 64, 64, 0, 0, 0, 0, 64, 64, 64, 64, 64, 64, 0, 0, 0, 0, 64, 64, 64, 64, 64, 64, 0, 0, 0, 0, 64, 64, 64, 64, 64, 64, 0, 0, 0, 0, 64, 64, 64, 64, 64, 64, 0, 0, 0, 0, 64, 64, 64, 64, 64, 64, 0, 0, 0, 0, 64, 64, 64, 64, 64, 64, 0, 0, 0, 0, 64, 64, 64, 64, 64, 64, 0, 0, 0, 0, 64, 64, 64, 64, 64, 64, 0, 0, 0, 0, 64, 64, 64, 64, 64, 64, 0, 0, 0, 0, 64, 64, 64, 64, 64, 64, 0, 0, 0, 0, 64, 64, 64, 64, 64, 64, 0, 0, 0, 0, 64, 64, 64, 64, 64, 64, 0, 0, 0, 0, 64, 64, 64, 64, 64, 64, 0, 0, 0, 0, 64, 64, 64, 64, 64, 64, 0, 0, 0, 0, 64, 64, 64, 64, 64, 64, 0, 0, 0, 0, 64, 64, 64, 64, 64, 64, 0, 0, 0, 0, 64, 64, 64, 64, 64, 64, 0, 0, 0, 0, 64, 64, 64, 64, 64, 64, 0, 0, 0, 0, 64, 64, 64, 64, 64, 64, 0, 0, 0, 0, 64, 64, 64, 64, 64, 64, 0, 0, 0, 0, 64, 64, 64, 64, 64, 64, 0, 0, 0, 0, 64, 64, 64, 64, 64, 64, 0, 0, 0, 0, 64, 64, 64, 64, 64, 64, 0, 0, 0, 0, 64, 64, 64, 64, 64, 64, 0, 0, 0, 0, 64, 64, 64, 64, 64, 64, 0, 0, 0, 0, 64, 64, 64, 64, 64, 64, 0, 0, 0, 0, 64, 64, 64, 64, 64, 64, 0, 0, 0, 0, 64, 64, 64, 64, 64, 64, 0, 0, 0, 0, 64, 64, 64, 64, 64, 64, 0, 0, 0, 0, 64, 64, 64, 64, 64, 64, 0, 0, 0, 0, 64, 64, 64, 64, 64, 64, 0, 0, 0, 0, 64, 64, 64, 64, 64, 64, 0, 0, 0, 0, 64, 64, 64, 64, 64, 64, 0, 0, 0, 0, 64, 64, 64, 27, 12, 27, 83}...)
	}
	_ = writeHorizontalDashLine

	// Init
	template = append(template, generate.Init()...)

	if attribute.OpenDrawer {
		// Cash Drawer
		template = append(template, generate.Cash()...)
	}

	// Set to Standard Mode
	template = append(template, generate.SetStandardMode()...)

	// Set print center
	template = append(template, generate.SetAlign("center")...)

	// Store Name
	template = append(template, []byte(attribute.StoreName)...)

	// Line Feed
	template = append(template, generate.Linefeed()...)

	// Address
	template = append(template, []byte(attribute.Address)...)

	// Line Feed
	template = append(template, generate.Linefeed()...)

	// Phone Number
	template = append(template, []byte(attribute.PhoneNumber)...)

	// Line Feed
	template = append(template, generate.Linefeed()...)

	// Horizontal Line
	writeHorizontalLine()

	// Align Left
	template = append(template, generate.SetAlign("left")...)

	// Line Feed
	template = append(template, generate.Linefeed()...)

	// Date Label
	template = append(template, []byte("Tanggal    : ")...)

	// Date
	template = append(template, []byte(attribute.Date)...)

	// Line Feed
	template = append(template, generate.Linefeed()...)

	// Cashier Label
	template = append(template, []byte("Kasir      : ")...)

	// Counter
	template = append(template, []byte(attribute.Cashier)...)

	// Line Feed
	template = append(template, generate.Linefeed()...)

	// Horizontal Line
	writeHorizontalLine()

	// Items List
	for _, item := range attribute.Items {
		itemWrapNames := item.WrapNames()

		// Set Bold
		template = append(template, generate.SetEmphasize(1)...)

		// First Row of each item
		template = append(template, []byte(fmt.Sprintf("%s%s", fmt.Sprintf("%-26s", strings.Trim(itemWrapNames[0], " ")), fmt.Sprintf("%19s", item.TotalPrice)))...)

		// Line Feed
		template = append(template, generate.Linefeed()...)

		// Next N row of item name wrapping
		if len(itemWrapNames) > 1 {
			for _, v := range itemWrapNames[1:] {
				// Item Wrapped Name
				template = append(template, []byte(v)...)

				// Line Feed
				template = append(template, generate.Linefeed()...)
			}
		}

		// Set Unbold
		template = append(template, generate.SetEmphasize(0)...)

		// Price x Qty
		template = append(template, []byte(fmt.Sprintf("%s x %s", item.Qty, item.PricePerQty))...)

		// Line Feed
		template = append(template, generate.Linefeed()...)
	}

	// Horizontal Line
	writeHorizontalLine()

	if attribute.SubTotal != nil {
		// Sub Total
		template = append(template, []byte(fmt.Sprintf("%-25sIDR%17s", "Sub Total", *attribute.SubTotal))...)

		// Line Feed
		template = append(template, generate.Linefeed()...)
	}

	if attribute.DiscountAmount != nil {
		// Discount
		template = append(template, []byte(fmt.Sprintf("%-25sIDR%17s", "Discount", *attribute.DiscountAmount))...)

		// Line Feed
		template = append(template, generate.Linefeed()...)
	}

	// Set Bold
	template = append(template, generate.SetEmphasize(1)...)

	// Grand Total
	template = append(template, []byte(fmt.Sprintf("%-25sIDR%17s", "Grand Total", attribute.GrandTotal))...)

	// Line Feed
	template = append(template, generate.Linefeed()...)

	// Set Unbold
	template = append(template, generate.SetEmphasize(0)...)

	if attribute.IsCash {
		// Paid
		template = append(template, []byte(fmt.Sprintf("%-25sIDR%17s", "Tunai", attribute.Paid))...)

		// Line Feed
		template = append(template, generate.Linefeed()...)

		// Paid
		template = append(template, []byte(fmt.Sprintf("%-25sIDR%17s", "Kembali", attribute.Change))...)
	} else {
		// Paid
		template = append(template, []byte(fmt.Sprintf("%-25sIDR%17s", "Non Tunai", attribute.Paid))...)
	}

	// Line Feed
	template = append(template, generate.Linefeed()...)

	// Line Feed
	template = append(template, generate.Linefeed()...)

	// cut
	template = append(template, generate.Cut()...)

	return []uint8(template)
}
