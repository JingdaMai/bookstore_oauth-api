package cassandra

import (
	"github.com/gocql/gocql"
	"log"
)

func init() {
	// connect to Cassandra cluster
	cluster:=gocql.NewCluster("127.0.0.1")
	cluster.Keyspace = "oatuh"
	cluster.Consistency = gocql.Quorum
	session, err := cluster.CreateSession()
	if err != nil {
		panic(err)
	}
	log.Println("cassandra connection successfully created")
	defer session.Close()
}