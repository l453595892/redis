// +build !go1.7

package redis

import (
	"github.com/l453595892/redis/internal/pool"
)

type baseClient struct {
	connPool pool.Pooler
	opt      *Options

	process func(Cmder) error
	onClose func() error // hook called when client is closed
}
