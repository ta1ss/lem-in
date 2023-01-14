package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	data, err := readData()
	if err != nil {
		fmt.Println(err)
		return
	}

	printOutput(data)
}

func readData() ([]string, error) {
	if len(os.Args) != 2 {
		return nil, fmt.Errorf("ERROR: incorrect number of arguments.\ninput format: go run . example00.txt")
	}

	path := os.Args[1]
	file, err := os.Open("examples/" + path)
	if err != nil {
		return nil, fmt.Errorf("ERROR: failed to open file: %v", err)
	}
	defer file.Close()

	var data []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data = append(data, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("ERROR: failed to read file: %v", err)
	}

	return data, scanner.Err()
}

func printOutput(data []string) {
	antsNr, allRooms, allConnections := filterData(data)
	if antsNr <= 0 {
		fmt.Println("ERROR: Invalid number of ants. Must be > 0")
		return
	}

	g := newGraph()

	// add rooms to the graph
	roomIDs := make(map[string]int)
	for id, room := range allRooms {
		roomIDs[room] = id
		g.addRoom(id, room)
	}

	// add connections to the graph
	for _, conn := range allConnections {
		parts := strings.Split(conn, "-")
		id1 := roomIDs[parts[0]]
		id2 := roomIDs[parts[1]]
		g.addConnection(id1, id2)
	}

	// assign start and end rooms
	startRoom := g.rooms[0]
	endRoom := g.rooms[len(g.rooms)-1]

	paths := g.findPaths(startRoom, endRoom)
	if len(paths) == 0 {
		fmt.Printf("ERROR: No path from room <%s> to room <%s> found.\n", startRoom.name, endRoom.name)
		return
	}

	validPaths := findCompatiblePaths(paths)
	bestPath := pathAssign(paths, validPaths, antsNr)

	path := os.Args[1]
	bytes, err := ioutil.ReadFile("examples/" + path)
	if err != nil {
		fmt.Println(err)
		return
	}
	content := string(bytes)
	fmt.Println(content)
	fmt.Println()
	printAntSteps(paths, bestPath)
}
