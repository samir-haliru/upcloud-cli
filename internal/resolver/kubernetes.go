package resolver

import (
	"context"

	"github.com/UpCloudLtd/upcloud-cli/v3/internal/service"

	"github.com/UpCloudLtd/upcloud-go-api/v8/upcloud/request"
)

// CachingKubernetes implements resolver for Kubernetes clusters, caching the results
type CachingKubernetes struct{}

// make sure we implement the ResolutionProvider interface
var _ ResolutionProvider = CachingKubernetes{}

// Get implements ResolutionProvider.Get
func (s CachingKubernetes) Get(ctx context.Context, svc service.AllServices) (Resolver, error) {
	clusters, err := svc.GetKubernetesClusters(ctx, &request.GetKubernetesClustersRequest{})
	if err != nil {
		return nil, err
	}
	return func(arg string) Resolved {
		rv := Resolved{Arg: arg}
		for _, cluster := range clusters {
			rv.AddMatch(cluster.UUID, MatchTitle(arg, cluster.Name))
			rv.AddMatch(cluster.UUID, MatchUUID(arg, cluster.UUID))
		}
		return rv
	}, nil
}

// PositionalArgumentHelp implements resolver.ResolutionProvider
func (s CachingKubernetes) PositionalArgumentHelp() string {
	return "<UUID/Name...>" //nolint:goconst
}
