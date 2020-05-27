package sysbench

import (
	"errors"
	"os/exec"
	"regexp"
	"strconv"
)

func RunCPUTest(sysbenchPath string, args... string) (float64, error) {
	args = append([]string{"cpu", "run"}, args...)

	cmd := exec.Command(sysbenchPath, args...)
	output, err := cmd.Output()
	if err != nil {
		return 0, err
	}

	return extractEventsPerSecond(output)
}

var eventsPerSecondRegexp = regexp.MustCompile(`events per second:\s+(\S.+)`)

func extractEventsPerSecond(output []byte) (float64, error) {
	matches := eventsPerSecondRegexp.FindSubmatch(output)
	if len(matches) < 2 {
		return 0, errors.New("unable to extract events per second from output")
	}
	return strconv.ParseFloat(string(matches[1]), 64)
}