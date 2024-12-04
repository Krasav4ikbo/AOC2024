package main

func Part2() {
	for i, row := range Table {
		for j, line := range row {
			if line == "A" {
				searchCross(i, j)
			}
		}
	}
}

func searchCross(i, j int) {
	if i+1 <= len(Table) && j+1 <= len(Table[i]) &&
		i-1 >= 0 && j-1 >= 0 {
		if ((Table[i+1][j+1] == "M" || Table[i+1][j+1] == "S") &&
			(Table[i-1][j-1] == "M" || Table[i-1][j-1] == "S") &&
			(Table[i+1][j+1] != Table[i-1][j-1])) &&
			((Table[i+1][j-1] == "M" || Table[i+1][j-1] == "S") &&
				(Table[i-1][j+1] == "M" || Table[i-1][j+1] == "S") &&
				(Table[i+1][j-1] != Table[i-1][j+1])) {
			PartTwoResult++
		}
	}

}
