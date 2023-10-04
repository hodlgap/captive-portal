//go:generate mockgen -typed -package=auth_mock -destination=mock/provider.go . Provider

package auth

import (
	"context"
	"fmt"

	"github.com/newrelic/go-agent/v3/newrelic"
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
	return popAllClients(newrelic.NewContext(ctx, newrelic.FromContext(ctx)), p.redisCli, gatewayHash)
}

func popAllClients(ctx context.Context, redisCli *redis.Client, gatewayHash string) ([]string, error) {
	rhids, err := listClients(ctx, redisCli, gatewayHash)
	if err != nil {
		return nil, err
	}

	return rhids, deleteGateway(ctx, redisCli, gatewayHash)
}

func (p *provider) ListClients(ctx context.Context, gatewayHash string) ([]string, error) {
	return listClients(newrelic.NewContext(ctx, newrelic.FromContext(ctx)), p.redisCli, gatewayHash)
}

func listClients(ctx context.Context, redisCli *redis.Client, gatewayHash string) ([]string, error) {
	rhids, err := redisCli.LRange(ctx, gatewayHash, 0, -1).Result()
	if err != nil && !errors.Is(err, redis.Nil) {
		return nil, errors.WithStack(err)
	}

	return rhids, nil
}

func (p *provider) AddPolicy(ctx context.Context, gatewayHash, rhid string, policy *ClientPolicy) error {
	return addPolicy(newrelic.NewContext(ctx, newrelic.FromContext(ctx)), p.redisCli, gatewayHash, rhid, policy)
}

func addPolicy(ctx context.Context, redisCli *redis.Client, gatewayHash, rhid string, policy *ClientPolicy) error {
	if err := redisCli.LPush(ctx, gatewayHash, rhid).Err(); err != nil {
		return errors.WithStack(err)
	}

	return errors.WithStack(redisCli.Set(ctx, toClientKey(gatewayHash, rhid), policy.toOpenNDSFormat(), 0).Err())
}

func (p *provider) DeletePolicies(ctx context.Context, gatewayHash string, rhids ...string) error {
	if len(rhids) == 0 {
		return nil
	}

	return deletePolicies(newrelic.NewContext(ctx, newrelic.FromContext(ctx)), p.redisCli, gatewayHash, rhids...)
}

func deletePolicies(ctx context.Context, redisCli *redis.Client, gatewayHash string, rhids ...string) error {
	if len(rhids) == 0 {
		return nil
	}

	for i, rhid := range rhids {
		rhids[i] = toClientKey(gatewayHash, rhid)
	}

	result, err := redisCli.Del(ctx, rhids...).Result()
	if err != nil {
		return errors.WithStack(err)
	}
	if result != int64(len(rhids)) {
		return errors.Errorf("failed to delete all clients in %s. expected: %d, deleted: %d", gatewayHash, len(rhids), result)
	}

	return nil
}

func (p *provider) DeleteGateway(ctx context.Context, gatewayHash string) error {
	return deleteGateway(newrelic.NewContext(ctx, newrelic.FromContext(ctx)), p.redisCli, gatewayHash)
}

func deleteGateway(ctx context.Context, redisCli *redis.Client, gatewayHash string) error {
	return errors.WithStack(redisCli.Del(ctx, gatewayHash).Err())
}

func (p *provider) ListPolicies(ctx context.Context, gatewayHash string, rhids ...string) ([]string, error) {
	if len(rhids) == 0 {
		return nil, nil
	}

	return listPolicies(newrelic.NewContext(ctx, newrelic.FromContext(ctx)), p.redisCli, gatewayHash, rhids...)
}

func listPolicies(ctx context.Context, redisCli *redis.Client, gatewayHash string, rhids ...string) ([]string, error) {
	if len(rhids) == 0 {
		return nil, nil
	}

	for i, rhid := range rhids {
		rhids[i] = toClientKey(gatewayHash, rhid)
	}

	policies, err := redisCli.MGet(ctx, rhids...).Result()
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
