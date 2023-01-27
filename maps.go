package main

import (
	"fmt"
)

const (
	BestFriend = 1
	Junior     = 2
)

type MyStructure struct {
	field1 string
	field2 string
}

type MyData = map[string]MyStructure

type File = []bool
type ChessBoard = map[string]File
type Alphabets = []rune

var myMap = make(map[string]int)

var myMap2 = map[string]string{}

func main() {

	myAlphabets := Alphabets{'a', 'b', 'c', 'd', 'e'}

	meChessBoard := ChessBoard{
		"A": File{},
		"B": File{},
	}

	k := MyData{
		"A": MyStructure{field1: "sasa", field2: "werwer"},
	}

	fmt.Println(k["A"].field1)

	fmt.Println(meChessBoard["A"][0])
	// l := meChessBoard["A"][0]

	var myMap3 = map[string]string{
		"myYoungerBrothersLife": "my rules until mom steps in",
	}
	myMap3["myLife"] = "my rules"
	i := 0
	ChangeInt(i)
	changeMap(myMap3)
	fmt.Println("myMap: ", myMap3, " i: ", i)

	rahulfamily := make(map[string]string)

	//  rahul family tree
	// key => value
	// map is nothing but a collection of key-value pair
	// one key has only one value
	// what happens when you assign the same key to more than values
	// teh value associate to the key is updated
	// map works like a pointer
	rahulfamily["brother"] = "varun"
	rahulfamily["father"] = "sanjeev"
	rahulfamily["sister"] = "priya"
	rahulfamily["mother"] = "soni"
	rahulfamily["BROTHER"] = "asdsad"

	narenSocialCircle := make(map[int]string)

	narenSocialCircle[BestFriend] = "amit"
	narenSocialCircle[Junior] = "aditya"

	programmingLanguages := map[string]int{
		"python":     0,
		"golang":     1,
		"ruby":       2,
		"javaScript": 3,
		"c#":         4,
		"Kotlin":     5,
		"c":          6,
		"java":       7,
		"c++":        8,
		"rust":       9,
		"assembly":   100000000,
	}

	programmingLanguages["html"] = -1
	programmingLanguages["css"] = 100

	// delete(nameOfMap, key_to_delete)
	delete(programmingLanguages, "java")

	for key := range programmingLanguages {
		delete(programmingLanguages, key)
	}

	programmingLanguages = make(map[string]int)

	// check for a key
	// value, ok:= map[key]
	valueAssociatedWithKey, ok := programmingLanguages["html"]

	_, IsExists := programmingLanguages["html"]
	fmt.Println(IsExists)
	// fmt.Println(_)

	if ok {
		programmingLanguages["html"] += 1
		fmt.Println(valueAssociatedWithKey)
	}

	for key, val := range programmingLanguages {
		fmt.Println(key, val)
	}

	for key := range programmingLanguages {
		fmt.Println(key)
	}

	for _, val := range programmingLanguages {
		fmt.Println(val)
	}

}

func changeMap(input1 map[string]string) {
	input2 := input1
	input2["myLife"] = "mommy's rule"
}

func ChangeInt(i int) {
	i = 8
}
