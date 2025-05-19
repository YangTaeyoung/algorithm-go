package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

type Schedule struct {
	Start int
	End   int
}

type Schedules []Schedule

// https://www.acmicpc.net/problem/1931 회의실 배정

func main() {
	writer := bufio.NewWriter(os.Stdout)
	reader := bufio.NewReader(os.Stdin)
	defer writer.Flush()

	var N int

	if _, err := fmt.Fscan(reader, &N); err != nil {
		log.Fatal(err)
	}

	schedules := make(Schedules, N)
	for i := 0; i < N; i++ {
		if _, err := fmt.Fscan(reader, &schedules[i].Start, &schedules[i].End); err != nil {
			log.Fatal(err)
		}
	}

	sort.Slice(schedules, func(i, j int) bool {
		if schedules[i].End == schedules[j].End {
			// 시작점과 끝점이 같은 경우가 있기 때문에, 무시하지 않기 위해 추가함. 1,4, 4,4가 있는 경우 4,4 가 먼저 오면 1,4는 넣지 못하게 된다.
			return schedules[i].Start < schedules[j].Start
		}
		return schedules[i].End < schedules[j].End
	})

	maximumSchedules := make(Schedules, 0)
	for _, schedule := range schedules {
		if len(maximumSchedules) > 0 {
			lastSchedule := maximumSchedules[len(maximumSchedules)-1]
			if lastSchedule.End > schedule.Start {
				continue
			}
		}

		maximumSchedules = append(maximumSchedules, schedule)
	}

	if _, err := fmt.Fprintln(writer, len(maximumSchedules)); err != nil {
		log.Fatal(err)
	}
}
