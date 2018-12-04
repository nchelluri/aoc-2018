package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type action uint8

const (
	beginsShift = action(iota)
	fallsAsleep
	wakesUp
)

type log struct {
	timestamp string
	guardNum  uint
	action    action
}

type logs []*log

func (logs logs) Len() int {
	return len(logs)
}

func (logs logs) Swap(i, j int) {
	logs[i], logs[j] = logs[j], logs[i]
}

func (logs logs) Less(i, j int) bool {
	return strings.Compare(logs[i].timestamp, logs[j].timestamp) < 0
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var logs logs

	// [1518-11-01 00:00] Guard #10 begins shift
	// [1518-11-01 00:05] falls asleep
	// [1518-11-01 00:25] wakes up
	logRE := regexp.MustCompile(`\[(\d{4}-\d{2}-\d{2} \d{2}:\d{2})] (Guard #(\d+) begins shift|falls asleep|wakes up)`)
	for scanner.Scan() {
		matches := logRE.FindStringSubmatch(scanner.Text())

		if len(matches) < 3 {
			panic("error parsing log line")
		}

		timestamp := matches[1]
		guardNum := uint(0)

		var action action
		if matches[2] == "falls asleep" {
			action = fallsAsleep
		} else if matches[2] == "wakes up" {
			action = wakesUp
		} else {
			action = beginsShift
			guardNumUint64, err := strconv.ParseUint(matches[3], 10, 0)
			guardNum = uint(guardNumUint64)
			if err != nil {
				panic(err)
			}
		}

		logs = append(logs, &log{
			timestamp: timestamp,
			guardNum:  guardNum,
			action:    action,
		})
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	sort.Sort(logs)

	// > only the minute portion (00 - 59) is relevant for those events.
	// Assumption: The guard always wakes up after falling asleep.
	minutesAsleep := make(map[uint]uint)
	guardNum := uint(0)
	minuteAsleep := uint(0)
	timestampRE := regexp.MustCompile(`\d{4}-\d{2}-\d{2} \d{2}:(\d{2})`)
	for _, log := range logs {
		if log.guardNum != uint(0) {
			guardNum = log.guardNum
		} else {
			log.guardNum = guardNum
		}

		matches := timestampRE.FindStringSubmatch(log.timestamp)
		if len(matches) < 2 {
			panic("error parsing timestamp")
		}

		minuteOfActionUint64, err := strconv.ParseUint(matches[1], 10, 0)
		if err != nil {
			panic(err)
		}
		minuteOfAction := uint(minuteOfActionUint64)

		if log.action == fallsAsleep {
			minuteAsleep = minuteOfAction
		} else if log.action == wakesUp {
			minutesAsleep[log.guardNum] += minuteOfAction - minuteAsleep
		}
	}

	guardNumWithMaxMinutesAsleep := uint(0)
	maxMinutesAsleep := uint(0)
	for guardNum, minutesAsleep := range minutesAsleep {
		if minutesAsleep > maxMinutesAsleep {
			maxMinutesAsleep = minutesAsleep
			guardNumWithMaxMinutesAsleep = guardNum
		}
	}

	maxAsleepGuardMinutesAsleep := [60]uint{}
	for _, log := range logs {
		if log.guardNum != guardNumWithMaxMinutesAsleep {
			continue
		}

		matches := timestampRE.FindStringSubmatch(log.timestamp)
		if len(matches) < 2 {
			panic("error parsing timestamp")
		}

		minuteOfActionUint64, err := strconv.ParseUint(matches[1], 10, 0)
		if err != nil {
			panic(err)
		}
		minuteOfAction := uint(minuteOfActionUint64)

		if log.action == fallsAsleep {
			minuteAsleep = minuteOfAction
		} else if log.action == wakesUp {
			for i := minuteAsleep; i < minuteOfAction; i++ {
				maxAsleepGuardMinutesAsleep[i]++
			}
		}
	}

	minuteMostCommonlyAsleepSleepTimes := uint(0)
	minuteMostCommonlyAsleepIndex := 0

	for minute, numTimesAsleep := range maxAsleepGuardMinutesAsleep {
		if numTimesAsleep > minuteMostCommonlyAsleepSleepTimes {
			minuteMostCommonlyAsleepSleepTimes = numTimesAsleep
			minuteMostCommonlyAsleepIndex = minute
		}
	}

	fmt.Println(uint(minuteMostCommonlyAsleepIndex) * guardNumWithMaxMinutesAsleep)
}
