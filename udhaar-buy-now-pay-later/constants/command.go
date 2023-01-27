package constants

var ExitCommand = map[string]bool{
	string(27): true, // escape key
	"exit":     true,
	"quit":     true,
	"stop":     true,
}
