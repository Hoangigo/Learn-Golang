package utils

func CountRectangles(rectangles [][]int) int {
	count := 0
	visited := make([][]bool, len(rectangles))
	for i := range visited {
		visited[i] = make([]bool, len(rectangles[i]))
	}

	for i := 0; i < len(rectangles); i++ {
		for j := 0; j < len(rectangles[i]); j++ {
			if rectangles[i][j] == 1 && !visited[i][j] {
				count++
				exploreRectangles(rectangles, i, j, visited)
			}
		}
	}

	return count
}

func exploreRectangles(rectangles [][]int, i, j int, visited [][]bool) {
	if i < 0 || i >= len(rectangles) || j < 0 || j >= len(rectangles[i]) || rectangles[i][j] == 0 || visited[i][j] {
		return
	}

	visited[i][j] = true
	exploreRectangles(rectangles, i, j+1, visited)
	exploreRectangles(rectangles, i+1, j, visited)
}
