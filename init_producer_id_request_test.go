package sarama

import (
	"testing"
)

var (
	initProducerRequestNoTransactionID = []byte{
		0xFF, 0xFF, // Null transactional ID.
		0x00, 0x00, 0x00, 0x00, // Transaction timeout.
	}
)

func TestProducerIdRequest(t *testing.T) {
	request := new(InitProducerIDRequest)
	testRequest(t, "no transaction Id", request, initProducerRequestNoTransactionID)
}
