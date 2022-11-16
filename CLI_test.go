package poker

import (
	"strings"
	"testing"
)

func TestCLI(t *testing.T) {
  in := strings.NewReader("Chris wins\n")
	playerStore := &StubPlayerStore{}

	cli := &CLI{playerStore, in}
	cli.PlayPoker()

  assertPlayerWin(t, playerStore, "Chris")
}

func assertPlayerWin(t testing.TB, store *StubPlayerStore, winner string) {
  t.Helper()

	if len(store.winCalls) != 1 {
		t.Fatalf("got %d calls to RecordWin want %d", len(store.winCalls), 1)
	}

	if store.winCalls[0] != winner {
		t.Errorf("didn't record correct winner got %q want %q", store.winCalls[0], winner)
	}
}