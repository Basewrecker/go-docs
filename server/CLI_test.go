package poker

import (
	"fmt"
	poker "g-test/server"
	"strings"
	"testing"
	"time"
)

func TestCLI(t *testing.T) {
	t.Run("record chris win from user input", func(t *testing.T) {
		in := strings.NewReader("Chris wins\n")
		playerStore := &StubPlayerStore{}

		cli := &CLI{playerStore, in}
		cli.PlayPoker()

		assertPlayerWin(t, playerStore, "Chris")
	})

	t.Run("record cleo win from user input", func(t *testing.T) {
		in := strings.NewReader("Cleo wins \n")
		playerStore := &StubPlayerStore{}

		cli := &CLI{playerStore, in}
		cli.PlayPoker()
		assertPlayerWin(t, playerStore, "CLEO")
	})

	t.Run("It schedules printing of blind values", func(t *testing.T) {
		in := strings.NewReader("Chris wins\n")
		playerStore := &poker.StubPlayerStore{}
		blindAlerter := &SpyBlindAlerter{}

		cli := poker.NewCLI(playerstore, in, blindAlerter)
		cli.playPoker()

		if len(blindAlerter.alerts) != 1 {
			t.Fatal("expected a blind alert to be passed or scheduled")
		}

		t.Run("it schedules printing of blind values", func(t *testing.T) {
			in := strings.NewReader("Chris wins\n")
			playerStore := &poker.StubPlayerStore{}
			blindAlerter := &SpyBlindAlerter{}

			cli := poker.NewCLI(playerStore, in, blindAlerter)
			cli.PlayPoker()

			cases := []struct {
				expectedScheduleTime time.Duration
				expectedAmount       int
			}{
				{0 * time.Second, 100},
				{10 * time.Minute, 200},
				{20 * time.Minute, 300},
				{30 * time.Minute, 400},
				{40 * time.Minute, 500},
				{50 * time.Minute, 600},
				{60 * time.Minute, 800},
				{70 * time.Minute, 1000},
				{80 * time.Minute, 2000},
				{90 * time.Minute, 4000},
				{100 * time.Minute, 8000},
			}

			for i, c := range cases {
				t.Run(fmt.Sprintf("%d scheduled for %v", c.expectedAmount, c.expectedScheduleTime), func(t *testing.T) {

					if len(blindAlerter.alerts) <= i {
						t.Fatalf("alert %d was not scheduled %v", i, blindAlerter.alerts)
					}

					alert := blindAlerter.alerts[i]

					amountGot := alert.amount
					if amountGot != c.expectedAmount {
						t.Errorf("got amount %d, want %d", amountGot, c.expectedAmount)
					}

					gotScheduledTime := alert.scheduledAt
					if gotScheduledTime != c.expectedScheduleTime {
						t.Errorf("got scheduled time of %v, want %v", gotScheduledTime, c.expectedScheduleTime)
					}
				})
			}
		})
	}),
	t.Run("it prompts the user to enter the number of players", func(t *testing.T) {
		stdout := &bytes.Buffer{}
		cli := poker.NewCLI(dummyPlayerStore, dummyStdIn, stdout, dummyBlindAlerter)
		cli.PlayPoker()

		got := stdout.String()
		want := "Please enter the number of players: "

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}
