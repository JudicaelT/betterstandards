package assert

func NotNil(value any) {
	Condition(
		value != nil,
		"Failed asserting that value is not nil",
	)
}
