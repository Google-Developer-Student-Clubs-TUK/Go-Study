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
	case "range":
		rangeFunction()
	case "functions":
		functions()
	case "multiple_return_values":
		multipleReturnValues()
	case "variadic_functions":
		variadicFunctions()
	case "closures":
		closures()
	case "recursion":
		recursion()
	case "pointers":
		pointers()
	case "strings_and_runes":
		stringsAndRunes()
	case "structs":
		structs()
	case "methods":
		methods()
	case "interfaces":
		interfaces()
	case "struct_embedding":
		structEmbedding()
	case "generics":
		generics()
	case "errors":
		errorsFunction()
	case "goroutines":
		goroutines()
	case "channels":
		channels()
	case "channel_buffering":
		channelBuffering()
	case "channel_synchronization":
		channelSynchronization()
	case "Channel_directions":
		channelDirections()
	case "select":
		selectFunction()
	case "timeouts":
		timeouts()
	case "non-Blocking_channel_operations":
		nonBlockingChannelOperations()
	case "closing_channels":
		closingChannels()
	case "range_over_channels":
		rangeOverChannels()
	case "timers":
		timers()
	case "tickers":
		tickers()
	case "worker_pools":
		workerPools()
	case "wait_groups":
		waitGroups()
	case "rate_limiting":
		rateLimiting()
	case "atomic_counters":
		atomicCounters()
	case "mutexes":
		mutexes()
	case "stateful_goroutines":
		statefulGoroutines()
	case "sorting":
		sorting()
	case "sorting_by_functions":
		sortingByFunctions()
	case "panic":
		panicFunction()
	case "defer":
		deferFuntion()
	case "recover":
		recoverFunction()
	case "string_functions":
		stringFunctions()
	case "string_formatting":
		stringFormatting()
	default:
		println("올바른 파일명을 입력해주세요")
	}
}
