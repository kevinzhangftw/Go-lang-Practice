package a1
	
import (
    "fmt"
	// "reflect"
	"strconv"
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
	seq := []int{0}
	seqString := strconv.FormatInt(int64(i), 2)
	
	if len(seqString) == length {
		seq[0], _ = strconv.Atoi(seqString)
		fmt.Println("len(seqString) == length seq[0]: %v", seq[0])
	}else{
		lenDiff := length - len(seqString)
		for i := 0; i < lenDiff; i++ {
			seqString = string('0') + seqString
		}
		fmt.Println("len(seqString) != length seq[0]: %v", seqString)
		seq[0], _ = strconv.Atoi(seqString)

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