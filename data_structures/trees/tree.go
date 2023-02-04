package main

import (
	"errors"
	"fmt"
	"log"

	"github.com/spf13/cast"
)

type Node struct {
	Id       int
	Children []*Node
}

func Build(input []int) (*Node, error) {
	if len(input) == 0 {
		return nil, errors.New("empty input")
	}

	var parent *Node

	for index, val := range input {
		if index == 0 {
			parent = &Node{
				Id: val,
			}
			continue
		}
		parent.Children = append(parent.Children, &Node{Id: val})
	}

	return parent, nil
}

func PrintTree(input *Node) {
	fmt.Println(input.Id)

	for _, val := range input.Children {
		fmt.Print(cast.ToString(val.Id) + " ")
	}
	fmt.Println()
}

func main() {
	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	tree, err := Build(input)
	if err != nil {
		log.Fatal(err)
	}

	PrintTree(tree)
}
