package hot100

import (
	"fmt"
	"testing"
)

func TestTwoSum(t *testing.T) {
	nums :=[]int{3,2,4}
	ans := TwoSum(nums,6)
	fmt.Println(ans)
}
