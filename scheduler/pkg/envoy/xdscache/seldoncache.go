package xdscache

import (
	"fmt"

	"github.com/envoyproxy/go-control-plane/pkg/cache/types"
	"github.com/seldonio/seldon-core/scheduler/pkg/envoy/resources"
)

const (
	DefaultListenerAddress        = "0.0.0.0"
	DefaultListenerPort    uint32 = 9000
)

type SeldonXDSCache struct {
	Listeners map[string]resources.Listener
	Routes    map[string][]resources.Route
	Clusters  map[string]resources.Cluster
}

func NewSeldonXDSCache() *SeldonXDSCache {
	return &SeldonXDSCache{
		Listeners: make(map[string]resources.Listener),
		Clusters:  make(map[string]resources.Cluster),
		Routes:    make(map[string][]resources.Route),
	}
}

func (xds *SeldonXDSCache) ClusterContents() []types.Resource {
	var r []types.Resource

	for _, c := range xds.Clusters {
		endpoints := make([]resources.Endpoint, 0, len(c.Endpoints))
		for _, value := range c.Endpoints { // Likely to be small (<100?) as is number of model replicas
			endpoints = append(endpoints, value)
		}
		r = append(r, resources.MakeCluster(c.Name, endpoints, c.Grpc))
	}

	return r
}

func (xds *SeldonXDSCache) RouteContents() []types.Resource {

	var routesArray []resources.Route
	for _, r := range xds.Routes { //This could be very large as is equal to number of models (100k?)
		routesArray = append(routesArray, r...)
	}

	return []types.Resource{resources.MakeRoute(routesArray)}
}

func (xds *SeldonXDSCache) ListenerContents() []types.Resource {
	var r []types.Resource

	for _, l := range xds.Listeners {
		r = append(r, resources.MakeHTTPListener(l.Name, l.Address, l.Port))
	}

	return r
}

//Note: We don;t use endpoints at present as Envoy does not allow strict_dns with EDS
func (xds *SeldonXDSCache) EndpointsContents() []types.Resource {
	var r []types.Resource

	for _, c := range xds.Clusters {
		endpoints := make([]resources.Endpoint, 0, len(c.Endpoints))
		for _, value := range c.Endpoints {
			endpoints = append(endpoints, value)
		}
		r = append(r, resources.MakeEndpoint(c.Name, endpoints))
	}

	return r
}

func (xds *SeldonXDSCache) AddListener(name string) {
	xds.Listeners[name] = resources.Listener{
		Name:    name,
		Address: DefaultListenerAddress,
		Port:    DefaultListenerPort,
	}
}

func (xds *SeldonXDSCache) AddRoute(name, modelName string, httpClusterName string, grpcClusterName string, logPayloads bool, trafficPercent uint32, version uint32) {
	xds.Routes[name] = append(xds.Routes[name], resources.Route{
		Name:           name,
		Host:           modelName,
		HttpCluster:    httpClusterName,
		GrpcCluster:    grpcClusterName,
		LogPayloads:    logPayloads,
		TrafficPercent: trafficPercent,
		Version:        version,
	})
}

func (xds *SeldonXDSCache) AddCluster(name string, route string, isGrpc bool) {
	cluster, ok := xds.Clusters[name]
	if !ok {
		cluster = resources.Cluster{
			Name:      name,
			Endpoints: make(map[string]resources.Endpoint),
			Routes:    make(map[string]bool),
			Grpc:      isGrpc,
		}
	}
	cluster.Routes[route] = true
	xds.Clusters[name] = cluster
}

func (xds *SeldonXDSCache) RemoveRoutes(modelName string) error {
	routeList, ok := xds.Routes[modelName]
	if !ok {
		return nil
	}
	delete(xds.Routes, modelName)
	for _, route := range routeList {
		httpCluster, ok := xds.Clusters[route.HttpCluster]
		if !ok {
			return fmt.Errorf("Can't find http cluster for model %s", modelName)
		}
		grpcCluster, ok := xds.Clusters[route.GrpcCluster]
		if !ok {
			return fmt.Errorf("Can't find grpc cluster for model %s", modelName)
		}
		delete(httpCluster.Routes, modelName)
		delete(grpcCluster.Routes, modelName)
		if len(httpCluster.Routes) == 0 {
			delete(xds.Clusters, route.HttpCluster)
			delete(xds.Clusters, route.GrpcCluster)
		} else {
			xds.Clusters[route.HttpCluster] = httpCluster
			xds.Clusters[route.GrpcCluster] = grpcCluster
		}
	}
	return nil
}

func (xds *SeldonXDSCache) AddEndpoint(clusterName, upstreamHost string, upstreamPort uint32) {
	cluster := xds.Clusters[clusterName]
	k := fmt.Sprintf("%s:%d", upstreamHost, upstreamPort)
	cluster.Endpoints[k] = resources.Endpoint{
		UpstreamHost: upstreamHost,
		UpstreamPort: upstreamPort,
	}

	xds.Clusters[clusterName] = cluster
}