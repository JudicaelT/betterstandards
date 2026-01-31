package assert

func True(value bool) {
	Condition(
		value == true,
		"Failed asserting that value is true",
	)
}
