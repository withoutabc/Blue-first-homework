/*猜数规则：系统随机生成一个四个数字（数字范围是0-9）都不相同的，然后
用户猜测一个四个数字都不相同的四位数，若猜测错误，则提示正确的数字
个数和数字处在正确位置的个数，并进行下一次猜测，直到猜测成功
*/
//此程序中4可以变为n作为推广，即未必一定是四位数
package main

import (
	"fmt"
	"math/rand"
	"sort"
	"strconv"
	"strings"
	"time"
)

func main() {
	var (
		x            int
		secretNumber int
		s2           []rune
		b            bool
		sum          = 0
	)
	s4 := make([]int, 0)
	//随机生成一个数字，保证四个数字都不重复
	for {
		rand.Seed(time.Now().UnixNano())
		secretNumber = rand.Intn(8999) + 1000
		s1 := strconv.Itoa(secretNumber)
		s2 = []rune(s1)
		for j := 0; j < 4; j++ {
			s4 = append(s4, int(s2[j]))
		}
		b = findDuplicates(s4)
		s4 = []int{}
		if b == true {
			break
		}
	}
	//fmt.Println(secretNumber)
	fmt.Println("Please input your guess")
	for {
		var guess string
		_, err := fmt.Scan(&guess)
		s3 := []rune(guess)
		//不是4个数字则报错
		if len(s3) != 4 {
			fmt.Println("error!")
			continue
		}
		for j := 0; j < 4; j++ {
			s4 = append(s4, int(s3[j]))
		}
		c := findDuplicates(s4)
		s4 = []int{}
		//输入的数有重复数字则报错
		if c == false {
			fmt.Println("error!")
			continue
		}
		if err != nil {
			fmt.Println("Invalid input. Please enter an integer value")
			continue
		}
		fmt.Println("You guess is", guess)
		//计算正确数字的个数
		for k := 0; k < 4; k++ {
			y := strings.Count(string(s2), string(s3[k]))
			sum += y
		}
		//计算正确位置的个数
		for i := 0; i < 4; i++ {
			if s2[i] == s3[i] {
				x++
			}
		}
		fmt.Println("The count of right number is", sum)
		fmt.Println("The count of correct location of number is", x)
		//成功时退出程序
		if x == 4 {
			fmt.Println("Correct, you Legend!")
			break
		}
		x = 0
		sum = 0
	}
}

// 函数：排序看相邻的数字是否相同
func findDuplicates(nums []int) bool {
	sort.Ints(nums)
	var i int
	for i = 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			return false
		}
	}
	return true
}
