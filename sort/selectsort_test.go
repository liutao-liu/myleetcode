package sort

import (
	"reflect"
	"testing"
)

func Test_sortArray(t *testing.T){
	if !reflect.DeepEqual([]int{1,2,3,5},sortArray([]int{5,2,3,1})){
		t.Fatalf("%v",sortArray([]int{5,2,3,1}))
	}
	if !reflect.DeepEqual([]int{0,0,1,1,2,5},sortArray([]int{5,1,1,2,0,0})){
		t.Fatalf("%v",sortArray([]int{5,1,1,2,0,0}))
	}
}