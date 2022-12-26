package minecraft

import (
	"context"
	"errors"
	"github.com/sandertv/go-raknet"
	"net"
)

// RakNet is an implementation of a RakNet v10 Network.
type RakNet struct{}

// DialContext ...
func (r RakNet) DialContext(ctx context.Context, address string) (net.Conn, error) {
	protocol, ok := ctx.Value("protocol").(byte)
	if !ok {
		return nil, errors.New("protocol not defined")
	}
	return raknet.DialContext(ctx, address, protocol)
}

// PingContext ...
func (r RakNet) PingContext(ctx context.Context, address string) (response []byte, err error) {
	return raknet.PingContext(ctx, address)
}

// Listen ...
func (r RakNet) Listen(address string) (NetworkListener, error) {
	return raknet.Listen(address)
}

// init registers the RakNet network.
func init() {
	RegisterNetwork("raknet", RakNet{})
}
