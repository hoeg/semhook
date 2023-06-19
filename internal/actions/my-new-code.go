package actions

type AntiPatternStruct struct {
	Name string
}

func myNiceAntiPattern() AntiPatternStruct {
	return AntiPatternStruct{Name: "test"}
}
