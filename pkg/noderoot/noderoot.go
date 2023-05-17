package noderoot

// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //

import (
	"fmt"

	"github.com/MakaroffAV/thesis-blockchain-node-base/pkg/command"
	"github.com/gorilla/websocket"
)

// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //

const (

	// urlNodeRoot is url for
	// connecting to root blockchain node and tracking open blockchain nodes
	urlNodeRoot = "ws://thesis-logistics.ru:2605"
)

// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //

// OpenNodesTracker is function for
// detecting open nodes in blockchain network
func OpenNodesTracker() {

	// open connection to root
	// node via web sockets connection
	conn, _, err := websocket.DefaultDialer.Dial(urlNodeRoot, nil)
	if err != nil {
		fmt.Println("failed to connect to root node")
		return
	}

	// infinite loop for
	// tracking open nodes in blockchain
	for {

		// fetch message from root node
		_, msg, msgErr := conn.ReadMessage()
		if msgErr != nil {
			fmt.Println("failed to read message from root node")
			return
		}

		// parse message from root node
		command.Process(msg)

	}

}

// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //
