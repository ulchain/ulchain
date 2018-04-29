package epvdpos

import (
	"github.com/epvchain/go-epvchain/common"
	"github.com/epvchain/go-epvchain/consensus"
	"github.com/epvchain/go-epvchain/core/types"
	"github.com/epvchain/go-epvchain/rpc"
)

type EAPI struct {
	chain  consensus.ChainReader
	dpos *DPos
}

func (e *EAPI) GetArchive(number *rpc.BlockNumber) (*Archive, error) {
	var header *types.Header
	if number == nil || *number == rpc.LatestBlockNumber {
		header = e.chain.CurrentHeader()
	} else {
		header = e.chain.GetHeaderByNumber(uint64(number.Int64()))
	}
	if header == nil {
		return nil, errUnknownBlock
	}
	return e.dpos.archive(e.chain, header.Number.Uint64(), header.Hash(), nil)
}

func (e *EAPI) Propose(address common.Address, auth bool) {
	e.dpos.lock.Lock()
	defer e.dpos.lock.Unlock()

	e.dpos.proposals[address] = auth
}

func (e *EAPI) Proposals() map[common.Address]bool {
	e.dpos.lock.RLock()
	defer e.dpos.lock.RUnlock()

	proposals := make(map[common.Address]bool)
	for address, auth := range e.dpos.proposals {
		proposals[address] = auth
	}
	return proposals
}

func (e *EAPI) GetArchiveAtHash(hash common.Hash) (*Archive, error) {
	header := e.chain.GetHeaderByHash(hash)
	if header == nil {
		return nil, errUnknownBlock
	}
	return e.dpos.archive(e.chain, header.Number.Uint64(), header.Hash(), nil)
}

func (e *EAPI) GetSigners(number *rpc.BlockNumber) ([]common.Address, error) {
	var header *types.Header
	if number == nil || *number == rpc.LatestBlockNumber {
		header = e.chain.CurrentHeader()
	} else {
		header = e.chain.GetHeaderByNumber(uint64(number.Int64()))
	}
	if header == nil {
		return nil, errUnknownBlock
	}
	archive, err := e.dpos.archive(e.chain, header.Number.Uint64(), header.Hash(), nil)
	if err != nil {
		return nil, err
	}
	return archive.signers(), nil
}

func (e *EAPI) Discard(address common.Address) {
	e.dpos.lock.Lock()
	defer e.dpos.lock.Unlock()

	delete(e.dpos.proposals, address)
}

func (e *EAPI) GetSignersAtHash(hash common.Hash) ([]common.Address, error) {
	header := e.chain.GetHeaderByHash(hash)
	if header == nil {
		return nil, errUnknownBlock
	}
	archive, err := e.dpos.archive(e.chain, header.Number.Uint64(), header.Hash(), nil)
	if err != nil {
		return nil, err
	}
	return archive.signers(), nil
}
