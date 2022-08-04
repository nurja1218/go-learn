package learnPart

type calculator func(int, int) int

type calcType string 
const (
	Add= "Add"
	Multi= "Multi"
	Subtr= "Subtr"
	Divis= "Divis"
)

func Calc(t calcType, a int, b int) int {
	addtion := func (a int, b int) int {
		return a + b
	}
	
	multiply := func (a int, b int) int {
		return a * b
	}

	division := func (a int, b int) int {
		return a / b
	}

	subtraction := func (a int, b int) int {
		return a - b
	}
	
	switch {
		case t == Add:
			return addtion(a, b)
		case t == Multi:
			return multiply(a, b)
		case t == Divis:
			return division(a, b)
		case t == Subtr:
			return subtraction(a, b)
		default:
			return 0
	}
}

