package dialer_test

import (
	"context"
	"github.com/fx2y/hands-on-software-engineering-with-golang/dialer"
	"github.com/juju/clock"
	"github.com/juju/clock/testclock"
	"log"
	"net"
	"testing"
	"time"
)

func TestRetryingDialerWithRealClock(t *testing.T) {
	log.SetFlags(0)

	// Dial a random local port that nothing is listening on.
	clk := clock.WallClock
	d := dialer.NewRetryingDialer(context.Background(), clk, net.Dial, 10)
	_, err := d.Dial("tcp", "127.0.0.1:65000")
	if err != dialer.ErrMaxRetriesExceeded {
		t.Fatalf("expected to get ErrMaxRetriesExceeded; got %v", err)
	}
}

func TestRetryingDialerWithFakeClock(t *testing.T) {
	log.SetFlags(0)

	doneCh := make(chan struct{})
	defer close(doneCh)
	clk := testclock.NewClock(time.Now())
	go func() {
		for {
			select {
			case <-doneCh: // test completed; exit go-routine
				return
			default:
				clk.Advance(1 * time.Minute)
			}
		}
	}()

	// Dial a random local port that nothing is listening on.
	d := dialer.NewRetryingDialer(context.Background(), clk, net.Dial, 10)
	_, err := d.Dial("tcp", "127.0.0.1:650000")
	if err != dialer.ErrMaxRetriesExceeded {
		t.Fatalf("expected to get ErrMaxRetriesExceeded; got %v", err)
	}
}
