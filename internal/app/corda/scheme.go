package corda

import (
	"crypto/md5"
	"fmt"
	"os"
)

//Node ...
type Node struct {
	Config *NodeConf
}

//GetTmpDir ...
func (node *Node) GetTmpDir() string {
	return fmt.Sprintf(os.Getenv("STOAGE_PATH")+"/tmp/%x", md5.Sum([]byte(node.Config.MyLegalName)))
}

//GetHomeDir ...
func (node *Node) GetHomeDir() string {
	return fmt.Sprintf(os.Getenv("STOAGE_PATH")+"/data/%x", md5.Sum([]byte(node.Config.MyLegalName)))
}

//NewNode ...
func NewNode(conf *NodeConf) *Node {
	return &Node{
		Config: conf,
	}
}

//NodeConf ...
type NodeConf struct {
	P2PAddress  string `form:"p2pAddress"`
	MyLegalName string `form:"myLegalName"`
}

//UniqueName return node unique name
func (n *NodeConf) UniqueName() string {
	return fmt.Sprintf("%x", md5.Sum([]byte(n.MyLegalName)))
}
