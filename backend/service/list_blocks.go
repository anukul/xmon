package service

import (
	"context"
	"github.com/ethereum/go-ethereum"
	"math/big"

	connect_go "github.com/bufbuild/connect-go"

	"github.com/anukul/xmon/backend/constants"
	"github.com/anukul/xmon/backend/proto"
)

func (s *Service) ListBlocks(context.Context, *connect_go.Request[proto.ListBlocksRequest]) (*connect_go.Response[proto.ListBlocksResponse], error) {
	var lowestBlock, highestBlock *big.Int
	if err := s.db.Get(constants.LowestBlock, &lowestBlock); err != nil {
		s.logger.Err(err).Send()
		return nil, err
	}
	if err := s.db.Get(constants.HighestBlock, &highestBlock); err != nil {
		s.logger.Err(err).Send()
		return nil, err
	}
	var blocks []*proto.ListBlocksResponse_BlockData
	for blockNumber := highestBlock.Int64(); blockNumber >= lowestBlock.Int64(); blockNumber-- {
		nodeHashes := map[string]string{}
		for _, node := range s.monitor.Nodes {
			// hashes for some block numbers may not be cached due to binary search
			blockHash, err := node.Execution.Client.GetBlockHash(big.NewInt(blockNumber))
			if err != nil && err != ethereum.NotFound {
				s.logger.Err(err).Send()
				return nil, err
			}
			nodeHashes[node.Name] = blockHash
		}
		hashes := map[string]bool{}
		for _, hash := range nodeHashes {
			hashes[hash] = true
		}
		blocks = append(blocks, &proto.ListBlocksResponse_BlockData{
			BlockNumber: blockNumber,
			Hashes:      nodeHashes,
			// block is synced if all nodes report a singular hash for this block
			// block is not synced if hash is empty (not found) in all nodes
			IsSynced: len(hashes) == 1 && !hashes[""],
		})
	}
	return connect_go.NewResponse(&proto.ListBlocksResponse{Blocks: blocks}), nil
}
