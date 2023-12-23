package main

func Day17Part1(data string) int {
	return 0
}

func Day17Part2(data string) int {
	return 0
}

// // Cannot go in a straight line
// // Find the shortest path
// // Different directions have different costs
// // Need to include the limitation of 3 steps forward

// // From a given starting point, compute the distance to the neighboring nodes (LRUD).
// // For each neighboring node, compute the minimum weight to reach that point, and the
// // direction from which it's coming.
// // Use the direction from which it's coming to force turns.

// const MAX_UINT32 uint32 = ^uint32(0)

// // start : (x, y) starting point.
// // end : (x, y) ending point.
// // graph : (x, y) weight. Neighbors are off by one in x xor y.
// func modifiedDjikstra(start, end [2]uint8, graph map[[2]uint8]uint8) int {
// 	distances := make(map[[2]uint8]uint32, len(graph))
// 	directions := make(map[[2]uint8][]uint8, len(graph))
// 	unvisited := make([][2]uint8, len(graph))

// 	i := 0
// 	for node, _ := range graph {
// 		distances[i] = MAX_UINT32
// 		unvisited[i] = node
// 		i++
// 	}
// }

// type Node struct {
// 	x, y, weight      uint8
// 	distanceFromStart uint32
// 	isVisited         bool
// }

// func (self Node) isNeighbor(other Node) bool {
// 	return (
// 		math.Abs(x
// 	)
// }
