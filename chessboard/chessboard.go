package chessboard

const (
	maxRank = 8
	minRank = 1
)

// Declare a type named File which stores if a square is occupied by a piece - this will be a slice of bools
type File []bool

// Declare a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"
type Chessboard map[string]File

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) int {
	return countOccupied(cb[file])
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
	if rank < minRank || rank > maxRank {
		return 0
	}

	count := 0
	for _, file := range cb {
		if file[rank-1] {
			count++
		}
	}
	return count
}

// CountAll should count how many squares are present in the chessboard.
// We could just return "64", but since this is a for...range exercise,
// we are going to use for...range.
func CountAll(cb Chessboard) int {
	count := 0
	for _, file := range cb {
		for range file {
			count++
		}
	}
	return count
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
	count := 0
	for _, file := range cb {
		count += countOccupied(file)
	}
	return count
}

// helper function
func countOccupied(file File) int {
	count := 0
	for _, occupied := range file {
		if occupied {
			count++
		}
	}
	return count
}
