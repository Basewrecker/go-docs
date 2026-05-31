// To print
// 3
// 2
// 1
// Go! (1000ms wait between each, iteration and mocking to be used)

package countdown

import (
	"fmt"
	"io"
	"os"
)

// declaring universal consts

const StartingNumber = 3
const FinalWord = "Go!"

type DefaultSleeper struct{}

type SpyCountdownOperations struct {
	Calls []string
}

func (s *SpyCountdownOperations) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *SpyCountdownOperations) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

const write = "write"
const sleep = "sleep"

func Countdown(out io.Writer, sleeper Sleeper) {
	for i := StartingNumber; i > 0; i = i - 1 {
		fmt.Fprintln(out, i)
		sleeper.Sleep()
	}
	fmt.Fprint(out, FinalWord)
}

type Sleeper interface {
	Sleep()
}

func main() {
	sleeper := &DefaultSleeper{}
	Countdown(os.Stdout, sleeper)
}

func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}
