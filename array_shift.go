package main

import "fmt"

func main() {
	arr := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	//var newArr [10]int
	var num int = 4
	var temp int
	var direction string
	fmt.Printf("Position: ")
	fmt.Scanf("%d", &num)
	fmt.Scanf(" %c")
	fmt.Printf("Left or right: ")
	fmt.Scanf("%s", &direction)
	if direction == "left" {
		for i := 0; i < (num); i++ {

			for pos, val := range arr {
				temp = arr[(pos+1)%10]
				arr[(pos+1)%10] = val
				arr[pos] = temp
			}
			//newArr[(10+(pos-num))%10] = val
		}
	}
	fmt.Println(arr)
}

// arr1:=[]int{1,2,3,4,5,6,7,8,9} =>Right{9,1,2,3,4,5,6,7,8} // left{2,3,4,5,6,7,8,9,1}
// 1)pos  4
// 2)direct(left or right)left
// Program should shift the entire array either left or right
//6,7,8,9,0,1,2,3,4,5
//4, 5, 6, 7, 8, 9, 0, 1, 2, 3

//right
