package app

// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //

import (
	"github.com/MakaroffAV/thesis-blockchain-node-base/pkg/noderoot"
)

// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //

func RunBaseNode() {

	go noderoot.OpenNodesTracker()

}

// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //
