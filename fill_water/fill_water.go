package fill_water

import "container/list"

func FillWaterCount(env []int) int {

	if len(env) <= 2 {
		return 0
	}

	total := 0
	stack := list.New()
	start := 0
	largest := 0

	for i := 1; i < len(env); i++ {

		if env[i] < env[start] {
			stack.PushBack(i)
		}

		if env[i] > env[start] {
		}

		// if env[i] > env[largest] {
		// 	largest = i
		// } else {
		// 	stack.PushBack(i)
		// }

		// // ignore := i - largest
		// // ignored := 0
		// for stack.Len() > 0 {
		// 	key := stack.Remove(stack.Back()).(int)
		// 	size := env[size]
		// 	// if ignored < ignore {
		// 	// 	continue
		// 	// }
		// 	log.Printf("Adding %d", env[largest]-size)
		// 	total += env[largest] - size
		// }
	}

	return total
}
