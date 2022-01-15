package chapter09

type Item struct {
	Name   string
	Value  int
	Weight int
}

// Knapsack is a popular problem
func Knapsack(items []Item, limit int) int {
	// why +1? Because we add a row with 0-s that can akt as a minimum value for
	// the first items.
	cell := make([][]int, len(items)+1)
	for i := 0; i < len(items)+1; i++ {
		cell[i] = make([]int, limit+1)
	}
	// the first row is left empty with 0s, so it can be the `max` in case of the first item.
	// Take care to start from 1 for row and col as well.
	for i := 1; i <= len(items); i++ {
		for j := 1; j <= limit; j++ {
			item := items[i-1]
			if item.Weight > j {
				cell[i][j] = cell[i-1][j]
			} else {
				cell[i][j] = max(cell[i-1][j], cell[i-1][j-item.Weight]+item.Value)
			}
		}
	}
	return cell[len(items)][limit]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
