package FiboCache
func fib(n int) int {
    baseCases := map[int]int{
        1: 0,
        2: 1,
    }
    return computeCache(n, baseCases)
}

func computeCache(n int, cache map[int]int) int {
    if val, found := cache[n]; found {
        return val
    }
    cache[n] = computeCache(n-1, cache) + computeCache(n-2, cache)
    return cache[n]
}