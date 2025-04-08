package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var N, M int
	fmt.Fscan(reader, &N, &M)

	houses := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Fscan(reader, &houses[i])
	}
	heaters := make([]int, M)
	for j := 0; j < M; j++ {
		fmt.Fscan(reader, &heaters[j])
	}

	sort.Ints(houses)
	sort.Ints(heaters)

	if isGood(0, houses, heaters, N, M) {
		fmt.Println(0)
		return
	}

	left, right := 0, int(1e9)
	answer := right

	for left <= right {
		mid := (left + right) / 2
		if isGood(mid, houses, heaters, N, M) {
			answer = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	fmt.Println(answer)
}

func isGood(radius int, houses, heaters []int, N, M int) bool {
	i, j := 0, 0

	for i < N && j < M {
		if heaters[j]-radius <= houses[i] && houses[i] <= heaters[j]+radius {
			i++
		} else {
			j++
		}
	}

	return i == N
}
