package main

import "github.com/unixpickle/serializer"

const (
	serializerTypePrefix = "github.com/unixpickle/char-rnn"
	serializerTypeLSTM   = serializerTypePrefix + "LSTM"
	serializerTypeGRU    = serializerTypePrefix + "GRU"
)

func init() {
	serializer.RegisterDeserializer(serializerTypeLSTM, DeserializeLSTM)
	serializer.RegisterDeserializer(serializerTypeGRU, DeserializeGRU)
}
