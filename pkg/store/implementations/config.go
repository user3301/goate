package implementations

import (
	"context"
	"fmt"

	"github.com/user3301/goate/pkg/store/implementations/mysqlstore"

	"github.com/user3301/goate/pkg/store/implementations/localstore"

	"github.com/user3301/goate/pkg/store"
)

type Config struct {
	MysqlConfig *mysqlstore.Config `yaml:"mysql, omitempty"`
}

func (c Config) StoreFromConfig(ctx context.Context) (store.UserStorer, error) {
	if c.MysqlConfig == nil {
		return localstore.NewLocalUserStore(), nil
	}
	switch {
	case c.MysqlConfig != nil:
		return c.MysqlConfig.MySQLStore()
	default:
		return nil, fmt.Errorf("unknown store type %v", c)
	}
}
