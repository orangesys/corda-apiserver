package corda

import (
	"crypto/md5"
	"fmt"
)

//NodeConf ...
type NodeConf struct {
	P2PAddress  string `form:"p2pAddress"`
	MyLegalName string `form:"myLegalName"`
}

//UniqueName return node unique name
func (n *NodeConf) UniqueName() string {
	return fmt.Sprintf("%x", md5.Sum([]byte(n.MyLegalName)))
}
