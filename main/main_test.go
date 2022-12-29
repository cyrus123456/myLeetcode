package main

import (
	"log"
	"reflect"
	"testing"
)

// 1488. 避免洪水泛滥
func TestAvoidFlood(t *testing.T) {
	t.Helper()
	log.SetFlags(log.Lshortfile)

	// if ans := avoidFlood([]int{1, 2, 3, 4}); !reflect.DeepEqual(ans, []int{-1, -1, -1, -1}) {
	// 	t.Errorf("应该是[-1, -1, -1, -1, -1], 却是 %d ", ans)
	// }
	// if ans := avoidFlood([]int{1, 2, 0, 0, 2, 1}); !reflect.DeepEqual(ans, []int{-1, -1, 2, 1, -1, -1}) {
	// 	t.Errorf("应该是[-1, -1, 2, 1, -1, -1], 却是 %d ", ans)
	// }
	if ans := avoidFlood([]int{1, 2, 0, 1, 2}); !reflect.DeepEqual(ans, []int{}) {
		t.Errorf("应该是[], 却是 %d ", ans)
	}
}
