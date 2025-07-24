package main

import (
	"fmt"
	"strings"
)

type celsius float64

func (c celsius) fahrenheit() fahrenheit {
	return fahrenheit((c * 9.0 / 5.0) + 32.0)
}

type fahrenheit float64

func (f fahrenheit) celsius() celsius {
	return celsius((f - 32.0) * 5.0 / 9.0)
}

type usedFunc func(i float64) (string, string)

func main() {
	fmt.Println("from ºC to ºF")
	drawTable("ºC", "ºF", 5, -40, 100, ctof)
	fmt.Println("from ºF to ºC")
	drawTable("ºF", "ºC", 5, -40, 100, ftoc)
}

func drawTable(hdr1, hdr2 string, step, from, to float64, data usedFunc) {
	header := fmt.Sprintf("| %-10v| %-10v|", hdr1, hdr2)
	tableBorder := strings.Repeat("=", len(header))

	fmt.Println(tableBorder)
	fmt.Println(header)
	fmt.Println(tableBorder)

	for val := from; val <= to; val += step {
		i1, i2 := data(val)
		dataRow := fmt.Sprintf("| %-10.6v| %-10.6v|", i1, i2)
		fmt.Println(dataRow)
		fmt.Println(tableBorder)
	}

}

func ctof(i float64) (string, string) {
	c := celsius(i)
	f := c.fahrenheit()

	return fmt.Sprintf("%.2f", c), fmt.Sprintf("%.2f", f)
}

func ftoc(i float64) (string, string) {
	f := fahrenheit(i)
	c := f.celsius()
	return fmt.Sprintf("%.2f", f), fmt.Sprintf("%.2f", c)
}

// Google Gemini's version
// better formated output

//package main
//
//import (
//	"fmt"
//	"strings"
//)
//
//// Define custom types for temperature units.
//// Using float64 as the underlying type is appropriate for precision.
//type Celsius float64
//type Fahrenheit float64
//
//// Celsius methods
//func (c Celsius) Fahrenheit() Fahrenheit {
//	return Fahrenheit((c * 9.0 / 5.0) + 32.0)
//}
//
//// Stringer interface for Celsius for default string representation.
//func (c Celsius) String() string {
//	return fmt.Sprintf("%.2fºC", c) // Format to 2 decimal places and add unit
//}
//
//// Fahrenheit methods
//func (f Fahrenheit) Celsius() Celsius {
//	return Celsius((f - 32.0) * 5.0 / 9.0)
//}
//
//// Stringer interface for Fahrenheit for default string representation.
//func (f Fahrenheit) String() string {
//	return fmt.Sprintf("%.2fºF", f) // Format to 2 decimal places and add unit
//}
//
//// ConverterFunc is a function type that takes a float64 input
//// and returns two strings for table display.
//type ConverterFunc func(input float64) (string, string)
//
//func main() {
//	fmt.Println("🌡️ Temperature Conversion Tables 🌡️")
//
//	// Convert Celsius to Fahrenheit
//	fmt.Println("\n--- Celsius to Fahrenheit ---")
//	drawTable("Celsius", "Fahrenheit", 10, -40, 100, cToFConverter) // Increased step for brevity
//
//	// Convert Fahrenheit to Celsius
//	fmt.Println("\n--- Fahrenheit to Celsius ---")
//	drawTable("Fahrenheit", "Celsius", 10, -40, 100, fToCConverter) // Increased step for brevity
//}
//
//// drawTable generates and prints a formatted table.
//// It takes header strings, step, range, and a conversion function.
//func drawTable(header1, header2 string, step, from, to float64, converter ConverterFunc) {
//	// Define column widths for better alignment.
//	const colWidth = 15
//	headerFormat := fmt.Sprintf("| %%-%dv| %%-%dv|", colWidth, colWidth)
//	rowFormat := fmt.Sprintf("| %%-%vs| %%-%vs|", colWidth, colWidth)
//
//	header := fmt.Sprintf(headerFormat, header1, header2)
//	tableBorder := strings.Repeat("-", len(header))
//
//	fmt.Println(tableBorder)
//	fmt.Println(header)
//	fmt.Println(tableBorder)
//
//	for val := from; val <= to; val += step {
//		strVal1, strVal2 := converter(val)
//		fmt.Printf(rowFormat, strVal1, strVal2)
//		fmt.Println() // Newline for each row
//	}
//	fmt.Println(tableBorder)
//}
//
//// cToFConverter converts Celsius to Fahrenheit for table display.
//func cToFConverter(input float64) (string, string) {
//	c := Celsius(input)
//	f := c.Fahrenheit()
//	return fmt.Sprintf("%.2f", c), fmt.Sprintf("%.2f", f) // Format output consistently
//}
//
//// fToCConverter converts Fahrenheit to Celsius for table display.
//func fToCConverter(input float64) (string, string) {
//	f := Fahrenheit(input)
//	c := f.Celsius()
//	return fmt.Sprintf("%.2f", f), fmt.Sprintf("%.2f", c) // Format output consistently
//}
