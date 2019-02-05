// Missionaries and Cannibals for CSI 380
// This program solves the classic Missionaries and Cannibals problem

package main

import "fmt"

// A representation of the state of the game
type position struct {
	boatOnWestBank   bool // true is west bank, false is east bank
	westMissionaries int  // west bank missionaries
	westCannibals    int  // west bank cannibals
	eastMissionaries int  // east bank missionaries
	eastCannibals    int  // east bank cannibals
}

// Is this a legal position? In particular, does it have
// more cannibals than missionaries on either bank? Because that is illegal.
func valid(pos position) bool {
	//fmt.Println(pos.westMissionaries)
	if pos.westMissionaries > 3 || pos.westCannibals > 3 || pos.eastMissionaries > 3 || pos.eastCannibals > 3{
		return false
	}
	if pos.eastMissionaries == 0 || pos.westMissionaries == 0 {
		return true
	}
	if pos.westCannibals <= pos.westMissionaries && pos.eastCannibals <= pos.eastMissionaries{
		return true
	}else {
		return false
	}
}

// What are all of the next positions we can go to legally from the current position
// Returns nil if there are no valid positions
func (pos position) successors() []position {
	temp := []position{}
	var m int = 0
	var c int = 0
	if pos.boatOnWestBank == true{
		m = pos.westMissionaries
		c = pos.westCannibals
	}else {
		m = pos.eastMissionaries
		c = pos.eastCannibals
	}

	for i := 0; i <= m; i++{
		for j := 0; j <= c; j++{
			if i+j <= 2 && i+j != 0 {
				test := pos
				if pos.boatOnWestBank == true {
					test.eastMissionaries += i
					test.westMissionaries -= i
					test.eastCannibals += j
					test.westCannibals -= j
					test.boatOnWestBank = false
				} else {
					test.westMissionaries += i
					test.eastMissionaries -= i
					test.westCannibals += j
					test.eastCannibals -= j
					test.boatOnWestBank = true
				}
				if valid(test) {
					temp = append(temp, test)
				}
			}
		}
	}
	//fmt.Println()
	return temp
}

// A recursive depth-first search that goes through to find the goal and returns the path to get there
// Returns nil if no solution found
func dfs(start position, goal position, solution []position, visited map[position]bool) []position {
	// Check if start position == goal
	if start == goal {
		return solution
	}
	// Get all successors
	for possible := start.successors(); len(possible) > 0; {
		// If not visted -> Do next thing
		choice := possible[0]
		if !visited[choice]{
			//fmt.Println(possible)
			temp_solution := append(solution, choice)
			visited[choice] = true
			anwser := dfs(choice, goal, temp_solution, visited)
			if anwser != nil{
				return anwser

			}
		}
		possible = possible[1:]
	}
	// If naw, return fam
	return nil
}

func main() {
	start := position{boatOnWestBank: true, westMissionaries: 3, westCannibals: 3, eastMissionaries: 0, eastCannibals: 0}
	goal := position{boatOnWestBank: false, westMissionaries: 0, westCannibals: 0, eastMissionaries: 3, eastCannibals: 3}
	solution := dfs(start, goal, []position{start}, make(map[position]bool))
	fmt.Println(solution)
}
