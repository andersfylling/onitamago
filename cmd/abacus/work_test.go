package main

import (
	"testing"

	"github.com/andersfylling/onitamago"
)

func TestWorkers(t *testing.T) {
	cards := onitamago.Deck(onitamago.DeckOriginal)
	configs := onitamago.GenCardConfigs(cards)

	startWork(3, 6, configs)

}