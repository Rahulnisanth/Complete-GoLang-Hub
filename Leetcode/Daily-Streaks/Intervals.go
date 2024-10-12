package main

import (
    "sort"
)

func minGroups(intervals [][]int) int {
    events := [][]int{}

    for _, interval := range intervals {
        events = append(events, []int{interval[0], 1})
        events = append(events, []int{interval[1] + 1, -1})   
    }

    sort.Slice(events, func(i, j int) bool {
        if events[i][0] == events[j][0] {
            return events[i][1] < events[j][1]
        }
        return events[i][0] < events[j][0]
    })

    currentGroups, maxGroups := 0, 0

    for _, event := range events {
        currentGroups += event[1]
        if currentGroups > maxGroups {
            maxGroups = currentGroups
        }
    }

    return maxGroups
}