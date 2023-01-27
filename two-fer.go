package main

import "fmt"

type MyStruct1 struct {
	ss int
}

type MyStruct2 struct {
	ss int
}

func ShareWith(name interface{}) string {
	default_message := "One for %v, One for me"
	switch v := name.(type) {
	case MyStruct1:
		// do the stuff you want to do with route1
		return ""
	case MyStruct2:
		// do the stuff you wnat to do with route2
		return ""
	case string:
		return fmt.Sprintf(default_message, v) // fmt.Sprintf( "One for %v, One for me", v)
	default:
		return fmt.Sprintf(default_message, "me")
	}
}
