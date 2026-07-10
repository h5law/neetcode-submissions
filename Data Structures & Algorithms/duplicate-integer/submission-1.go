func hasDuplicate(nums []int) bool {
    d := make(map[int]int, len(nums))
    for _, v := range nums {
        d[v] += 1
        if d[v] > 1 {
            return true
        }
    }
    return false
}
