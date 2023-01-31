package constants

import "fmt"

const (
	Reset  = "\033[1m"
	Red    = "\033[31m" // for errors
	Green  = "\033[32m" // for success messages
	Yellow = "\033[33m" // for welcome message
	Blue   = "\033[34m" // for execution_time
	Purple = "\033[35m"
	Cyan   = "\033[36m"
	Gray   = "\033[37m"
	White  = "\033[97m" // for input
)

var UdhaarAppSymbol = fmt.Sprintf(`%v
       |     ____      |
       |      |  \     |
       |     _|__/     |
        \             /
         \           / 
          \_________/	
				%v `, Yellow, White)
