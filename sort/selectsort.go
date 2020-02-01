package sort

func sortArray(nums []int)[]int{
	return sortArray_quick(nums)
}

func sortArray_select(nums []int) []int {
	var min,i,j,tmp int
	for i=range nums{
		for j = i;j< len(nums);j++{
			if nums[j]<nums[min]{
				min = j
			}
		}
		tmp = nums[i]
		nums[i] = nums[min]
		nums[min]=tmp
		min = i+1
	}
	return nums
}


func sortArray_quick(nums []int) []int {
	quick(nums,0, len(nums)-1)
	return nums
}
/*
*快排
*找到p的位置，每次递归要遍历该区间的所有数据，类似二分查找，所以就出现了n log n的时间复杂度
*相比与冒泡排序，两者每次遍历都会找到一个元素的位置，但是快排每次交换跨度比冒泡大，所以效率也高；
*快排最差情况：已经是正序、逆序的情况下，与冒泡排序的效率就相同了；
 */
func quick(nums []int,l,r int){
	if l>=r{
		return
	}
	var p,i,j int
	i = l
	j = r
	p = nums[l]
	for ;i<j;{
		for ;i<j&&nums[j]>= p;{
			j--
		}
		for ;i<j&&nums[i]<= p;{
			i++
		}
		if i<j{
			nums[l]=nums[j]
			nums[j] = nums[i]
			nums[i]= nums[l]
		}
	}
	nums[l] = nums[i]
	nums[i]= p
	quick(nums,l,j-1)
	quick(nums,j+1,r)
}