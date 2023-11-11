package p2p

// Message holds any arbitrary data that is being sent over each transport between 2
// nodes in the network.
type Message struct {
	Payload []byte
}
