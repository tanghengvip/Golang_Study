//golang练习题1：
//实现：输入任意一串以空格分隔的数字，统计这串数字代表的月份所属季度出现的次数，
//并把相应的月份按正序使用英文简写输出出来（同一个月份需要去重，只输出一次）
//例如输入：1 12 6 10 11 3 6
//输出： 第一季度：2次 （Jan、Mar） 第二季度：2次（June） 第三季度：0次（） 第四季度：3次（Oct、Nov、Dec）
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	month := [...]string{
		"NO", "Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"}

	//fmt.Println(month[1])
	Reader := bufio.NewReader(os.Stdin)
	fmt.Println("输入月份，空格分隔")
	input, _ := Reader.ReadString('\n')
	input = strings.TrimSpace(input)
	inputs := strings.Split(input, " ") //根据空格分割，放在inputs字符串数组里面

	cnts1, cnts2, cnts3, cnts4 := 0, 0, 0, 0
	str1, str2, str3, str4 := "", "", "", ""
	for _, num := range inputs {
		numInt, err := strconv.Atoi(num)
		if err != nil {
			fmt.Println("别乱写！", err)
			break
		} else if numInt < 4 {
			cnts1++
			str1 += month[numInt] + " "
		} else if numInt < 7 {
			cnts2++
			str2 += month[numInt] + " "
		} else if numInt < 10 {
			cnts3++
			str3 += month[numInt] + " "
		} else {
			cnts4++
			str4 += month[numInt] + " "
		}
	}
	fmt.Println("第一季度：", cnts1, str1)
	fmt.Println("第二季度：", cnts2, str2)
	fmt.Println("第三季度：", cnts3, str3)
	fmt.Println("第四季度：", cnts4, str4)
}
