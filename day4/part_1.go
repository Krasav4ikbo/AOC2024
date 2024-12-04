package main

func Part1() {
	for i, row := range Table {
		for j, line := range row {
			if line == "X" {
				for way := range 8 {
					search(i, j, 0, way)
				}
			}
		}
	}
}

func search(i, j, current, way int) {
	current++

	if current == len(Letters) {
		PartOneResult++
		return
	}

	switch way {
	case 0:
		i++
		break
	case 1:
		i--
		break
	case 2:
		j++
		break
	case 3:
		j--
		break
	case 4:
		i++
		j++
		break
	case 5:
		i++
		j--
		break
	case 6:
		i--
		j++
		break
	case 7:
		i--
		j--
		break
	}

	if "" != Table[i][j] && Table[i][j] == Letters[current] {
		search(i, j, current, way)
	}
}
