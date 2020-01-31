package jobSchedule

import (
	"reflect"
	"testing"
)


func Test_JobSchedule(t *testing.T){


	if !reflect.DeepEqual(66,jobScheduling([]int{43,13,36,31,40,5,47,13,28,16,2,11},[]int{44,22,41,41,47,13,48,35,48,26,21,39},[]int{8,20,3,19,16,8,11,13,2,15,1,1})){
	t.Fatalf("%d",jobScheduling([]int{43,13,36,31,40,5,47,13,28,16,2,11},[]int{44,22,41,41,47,13,48,35,48,26,21,39},[]int{8,20,3,19,16,8,11,13,2,15,1,1}))
	}
	if !reflect.DeepEqual(120,jobScheduling([]int{1,2,3,3},[]int{3,4,5,6},[]int{50,10,40,70})){
		t.Fail()
	}
	if !reflect.DeepEqual(120,jobScheduling([]int{1,2,3,3},[]int{3,4,5,6},[]int{50,10,40,70})){
		t.Fail()
	}
}

//func Test_radixSort(t *testing.T){
//	jobs:=[]*Job{
//		{1,44,3},
//		{1,22,3},
//		{1,41,3},
//		{1,41,3},
//		{1,47,3},
//		{1,13,3},
//		{1,48,3},
//		{1,35,3},
//		{1,48,3},
//		{1,26,3},
//		{1,21,3},
//		{1,39,3},
//	}
//	radixsort(jobs)
//	for i:=range jobs{
//		fmt.Println(jobs[i].end)
//	}
//
//}

func Benchmark_JobSchedule(b *testing.B){
	b.ResetTimer()
	for i:=0;i<b.N;i++{
		jobScheduling([]int{43,13,36,31,40,5,47,13,28,16,2,11},[]int{44,22,41,41,47,13,48,35,48,26,21,39},[]int{8,20,3,19,16,8,11,13,2,15,1,1})
	}
}

func Test_binary_select(t *testing.T){
	if !reflect.DeepEqual(2,selectPreJob([]int{1,2,3,5},2)){
		t.Fatalf("%d",selectPreJob([]int{1,2,3,5},2))
	}
	//fmt.Println(selectPreJob([]int{1,2,4,6},3))
	if !reflect.DeepEqual(2,selectPreJob([]int{1,2,4,6},3)){
		t.Fatalf("%d",selectPreJob([]int{1,2,4,6},3))
	}
	if !reflect.DeepEqual(8,selectPreJob([]int{1,2,3,5,8,13,21,33,54},34)){
		t.Fatalf("%d",selectPreJob([]int{1,2,3,5,8,13,21,33,54},34))
	}
	if !reflect.DeepEqual(8,selectPreJob([]int{1,2,3,5,8,13,21,33,54},33)){
		t.Fatalf("%d",selectPreJob([]int{1,2,3,5,8,13,21,33,54},33))
	}
	if !reflect.DeepEqual(2,selectPreJob([]int{1,2,3,5,8,13,21,33,54},2)){
		t.Fatalf("%d",selectPreJob([]int{1,2,3,5,8,13,21,33,54},2))
	}
	if !reflect.DeepEqual(4,selectPreJob([]int{1,2,3,5,8,13,21,33,54},7)){
		t.Fatalf("%d",selectPreJob([]int{1,2,3,5,8,13,21,33,54},7))
	}
	if !reflect.DeepEqual(9,selectPreJob([]int{1,2,3,5,8,13,21,33,54},54)){
		t.Fatalf("%d",selectPreJob([]int{1,2,3,5,8,13,21,33,54},54))
	}
	if !reflect.DeepEqual(9,selectPreJob([]int{1,2,3,5,8,13,21,33,54},59)){
		t.Fatalf("%d",selectPreJob([]int{1,2,3,5,8,13,21,33,54},59))
	}

}