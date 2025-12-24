package closure

func ClosureDemonstrationFunc() func(n *int) int {
	a := 10

	return func(n *int) int {
		return a + *n
	}
}

/*

Its like JS closure where inner function has access to outer function variables even after outer function has returned

A closure means:

	A function
	That captures variables from its surrounding lexical scope
	And can use them even after the outer function has finished execution

*/
