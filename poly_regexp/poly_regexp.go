package poly_regexp
import "fmt"
import s "strings"
import "math"
import "strconv"
import "regexp"
var print = fmt.Println

func CreateTermsByRegExp(pol string) []string {

	pol = s.Replace(pol, " ", "", -1)
	var pattern = "([+-]?\\d*(?:\\.?\\d*))x(\\^(\\d*))?|([+-]\\d*(?:\\.?\\d*))"
	pat, _ := regexp.Compile(pattern)
	return pat.FindAllString(pol, -1)
}

func EvaluateTerm(singleTerm string , val float64) float64 {
	splittedCoeffAndPower := DetermineTypeOfTermForSplitting(singleTerm)
	coeff, exp := ConvertTermFromStringToDouble(splittedCoeffAndPower)
	return coeff * (math.Pow(val , exp))
}

func DetermineTypeOfTermForSplitting(singleTerm string) []string {
	if s.Contains(singleTerm, "^") == true {
		if s.HasPrefix(singleTerm, "x") == true || s.HasPrefix(singleTerm, "-x") == true || s.HasPrefix(singleTerm, "+x") == true {
			singleTerm = s.Replace(singleTerm, "x", "1x", -1)
		} 
	} else {
		if s.Contains(singleTerm , "x") == true {
			singleTerm = s.Replace(singleTerm, "x", "x^1", -1)
		} else {
			var appender string = "x^0"
			singleTerm = s.Join([]string{singleTerm, appender} , "")
		}
	} 
	return SplitTermIntoCoeffAndPower(singleTerm)
}

func SplitTermIntoCoeffAndPower(singleTerm string) []string {

	if singleTerm == "x^1" || singleTerm == "-x^1" {
		singleTerm = s.Replace(singleTerm, "x", "1", -1)
		return s.Split(singleTerm, "^")
	} 
	return s.Split(singleTerm, "x^")
}

func ConvertTermFromStringToDouble(splittedCoeffAndPower []string) (coeff, exp float64) {
	var coeffAndExpArray = []float64{}
	for _,i := range splittedCoeffAndPower {
		flt, err := strconv.ParseFloat(i , 64)
		if err != nil {
			panic(err)
		}
		coeffAndExpArray = append(coeffAndExpArray, flt)
	}
	coeff = coeffAndExpArray[0]
	exp = coeffAndExpArray[1]
	return coeff, exp
}

func CalculateFinalResult(monomialsArray []string, val float64) float64 {
	var finalRes float64 = 0.0

	for i := 0; i < len(monomialsArray); i++ {
		finalRes = finalRes + EvaluateTerm(monomialsArray[i] , val)	
	}
    return finalRes
}

func main() {
    pol := "-x"
    val := 1.0
    monomialsArray := CreateTermsByRegExp(pol)
    CalculateFinalResult(monomialsArray , val)
}