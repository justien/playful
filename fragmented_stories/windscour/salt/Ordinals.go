package salt

/*
this function takes a number, and returns the ordinal suffix
of that number
*/

func Ordinals(i int) string {

	if i == 1 {
		return "st"
	} else if i == 2 {
		return "nd"
	} else if i == 3 {
		return "rd"
	} else {
		return "th"
	}
}
