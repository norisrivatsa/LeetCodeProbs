func moveZeroes(nums []int)  {
    l := 0
    r := 0
    temp := 0
    if len(nums) == 1 {
        return 
    }
    for l < len(nums)  {
        
        if nums[l] == 0 {
            if r >= len(nums) {
                    break
            }else if nums[r] == 0{
                    r = r + 1 
                    continue
            } else {
                temp = nums[r]
                nums[r] = nums[l]
                nums[l] = temp
                l = l + 1
                r = l + 1
            }
        }  else {
            l = l + 1
            r = l + 1
        }
        
    }
    return 
}