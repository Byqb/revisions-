<----- atoi ----->
package main
func Atoi(s string) (int, bool) {
	sign := 1
	res := 0
	err := false

	if s[0] == '-' {
		sign = -1
		s = s[1:]
	}
	for _, e := range s {
		if e >= '0' && e <= '9' {
			res = res*10 + int(e-48)
		} else {
			err = true
		}
	}
	return res * sign, err
}

<----- itoa ----->
package main
func Itoa(i int) string {
	if i == 0 {
		return "0"
	}
	res := ""
	for i > 0 {
		res = string((i%10 + 48)) + res
		i /= 10
	}
	if i < 0 {
		return "-" + Itoa(-i)
	} else {
		return res
	}
}

<----- alphamirror ----->
package main
func Alphamirror(s string) string {
	res := ""
	for _, e := range s {
		if e >= 'A' && e <= 'Z' {
			res += string('A' + 'Z' - e)
		} else if e >= 'a' && e <= 'z' {
			res += string('a' + 'z' - e)
		} else {
			res += string(e)
		}
	}
	return res
}
----------or------------
package main
import ("os""github.com/01-edu/z01")
func main() {
	if len(os.Args[1:]) != 1 {
		z01.PrintRune(10)
		return
	} else {
		for _, l := range os.Args[1] {
			if l == 'A' || l == 'a' {
				l = l + 25
				z01.PrintRune(l)
			} else if l == 'B' || l == 'b' {
				l = l + 23
				z01.PrintRune(l)
			} else if l == 'C' || l == 'c' {
				l = l + 21
				z01.PrintRune(l)
			} else if l == 'D' || l == 'd' {
				l = l + 19
				z01.PrintRune(l)
			} else if l == 'E' || l == 'e' {
				l = l + 17
				z01.PrintRune(l)
			} else if l == 'F' || l == 'f' {
				l = l + 15
				z01.PrintRune(l)
			} else if l == 'G' || l == 'g' {
				l = l + 13
				z01.PrintRune(l)
			} else if l == 'H' || l == 'h' {
				l = l + 11
				z01.PrintRune(l)
			} else if l == 'I' || l == 'i' {
				l = l + 9
				z01.PrintRune(l)
			} else if l == 'J' || l == 'j' {
				l = l + 7
				z01.PrintRune(l)
			} else if l == 'K' || l == 'k' {
				l = l + 5
				z01.PrintRune(l)
			} else if l == 'L' || l == 'l' {
				l = l + 3
				z01.PrintRune(l)
			} else if l == 'M' || l == 'm' {
				l = l + 1
				z01.PrintRune(l)
			} else if l == 'N' || l == 'n' {
				l = l - 1
				z01.PrintRune(l)
			} else if l == 'O' || l == 'o' {
				l = l - 3
				z01.PrintRune(l)
			} else if l == 'P' || l == 'p' {
				l = l - 5
				z01.PrintRune(l)
			} else if l == 'Q' || l == 'q' {
				l = l - 7
				z01.PrintRune(l)
			} else if l == 'R' || l == 'r' {
				l = l - 9
				z01.PrintRune(l)
			} else if l == 'S' || l == 's' {
				l = l - 11
				z01.PrintRune(l)
			} else if l == 'T' || l == 't' {
				l = l - 13
				z01.PrintRune(l)
			} else if l == 'U' || l == 'u' {
				l = l - 15
				z01.PrintRune(l)
			} else if l == 'V' || l == 'v' {
				l = l - 17
				z01.PrintRune(l)
			} else if l == 'W' || l == 'w' {
				l = l - 19
				z01.PrintRune(l)
			} else if l == 'X' || l == 'x' {
				l = l - 21
				z01.PrintRune(l)
			} else if l == 'Y' || l == 'y' {
				l = l - 23
				z01.PrintRune(l)
			} else if l == 'Z' || l == 'z' {
				l = l - 25
				z01.PrintRune(l)
			} else {
				z01.PrintRune(l)
			}
		}
	}
	z01.PrintRune(10)
}


<----- chunk ----->
package main
import "fmt"
func Chunk(a []int, ch int) {
	if ch <= 0 {
		return
	}
	var slice []int
	result := make([][]int, 0, len(a)/ch+1)
	for len(a) >= ch {
		slice, a = a[:ch], a[ch:]
		result = append(result, slice)
	}
	if len(a) > 0 {
		result = append(result, a)
	}
	fmt.Println(result)
}

<-----displayfirstparam----->

package main
import (
	"os"
	"github.com/01-edu/z01"
)
func main() {
	if len(os.Args) > 1 {
		firstArg := os.Args[1]
		for _, value := range firstArg {
			z01.PrintRune(value)
		}
		z01.PrintRune('\n')
	}
}
<----- displaylastparam ----->

package main
import (
	"os"
	"github.com/01-edu/z01"
)
func main() {
	printstr(os.Args[len(os.Args)-1])
	z01.PrintRune('\n')
	for _, e := range s {
		z01.PrintRune(e)
	}
}
<----- printbits ----->
package main

