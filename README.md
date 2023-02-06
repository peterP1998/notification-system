# Notification-system

## Initial steps
### 1. Run kafka on your local machine:
First you need to have docker installed on your machine. Here is a
documentation how to install docker on your machine: .
#### 1.1. Pull kafka image from docker registry:
```
docker pull spotify/kafka
```
#### 1.2. Run kafka from the docker image as container:
```
docker run -p 2181:2181 -p 9092:9092 --name kafka-docker-container --env ADVERTISED_HOST=127.0.0.1 --env ADVERTISED_PORT=9092 spotify/kafka
```

#### 1.3. Exec to container
The first step is to get the container id
```
docker ps| grep kafka-docker-container
```
Next step is to run docker exec
```
docker exec -it <container-id> /bin/bash
```
#### 1.3. Create kafka topics:
Kafka topic for sms notifications:

```
/opt/kafka_2.11-0.10.1.0/bin/kafka-topics.sh --create --topic sms-notification-topic --partitions 1 --replication-factor 1 --zookeeper localhost:2181
```

Kafka topic for slack notifications:
```
/opt/kafka_2.11-0.10.1.0/bin/kafka-topics.sh --create --topic slack-notification-topic --partitions 1 --replication-factor 1 --zookeeper localhost:2181
```

Kafka topic for email notifications:
```
/opt/kafka_2.11-0.10.1.0/bin/kafka-topics.sh --create --topic email-notification-topic --partitions 1 --replication-factor 1 --zookeeper localhost:2181
```
