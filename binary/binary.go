package binary

import (
	"math/bits"
)

func NewBinaryRep(num int) *BinaryRep {
	return &BinaryRep{uint32(num)}
}

type BinaryRep struct {
	num uint32
}

func (b *BinaryRep) GetBitSumArray() []int {

	bitSumArry := make([]int, b.num)

	for i := range b.num {
		count := bits.OnesCount32(i)
		bitSumArry[i] = count & 1
	}

	return bitSumArry
}
