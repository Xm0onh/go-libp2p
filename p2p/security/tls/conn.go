package libp2ptls

import (
	"crypto/tls"

	ci "github.com/libp2p/go-libp2p/core/crypto"
	"github.com/libp2p/go-libp2p/core/network"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-libp2p/core/sec"
)

type conn struct {
	*tls.Conn

	localPeer peer.ID
	privKey   ci.PrivKey

	remotePeer      peer.ID
	remotePubKey    ci.PubKey
	connectionState network.ConnectionState
}

var _ sec.SecureConn = &conn{}

func (c *conn) LocalPeer() peer.ID {
	return c.localPeer
}

func (c *conn) LocalPrivateKey() ci.PrivKey {
	return c.privKey
}

func (c *conn) RemotePeer() peer.ID {
	return c.remotePeer
}

func (c *conn) RemotePublicKey() ci.PubKey {
	return c.remotePubKey
}

func (c *conn) ConnState() network.ConnectionState {
	return c.connectionState
}

func (c *conn) SetConnState(state network.ConnectionState) {
	c.connectionState = state
}
