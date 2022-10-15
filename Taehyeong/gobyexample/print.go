package gobyexample

func PrintExample(s string) {
	switch s {
	case "hello_world":
		helloWorld()
	case "values":
		values()
	case "variables":
		variables()
	case "constants":
		constants()
	case "for":
		foFunction()
	case "if_else":
		ifElse()
	case "switch":
		switchFunction()
	case "arrays":
		arrays()
	case "slices":
		slices()
	case "maps":
		maps()
		// case "range":
		// 	range()
		// case "functions":
		// 	functions()
		// case "multiple_return_values":
		// 	multipleReturnValues()
		// case "variadic_functions":
		// 	variadicFunctions()
		// case "closures":
		// 	closures()
		// case "recursion":
		// 	recursion()
		// case "pointers":
		// 	pointers()
		// case "strings_and_runes":
		// 	stringsAndRunes()
		// case "structs":
		// 	structs()
		// case "methods":
		// 	methods()
		// case "interfaces":
		// 	interfaces()
		// case "struct_embedding":
		// 	structEmbedding()
		// case "generics":
		// 	generics()
	}
}
