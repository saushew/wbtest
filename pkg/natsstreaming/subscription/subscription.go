package subscription

import (
	"fmt"
	"time"

	"github.com/nats-io/stan.go"
	"github.com/saushew/wb_testtask/pkg/natsstreaming"
)

// Subscription .
type Subscription struct {
	conn  *natsstreaming.Connection
	subs  *stan.Subscription
	error chan error
	stop  chan struct{}

	timeout time.Duration
}

// New .
func New(subject, clusterID, clientID, URL string, hndlr stan.MsgHandler) (*Subscription, error) {
	s := &Subscription{
		error: make(chan error),
		stop:  make(chan struct{}),

		timeout: 2 * time.Second,
	}

	conn, err := natsstreaming.New(clusterID, clientID, URL, s.error)
	if err != nil {
		return nil, fmt.Errorf("subscription - New - natsstreaming.New: %w", err)
	}

	subscription, err := conn.Connection.Subscribe(subject, hndlr)
	if err != nil {
		return nil, fmt.Errorf("subscription - New - conn.Connection.Subscribe: %w", err)
	}

	s.conn = conn
	s.subs = &subscription

	return s, nil
}

// Notify -.
func (s *Subscription) Notify() <-chan error {
	return s.error
}

// Shutdown -.
func (s *Subscription) Shutdown() error {
	select {
	case <-s.error:
		return nil
	default:
	}

	close(s.stop)
	time.Sleep(s.timeout)

	err := s.conn.Connection.Close()
	if err != nil {
		return fmt.Errorf("natsstreaming - Subscription - Shutdown - s.conn.Connection.Close: %w", err)
	}

	return nil
}
