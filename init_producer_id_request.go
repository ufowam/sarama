package sarama

type InitProducerIDRequest struct {
}

func (r *InitProducerIDRequest) encode(pe packetEncoder) error {
	// TODO No support for transactions for now.
	err := pe.putNullableString(nil)
	if err != nil {
		return err
	}

	pe.putInt32(0)

	return nil
}

func (r *InitProducerIDRequest) decode(pd packetDecoder, version int16) (err error) {
	_, err = pd.getNullableString()
	if err != nil {
		return err
	}

	_, err = pd.getInt32()
	if err != nil {
		return err
	}

	return nil
}

func (r *InitProducerIDRequest) key() int16 {
	return 22
}

func (r *InitProducerIDRequest) version() int16 {
	return 0
}

func (r *InitProducerIDRequest) requiredVersion() KafkaVersion {
	return V0_11_0_0
}
