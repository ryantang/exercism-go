package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	seen := 0
	for i := 0; i < len(birdsPerDay); i++ {
		seen += birdsPerDay[i]
	}
	return seen
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	seen := 0
	startOfWeek := 7 * (week - 1)

	for i := startOfWeek; i < startOfWeek+7; i++ {
		seen += birdsPerDay[i]
	}
	return seen
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	for i := 0; i < len(birdsPerDay); i += 2 {
		birdsPerDay[i]++
	}
	return birdsPerDay
}
