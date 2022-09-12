package resolver

import (
	"context"
	"net"
)

var overrideNames = map[string]string{}

func LookupIPWithOverride(ctx context.Context, network, host string) ([]net.IP, error) {

	newHost, ok := overrideNames[host]
	if ok {
		host = newHost
	}

	return net.DefaultResolver.LookupIP(ctx, network, host)
}

func OverrideName(original, replacement string) {

	m := map[string]string{}
	for k, v := range overrideNames {
		m[k] = v
	}
	m[original] = replacement
	overrideNames = m
}
