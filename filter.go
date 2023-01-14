package main

import (
	"regexp"
	"strconv"
	"strings"
)

// take data as input and sort through. Find ants number, start and end room,
// additional rooms, connections between rooms.
func filterData(data []string) (antsNr int, allRooms, connections []string) {
	antsNr, _ = strconv.Atoi(data[0])
	var startRoom []string
	startIdx := -1
	for i, line := range data {
		if strings.Contains(line, "##start") {
			startIdx = i
			break
		}
	}
	if startIdx != -1 {
		temp := data[startIdx+1]
		parts := strings.Split(temp, " ")
		startRoom = append(startRoom, parts[0])
	}

	var endRoom string
	endIdx := -1
	for i, line := range data {
		if strings.Contains(line, "##end") {
			endIdx = i
			break
		}
	}
	if endIdx != -1 {
		temp := data[endIdx+1]
		parts := strings.Split(temp, " ")
		endRoom = parts[0]
	}

	var rooms []string
	for i, line := range data {
		if i == startIdx+1 || i == endIdx+1 {
			continue
		}
		if strings.Contains(line, " ") {
			parts := strings.Split(line, " ")
			rooms = append(rooms, parts[0])
		}
	}
	allRooms = append(startRoom, rooms...)
	allRooms = append(allRooms, endRoom)
	re := regexp.MustCompile(`\b\w+-\w+\b`)
	for _, line := range data {
		matches := re.FindAllString(line, -1)
		connections = append(connections, matches...)
	}

	return antsNr, allRooms, connections
}
