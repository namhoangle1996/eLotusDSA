package main

import "fmt"

// 1 <= n <= 16
// calculate XOR https://xor.pw/#
// Time complexity : O(2^n)
func grayCode(n int) (res []int) {
	// loop i := 0 => 2^n - 1
	for i := 0; i < (1 << n); i++ {

		// XOR i and ( i divided by 2, i times ( i>>1) )
		res = append(res, i^(i>>1))
	}

	return res
}

// Input: nums1 = [0,0,0,0,0], nums2 = [0,0,0,0,0] Output: 5
// Time complexity : O(n * m),
func findLength(nums1 []int, nums2 []int) (res int) {

	// Create a slice that store repeated sub-slice
	n := len(nums1)
	m := len(nums2)
	tmpSlice := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		tmpSlice[i] = make([]int, m+1)
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if nums1[i-1] == nums2[j-1] {

				// If two elements are equal,
				// increase the length of the common segment by 1 and update the tmpSlice at position [i][j]
				tmpSlice[i][j] = tmpSlice[i-1][j-1] + 1

				// increase the response
				if tmpSlice[i][j] > res {
					res = tmpSlice[i][j]
				}
			} else {
				// If the values of the two elements differ ,
				// find the table that has the largest value at [i-1][j] and [i][j-1] places,
				// then update the tmpSlice
				tmpSlice[i][j] = getMaxInt(tmpSlice[i-1][j], tmpSlice[i][j-1])
			}
		}
	}

	return res
}

func getMaxInt(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func sumOfDistancesInTree(n int, edges [][]int) (res []int) {
	// I have not been able to resolve this issue :D
	return
}

func main() {
	fmt.Println("get gray code ", grayCode(3))
	fmt.Println("get maximum length of repeated subarray  ", findLength([]int{0, 1, 2}, []int{1, 2, 3, 0}))
	fmt.Println("find sumOfDistancesInTree", sumOfDistancesInTree(2, [][]int{}))

}
