package monitor

import (
	"github.com/anukul/xmon/backend/db"
	el_test "github.com/anukul/xmon/backend/execution_client/test"
	"math/big"
	"sync"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/anukul/xmon/backend/common"
	"github.com/anukul/xmon/backend/config"
	"github.com/anukul/xmon/backend/consensus_client/beacon_api"
	"github.com/anukul/xmon/backend/execution_client/json_rpc"
	"github.com/anukul/xmon/backend/proto"
)

type NodeClient interface {
	Version() (string, error)
}

type ExecutionClient interface {
	NodeClient
	BlockNumber() (*big.Int, error)
	GetBlockNumber() (*big.Int, error)
	BlockHash(blockNumber *big.Int) (string, error)
	GetBlockHash(blockNumber *big.Int) (string, error)
}

type ConsensusClient interface {
	NodeClient
	LatestSlot() (uint64, string, error)
	Slot(slotNumber *big.Int) (uint64, string, error)
}

type Client[T NodeClient] struct {
	Status   proto.ListNodesResponse_Client_Status
	Version  string
	LastPing time.Time
	Client   T

	statusMu sync.Mutex
}

type Node struct {
	Name      string
	Execution *Client[ExecutionClient]
	Consensus *Client[ConsensusClient]
}

type Monitor struct {
	logger *zerolog.Logger
	db     common.DatabaseClient
	Nodes  []*Node
}

func NewMonitor(cfg *config.Config, db *db.Database) (*Monitor, error) {
	logger := log.With().Str("component", "monitor").Logger()

	nodes := make([]*Node, 0)
	for _, n := range cfg.Nodes {
		executionClient, err := json_rpc.NewClient(n.ExecutionLayerURL, n.AuthHeaders, db.WithScope(n.Name))
		if err != nil {
			return nil, err
		}
		consensusClient, err := beacon_api.NewClient(n.ConsensusLayerURL, n.AuthHeaders)
		if err != nil {
			return nil, err
		}

		executionClientVersion, err := executionClient.Version()
		if err != nil {
			return nil, err
		}
		consensusClientVersion, err := consensusClient.Version()
		if err != nil {
			return nil, err
		}

		nodes = append(nodes, &Node{
			Name: n.Name,
			Execution: &Client[ExecutionClient]{
				Client:  executionClient,
				Version: executionClientVersion,
			},
			Consensus: &Client[ConsensusClient]{
				Client:  consensusClient,
				Version: consensusClientVersion,
			},
		})
	}

	return &Monitor{
		logger: &logger,
		Nodes:  nodes,
		db:     db,
	}, nil
}

func NewTestMonitor(cfg *config.Config, db *db.Database) (*Monitor, error) {
	logger := log.With().Str("component", "monitor").Logger()

	nodes := make([]*Node, 0)
	for node, chain := range cfg.TestNodes {
		executionClient := el_test.NewClient(chain, db.WithScope(node))
		executionClientVersion, err := executionClient.Version()
		if err != nil {
			return nil, err
		}

		nodes = append(nodes, &Node{
			Name: node,
			Execution: &Client[ExecutionClient]{
				Client:  executionClient,
				Version: executionClientVersion,
			},
		})
	}

	return &Monitor{
		logger: &logger,
		db:     db,
		Nodes:  nodes,
	}, nil
}
