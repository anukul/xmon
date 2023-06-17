package monitor

import (
	"math"
	"math/big"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/samber/lo"

	"github.com/anukul/xmon/backend/constants"
	"github.com/anukul/xmon/backend/proto"
	"github.com/anukul/xmon/backend/utils"
)

func (m *Monitor) Refresh() {
	m.logger.Debug().Msg("refreshing")

	// update block number, status, last ping
	m.pingNodes()

	lowestBlockNumber := big.NewInt(math.MaxInt)
	highestBlockNumber := big.NewInt(0)

	healthyNodes := lo.Filter(m.Nodes, func(node *Node, _ int) bool {
		if node.Execution.Status == proto.ListNodesResponse_Client_UNHEALTHY {
			return false
		}
		if blockNumber, err := node.Execution.Client.BlockNumber(); err == nil {
			lowestBlockNumber = utils.MinBigInt(lowestBlockNumber, blockNumber)
			highestBlockNumber = utils.MaxBigInt(highestBlockNumber, blockNumber)
		}
		return true
	})

	// for each pair of nodes, find the highest block number with the same block hash
	matchingBlockNumbers := utils.ProcessPairs(healthyNodes, m.processNodePair)

	lowestMatchingBlockNumber := big.NewInt(math.MaxInt)
	for _, matchingBlockNumber := range matchingBlockNumbers {
		if matchingBlockNumber == nil {
			// processor error
			continue
		}
		lowestMatchingBlockNumber = utils.MinBigInt(lowestMatchingBlockNumber, matchingBlockNumber)
	}
	if lowestMatchingBlockNumber.Cmp(big.NewInt(math.MaxInt)) != 0 {
		lowestBlockNumber = lowestMatchingBlockNumber
	}

	if err := m.updateBlockLimits(lowestBlockNumber, highestBlockNumber); err != nil {
		m.logger.Err(err).Send()
	}
}

func (m *Monitor) pingNodes() {
	var wg sync.WaitGroup
	for nodeIndex, node := range m.Nodes {
		wg.Add(1)
		go func(n *Node, i int) {
			defer wg.Done()
			n.Execution.statusMu.Lock()
			defer n.Execution.statusMu.Unlock()

			latestBlockNumber, err := n.Execution.Client.GetBlockNumber()
			if err != nil {
				n.Execution.Status = proto.ListNodesResponse_Client_UNHEALTHY
				m.logger.Err(err).Send()
				return
			}
			_, err = n.Execution.Client.GetBlockHash(latestBlockNumber)
			if err != nil {
				n.Execution.Status = proto.ListNodesResponse_Client_UNHEALTHY
				m.logger.Err(err).Send()
				return
			}

			n.Execution.Status = proto.ListNodesResponse_Client_HEALTHY
			n.Execution.LastPing = time.Now()
		}(node, nodeIndex)
	}
	wg.Wait()
}

func (m *Monitor) processNodePair(a, b *Node) *big.Int {
	logger := m.logger.With().Str("process_id", uuid.NewString()).Logger()
	logger.Info().Msgf("processing %s and %s", a.Name, b.Name)
	latestBlockNumberA, err := a.Execution.Client.BlockNumber()
	if err != nil {
		logger.Err(err).Send()
		return nil
	}
	logger.Debug().Msgf(`%s is at %d`, a.Name, latestBlockNumberA)
	latestBlockNumberB, err := b.Execution.Client.BlockNumber()
	if err != nil {
		logger.Err(err).Send()
		return nil
	}
	logger.Debug().Msgf(`%s is at %d`, b.Name, latestBlockNumberB)
	latestCommonBlockNumber := utils.MinBigInt(latestBlockNumberA, latestBlockNumberB)
	areExecutionClientsSynced, err := m.areExecutionClientsSynced(a, b, latestCommonBlockNumber)
	if err != nil {
		logger.Err(err).Send()
		return nil
	}
	if areExecutionClientsSynced {
		return latestCommonBlockNumber
	}
	// nodes not synced at latest block, find older block
	var lowestBlockNumber *big.Int
	if err := m.db.Get(constants.LowestBlock, &lowestBlockNumber); err != nil {
		logger.Err(err).Send()
		lowestBlockNumber = big.NewInt(0)
	}
	matchingBlockNumber := utils.BinarySearch(int(lowestBlockNumber.Int64()), int(latestCommonBlockNumber.Int64()), func(blockNumber int) bool {
		res, err := m.areExecutionClientsSynced(a, b, big.NewInt(int64(blockNumber)))
		if err != nil {
			logger.Err(err).Send()
			return false
		}
		return res
	})
	logger.Warn().Msgf(`%s and %s have different hash on %d, matching on %d`, a.Name, b.Name, latestCommonBlockNumber, matchingBlockNumber)
	a.Execution.statusMu.Lock()
	defer a.Execution.statusMu.Unlock()
	a.Execution.Status = proto.ListNodesResponse_Client_HEALTHY
	return big.NewInt(int64(matchingBlockNumber))
}

func (m *Monitor) areExecutionClientsSynced(a, b *Node, blockNumber *big.Int) (bool, error) {
	blockHashA, err := a.Execution.Client.GetBlockHash(blockNumber)
	if err != nil {
		return false, err
	}
	blockHashB, err := b.Execution.Client.GetBlockHash(blockNumber)
	if err != nil {
		return false, err
	}
	return blockHashA == blockHashB, nil
}

func (m *Monitor) updateBlockLimits(lowestBlockNumber, highestBlockNumber *big.Int) error {
	if err := m.db.Set(constants.LowestBlock, lowestBlockNumber); err != nil {
		return err
	}
	if err := m.db.Set(constants.HighestBlock, highestBlockNumber); err != nil {
		return err
	}
	return nil
}
