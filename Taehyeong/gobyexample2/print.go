package gobyexample2

func PrintExample(s string) {
	switch s {
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
