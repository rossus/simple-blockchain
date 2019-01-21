package tutorial

import (
	"strconv"
)

func RouteToLesson(args []string) {
	var n int

	if len(args) >= 2 {
		choice, err := strconv.ParseInt(args[1], 10, 64)
		if err != nil {
			n = 1
		} else {
			n = int(choice)
		}
	} else {
		n = 1
	}

	switch n {
	case 2:
		lesson2main()
		break
	default:
		lesson1main()
	}
}
