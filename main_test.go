package main
import "testing"

func TestCalculateFinalResult(t *testing.T) {

	print(len(testMap))

	for key, value := range testMap {
		var testPol = key
		var testMonArray = createTermsByRegExp(testPol)
		var testFinRes = calculateFinalResult(testMonArray, testVal)
		if(testFinRes != value) {
			t.Error("Status: Error")
			return
		} else {

			print("Status: OK")
		}

		if len(testMonArray) <= 0 {
			t.Error("Status: Invalid Polynomial")
			return
		}
	}
}

type testpair struct {
	testVal float64
	testMap map[string]float64
}

var testVal = 1.9
var testMap = map[string]float64 {
"-x" : -1.9, 
"x^1" : 1.9, 
"x^11" : 1164.9025889821899,
"+x^12" : 2213.3149190661607,
"-2x^10" : -1226.2132515602, 
"+100" : 100.000, 
"-100.001" : -100.001, 
".2x^3 + .2x + .2" : 1.9517999999999998, 
"-2x^3-2x-2" : -19.517999999999997, 
"+1x^2 + 3x^3 + 5x^5 + 7x^7 +1.17" : 774.8721672999999, 
"1x^2+3x^3+5x^5+7x^7+1.17+0" : 774.8721672999999,
"- 40x^4 + 30x^3 + x^5 -20x^2 +10x+81.3" : -262.65301,
//"cat" : 0.000,
}

