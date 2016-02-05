package main
import "testing"

func TestUnit(t *testing.T) {
	for i:=0; i<len(pol); i++ {
		var testMonArray = createTermsByRegExp(pol[i])
		if len(testMonArray) <= 0 || testMonArray[0] == " " {
			t.Error("this: is Invalid Polynomial")
			return
		}
		print("final result:", calculateFinalResult(testMonArray , val))
		print("this: is OK\n")
	} 
}

type testpair struct {
	pol []float64
	val float64
}
var val = 1.0

var pol = []string {
"-x", 
"-x^1", 
"+x^12",
"-2x^10", 
"+100", 
"-100.001", 
".2x^3 + .2x + .2", 
"-2x^3-2x-2", 
"-x^101 + 2", 
"+1x^2 + 3x^3 + 5x^5 + 7x^7 +1.17", 
"1x^2+3x^3+5x^5+7x^7+0+1",
"cats",
"- 40x^4 + 30x^3 + x^5 -20x^2 +10x+81.3",
}

