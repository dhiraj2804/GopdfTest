package main

import (
	"fmt"
	"log"
	"math"
	globals "practice/gopdf/GopdfTest/global"
	"strings"
	"time"

	"github.com/signintech/gopdf"
)

type ReadyStockOrderDetailsForInvoice struct {
	OrderID   string `json:"order_id"`
	Brand     string `json:"brand"`
	OuID      int    `json:"ou"`
	OrderDate string `json:"order_date"`
	InvoiceID string `json:"invoice_id"`

	// Customer Details
	CustomerID        string `json:"customer_id"`
	CustomerName      string `json:"customer_name"`
	CustomerGSTNumber string `json:"customer_gst_number"`

	Status          string               `json:"status"`
	TotalCost       float64              `json:"total_cost"`
	DeliveryCost    float64              `json:"delivery_cost"`
	GstNumber       string               `json:"gst_number"`
	IsPreOrder      bool                 `json:"is_preorder"`
	Discount        float64              `json:"discount"`
	DiscountPercent int                  `json:"discount_percent"`
	ItemsSubtotal   float64              `json:"items_subtotal"`
	StoreName       string               `json:"store_name"`
	Address         string               `json:"address"`
	Phone           string               `json:"phone"`
	Products        []ProductsForInvoice `json:"products"`

	//Salesman Details
	SalesmanName        string `json:"salesman_name"`
	SalesmanPhoneNumber string `json:"salesman_phone"`
}

type ProductsForInvoice struct {
	CGST               float64 `json:"cgst"`
	Count              int     `json:"count"`
	CurrentPrice       float64 `json:"current_price"`
	HsnCode            string  `json:"hsn_code"`
	Images             string  `json:"images"`
	Mrp                float64 `json:"mrp"`
	Name               string  `json:"name"`
	Price              float64 `json:"price"`
	ProductDescription string  `json:"product_description"`
	SGST               float64 `json:"sgst"`
	Sku                string  `json:"sku"`
}

// func main() {

// 	pdf := gopdf.GoPdf{}
// 	mm6ToPx := 22.68

// 	// Base trim-box
// 	pdf.Start(gopdf.Config{
// 		PageSize: *gopdf.PageSizeA4, //595.28, 841.89 = A4
// 		TrimBox:  gopdf.Box{Left: mm6ToPx, Top: mm6ToPx, Right: 595 - mm6ToPx, Bottom: 842 - mm6ToPx},
// 	})

// 	// Page trim-box
// 	opt := gopdf.PageOption{
// 		PageSize: gopdf.PageSizeA4, //595.28, 841.89 = A4
// 		TrimBox:  &gopdf.Box{Left: mm6ToPx, Top: mm6ToPx, Right: 595 - mm6ToPx, Bottom: 842 - mm6ToPx},
// 	}
// 	pdf.AddPageWithOption(opt)
// 	err := pdf.AddTTFFont("fira-font", "./view/my_font.ttf")
// 	if err != nil {
// 		log.Print(err.Error())
// 		return
// 	}
// 	// if err := pdf.AddTTFFont("wts11", "../ttf/wts11.ttf"); err != nil {
// 	//     log.Print(err.Error())
// 	//     return
// 	// }

// 	if err := pdf.SetFont("fira-font", "", 14); err != nil {
// 		log.Print(err.Error())
// 		return
// 	}

// 	pdf.Cell(nil, "Hi dhirajk uam fhfdsjlajkf dfhadsfj;ajfalkfj sdkdlhfdalkfjljsf dfskfalkjdlfkja flkhs")
// 	pdf.WritePdf("hello.pdf")
// }

