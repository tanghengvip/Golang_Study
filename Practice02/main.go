//练习题
//定义一个三角形struct，并实现以下方法：
//（方法1）判断三角形是否存在、
//（方法2）打印三角形的周长和面积。
//创建该三角形的实例对象，并调用以上两个方法。
//（额外要求：需定义一个计算n边形周长的通用方法，可以用来嵌套到三角形struct中的方法2来使用，此额外要求的具体实现方法可参考视频课程13节的内容）
package main

import (
	"fmt"
	"math"
)

type triangle struct {
	A, B, C float64
}

func main() {
	triangle := triangle{1, 3, 3}
	triangle.TriangleCheck()
}
func (striange triangle) TriangleCheck() {
	if striange.IsNotStriange() {
		fmt.Println("不是三角形")
	} else {
		Perimeter, Area := striange.result()
		fmt.Printf("周长 = %v\n面积 = %v", Perimeter, Area)
	}
}

func (striange triangle) IsNotStriange() bool { //非三角形判断
	if striange.A+striange.B <= striange.C && striange.A+striange.C <= striange.B && striange.B+striange.C <= striange.A {
		return true
	}
	return false
}
func (striange triangle) result() (float64, float64) {
	var Perimeter = striange.A + striange.B + striange.C //三角形周长计算
	var p = (Perimeter) / 2
	var Area = math.Sqrt(p * (p - striange.A) * (p - striange.B) * (p - striange.C)) //三角形面积计算
	return Perimeter, Area
}
