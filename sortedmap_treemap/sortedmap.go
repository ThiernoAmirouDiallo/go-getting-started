package sortedmap_treemap

import (
	"fmt"
	"github.com/gobs/sortedmap"
)

func SortedMapExample() {
	unsorted := map[string]interface{}{
		"b": 2.0,
		"a": 1,
		"c": true,
		"e": nil,
		"d": "four",
	}

	fmt.Println(sortedmap.AsSortedMap(unsorted))

	//sorted := sortedmap.AsSortedMap(unsorted)
	var sorted sortedmap.SortedMap[string, interface{}] = sortedmap.AsSortedMap(unsorted)
	for _, ele := range sorted {
		fmt.Println(ele) // fmt.Printf("Key: %s, Value: %s\n", ele.Key, ele.Value)
	}
}