func main() {
	res := ReadyStockOrderDetailsForInvoice{
		OrderID:           "LYE0000133A00049",
		Brand:             "",
		OuID:              11,
		OrderDate:         "2020-12-02 12:43:18",
		InvoiceID:         "11T/10031/2021",
		CustomerID:        "",
		CustomerName:      "",
		CustomerGSTNumber: "",
		Status:            "confirmed",
		TotalCost:         1010,
		DeliveryCost:      0,
		GstNumber:         "",
		IsPreOrder:        false,
		Discount:          101,
		DiscountPercent:   0,
		ItemsSubtotal:     0,
		StoreName:         "Ramco",
		Address:           "43,43/4, Basant Ave Rd, Padmanabha Nagar, Adyar, Chennai, Tamil Nadu 600020, India",
		Phone:             "4561231251",
		Products: []ProductsForInvoice{
			{
				CGST:               2.5,
				Count:              1,
				CurrentPrice:       10,
				HsnCode:            "",
				Images:             "",
				Mrp:                10,
				Name:               "Hide & Seek",
				Price:              10,
				ProductDescription: "",
				SGST:               2.5,
				Sku:                "SITTEST01",
			},
			{
				CGST:               9,
				Count:              2,
				CurrentPrice:       500,
				HsnCode:            "",
				Images:             "",
				Mrp:                500,
				Name:               "MRF Bat",
				Price:              1000,
				ProductDescription: "MRF Bat",
				SGST:               9,
				Sku:                "BAT01",
			},
			{
				CGST:               9,
				Count:              2,
				CurrentPrice:       500,
				HsnCode:            "",
				Images:             "",
				Mrp:                500,
				Name:               "MRF Bat",
				Price:              1000,
				ProductDescription: "MRF Bat",
				SGST:               9,
				Sku:                "BAT01",
			},
			{
				CGST:               9,
				Count:              2,
				CurrentPrice:       500,
				HsnCode:            "",
				Images:             "",
				Mrp:                500,
				Name:               "MRF Bat",
				Price:              1000,
				ProductDescription: "MRF Bat",
				SGST:               9,
				Sku:                "BAT01",
			}, {
				CGST:               9,
				Count:              2,
				CurrentPrice:       500,
				HsnCode:            "",
				Images:             "",
				Mrp:                500,
				Name:               "MRF Bat",
				Price:              1000,
				ProductDescription: "MRF Bat",
				SGST:               9,
				Sku:                "BAT01",
			}, {
				CGST:               9,
				Count:              2,
				CurrentPrice:       500,
				HsnCode:            "",
				Images:             "",
				Mrp:                500,
				Name:               "MRF Bat",
				Price:              1000,
				ProductDescription: "MRF Bat",
				SGST:               9,
				Sku:                "BAT01",
			}, {
				CGST:               9,
				Count:              2,
				CurrentPrice:       500,
				HsnCode:            "",
				Images:             "",
				Mrp:                500,
				Name:               "MRF Bat",
				Price:              1000,
				ProductDescription: "MRF Bat",
				SGST:               9,
				Sku:                "BAT01",
			}, {
				CGST:               9,
				Count:              2,
				CurrentPrice:       500,
				HsnCode:            "",
				Images:             "",
				Mrp:                500,
				Name:               "MRF Bat",
				Price:              1000,
				ProductDescription: "MRF Bat",
				SGST:               9,
				Sku:                "BAT01",
			}, {
				CGST:               9,
				Count:              2,
				CurrentPrice:       500,
				HsnCode:            "",
				Images:             "",
				Mrp:                500,
				Name:               "MRF Bat",
				Price:              1000,
				ProductDescription: "MRF Bat",
				SGST:               9,
				Sku:                "BAT01",
			}, {
				CGST:               9,
				Count:              2,
				CurrentPrice:       500,
				HsnCode:            "",
				Images:             "",
				Mrp:                500,
				Name:               "MRF Bat",
				Price:              1000,
				ProductDescription: "MRF Bat",
				SGST:               9,
				Sku:                "BAT01",
			}, {
				CGST:               9,
				Count:              2,
				CurrentPrice:       500,
				HsnCode:            "",
				Images:             "",
				Mrp:                500,
				Name:               "MRF Bat",
				Price:              1000,
				ProductDescription: "MRF Bat",
				SGST:               9,
				Sku:                "BAT01",
			}, {
				CGST:               9,
				Count:              2,
				CurrentPrice:       500,
				HsnCode:            "",
				Images:             "",
				Mrp:                500,
				Name:               "MRF Bat",
				Price:              1000,
				ProductDescription: "MRF Bat",
				SGST:               9,
				Sku:                "BAT01",
			},
		},

		SalesmanName:        "",
		SalesmanPhoneNumber: "",
	}

	//fmt.Printf("%+v", res)
	// variables
	var (
		leftMargin float64 = 10
		//rightmargin float64 = 10
		uppermargin float64 = 10
		curY        float64 = 20
		curX        float64 = leftMargin
		tableGap    float64 = 15
		pageWindth  float64 = 595 //W: 595, H: 842
		invoiceX    float64 = curX + 200
		customerX   float64 = invoiceX + 170
		//pageHeight  float64 = 1008
	)

	loc, _ := time.LoadLocation("Asia/Kolkata")
	now := time.Now().In(loc)
	// sample pdf
	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: *gopdf.PageSizeA4}) //W: 720, H: 1008,
	pdf.AddPage()
	err := pdf.AddTTFFont("roboto_normal", "./view/my_font_normal.ttf")
	if err != nil {
		log.Print(err.Error())
		return
	}

	err = pdf.AddTTFFont("roboto_bold", "./view/my_font_bold.ttf")
	if err != nil {
		log.Print(err.Error())
		return
	}

	// itsstring := "dhiraj,kumar,mathura,tudu allu,ajush"
	// s := strings.Split(itsstring, ",")
	// for _, val := range s {
	// 	fmt.Println(val)
	// }

	err = pdf.SetFont("roboto_bold", "", 12)
	if err != nil {
		log.Print(err.Error())
		return
	}

	// main heading
	taxInvoiceX := 250.0
	pdf.SetX(taxInvoiceX) //move current location
	pdf.SetY(uppermargin)
	pdf.Cell(nil, "Tax Invoice")

	curY += 20
	invoiceY := curY
	customerY := curY

	// create column cell

	err = pdf.SetFont("roboto_bold", "", 10)
	if err != nil {
		log.Print(err.Error())
		return
	}

	pdf.Image("./view/logo.png", leftMargin, 30, nil)
	// logo column
	pdf.SetX(curX) //move current location
	pdf.SetY(curY)

	pdf.Cell(nil, "LYNKS LOGISTICS LIMITED")
	curY += 15

	err = pdf.SetFont("roboto_normal", "", 8)
	if err != nil {
		log.Print(err.Error())
		return
	}
	lynkdata, err := globals.GetLynkAddressDetails(res.StoreName)

	pdf.SetX(curX) //move current location
	pdf.SetY(curY)
	width, _ := pdf.MeasureTextWidth(lynkdata.Address)
	//pdf.SplitText()
	splitdata := strings.SplitAfter(lynkdata.Address, ",")

	index := 0
	i := 0

	for index < len(splitdata) {
		//val := splitdata[index]
		var currLen []string
		added := 0.0

		if width > invoiceX-curX {

			for i = index; i < len(splitdata); i++ {
				if added < 40 {
					currLen = append(currLen, splitdata[i])
					added += float64(len(splitdata[i]))
					index = i + 1
				} else {
					index = i
					break
				}
			}
			pdf.SetX(curX) //move current location
			pdf.SetY(curY)
			pdf.Cell(nil, fmt.Sprintf("%s", strings.Join(currLen, "")))
			curY += 10
			width = width - float64(len(strings.Join(currLen, "")))

		} else {

			pdf.SetX(curX) //move current location
			pdf.Br(10)
			pdf.Cell(nil, fmt.Sprintf("%s", strings.Join(currLen, "")))
			curY += 10
			width = width - float64(len(strings.Join(currLen, "")))
			break

		}

	}

	// pdf.Cell(nil, fmt.Sprintf("%s", lynkdata.Address))
	// curY += 15

	// pdf.SetX(0 + leftMargin)
	// pdf.SetY(45)
	// pdf.Cell(nil, "MYLAPORE, Chennai, Tamil Nadu, 600004")

	//lynk details
	// pdf.SetX(0 + leftMargin)
	// pdf.SetY(45)
	// pdf.Cell(nil, "MYLAPORE, Chennai, Tamil Nadu, 600004")

	err = pdf.SetFont("roboto_bold", "", 8)
	// if err != nil {
	// 	log.Print(err.Error())
	// 	return
	// }
	pdf.SetX(curX)
	pdf.SetY(curY)
	pdf.Cell(nil, fmt.Sprintf("GST No:"))

	err = pdf.SetFont("roboto_normal", "", 8)
	pdf.SetX(curX + 35)
	pdf.Cell(nil, fmt.Sprintf("%s", lynkdata.GSTN))
	curY += 15

	//invoice details column
	curX += 200
	pdf.SetX(invoiceX)
	pdf.SetY(invoiceY)
	err = pdf.SetFont("roboto_bold", "", 8)
	pdf.Cell(nil, fmt.Sprintf("Invoice No:"))
	err = pdf.SetFont("roboto_normal", "", 8)
	pdf.SetX(curX + 45)
	pdf.Cell(nil, fmt.Sprintf("%s", res.InvoiceID))
	curY += 15

	invoiceY += 10

	pdf.SetX(curX)
	pdf.SetY(invoiceY)
	err = pdf.SetFont("roboto_bold", "", 8)
	pdf.Cell(nil, fmt.Sprintf("Invoice Date:"))
	err = pdf.SetFont("roboto_normal", "", 8)
	pdf.SetX(curX + 50)
	pdf.Cell(nil, fmt.Sprintf("%s", now.Format(time.RFC3339)))

	invoiceY += 10

	pdf.SetX(curX)
	pdf.SetY(invoiceY)
	err = pdf.SetFont("roboto_bold", "", 8)
	pdf.Cell(nil, fmt.Sprintf("Order ID:"))
	err = pdf.SetFont("roboto_normal", "", 8)
	pdf.SetX(curX + 40)
	pdf.Cell(nil, fmt.Sprintf("%s", res.OrderID))
	curX += 170

	// customer details
	pdf.SetX(customerX)
	pdf.SetY(customerY)
	err = pdf.SetFont("roboto_bold", "", 8)
	pdf.Cell(nil, fmt.Sprintf("Customer Name: %s", res.CustomerName))
	err = pdf.SetFont("roboto_normal", "", 8)
	pdf.SetX(curX + 65)
	pdf.Cell(nil, fmt.Sprintf("%s", res.CustomerName))

	customerY += 10

	pdf.SetX(curX)
	pdf.SetY(customerY)
	err = pdf.SetFont("roboto_bold", "", 8)
	pdf.Cell(nil, fmt.Sprintf("State Code & Name:"))
	err = pdf.SetFont("roboto_normal", "", 8)
	pdf.SetX(curX + 75)
	pdf.Cell(nil, fmt.Sprintf("%s %s", lynkdata.StateCode, lynkdata.StateName))
	customerY += 10

	pdf.SetX(curX)
	pdf.SetY(customerY)
	// pdf.Cell(nil, "Customer Address:")
	splitdata = strings.SplitAfter(res.Address, ",")
	index = 0
	i = 0

	for index < len(splitdata) {
		// val := splitdata[index]
		var currLen []string
		added := 0.0
		// fmt.Println(val)
		// fmt.Println(pageWindth - leftMargin - customerX)
		// fmt.Println(width)
		if width > invoiceX-curX {

			for i = index; i < len(splitdata); i++ {
				if i == 0 {
					currLen = append(currLen, "Customer Address: ")
					added += float64(len("Customer Address: "))
				}
				if added < 40 {
					currLen = append(currLen, splitdata[i])
					added += float64(len(splitdata[i]))
					index = i + 1
				} else {
					index = i
					break
				}
			}

			//move current location
			pdf.SetY(customerY)
			pdf.SetX(curX)
			pdf.Cell(nil, fmt.Sprintf("%s", strings.Join(currLen, "")))
			customerY += 10
			width = width - float64(len(strings.Join(currLen, "")))

		} else {

			pdf.SetX(curX) //move current location
			pdf.SetY(curY)
			pdf.Cell(nil, fmt.Sprintf("%s", strings.Join(currLen, "")))
			customerY += 10
			width = width - float64(len(strings.Join(currLen, "")))
			break

		}

	}

	// customerY += 10

	pdf.SetX(curX)
	pdf.SetY(customerY)
	err = pdf.SetFont("roboto_bold", "", 8)
	pdf.Cell(nil, fmt.Sprintf("Customer Phone:"))
	err = pdf.SetFont("roboto_normal", "", 8)
	pdf.SetX(curX + 65)
	pdf.Cell(nil, fmt.Sprintf("%s", res.Phone))

	curY = math.Max(curY, math.Max(customerY, invoiceY)) + 15
	// products table
	pdf.SetLineWidth(0.2)
	pdf.SetLineType("solid")
	//pdf.SetLineType("dotted")
	pdf.Line(leftMargin, curY, pageWindth-leftMargin, curY)
	curY += 10
	curX = leftMargin

	slnoX := curX
	//sl.no
	pdf.SetX(slnoX)
	pdf.SetY(curY)
	err = pdf.SetFont("roboto_bold", "", 8)
	pdf.Cell(nil, "S.No") // 15
	curX += 25
	hsncodeX := curX
	//HSN Code
	pdf.SetX(hsncodeX)
	pdf.SetY(curY)
	pdf.Cell(nil, "HSN Code") // 25
	curX += 115
	prodnameX := curX
	//Product Name
	pdf.SetX(prodnameX)
	pdf.SetY(curY)
	pdf.Cell(nil, "Product") // 300
	curX += 100
	qtyX := curX
	// qty
	pdf.SetX(qtyX)
	pdf.SetY(curY)
	pdf.Cell(nil, "Qty") // 20
	curX += 20
	mrpX := curX
	// mrp
	pdf.SetX(mrpX)
	pdf.SetY(curY)
	pdf.Cell(nil, "MRP") // 20
	curX += 35
	rateX := curX

	// rate
	pdf.SetX(rateX)
	pdf.SetY(curY)
	pdf.Cell(nil, "Rate") // 20
	curX += 35
	taxableX := curX

	// taxable amount
	pdf.SetX(taxableX)
	pdf.SetY(curY)
	pdf.Cell(nil, "Taxable") // 20

	pdf.SetX(taxableX)
	pdf.SetY(curY + 10)
	pdf.Cell(nil, "Amount") // 20
	curX += 60
	cgstX := curX

	// CGST
	pdf.SetX(cgstX - 10)
	pdf.SetY(curY)
	pdf.Cell(nil, "CGST") // 50
	cgstperX := curX - 20
	cgstperY := curY + 20

	pdf.SetX(cgstperX)
	pdf.SetY(cgstperY)
	pdf.Cell(nil, "%") // 50
	cgstamtX := curX + 10

	pdf.SetX(cgstamtX)
	pdf.SetY(cgstperY)
	pdf.Cell(nil, "Amt") // 50
	curX += 65
	sgstX := curX

	// SGST
	pdf.SetX(sgstX)
	pdf.SetY(curY)
	pdf.Cell(nil, "SGST") // 50
	sgstperX := curX - 15
	sgstperY := curY + 20

	pdf.SetX(sgstperX)
	pdf.SetY(sgstperY)
	pdf.Cell(nil, "%") // 50
	sgstamtX := curX + 20

	pdf.SetX(sgstamtX)
	pdf.SetY(sgstperY)
	pdf.Cell(nil, "Amt") // 50
	curX += 70
	totalX := curX

	// total amount
	pdf.SetX(totalX)
	pdf.SetY(curY)
	pdf.Cell(nil, "Total") // 50

	pdf.SetX(totalX)
	pdf.SetY(curY + 10)
	pdf.Cell(nil, "Amt") // 50

	maxY := math.Max(sgstperY, cgstperY) + 10
	pdf.Line(leftMargin, maxY, pageWindth-leftMargin, maxY)
	curY = maxY + 5

	// CGST
	// Count
	// CurrentPrice
	// Images
	// Mrp
	// Price
	// ProductDescription
	// SGST
	// Sku
	//
	totalRows := len(res.Products)
	var rate, taxableAmt, cgstper, cgstamt, sgstper, sgstamt, totalamt, prodtotalAmount float64
	var qty int
	err = pdf.SetFont("roboto_normal", "", 8)
	for i := 0; i < totalRows; i++ {

		rate = res.Products[i].Price
		qty = res.Products[i].Count
		taxableAmt = float64(qty) * rate
		cgstper = res.Products[i].CGST
		cgstamt = float64(cgstper/100) * taxableAmt
		sgstper = res.Products[i].SGST
		sgstamt = taxableAmt * float64(sgstper/100)
		totalamt = taxableAmt + cgstamt + sgstamt
		pdf.SetX(slnoX)
		pdf.SetY(curY)
		pdf.Cell(nil, fmt.Sprintf("%d", i+1)) // 15

		pdf.SetX(hsncodeX)
		pdf.SetY(curY)
		pdf.Cell(nil, fmt.Sprintf("%s", res.Products[i].HsnCode)) // 25

		pdf.SetX(prodnameX - 60)
		pdf.SetY(curY)
		pdf.Cell(nil, fmt.Sprintf("%s", res.Products[i].Name)) // 300

		// qty
		pdf.SetX(qtyX)
		pdf.SetY(curY)
		pdf.Cell(nil, fmt.Sprintf("%d", qty)) // 20

		// mrp
		pdf.SetX(mrpX)
		pdf.SetY(curY)
		pdf.Cell(nil, fmt.Sprintf("%0.1f", res.Products[i].Mrp)) // 20

		// rate
		pdf.SetX(rateX)
		pdf.SetY(curY)
		pdf.Cell(nil, fmt.Sprintf("%0.1f", rate)) // 20

		// taxable amount
		pdf.SetX(taxableX)
		pdf.SetY(curY)
		pdf.Cell(nil, fmt.Sprintf("%0.1f", taxableAmt)) // 20

		// CGST
		pdf.SetX(cgstperX)
		pdf.SetY(curY)
		pdf.Cell(nil, fmt.Sprintf("%0.1f", cgstper)) // 50

		pdf.SetX(cgstamtX)
		pdf.SetY(curY)
		pdf.Cell(nil, fmt.Sprintf("%0.1f", cgstamt)) // 50

		// SGST
		pdf.SetX(sgstperX)
		pdf.SetY(curY)
		pdf.Cell(nil, fmt.Sprintf("%0.1f", sgstper)) // 50

		pdf.SetX(sgstamtX)
		pdf.SetY(curY)
		pdf.Cell(nil, fmt.Sprintf("%0.1f", sgstamt)) // 50

		// total amount
		pdf.SetX(totalX)
		pdf.SetY(curY)
		pdf.Cell(nil, fmt.Sprintf("%0.2f", totalamt)) // 50

		// curY += tableGap

		pdf.Line(leftMargin, curY+10, pageWindth-leftMargin, curY+10)
		curY = curY + tableGap
		prodtotalAmount = prodtotalAmount + totalamt
	}

	//GST Summary
	//pdf.Line(leftMargin+30, curY+15, sgstperX, curY+15)
	curY += 40
	curX = leftMargin

	//Populate value in GST Summary

	gstSumX := curX + 30
	pdf.SetX(gstSumX)
	pdf.SetY(curY)
	err = pdf.SetFont("roboto_bold", "", 8)

	pdf.Cell(nil, "GST Summary") // 150
	curX += gstSumX + 140
	cgstSumX := curX

	pdf.SetX(cgstSumX)
	pdf.SetY(curY)
	pdf.Cell(nil, "CGST") // 100
	cgsttaxableSumX := cgstSumX - 60

	pdf.SetX(cgsttaxableSumX)
	pdf.SetY(curY + 15)
	pdf.Cell(nil, "Taxable Amt") // 100
	cgsttaxamtSumX := cgstSumX + 20

	pdf.SetX(cgsttaxamtSumX)
	pdf.SetY(curY + 15)
	pdf.Cell(nil, "Tax Amt") // 100
	sgstSumX := curX + 140

	pdf.SetX(sgstSumX)
	pdf.SetY(curY)
	pdf.Cell(nil, "SGST") // 100
	sgstY := curY
	sgsttaxableSumX := sgstSumX - 60

	pdf.SetX(sgsttaxableSumX)
	pdf.SetY(curY + 15)
	pdf.Cell(nil, "Taxable Amt") // 100
	sgsttaxSumX := sgstSumX + 30

	pdf.SetX(sgsttaxSumX)
	pdf.SetY(curY + 15)
	pdf.Cell(nil, "Tax Amt") // 100
	pdf.Line(gstSumX-20, curY+10, sgsttaxSumX+30, curY+10)
	pdf.Line(gstSumX-20, curY+30, sgsttaxSumX+30, curY+30)

	//pdf.Line(leftMargin+50, curY+40, 470, curY+40)

	err = pdf.SetFont("roboto_normal", "", 8)
	curY += 30
	for i := 0; i < totalRows; i++ {
		rate = res.Products[i].Price
		qty = res.Products[i].Count
		taxableAmt = float64(qty) * rate
		cgstper = res.Products[i].CGST
		cgstamt = float64(cgstper/100) * taxableAmt
		sgstper = res.Products[i].SGST
		sgstamt = taxableAmt * float64(sgstper/100)

		pdf.SetX(gstSumX)
		pdf.SetY(curY + 5)
		pdf.Cell(nil, fmt.Sprintf("%0.2f%s", cgstper+sgstper, "%")) // 150

		pdf.SetX(cgsttaxableSumX)
		pdf.SetY(curY + 5)
		pdf.Cell(nil, fmt.Sprintf("%0.2f", taxableAmt)) // 100

		pdf.SetX(cgsttaxamtSumX)
		pdf.SetY(curY + 5)
		pdf.Cell(nil, fmt.Sprintf("%0.2f", cgstamt)) // 100

		pdf.SetX(sgsttaxableSumX)
		pdf.SetY(curY + 5)
		pdf.Cell(nil, fmt.Sprintf("%0.2f", taxableAmt)) // 100

		pdf.SetX(sgsttaxSumX)
		pdf.SetY(curY + 5)
		pdf.Cell(nil, fmt.Sprintf("%0.2f", sgstamt)) // 100
		curY += 15
		pdf.Line(gstSumX-20, curY, sgsttaxSumX+30, curY)
	}
	curY += 10

	// total summary
	err = pdf.SetFont("roboto_bold", "", 8)
	pdf.SetX(sgsttaxSumX + 70)
	pdf.SetY(sgstY)
	pdf.Cell(nil, "Total") // 100

	err = pdf.SetFont("roboto_normal", "", 8)
	pdf.SetX(sgsttaxSumX + 150)
	pdf.SetY(sgstY)
	pdf.Cell(nil, fmt.Sprintf("₹%0.2f", prodtotalAmount)) // 100

	err = pdf.SetFont("roboto_bold", "", 8)
	pdf.SetX(sgsttaxSumX + 70)
	pdf.SetY(sgstY + 15)
	pdf.Cell(nil, "Delivery Cost") // 100

	err = pdf.SetFont("roboto_normal", "", 8)
	pdf.SetX(sgsttaxSumX + 150)
	pdf.SetY(sgstY + 15)
	pdf.Cell(nil, fmt.Sprintf("₹%0.2f", res.DeliveryCost)) // 100

	pdf.Line(sgsttaxSumX+70, sgstY+35, sgsttaxSumX+200, sgstY+35)

	err = pdf.SetFont("roboto_bold", "", 8)
	pdf.SetX(sgsttaxSumX + 70)
	pdf.SetY(sgstY + 40)
	pdf.Cell(nil, "Grand Total") // 100

	err = pdf.SetFont("roboto_normal", "", 8)
	pdf.SetX(sgsttaxSumX + 150)
	pdf.SetY(sgstY + 40)
	pdf.Cell(nil, fmt.Sprintf("₹%0.2f", prodtotalAmount+res.DeliveryCost)) // 100

	pdf.Line(sgsttaxSumX+70, sgstY+55, sgsttaxSumX+200, sgstY+55)
	// footer
	err = pdf.SetFont("roboto_normal", "", 8)
	if err != nil {
		log.Print(err.Error())
		return
	}
	pdf.SetX(110)
	pdf.SetY(curY + 30)
	pdf.Cell(nil, "This is a Computer Generated Invoice. For any queries/complaints, please call 1800-258-2424")
	pdf.WritePdf("./view/hello.pdf")
}
