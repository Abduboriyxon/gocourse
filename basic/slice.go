package basic

import (
	"fmt"
	"slices"
)

func main() {

	// var sliceName []type
	// var numbers []int
	// var numbers1 = []int{1,2,3}
	// numbers2 := []int{9,8,7}

	// slices := make([]int,5)
	a := [5]int{1,2,3,4,5}
	slices1 := a[1:4] // slicing an array
	fmt.Println(slices1)

	slices1= append(slices1, 6,7,8) // appending to a slice
	fmt.Println("Slice1:",slices1)

	sliceCopy := make([]int, len(slices1)) // copying a slice
	copy(sliceCopy, slices1)
	fmt.Println("CopySlice:",sliceCopy)

	var nilSlice []int // nil slice
	fmt.Println("NilSlice:",nilSlice, len(nilSlice), cap(nilSlice))

	for i, v := range slices1 {
		fmt.Printf("Index: %d, Value: %d\n", i, v)
	}

	// fmt.Println("Elemetn at index 3:", slices1[3])
	// slices1[3] = 100
	// fmt.Println("After updating index 3:", slices1[3])

	if slices.Equal(slices1, sliceCopy) {
		fmt.Println("slices1 and sliceCopy are equal")
	} else {
		fmt.Println("slices1 and sliceCopy are not equal")
	}

	twoD := make([][]int,3)
	for i:=0; i<3;i++ {
		innerlen:=i+1
		twoD[i] = make([]int, innerlen)
		for j:=0; j<innerlen; j++ {
			twoD[i][j] = i + j
			fmt.Printf("Adding value %d in outer slice at index %d, and in inner slice index of %d\n",i+j,i,j)
		}
	}
	fmt.Println("2D Slice:", twoD)

	slice2 := slices1[2:4]
	fmt.Println(slice2)
	fmt.Println("The capasity of slice2 is:", cap(slice2))
	fmt.Println("The length of slice2 is:", len(slice2))

}