package poker

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
	"time"
)

var dummySpyAlerter = &SpyBlindAlerter{}

const PlayerPrompt = "Please enter the amt of players"

type CLI struct {
	playerStore PlayerStore
	in          *bufio.Scanner
	out         io.Writer
	alerter     BlindAlerter
}

type SpyBlindAlerter struct {
	alerts []struct {
		scheduledAt time.Duration
		amount      int
	}
}

type BlindAlerter interface {
	ScheduleAlertAt(duration time.Duration, amount int)
}

func NewCLI(store PlayerStore, in io.Reader, out io.Writer, alerter BlindAlerter) *CLI {
	return &CLI{
		playerStore: store,
		in:          bufio.NewScanner(in),
		out:         out,
		alerter:     alerter,
	}
}

func NewGame(alerter BlindAlerter, store PlayerStore) *Game {
	return &Game{
		alerter: alerter,
		store:   store,
	}
}
func (cli *CLI) PlayPoker() {
	fmt.Fprint(cli.out, PlayerPrompt)
	numberOfPlayers, _ := strconv.Atoi(cli.readLine())
	cli.scheduleBlindAlerts(numberOfPlayers)
	userInput := cli.readLine()
	cli.playerStore.RecordWin(extractWinner(userInput))
}

func (cli *CLI) scheduleBlindAlerts(numberOfPlayers int) {
	blindIncrement := time.Duration(5+numberOfPlayers) * time.Minute

	blinds := []int{100, 200, 300, 400, 500, 600, 800, 1000, 2000, 4000, 8000}
	blindTime := 0 * time.Second
	for _, blind := range blinds {
		cli.alerter.ScheduleAlertAt(blindTime, blind)
		blindTime = blindTime + blindIncrement
	}
}

func extractWinner(userInput string) string {
	return strings.Replace(userInput, "wins", "", 1)
}
