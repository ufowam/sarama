package sarama

type InitProducerIdResponse struct {
	Err           KError
	ThrottleTime  int32
	ProducerID    int64
	ProducerEpoch int16
}

func (r *InitProducerIdResponse) encode(pe packetEncoder) error {
	pe.putInt32(r.ThrottleTime)
	pe.putInt16(int16(r.Err))
	pe.putInt64(r.ProducerID)
	pe.putInt16(r.ProducerEpoch)
	return nil
}

func (r *InitProducerIdResponse) decode(pd packetDecoder, version int16) (err error) {
	r.ThrottleTime, err = pd.getInt32()
	if err != nil {
		return err
	}

	tmp, err := pd.getInt16()
	if err != nil {
		return err
	}

	r.Err = KError(tmp)

	r.ProducerID, err = pd.getInt64()
	if err != nil {
		return err
	}

	r.ProducerEpoch, err = pd.getInt16()
	if err != nil {
		return err
	}

	return nil
}

func (r *InitProducerIdResponse) key() int16 {
	return 22
}

func (r *InitProducerIdResponse) version() int16 {
	return 0
}

func (r *InitProducerIdResponse) requiredVersion() KafkaVersion {
	return V0_11_0_0
}
