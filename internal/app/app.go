package app

// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //

import (
	"log"
	"net/http"

	"github.com/MakaroffAV/thesis-blockchain-node-base/internal/rts"
	"github.com/MakaroffAV/thesis-blockchain-node-base/pkg/noderoot"
)

// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //

func RunBaseNode() {

	go noderoot.OpenNodesTracker()

	for _, route := range rts.GetRoutes() {
		http.HandleFunc(route.Path, route.Handler)
	}

	log.Fatal(http.ListenAndServe("0.0.0.0:2605", nil))

}

// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //
