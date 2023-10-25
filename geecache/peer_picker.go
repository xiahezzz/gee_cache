package geecache

type PeerPicker interface {
	PickPeer(key string) (peer peerGetter, ok bool)
}

type peerGetter interface {
	Get(group string, key string) ([]byte, error)
}
