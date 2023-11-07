package app

import (
	"log"
	"net"
	"net/rpc"
	"testing"

	"github.com/Awesome-Sauces/abstract/crypto"
)

type Node struct {
	id    string
	key   crypto.KeyPair
	nodes []*Node
}

type MessageArgs struct {
	Sender    string
	Message   string
	Signature string
}

type MessageReply struct {
	Verifications int
}

func (node *Node) PropagateMessage(msg string) {

	signature := node.key.Sign(msg)

	verifications := 0

	for _, n := range node.nodes {
		args := &MessageArgs{
			Sender:    node.key.Address,
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

func (node *Node) ReceiveMessage(sender string, msg string, signature string) int {
	pk, err := crypto.RecoverPublicKeyFromSignature(msg, signature)

	if err != nil {
		log.Fatal(err)
	}

	if pk.Address() != sender {
		return 0
	}

	verifications := 1 //Start with 1 for the current node

	for _, n := range node.nodes {
		// Propagate the message to connected nodes and collect verifications
		verifications += n.ReceiveMessage(sender, msg, signature)
	}

	return verifications
}

func (node *Node) Link(n *Node) {
	node.nodes = append(node.nodes, n)
}

func (node *Node) PropagateMessageRPC(args *MessageArgs, reply *MessageReply) error {
	verifications := node.ReceiveMessage(args.Sender, args.Message, args.Signature)
	reply.Verifications = verifications
	return nil
}

func TestConsensus(t *testing.T) {
	t.Run("InitTest", func(t *testing.T) {

		node_a := &Node{
			key:   *crypto.NewKeyPair(),
			nodes: make([]*Node, 0),
			id:    "a",
		}

		node_b := &Node{
			key:   *crypto.NewKeyPair(),
			nodes: make([]*Node, 0),
			id:    "b",
		}

		node_c := &Node{
			key:   *crypto.NewKeyPair(),
			nodes: make([]*Node, 0),
			id:    "c",
		}

		node_a.Link(node_b)
		node_b.Link(node_c)

		node_a.PropagateMessage("Whatever")

	})
}
func startRPCServer(node *Node) {
	rpc.Register(node)
	l, err := net.Listen("tcp", ":1234")
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
