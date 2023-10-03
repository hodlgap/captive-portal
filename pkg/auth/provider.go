//go:generate mockgen -typed -package=auth_mock -destination=mock/provider.go . Provider

package auth

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	"github.com/redis/go-redis/v9"
)

// Provider is an interface for client(RHID) authentication
type Provider interface {
	PopAllClients(ctx context.Context, gatewayHash string) ([]string, error)
	ListClients(ctx context.Context, gatewayHash string) ([]string, error)
	AddPolicy(ctx context.Context, gatewayHash, rhid string, policy *ClientPolicy) error
	DeletePolicies(ctx context.Context, gatewayHash string, rhids ...string) error
	DeleteGateway(ctx context.Context, gatewayHash string) error
	ListPolicies(ctx context.Context, gatewayHash string, rhids ...string) ([]string, error)
}

type provider struct {
	redisCli *redis.Client
}

func toClientKey(gatewayHash, rhid string) string {
	return fmt.Sprintf("%s-%s", gatewayHash, rhid)
}

func (p *provider) PopAllClients(ctx context.Context, gatewayHash string) ([]string, error) {
	rhids, err := p.ListClients(ctx, gatewayHash)
	if err != nil {
		return nil, err
	}

	return rhids, p.DeleteGateway(ctx, gatewayHash)
}

func (p *provider) ListClients(ctx context.Context, gatewayHash string) ([]string, error) {
	rhids, err := p.redisCli.LRange(ctx, gatewayHash, 0, -1).Result()
	if err != nil && !errors.Is(err, redis.Nil) {
		return nil, errors.WithStack(err)
	}

	return rhids, nil
}

func (p *provider) AddPolicy(ctx context.Context, gatewayHash, rhid string, policy *ClientPolicy) error {
	if err := p.redisCli.LPush(ctx, gatewayHash, rhid).Err(); err != nil {
		return errors.WithStack(err)
	}

	return errors.WithStack(p.redisCli.Set(ctx, toClientKey(gatewayHash, rhid), policy.toOpenNDSFormat(), 0).Err())
}

func (p *provider) DeletePolicies(ctx context.Context, gatewayHash string, rhids ...string) error {
	if len(rhids) == 0 {
		return nil
	}

	for i, rhid := range rhids {
		rhids[i] = toClientKey(gatewayHash, rhid)
	}

	result, err := p.redisCli.Del(ctx, rhids...).Result()
	if err != nil {
		return errors.WithStack(err)
	}
	if result != int64(len(rhids)) {
		return errors.Errorf("failed to delete all clients in %s. expected: %d, deleted: %d", gatewayHash, len(rhids), result)
	}

	return nil
}

func (p *provider) DeleteGateway(ctx context.Context, gatewayHash string) error {
	return errors.WithStack(p.redisCli.Del(ctx, gatewayHash).Err())
}

func (p *provider) ListPolicies(ctx context.Context, gatewayHash string, rhids ...string) ([]string, error) {
	if len(rhids) == 0 {
		return nil, nil
	}

	for i, rhid := range rhids {
		rhids[i] = toClientKey(gatewayHash, rhid)
	}

	policies, err := p.redisCli.MGet(ctx, rhids...).Result()
	if err != nil && !errors.Is(err, redis.Nil) {
		return nil, errors.WithStack(err)
	}

	strPolicies := make([]string, len(policies))
	for i, policy := range policies {
		strPolicies[i] = policy.(string)
	}

	return strPolicies, nil
}

func NewProvider(redisCli *redis.Client) Provider {
	return &provider{redisCli: redisCli}
}
