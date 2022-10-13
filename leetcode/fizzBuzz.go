func fizzBuzz(n int) []string {
	sli := make([]int, n)
	sls := make([]string, n)
	for i := 0; i < n; i++ {
		sli[i] = i + 1
		sls[i] = strconv.Itoa(i + 1)
		if sli[i]%3 == 0 && sli[i]%5 == 0 {
			sls[i] = "FizzBuzz"
		} else if sli[i]%3 == 0 {
			sls[i] = "Fizz"
		} else if sli[i]%5 == 0 {
			sls[i] = "Buzz"
		}
	}
	return sls
}
