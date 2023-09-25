package raft

import (
	"encoding/json"
	"fileProcessing/config"
	domains "fileProcessing/internal/core/domain"
	"fileProcessing/internal/repositories/redis"
	raft2 "github.com/hashicorp/raft"
	raftboltdb "github.com/hashicorp/raft-boltdb"
	"io"
	"log"
	"net"
	"os"
	"path/filepath"
	"sync"
	"time"
)

type RaftCluster struct {
	config config.AppConfig
}

var (
	raftNode  *raft2.Raft
	redisRepo *redis.RedisRepository
)

type raftFSM struct {
}

func (rf *raftFSM) Snapshot() (raft2.FSMSnapshot, error) {
	//TODO implement later
	panic("implement later")
}

func (rf *raftFSM) Restore(snapshot io.ReadCloser) error {
	//TODO implement later
	panic("implement later")
}

func NewRaftCluster(conf config.AppConfig, redisRepo2 *redis.RedisRepository) *RaftCluster {
	redisRepo = redisRepo2
	return &RaftCluster{
		config: conf,
	}
}

func (rc *RaftCluster) CreateNewRaftCluster() *raft2.Raft {
	raftConf := raft2.DefaultConfig()
	raftConf.LocalID = raft2.ServerID("node1")
	logDir := "../raft-data"
	os.MkdirAll(logDir, 0700)
	logFile := filepath.Join(logDir, "raft-log.bolt")

	boltDB, err := raftboltdb.NewBoltStore(logFile)
	if err != nil {
		log.Fatalf("Error creating Bolt store: %v", err)
	}
	raftSnapShotRetain := 2
	snapshotStore, err := raft2.NewFileSnapshotStore("", raftSnapShotRetain, os.Stdout)
	if err != nil {
		log.Fatal(err)
	}
	// Create the Raft transport
	addr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:9000")
	if err != nil {
		log.Fatalf("Error resolving TCP address: %v", err)
	}
	transport, err := raft2.NewTCPTransport(addr.String(), addr, 3, 10*time.Second, os.Stdout)
	if err != nil {
		log.Fatalf("Error creating TCP transport: %v", err)
	}

	// Create the Raft node
	r, err := raft2.NewRaft(raftConf, &raftFSM{}, boltDB, boltDB, snapshotStore, transport)
	if err != nil {
		log.Fatalf("Error creating Raft node: %v", err)
	}
	configuration := raft2.Configuration{
		Servers: []raft2.Server{
			{
				ID:      raftConf.LocalID,
				Address: transport.LocalAddr(),
			},
		},
	}
	r.BootstrapCluster(configuration)
	raftNode = r
	return raftNode
}

func (rf *raftFSM) Apply(log *raft2.Log) interface{} {
	mu := sync.Mutex{}
	var file domains.File
	err := json.Unmarshal(log.Data, &file)
	if err != nil {
		return err
	}
	mu.Lock()
	defer mu.Unlock()
	err = redisRepo.Set(file.FileName, file.Size)
	if err != nil {
		return err
	}
	return nil
}
