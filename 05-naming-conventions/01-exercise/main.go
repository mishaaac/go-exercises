/*
Exercise 1: Idiomatic Go Naming Conventions Refactoring

Fix all naming convention violations in the provided code snippet to align
strictly with Go's idiomatic style guidelines (camelCase, PascalCase, initialisms/acronyms,
and package visibility rules).

Key Adjustments Made:
  - Constants: Replaced SCREAMING_SNAKE_CASE (MAX_LIMIT) with lowerCamelCase (maxLimit) or UpperCamelCase.
  - Initialisms: Standardized acronyms to consistent uppercase (userID instead of User_Id or userId).
  - Struct & Field Names: Replaced snake_case (user_profile) with unexported camelCase (userProfile)
    and unexported field names (userID, emailAddress).
  - Local Variables: Replaced PascalCase_With_Snake (Employee_Name) with idiomatic short declarations (employeeName := "Carlos").
*/

package main

import "fmt"

// Go constants do NOT use SCREAMING_SNAKE_CASE.
// Use camelCase for unexported constants (or PascalCase if exported).
const maxLimit = 100

// Struct names and fields use mixedCaps (camelCase / PascalCase).
// Acronyms like ID stay uppercase (userID, not userId or User_Id).
type userProfile struct {
	userID       int
	emailAddress string
}

func main() {
	// Local variables use lowerCamelCase and idiomatic short declaration (:=).
	employeeName := "Carlos"

	fmt.Println(employeeName)
}
