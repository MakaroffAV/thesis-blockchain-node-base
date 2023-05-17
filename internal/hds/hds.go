package hds

// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //

import (
	"fmt"
	"net/http"

	"github.com/MakaroffAV/thesis-blockchain-node-base/pkg/nodebase"
)

// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //

func OpenNodes(w http.ResponseWriter, r *http.Request) {

	baseNodes, baseNodesErr := nodebase.GetBaseNodes()
	if baseNodesErr != nil {
		fmt.Println(baseNodesErr)
	}

	w.Write(baseNodes)

}

// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //
