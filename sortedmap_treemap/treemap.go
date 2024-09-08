package sortedmap_treemap

import (
	"fmt"
	"github.com/emirpasic/gods/maps/treemap"
)

func TreeMapExample() {
	treeMap := treemap.NewWithStringComparator()
	treeMap.Put("b", 2.0)
	treeMap.Put("a", 1)
	treeMap.Put("c", true)
	treeMap.Put("e", nil)
	treeMap.Put("d", "four")
	treeMap.Put("g", "four")
	treeMap.Put("h", "four")

	fmt.Println(treeMap)
	minKey, minVal := treeMap.Min()
	fmt.Printf("Min {key: %s, value: %s}\n", minKey, minVal)

	ceilingCKey, _ := treeMap.Ceiling("c")
	fmt.Println("Ceiling(\"c\")", ceilingCKey)
	floorCKey, _ := treeMap.Floor("c")
	fmt.Println("Floor(\"c\")", floorCKey)

	ceilingFKey, _ := treeMap.Ceiling("f")
	fmt.Println("Ceiling(\"f\")", ceilingFKey)
	floorFKey, _ := treeMap.Floor("f")
	fmt.Println("Floor(\"f\")", floorFKey)

	// Iterate through the map and print the key-value pairs
	for it := treeMap.Iterator(); it.Next(); {
		fmt.Printf("Key: %s, Value: %v\n", it.Key(), it.Value())
	}
}
