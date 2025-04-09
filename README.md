# Kafka-Golang

This repository provides a boilerplate code for integrating Kafka with Golang.

## Partitions

### Kafka Producers:
- When producing data, you can **explicitly specify a partition** for the message. However, Kafka may override the partition assignment based on the message's key.
- Kafka typically assigns partitions based on a hashing function of the key.

### Kafka Consumers:
- Consumers in a group consume messages from multiple partitions.
- Kafka will assign partitions to consumers based on availability and the number of consumers in the group.

**Note**: Even if you explicitly specify a partition in your producer message, Kafka may assign a different partition depending on the key's hash. This is why consumers may also see messages from partitions other than the one explicitly specified during production.

## Consumer Groups

In Kafka, a **Consumer Group** is a group of consumers that share the same `GroupID`. The consumers in a group will consume messages from **multiple partitions** of a topic. Kafka ensures that each partition is consumed by only one consumer at a time within a group, and messages from multiple partitions are distributed across the consumers in the group.

- **GroupID**: A unique identifier for each consumer group.
- Each consumer group can have multiple consumers, with the number of consumers determined by the number of partitions in the Kafka topic.
- If the number of consumers exceeds the number of partitions, some consumers will remain idle.
- 
To check the consumer group and the number of consumers, use the following command:

```bash
./kafka-consumer-groups.sh --bootstrap-server localhost:9092 --describe --group {group-name}
