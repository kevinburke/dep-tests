package main

import "fmt"
import bracket "github.com/kevinburke/dep-tests/issue-2046/bracket[bracket/boo"

func main() {
	bracket.Do()
}
