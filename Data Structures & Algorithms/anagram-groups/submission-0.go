import "slices"

func groupAnagrams(strs []string) [][]string {
    arr := make([][]string, 0, len(strs))
    sets := make(map[string][]string)
    for _, str := range strs {
        s := []rune(str)
        slices.Sort(s)
        _, ok := sets[string(s)]
        if !ok {
            sets[string(s)] = make([]string, 0, len(strs))
        }
        sets[string(s)] = append(sets[string(s)], str)
    }
    for _, v := range sets {
        arr = append(arr, v)
    } 
    return arr
}
