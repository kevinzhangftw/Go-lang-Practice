package a1

// adapted from my past C++ homework https://github.com/kevinzhangftw/User-Account-Management-Hash-Table/blob/master/UserAccountManagement/hashtable.cpp
func countPrimes(n int) int {
	result := 0
	for i := 2; i <= n; i++ {
		if isPrime(i)==true {
			result++
		}
	}
	return result
}

func isPrime(n int) bool{
	result := true
	for i := 2; i < n; i++ {
		if n % i == 0 {
			result = false
			return result
		}	
	}
	return result
}