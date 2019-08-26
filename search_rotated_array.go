// Question :
// https://leetcode.com/problems/search-in-rotated-sorted-array/

// Suppose an array sorted in ascending order is rotated at some pivot unknown to you beforehand.
// (i.e., [0,1,2,4,5,6,7] might become [4,5,6,7,0,1,2]).
// You are given a target value to search. If found in the array return its index, otherwise return -1.
// You may assume no duplicate exists in the array.
// Eg, Input: nums = [4,5,6,7,0,1,2], target = 0
//     Output: 4

// my_algorithm :
// First find the pivot (the point where the array rotates or in other words, the start of the array - the smallest element)
// I find the pivot lineraly - 0(n)
// After I do, I have two chunks of the array - one from 0 to this pivot-1, another is pivot to end 
// My target will be in one of the two chunks if at all it is 
// To figure out which, I just need to compare my target to the index 0 element in the array 
// If target is lesser than that it'll obviously lie in (pivot to end) 
// If target is greater than first element, it obviously lies in (0  to pivot-1)
// So I just do a bin search on the appropriate set in O(log n)

// At the time of writing this solution had a 0ms runtime, beating 100% of go submissions in time, 50% go submissions in space

func findpivot(nums []int) int {
     i := 0
     pivot := -1
	
     for ; i<len(nums)-1 ; i++ {
        if(nums[i] > nums[i+1]) {
            pivot = i+1
        }
    }
    return pivot
}

func binarysearch(arr []int, target int) int {
    l := 0
    r := len(arr) - 1
    p := -1
    
    for l<=r {
        p = (l+r)/2
        if arr[p] == target  {
            return p
        } else if arr[p] < target {
            l = p+1 
        } else {
            r = p-1
        }
    }
    return -1
}

func search(nums []int, target int) int {
    
    if len(nums) == 0 {
        return -1
    }
    
    if len(nums) == 1 {
        if nums[0] == target {
            return 0
        } else {
            return -1
        }
    }
	
	pivot := -1
	pivot = findpivot(nums)
	
  //if an array not rotated (fully sorted in ascending order) we will have a -1 pivot. In which case we can simply send the array for a
  //binary search. And I am safely assuming that the array is not sorted ever in fully descending order. If it did I would need to add
  //code to somehow figure out that its fully desc and need to run like a reverse binary search. 
  
	if pivot == -1 {
		return binarysearch(nums, target)
	}  else {
        if target == nums[0] {
            return 0
        }
        if target < nums[0] {
            //we need to return index from the beginning since target lies in second set - so add pivot 
            if targetindex := binarysearch(nums[pivot:], target) ; targetindex != -1  {
                return targetindex + pivot
            }
        } else if target > nums[0] {
            return binarysearch(nums[0:pivot+1], target)
        } 
    } 
    return -1
}
