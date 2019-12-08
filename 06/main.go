package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func getInput() map[string]string {
	dict := make(map[string]string)
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	s := strings.Split(string(data), "\r\n")
	for i := range s {
		t := strings.Split(s[i], ")")
		dict[t[1]] = t[0]
	}

	return dict
}

func printMap(dict map[string]string) {
	for k, v := range dict {
		fmt.Println(k, ": ", v)
	}
}

func findPaths(dict map[string]string) int {
	cnt := 0
	for k, _ := range dict {
		cur := k
		for {
			if cur == "COM" {
				break
			}
			cur = dict[cur]
			cnt++
		}
	}
	return cnt
}

// State describes the current visited state
type State struct {
	node  string
	steps int
}

func reverseDict(dict map[string]string) map[string][]string {
	revDict := make(map[string][]string)
	for k, v := range dict {
		revDict[v] = append(revDict[v], k)
	}
	return revDict
}

func bfs(dict map[string]string) int {
	revDict := reverseDict(dict)

	seen := make(map[string]bool)
	cur := State{
		node:  "YOU",
		steps: 0,
	}
	next := make([]State, 0)
	next = append(next, cur)

	for {
		// Pop the top element
		cur = next[0]
		next = next[1:]

		// Check for seen
		if cur.node == "" || seen[cur.node] {
			continue
		}

		// Exit condition
		if cur.node == "SAN" {
			return cur.steps
		}

		// Set the state
		seen[cur.node] = true

		// Add the next states into the queue
		a := State{
			node:  dict[cur.node],
			steps: cur.steps + 1,
		}
		next = append(next, a)

		for i := range revDict[cur.node] {
			a = State{
				node:  revDict[cur.node][i],
				steps: cur.steps + 1,
			}
			next = append(next, a)
		}
	}
}

func main() {
	dict := getInput()
	//printMap(dict)
	fmt.Println(findPaths(dict))
	fmt.Println(bfs(dict) - 2)
}
