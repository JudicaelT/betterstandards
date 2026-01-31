package assert

func False(value bool) {
	Condition(
		value == false,
		"Failed asserting that value is false",
	)
}
