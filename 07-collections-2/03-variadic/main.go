package main

func prependValue(v string, args ...string) []string {
	var newArgs = []string{}
	newArgs = append(newArgs, v)
	return append(newArgs, args...)
}

func DebugLog(args ...string) []string {
	return append([]string{"[DEBUG]"}, args...)
}

func InfoLog(args ...string) []string {
	return prependValue("[INFO]", args...)
}

func ErrorLog(args ...string) []string {
	return prependValue("[ERROR]", args...)
}
