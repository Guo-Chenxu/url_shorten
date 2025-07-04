package utils

import (
	"math"
)

// 计算布隆过滤器的参数
func CalculateBloomFilterParams(n int64, p float64) (int64, int64) {
	if n <= 0 || p <= 0 || p >= 1 {
		return 0, 0
	}

	m := int64(-float64(n) * math.Log(p) / (math.Log(2) * math.Log(2)))
	m = findNextPowerOfTwo(m)
	k := int64(math.Ceil(float64(m) / float64(n) * math.Log(2)))
	return m, k
}

func findNextPowerOfTwo(n int64) int64 {
	if n <= 0 {
		return 1
	}
	return 1 << int64(math.Ceil(math.Log2(float64(n))))
}

// 创建哈希函数
func CreateHash(size int64) func(seed int64, value string) int64 {
	return func(seed int64, value string) int64 {
		var result int64 = 0
		for i := 0; i < len(value); i++ {
			result = result*seed + int64(value[i])
		}
		//length = 2^n 时，X % length = X & (length - 1)
		return result & (size - 1)
	}
}

func sieveOfEratosthenes(max int64) []int64 {
	sieve := make([]bool, max+1)
	for i := range sieve {
		sieve[i] = true
	}
	sieve[0], sieve[1] = false, false

	for i := int64(2); i*i <= max; i++ {
		if sieve[i] {
			for j := i * i; j <= max; j += i {
				sieve[j] = false
			}
		}
	}

	var primes []int64
	for i, prime := range sieve {
		if prime {
			primes = append(primes, int64(i))
		}
	}
	return primes
}

// 生成n个素数
func GeneratePrimes(n int64) []int64 {
	if n <= 0 {
		return []int64{}
	}

	// 估计素数的上限
	max := 2 * n * int64(math.Log(float64(n)))
	primes := sieveOfEratosthenes(max)
	for int64(len(primes)) < n {
		max *= 2
		primes = sieveOfEratosthenes(max)
	}

	return primes[:n]
}
