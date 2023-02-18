# lem-in

This project consists of building a digital version of an ant farm.
The program receives as an argument .txt file with instructions consisting of:
- number of ants
- rooms
- connections between rooms

The goal is to bring all ants to end room with as few moves as possible following set of constriction.

[Task Objective and Audit](https://github.com/01-edu/public/tree/master/subjects/lem-in)

## Data representation
> 4 <--- Number of `Ants`

> ##start <--- Next line after ##start indicates `start` room

> 0 0 3 <--- `Room name`, X coordinate, Y Coordinate

>2 2 5

>3 4 0

>##end <--- Next line after ##end indicates `end` room

>1 8 3

>0-2 <--- `Connections` between rooms

>2-3

>3-1

## How To Run
1. Clone the repo:

2. Input the following into the terminal:
 > `go run . <example of your choice.txt>`

 - The .txt files are in `/examples` directory

3. The output will display the contents of .txt file along with steps taken by ants to get to the `end` room.


### Output
>L1-2 <--- `Name` of the ant - moved into `room name`

>L1-3 L2-2

>L1-1 L2-3 L3-2

>L2-1 L3-3 L4-2

>L3-1 L4-3

>L4-1

## Quick testing
To run through all audit tests:
 > `bash testing.sh`


## Implementation
The program was built using standard `GO` packages.
DFS algorithm was used for finding connections between nodes and additional logic implemented for assigning unique paths and deploying ants. 

## Authors
*Viktor Veertee* & *Enri Suimets*
