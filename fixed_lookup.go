package resolver

import (
	"context"
	"net"
)

func fixedFunc(v string) func(context.Context, string, string) ([]net.IP, error) {
	return func(ctx context.Context, network, host string) ([]net.IP, error) {
		return net.DefaultResolver.LookupIP(ctx, network, v)
	}
}

func NewFixedResolver(fixedHost string) *net.Resolver {
	v := &MemResolver{
		LookupIP: fixedFunc(fixedHost),
	}
	return NewMemoryResolver(v)
}
