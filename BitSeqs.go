package a1
	
import (
    // "fmt"
	// "reflect"
)

func allBitSeqs(n int) [][]int{
	seqs := make([][]int, 0)
	if n<=0 {
		return seqs
	}




	return seqs
}

func isBitSeq(seq []int) bool {
    for _, b := range seq {
        if !(b == 0 || b == 1) {
            return false
        }
    }
    return true
}