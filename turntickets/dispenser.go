package turntickets

import (
	"sync/atomic"
)

type TurnNumberSequence struct {
	n int64
}

func (s *TurnNumberSequence) NextTurnNumber() int64 {
	atomic.AddInt64(&s.n, 1)
	return s.n
}

type Ticket struct {
	number int64
}

type TicketDispenser struct {
	seq *TurnNumberSequence
}

func (td TicketDispenser) GetTurnTicket() *Ticket {
	return &Ticket{number: td.seq.NextTurnNumber()}
}
