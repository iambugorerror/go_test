package main

import "fmt"

func main() {
    array := [4]int{10, 20, 30, 40}
    slice := array[0:2] // 切片包含数组的前两个元素 {10, 20}
    newSlice := append(slice, 50) // 切片追加元素 {10, 20, 50}
    newSlice[1] += 10 // 修改切片中的第二个元素

    fmt.Println("slice:", slice)     // 输出：slice: [10 30]
    fmt.Println("array:", array)     // 输出：array: [10 30 50 40]
    fmt.Println("newSlice:", newSlice) // 输出：newSlice: [10 30 50]
}
