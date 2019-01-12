package onitamago

import "math/bits"

const (
	NrOfPieceTypes Amount = 2
	NrOfPlayerPieces Amount = 5

	StudentsIndex Index = 0
	MasterIndex Index = 1
)


// LSB Least Significant Bit
func LSB(x Board) BoardIndex {
	return BoardIndex(bits.TrailingZeros64(x))
}

// NLSB Next Least Significant Bit
func NLSB(x *Board, i BoardIndex) BoardIndex {
	*x ^= 1 << i
	return LSB(*x)
}

func boardIndexToBoard(i BoardIndex) Board {
	return 1 << i
}