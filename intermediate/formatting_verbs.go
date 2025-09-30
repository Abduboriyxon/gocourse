package intermediate

import "fmt"

func main() {

	// --- General Formatting Verbs
	// %v	Prints the value in the default format
	// %#v	Prints the value in Go-syntax format
	// %T	Prints the type of the value
	// %%	Prints the % sign

	// i := 15.6
	// string := "Hello World"

	// fmt.Printf("%v\n", i)    // 15.6
	// fmt.Printf("%#v\n", i)   // 15.6
	// fmt.Printf("%T\n", i)    // float64
	// fmt.Printf("%v%%\n", i)       // 15.6%

	// fmt.Printf("%v\n", string)    // Hello World
	// fmt.Printf("%#v\n", string)   // "Hello World"
	// fmt.Printf("%T\n", string)    // string

	// --- Integer Formatting Verbs
	// 	%b	Base 2
	// %d	Base 10
	// %+d	Base 10 and always show sign
	// %o	Base 8
	// %O	Base 8, with leading 0o
	// %x	Base 16, lowercase
	// %X	Base 16, uppercase
	// %#x	Base 16, with leading 0x
	// %4d	Pad with spaces (width 4, right justified)
	// %-4d	Pad with spaces (width 4, left justified)
	// %04d	Pad with zeroes (width 4)

	// int := 18

	// fmt.Printf("%b\n", int)    // 10010 that is binary representation of 18
	// fmt.Printf("%d\n", int)    // 18 that is decimal representation of 18
	// fmt.Printf("%+d\n", int)   // +18 that is decimal representation of 18 with sign
	// fmt.Printf("%o\n", int)    // 22 that is octal representation of 18
	// fmt.Printf("%O\n", int)    // 0o22 that is octal representation of 18 with leading 0o
	// fmt.Printf("%x\n", int)    // 12 that is hexadecimal representation of 18
	// fmt.Printf("%X\n", int)    // 12 that is hexadecimal representation of 18 in uppercase
	// fmt.Printf("%#x\n", int)   // 0x12 that is hexadecimal representation of 18 with leading 0x
	// fmt.Printf("%4d\n", int)   //   18 (that is 18 right justified in width 4)
	// fmt.Printf("%-4d\n", int)  // 18  (that is 18 left justified in width 4)
	// fmt.Printf("%04d\n", int)  // 0018 (that is 18 padded with zeroes in width 4)


	// --- String Formatting Verbs
	// %s	Prints the value as plain string
	// %q	Prints the value as a double-quoted string
	// %8s	Prints the value as plain string (width 8, right justified)
	// %-8s	Prints the value as plain string (width 8, left justified)
	// %x	Prints the value as hex dump of byte values
	// % x	Prints the value as hex dump with spaces

	// txt := "World"

	// fmt.Printf("%s\n", txt)    // World
	// fmt.Printf("%q\n", txt)    // "World"
	// fmt.Printf("%8s\n", txt)   //    World (that is World right justified in width 8)
	// fmt.Printf("%-8s\n", txt)  // World    (that is World left justified in width 8)
	// fmt.Printf("%x\n", txt)    // 576f726c64 (that is hex dump of byte values)
	// fmt.Printf("% x\n", txt)   // 57 6f 72 6c 64 (that is hex dump of byte values with spaces)

	// --- Boolean Formatting Verbs
	// %t	Value of the boolean operator in true or false format (same as using %v)

	// t :=  n", f)    // false

	// 	--- Float Formatting Verbs
	// Verb	Description
	// %e	Scientific notation with 'e' as exponent
	// %f	Decimal point, no exponent
	// %.2f	Default width, precision 2
	// %6.2f Minimum Width 6, precision 2
	// %g	Exponent as needed, only necessary digits

	flt := 9.18

	fmt.Printf("%e\n", flt)    // 9.180000e+00
	fmt.Printf("%f\n", flt)    // 9.180000
	fmt.Printf("%.2f\n", flt)  // 9.18
	fmt.Printf("%6.2f\n", flt) //   9.18 (that is 9.18 right justified in width 6)
	fmt.Printf("%g\n", flt)    // 9.18

}