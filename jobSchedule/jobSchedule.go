package jobSchedule

import (
	"fmt"
	"math"
)

func jobScheduling(startTime []int, endTime []int, profit []int) int {
	jobs := genJob(startTime, endTime, profit)
	radixsort(jobs)

	incomes := []int{0}
	incomesEnd := []int{0}
	for i := range jobs {
		index := selectPreJob(incomesEnd, jobs[i].start)
		do_incomes := incomes[index-1] + jobs[i].profit
		if do_incomes <= incomes[len(incomes)-1] {
			continue
		}
		incomes = append(incomes, do_incomes)
		incomesEnd = append(incomesEnd, jobs[i].end)
	}

	return incomes[len(incomes)-1]
}

func selectPreJob(incomesEnd []int, start int) int {
	left, right, mid := 0, len(incomesEnd)-1, 0
	for {
		// mid向下取整
		mid = int(math.Floor(float64((left + right) / 2)))

		if incomesEnd[mid] > start {
			// 如果当前元素大于k，那么把right指针移到mid - 1的位置
			right = mid - 1
			if right < 0 {
				return mid
			}
			if incomesEnd[right] < start {
				return mid
			}

		} else if incomesEnd[mid] < start {
			// 如果当前元素小于k，那么把left指针移到mid + 1的位置
			left = mid + 1
			if left > len(incomesEnd)-1 {
				return left
			}
			if incomesEnd[left] > start {
				return left
			}

		} else {
			// 否则就是相等了，退出循环
			mid += 1
			break
		}
		if right-left == 1 && incomesEnd[left] < start && incomesEnd[right] > start {
			mid = right
			break
		}
		if incomesEnd[left] == start {
			return left + 1
		}
		if incomesEnd[right] == start {
			return right + 1
		}
		if right-left < 0 {
			fmt.Println("error select")
		}

		// 判断如果left大于right，那么这个元素是不存在的。返回-1并且退出循环
		if left == right {
			mid = left + 1
			break
		}
	}
	// 输入元素的下标
	return mid
}

func getLoopTimes(num int) int {
	count := 1
	temp := num / 10
	for temp != 0 {
		count++
		temp = temp / 10
	}
	return count
}

func getMaxNum(array []*Job) int {
	max := 0
	for i := 0; i < len(array); i++ {
		if max < array[i].end {
			max = array[i].end
		}
	}
	return max
}

func getNumInPos(num int, pos int) int {
	// 求桶的 index 的除数，如：798
	// 个位桶 index = (798 / 1) % 10 = 8
	// 十位桶 index = (798 / 10) % 10 = 9
	// 百位桶 index = (798 / 100) % 10 = 7
	temp := 1
	for i := 0; i < pos-1; i++ {
		temp *= 10
	}
	return (num / temp) % 10
}

const RADIX_10 = 10

/*
*求数据的最大位数,决定排序次数
 */
func maxbit(data []*Job) int {
	n := len(data)
	d := 1 //保存最大的位数
	p := 10
	for i := 0; i < n; i++ {
		for data[i].end >= p {
			p *= 10
			d++
		}
	}
	return d
}

/*
基数排序
0 计算最大比特位数bits，每位bit进行分散、编号、收集
1 分散：计数器，每次循环，统计该位上数字个数
2 编号：计数器中，count[j]=[j-1]+[j]，计算出位置（基数排序中，只比较当前位，其他位不考虑）
3 收集：根据计数器中排序的位置收集各个数字，收集的时候，要注意，从队尾开始遍历确认每个数值的位置（原因：后入计数器的肯定是前一个bit位大的，所以后入计数器的要先收集，这个是重点）
 */
func radixsort(data []*Job) {
	d := maxbit(data)
	tmp := make([]*Job, len(data), len(data))
	count := make([]int, 10, 10) //计数器
	var i, j, k int
	radix := 1
	for i = 1; i <= d; i++ { //进行d次排序
		for j = 0; j < 10; j++ {
			count[j] = 0 //每次分配前清空计数器
		}

		for j = 0; j < len(data); j++ {
			k = (data[j].end / radix) % 10 //统计每个桶中的记录数
			count[k]++
		}
		for j = 1; j < 10; j++ {
			count[j] = count[j-1] + count[j] //将tmp中的位置依次分配给每个桶
		}

		for j = len(data) - 1; j >= 0; j-- { //将所有桶中记录依次收集到tmp中

			k = (data[j].end / radix) % 10
			tmp[count[k]-1] = data[j]
			count[k]--
		}
		for j = 0; j < len(data); j++ {
			data[j] = tmp[j]
		} //将临时数组的内容复制到data中

		radix = radix * 10
	}

}
func radixSort(array []*Job) {
	// 分别为 0~9 的序列空间
	var radixArrays [RADIX_10][]*Job

	// 获取数组中的最大数
	maxNum := getMaxNum(array)
	// 获取最大数的位数，也是再分配的次数
	loopTimes := getLoopTimes(maxNum)
	// 对每一位进行分配
	for pos := 1; pos <= loopTimes; pos++ {
		for i := 0; i < RADIX_10; i++ {
			radixArrays[i] = make([]*Job, 0)
		}
		// 分配过程
		for i := 0; i < len(array); i++ {
			num := getNumInPos(array[i].end, pos)

			radixArrays[num] = append(radixArrays[num], array[i])
		}
		// 收集过程
		i := 0
		j := 0
		for ; i < RADIX_10; i++ {
			for k := 0; k < len(radixArrays[i]); k++ {
				array[j] = radixArrays[i][k]
				j++
			}
		}

	}
}

func genJob(startTime []int, endTime []int, profit []int) []*Job {
	jobs := make([]*Job, len(startTime))
	for i := range startTime {
		jobs[i] = &Job{startTime[i], endTime[i], profit[i]}
	}
	return jobs
}

type Job struct {
	start  int
	end    int
	profit int
}

func (j *Job) getStart() int {
	return j.start
}

