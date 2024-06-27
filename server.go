package demo

import (
	"context"
	"sync"
)

type Server interface {
	Run(ctx context.Context, wg *sync.WaitGroup) error
	RegisterHandlers()
}
