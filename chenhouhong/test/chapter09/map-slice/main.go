package main

import (
	"fmt"
	"sort"
)

func main() {
	monsters := make([]map[string]string, 2)
	for i := range monsters {
		monsters[i] = make(map[string]string, 2)
		monsters[i]["name"] = "Tom"
		monsters[i]["age"] = "90"
	}
	// monsters = append(monsters, make(map[string]string, 2))		错误？？

	newMonster := make(map[string]string, 2)
	newMonster["name"] = "Tom"
	newMonster["age"] = "90"
	monsters = append(monsters, newMonster)
	fmt.Printf("monsters: %v\n", monsters)

	fmt.Println()

	slice := []int{2, 4, 6, 7, 6}
	sort.Ints(slice)
	fmt.Printf("slice: %v\n", slice)

}
