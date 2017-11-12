package sarama

import (
	"testing"
)

var (
	initProducerResponseEmpty = []byte{
		0x00, 0x00, 0x00, 0x00, // Throttle time.
		0x00, 0x00, // Error code.
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // Producer ID.
		0x00, 0x00, // Producer epoch.
	}

	initProducerResponseWithPID = []byte{
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01,
		0x00, 0x00,
	}

	initProducerResponseWithError = []byte{
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x2A,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00,
	}

	initProducerResponseWithThrottleTime = []byte{
		0x00, 0x00, 0x00, 0x64,
		0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00,
	}

	initProducerResponseWithProducerEpoch = []byte{
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x01,
	}
)

func TestProducerIdResponseEmpty(t *testing.T) {
	var response *InitProducerIdResponse

	response = new(InitProducerIdResponse)
	testVersionDecodable(t, "empty response", response, initProducerResponseEmpty, 0)
	if response.ThrottleTime != 0 {
		t.Error("Decoding throttle time failed: 0 expected but found", response.ThrottleTime)
	}
	if response.Err != 0 {
		t.Error("Decoding error failed: 0 expected but found", response.Err)
	}
	if response.ProducerID != 0 {
		t.Error("Decoding producer ID failed: 0 expected but found", response.ProducerID)
	}
	if response.ProducerEpoch != 0 {
		t.Error("Decoding producer epoch failed: 0 expected but found", response.ProducerEpoch)
	}
}

func TestProducerIdResponseWithPID(t *testing.T) {
	var response *InitProducerIdResponse

	response = new(InitProducerIdResponse)
	testVersionDecodable(t, "producer ID is 1", response, initProducerResponseWithPID, 0)
	if response.ProducerID != 1 {
		t.Error("Decoding producer ID failed: 1 expected but found", response.ProducerID)
	}
}

func TestProducerIdResponseWithError(t *testing.T) {
	var response *InitProducerIdResponse

	response = new(InitProducerIdResponse)
	testVersionDecodable(t, "Error code is 42", response, initProducerResponseWithError, 0)
	if response.Err != 42 {
		t.Error("Decoding error code failed: 42 expected but found", response.Err)
	}
}

func TestProducerIdResponseWithThrottleTime(t *testing.T) {
	var response *InitProducerIdResponse

	response = new(InitProducerIdResponse)
	testVersionDecodable(t, "Throttle time is 100", response, initProducerResponseWithThrottleTime, 0)
	if response.ThrottleTime != 100 {
		t.Error("Decoding throttle time failed: 100 expected but found", response.ThrottleTime)
	}
}

func TestProducerIdResponseWithProducerEpoch(t *testing.T) {
	var response *InitProducerIdResponse

	response = new(InitProducerIdResponse)
	testVersionDecodable(t, "Producer epoch is 1", response, initProducerResponseWithProducerEpoch, 0)
	if response.ProducerEpoch != 1 {
		t.Error("Decoding producer epoch failed: 1 expected but found", response.ProducerEpoch)
	}
}
