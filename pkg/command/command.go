package command

import (
	"bytes"
	"encoding/gob"
	"fmt"

	"github.com/MakaroffAV/thesis-blockchain-node-base/pkg/nodebase"
)

// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //

// hNodeIps is private function for
func hNodeIps(message []byte) {

	var (
		buff bytes.Buffer

		j = map[string]bool{}
	)

	buff.Write(message[20:])

	decoder := gob.NewDecoder(&buff)
	err := decoder.Decode(&j)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(j)

	nodebase.UpdateBaseNodes(&j)
}

// ------------------------------------------------------------------------ //

func parseCommand(message []byte) string {

	var comm []byte

	for _, b := range message {
		if b != 0x0 {
			comm = append(comm, b)
		}
	}

	return string(comm)

}

// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //

func Process(msg []byte) {

	switch parseCommand(msg) {
	case "nodeIps":
		{
			hNodeIps(msg)
		}
	default:
		fmt.Println("passed undefined message")
	}

}

// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //
