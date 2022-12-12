package taifn

type KafkaProducer interface {
	ProduceFull(topic string, key string, value string, Headers map[string]string, brokers []string)
	ProduceWithHeaders(topic string, key string, value string, Headers map[string]string)
	Produce(topic string, key string, value string)
}