import "github.com/01-edu/z01"
func PrintBits(octet byte) {
	for i := 7; i >= 0; i-- {
		if octet&(1<<uint(i)) != 0 {
			z01.PrintRune('1')
		} else {
			z01.PrintRune('0')
		}
	}
}
<----- paramcount ----->

package main
import ("os" "github.com/01-edu/z01")
func main() {
	// z01.PrintRune()                                 This prints a rune
	// len(os.Args) - 1 + '0'                          This takes the length of arguments
	// rune(len(os.Args) - 1 + '0')                    This converts the length of arguments to a rune
	// z01.PrintRune(rune(len(os.Args) - 1 + '0'))     Everything together
	z01.PrintRune(rune(len(os.Args) - 1 + '0'))
	z01.PrintRune('\n')
}
<----- powerof2 ----->
package main
func isPowerOf2(n int) bool {
	return n != 0 && ((n & (n - 1)) == 0)
}


<----- wdmatch ----->
package main

import (
	"os"

	"github.com/01-edu/z01"
)

func ok(s1 string, s2 string) bool {
	runes1 := []rune(s1)
	runes2 := []rune(s2)
	var rest string
	count := 0
	for i := 0; i < len(runes1); i++ {
		for j := count; j < len(runes2); j++ {
			if runes1[i] == runes2[j] {
				rest += string(runes1[i])
				j = len(runes2) - 1
			}
			count++
		}
	}
	return s1 == rest
}

func main() {
	if len(os.Args) == 3 {
		if ok(os.Args[1], os.Args[2]) {
			for _, rng := range os.Args[1] {
				z01.PrintRune(rng)
			}
			z01.PrintRune('\n')
		}
	}
}

<----- swapbits ----->
package checkpoint

func SwapBits(octet byte) byte {
	return octet<<4 | octet>>4
}
<----- tabmult ----->
package main

import (
	"os"
	"strconv"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) == 2 {
		nbr, _ := strconv.Atoi(os.Args[1])
		Tabmult(nbr)
	}
}

func Tabmult(nbr int) {
	for i := 1; i <= 9; i++ {
		z01.PrintRune(rune(i + '0'))
		printstr(" x ")
		printstr(itoa(nbr))
		printstr(" = ")
		printstr(itoa(i * nbr))
		z01.PrintRune('\n')
	}
}
func printstr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}

func itoa(i int) string {
	result := ""

	if i == 0 {
		return "0"
	}

	for i > 0 {
		result = string(i%10+'0') + result
		i /= 10
	}
	if i < 0 {
		return "-" + itoa(-i)
	}
	return result

}
<----- paramcount ----->

package main
import (
	"os"
	"github.com/01-edu/z01"
)
func main() {
	z01.PrintRune(rune(len(os.Args) - 1 + '0'))
	z01.PrintRune('\n')
}

<----- lastword ----->

package main
import ("os""github.com/01-edu/z01")
func main() {
	if len(os.Args) != 2 {
		return
	}
	PrintStr(lastWord(os.Args[1]))
	z01.PrintRune('\n')
}
func rmEndSpace(s string) string {
	for s[len(s)-1] == ' ' {
		s = s[:len(s)-1]
	}
	return s
}
func lastWord(s string) string {
	s = rmEndSpace(s)
	result := ""
	for _, c := range s {
		if c != ' ' {
			result += string(c)
		} else {
			result = ""
		}
	}
	return result
}

<----- displayFirstParam ----->

package main
import ("os""github.com/01-edu/z01")
func main() {
	if len(os.Args) > 1 {
		firstArg := os.Args[1]
		for _, value := range firstArg {
			z01.PrintRune(value)
		}
		z01.PrintRune('\n')
	}
}

<-----fishandchips----->

package main
import "fmt"
func FishAndChips(n int) string {
	if n%2 == 0 && n%3 == 0 {
		return "fish and chips"
	}
	if n%2 == 0 {
		return "fish"
	}
	if n%3 == 0 {
		return "chips"
	}
	if n != n%2 && n != n%3 {
		return "error: non divisible"
	}
	if n%2 < 0 && n%3 < 0 {
		return "error: number is negative"
	}
	return ""
}

<----- recursivesum ----->

package main
func RecursiveSum(n int) int {
	if n == 0 {
		return 0
	}
	return n + RecursiveSum(n-1)
}

<----- reduceint ----->
package main
import "strconv"
// func ReduceInt(a []int, f func(int, int) int) {
// 	firstNum := a[0]
// 	secondNum := a[1]
// 	result := f(firstNum, secondNum)
// 	PrintStr(strconv.Itoa(result))
// }
// shorter version
func ReduceInt(a []int, f func(int, int) int) {
	PrintStr(strconv.Itoa(f(a[0], a[1])))

}
<----- rot13 ----->
package main
func Rot13(s string) string {
	res := ""
	for _, e := range s {
		if e >= 'A' && e <= 'Z' {
			res += string((e-'A'+13)%26 + 'A')
		} else if e >= 'a' && e <= 'z' {
			res += string((e-'a'+13)%26 + 'a')
		} else {
			res += string(e)
		}
	}
	return res
}

<-----  swapbits ----->
package main
func SwapBits(octet byte) byte {
	return octet<<4 | octet>>4
}
<-----  ----->

<-----  ----->

