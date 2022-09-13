package resolver

import (
	"context"
	"net"
)

var overrideNames = map[string]string{}

func LookupIPWithOverride(ctx context.Context, network, host string) ([]net.IP, error) {
	if newHost, ok := overrideNames[host]; ok {
		host = newHost
	}
	return net.DefaultResolver.LookupIP(ctx, network, host)
}

func Override(original, replacement string) {
	m := map[string]string{}
	for k, v := range overrideNames {
		m[k] = v
	}
	m[original] = replacement
	overrideNames = m
}

func Empty() {
	overrideNames = map[string]string{}
}

func Replace(m map[string]string) {
	overrideNames = m
}

func NewOverridableResolver() *net.Resolver {
	v := &MemResolver{
		LookupIP: LookupIPWithOverride,
	}
	return NewMemoryResolver(v)
}
