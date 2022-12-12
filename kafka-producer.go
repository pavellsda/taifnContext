package taifn

type KafkaProducer interface {
	ProduceFull(topic string, key string, value string, Headers map[string]string, brokers []string) error
	ProduceWithHeaders(topic string, key string, value string, Headers map[string]string) error
	Produce(topic string, key string, value string) error
}
