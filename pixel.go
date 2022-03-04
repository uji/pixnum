package pixnum

type point struct {
	x, y int
}

type pixel []point

var zero pixel = pixel{
	{x: 1, y: 0},
	{x: 0, y: 1},
	{x: 2, y: 1},
	{x: 0, y: 2},
	{x: 2, y: 2},
	{x: 0, y: 3},
	{x: 2, y: 3},
	{x: 1, y: 4},
}

var one pixel = pixel{
	{x: 1, y: 0},
	{x: 0, y: 1},
	{x: 1, y: 1},
	{x: 1, y: 2},
	{x: 1, y: 3},
	{x: 1, y: 4},
}

var two pixel = pixel{
	{x: 1, y: 0},
	{x: 0, y: 1},
	{x: 2, y: 1},
	{x: 2, y: 2},
	{x: 1, y: 3},
	{x: 0, y: 4},
	{x: 1, y: 4},
	{x: 2, y: 4},
}

var three pixel = pixel{
	{x: 0, y: 0},
	{x: 1, y: 0},
	{x: 2, y: 0},
	{x: 2, y: 1},
	{x: 1, y: 2},
	{x: 2, y: 3},
	{x: 0, y: 4},
	{x: 1, y: 4},
	{x: 2, y: 4},
}

var four pixel = pixel{
	{x: 0, y: 0},
	{x: 2, y: 0},
	{x: 0, y: 1},
	{x: 2, y: 1},
	{x: 0, y: 2},
	{x: 1, y: 2},
	{x: 2, y: 2},
	{x: 2, y: 3},
	{x: 2, y: 4},
}

var five pixel = pixel{
	{x: 0, y: 0},
	{x: 1, y: 0},
	{x: 2, y: 0},
	{x: 0, y: 1},
	{x: 0, y: 2},
	{x: 1, y: 2},
	{x: 2, y: 3},
	{x: 0, y: 4},
	{x: 1, y: 4},
}

var six pixel = pixel{
	{x: 1, y: 0},
	{x: 2, y: 0},
	{x: 0, y: 1},
	{x: 0, y: 2},
	{x: 1, y: 2},
	{x: 0, y: 3},
	{x: 2, y: 3},
	{x: 1, y: 4},
}

var seven pixel = pixel{
	{x: 0, y: 0},
	{x: 1, y: 0},
	{x: 2, y: 0},
	{x: 2, y: 1},
	{x: 2, y: 2},
	{x: 1, y: 3},
	{x: 1, y: 4},
}

var eight pixel = pixel{
	{x: 1, y: 0},
	{x: 0, y: 1},
	{x: 2, y: 1},
	{x: 1, y: 2},
	{x: 0, y: 3},
	{x: 2, y: 3},
	{x: 0, y: 4},
	{x: 1, y: 4},
	{x: 2, y: 4},
}

var nine pixel = pixel{
	{x: 0, y: 0},
	{x: 1, y: 0},
	{x: 2, y: 0},
	{x: 0, y: 1},
	{x: 2, y: 1},
	{x: 0, y: 2},
	{x: 1, y: 2},
	{x: 2, y: 2},
	{x: 2, y: 3},
	{x: 2, y: 4},
}

func getPoints(n int) pixel {
	switch n {
	case 0:
		return zero
	case 1:
		return one
	case 2:
		return two
	case 3:
		return three
	case 4:
		return four
	case 5:
		return five
	case 6:
		return six
	case 7:
		return seven
	case 8:
		return eight
	case 9:
		return nine
	default:
		panic("number is not one digit")
	}
}
