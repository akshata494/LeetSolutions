//Question : https://leetcode.com/problems/two-sum/

//My algorithm : for each num in the array, check if it's compliment (target - num) is in the hash 
//               if yes, return the index set. 
//               if not, add the number's compliemnt to hash 

func twoSum(nums []int, target int) []int {
    compliments := make(map[int]int)
    
    for i, value := range nums {
        if _, ok := compliments[value]; ok {
            j := compliments[value]
            return []int{j, i}
        } else {
            compliments[target-value] = i
        }
    }
    return []int{0,0}
}
