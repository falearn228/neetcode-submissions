/**
 * Definition of Interval:
 * type Interval struct {
 *    start int
 *    end   int
 * }
 */
import (
	"slices"
)

type z struct {
	time int
	delta int
}
// {0, 1} {8, -1}
func minMeetingRooms(intervals []Interval) int {
	events := make([]z, 0, 2*len(intervals))

	for _, it := range intervals {
		events = append(events, z{it.start, 1})
		events = append(events, z{it.end, -1})
	}

	slices.SortFunc(events, func(a, b z) int {
		if a.time == b.time {
			return a.delta - b.delta
		}
		return a.time - b.time
	})

	cur, maxRooms := 0, 0
	for _, e := range events {
		cur += e.delta
		maxRooms = max(maxRooms, cur)
	}
	return maxRooms
}
