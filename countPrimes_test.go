package a1
import "testing"

func TestCountPrimes(t *testing.T){
	x := countPrimes(-5)
	if x != 0 {
		t.Errorf("Expected primes to be 0, but it was %d instead.", x)
	}
	y := countPrimes(100)
	if y != 25 {
		t.Errorf("Expected primes to be 25, but it was %d instead.", y)
	}
	z := countPrimes(569)
	if z != 104 {
		t.Errorf("Expected primes to be 25, but it was %d instead.", z)
	}
	u := countPrimes(2)
	if u != 1 {
		t.Errorf("Expected primes to be 1, but it was %d instead.", u)
	}
	v := countPrimes(1)
	if v != 0 {
		t.Errorf("Expected primes to be 0, but it was %d instead.", v)
	} 
}

