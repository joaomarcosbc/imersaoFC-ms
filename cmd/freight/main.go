package freight

import (
	"database/sql"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/joaomarcosbc/imersaoFC-ms/pkg/kafka"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(host.docker.internal:3306)/routes?parseTime=true")

	if err != nil {
		panic(err)
	}

	defer db.Close()

	msgChan := make(chan *ckafka.Message)
	topics := []string{"routes"}
	servers := "host.docker.internal:9094"
	kafka.Consume(topics, servers, msgChan)
}
