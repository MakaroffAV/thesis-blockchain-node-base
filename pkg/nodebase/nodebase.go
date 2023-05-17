package nodebase

import (
	"encoding/json"
	"errors"
	"sync"
)

// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //

var (
	mu                  = &sync.Mutex{}
	baseNodes           = make(map[string]bool)
	errBaseNodesToBytes = errors.New("")
)

// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //

// GetOpenBaseNodes is public function for
// creating simple interface
func GetBaseNodes() ([]byte, error) {

	j, jErr := json.Marshal(baseNodes)
	if jErr != nil {
		return nil, errBaseNodesToBytes
	}

	return j, nil
}

// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //

// UpdateBaseNodes is public function for
// updating set of available nodes in blockchain
func UpdateBaseNodes(newOpenBaseNodes *map[string]bool) {

	// make updating
	// operation thread safe
	mu.Lock()
	defer mu.Unlock()

	// updating set of
	// open nodes in blockchain network
	baseNodes = *newOpenBaseNodes

}

// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //
