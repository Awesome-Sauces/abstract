package p2p

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net"
	"net/rpc"
	"sync"

	"github.com/Awesome-Sauces/abstract/crypto"
)

type Peer struct {
	Key   crypto.KeyPair
	Nodes []*Peer
	mu    sync.Mutex
}

type MessageArgs struct {
	Sender    string
	Message   string
	Signature string
}

type MessageReply struct {
	Verifications int
}

type P2PNetwork struct {
	Nodes map[string]*Peer
}

func NewPeer() *Peer {
	return &Peer{
		Key: *crypto.NewKeyPair(),
	}
}

func NewPeerFromKey(key string) *Peer {
	return &Peer{
		Key: *crypto.KeyPairFromPK(key),
	}
}

func NewP2PNetwork() *P2PNetwork {
	return &P2PNetwork{
		Nodes: make(map[string]*Peer),
	}
}

func (p2p *P2PNetwork) AddNode(node *Peer) {
	p2p.Nodes[node.Key.Address] = node
}

func (p2p *P2PNetwork) GetNode(address string) *Peer {
	return p2p.Nodes[address]
}

func (node *Peer) PropagateMessage(msg string) {
	signature := node.Key.Sign(msg)
	verifications := 0

	for _, n := range node.Nodes {
		args := &MessageArgs{
			Sender:    node.Key.Address,
			Message:   msg,
			Signature: signature,
		}
		reply := new(MessageReply)
		err := n.PropagateMessageRPC(args, reply)
		if err == nil {
			verifications += reply.Verifications
		}
	}

	log.Println("Verifications ", verifications)
}

func (node *Peer) ReceiveMessage(sender string, msg string, signature string) int {
	pk, err := crypto.RecoverPublicKeyFromSignature(msg, signature)

	if err != nil {
		log.Fatal(err)
	}

	if pk.Address() != sender {
		return 0
	}

	verifications := 1

	for _, n := range node.Nodes {
		verifications += n.ReceiveMessage(sender, msg, signature)
	}

	return verifications
}

func (node *Peer) Link(n *Peer) {
	node.mu.Lock()
	defer node.mu.Unlock()
	node.Nodes = append(node.Nodes, n)
}

func (node *Peer) PropagateMessageRPC(args *MessageArgs, reply *MessageReply) error {
	verifications := node.ReceiveMessage(args.Sender, args.Message, args.Signature)
	reply.Verifications = verifications
	return nil
}

func (p2p *P2PNetwork) SaveNetworkToJSON(filename string) error {
	data, err := json.Marshal(p2p)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(filename, data, 0644)
	if err != nil {
		return err
	}

	return nil
}

func LoadNetworkFromJSON(filename string) (*P2PNetwork, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var p2p P2PNetwork
	err = json.Unmarshal(data, &p2p)
	if err != nil {
		return nil, err
	}

	return &p2p, nil
}

func StartRPCServer(node *Peer, port string) {
	rpc.Register(node)
	l, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatal("listen error:", err)
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			continue
		}
		go rpc.ServeConn(conn)
	}
}
