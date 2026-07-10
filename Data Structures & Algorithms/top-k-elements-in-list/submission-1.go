func topKFrequent(nums []int, k int) []int {
    count := make(map[int]int, len(nums))
    freq := make([][]int, len(nums)+1)
    arr := make([]int, 0, k)
    for _, n := range nums {
        count[n]++
    }
    for num, cnt := range count {
        freq[cnt] = append(freq[cnt], num)
    }
    for i := len(freq) - 1; i > 0; i-- {
        for _, num := range freq[i] {
            arr = append(arr, num)
            if len(arr) == k {
                return arr
            }
        }
    }
    return arr
}
