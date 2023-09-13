package main

// import "fmt"

// func main() {
//     slice := []int{1, 2, 3, 4, 5}

//     for index, value := range slice {
//         fmt.Printf("Before - Index: %d, Value: %d, Address: %p\n", index, value, &value)
//     }

//     for index := range slice {
//         slice[index] += 10
//     }

//     for index, value := range slice {
//         fmt.Printf("After - Index: %d, Value: %d, Address: %p\n", index, value, &value)
//     }
// }
// /"在上面的例子中，我们首先使用 range 遍历切片 slice，并打印出每个值的地址。然后，我们对切片中的每个元素进行加 10 的操作。
// 在第二次使用 range 遍历切片时，我们再次打印出每个值的地址。可以观察到，重新赋值后，每个值的地址发生了变化，说明重新赋值后的值被分配到了不同的内存地址上。
// 需要注意的是，即使值的地址发生了变化，但切片中的元素仍然被修改了。这是因为我们通过索引访问和修改了切片中的元素，而不是仅仅修改迭代得到的值的拷贝。"/package main

import "fmt"

func main() {
    slice := []int{1, 2, 3, 4, 5}

    for index, value := range slice {
        fmt.Printf("Before - Index: %d, Value: %d, Address: %p\n", index, value, &value)
    }

    slice[2] += 10

    for index, value := range slice {
        fmt.Printf("After - Index: %d, Value: %d, Address: %p\n", index, value, &value)
    }
}
