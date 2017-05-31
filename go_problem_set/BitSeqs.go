package a1
	
import (
    // "fmt"
	// "reflect"
	// "strconv"
	"math"
)

func allBitSeqs(n int) [][]int{
	if n<=0 {
		emptySeqs := make([][]int, 0)
		return emptySeqs
	}

	length := int(math.Pow(2, float64(n))) 
	seqs := make([][]int, length)

	for i := 0; i < length; i++ {
		seqs[i]=getSequence(i, n)
	}

	return seqs
}

func getSequence(i int, length int) []int {
	seq := make([]int, length)
	for x := 0; x < length; x++ {
		bitIndex := x+1
		largestBitSize := int(math.Pow(2, float64(length-bitIndex)))

		// fmt.Println("largestBitSize:%v", largestBitSize)
		// fmt.Println("i:%v", i)

		if largestBitSize <= i {
			seq[x]=1
			i = i - largestBitSize
		}else{
			seq[x]=0
		}
	}
	return seq
}

func isBitSeq(seq []int) bool {
    for _, b := range seq {
        if !(b == 0 || b == 1) {
            return false
        }
    }
    return true
}