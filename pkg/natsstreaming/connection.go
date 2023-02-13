package natsstreaming

import (
	"fmt"
	"time"

	"github.com/nats-io/stan.go"
)

const (
	connectWait        = time.Second * 30
	pubAckWait         = time.Second * 30
	interval           = 10
	maxOut             = 5
	maxPubAcksInflight = 25
)

// Connection .
type Connection struct {
	Connection stan.Conn
}

//New .
func New(clusterID, clientID, URL string, errC chan error) (*Connection, error) {
	sc, err := stan.Connect(
		clusterID,
		clientID,
		stan.ConnectWait(connectWait),
		stan.PubAckWait(pubAckWait),
		stan.NatsURL(URL),
		stan.Pings(interval, maxOut),
		stan.SetConnectionLostHandler(func(_ stan.Conn, reason error) {
			errC <- fmt.Errorf("Connection lost, reason: %v", reason)
		}),
		stan.MaxPubAcksInflight(maxPubAcksInflight),
	)

	if err != nil {
		return nil, fmt.Errorf("natsstreaming - New - stan.Connection: %w", err)
	}

	conn := &Connection{
		Connection: sc,
	}

	return conn, nil
}
