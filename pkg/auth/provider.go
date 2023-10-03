package auth

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/pkg/errors"
	"github.com/redis/go-redis/v9"
)

// Provider is an interface for client(RHID) authentication
type Provider interface {
	ListClients(ctx context.Context, gateway string) ([]string, error)
	AddClient(ctx context.Context, gateway, rhid string, policy *ClientPolicy) error
	DeleteClient(ctx context.Context, gateway, rhid string) error
}

type provider struct {
	redisCli *redis.Client
}

func hash(s string) string {
	bs := sha256.Sum256([]byte(s))

	return hex.EncodeToString(bs[:])
}

func (p provider) ListClients(ctx context.Context, gateway string) ([]string, error) {
	//TODO implement me
	panic("implement me")
}

func (p provider) AddClient(ctx context.Context, gateway, rhid string, policy *ClientPolicy) error {
	gwHash := hash(gateway)
	if err := p.redisCli.LPush(ctx, gwHash, rhid).Err(); err != nil {
		return errors.WithStack(err)
	}

	key := fmt.Sprintf("%s-%s", gwHash, rhid)
	if err := p.redisCli.Set(ctx, key, policy.toOpenNDSFormat(), 0).Err(); err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (p provider) DeleteClient(ctx context.Context, gateway, rhid string) error {
	//TODO implement me
	panic("implement me")
}

func NewProvider(redisCli *redis.Client) Provider {
	return &provider{redisCli: redisCli}
}
