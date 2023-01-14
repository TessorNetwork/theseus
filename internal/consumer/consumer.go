// Package consumer contains interface of blocks consumer.
package consumer

import (
	"context"

	"github.com/TessorNetwork/go-api/health"
)

//go:generate mockgen -destination=./mock/consumer.go -package=consumer -source=consumer.go

// Consumer consumes blocks from furya blockchain.
type Consumer interface {
	health.Pinger

	Run(ctx context.Context) error
}
