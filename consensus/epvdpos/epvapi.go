package epvdpos

import (
	"github.com/epvchain/go-epvchain/common"
	"github.com/epvchain/go-epvchain/consensus"
	"github.com/epvchain/go-epvchain/core/types"
	"github.com/epvchain/go-epvchain/rpc"
)

type EAPI struct {
	chain  consensus.ChainReader
	epvdpos *EPVDpos
}

func (e *EAPI) GetSnapshot(number *rpc.BlockNumber) (*Snapshot, error) {
	var header *types.Header
	if number == nil || *number == rpc.LatestBlockNumber {
		header = e.chain.CurrentHeader()
	} else {
		header = e.chain.GetHeaderByNumber(uint64(number.Int64()))
	}
	if header == nil {
		return nil, errUnknownBlock
	}
	return e.epvdpos.snapshot(e.chain, header.Number.Uint64(), header.Hash(), nil)
}

func (e *EAPI) Propose(address common.Address, auth bool) {
	e.epvdpos.lock.Lock()
	defer e.epvdpos.lock.Unlock()

	e.epvdpos.proposals[address] = auth
}

func (e *EAPI) Proposals() map[common.Address]bool {
	e.epvdpos.lock.RLock()
	defer e.epvdpos.lock.RUnlock()

	proposals := make(map[common.Address]bool)
	for address, auth := range e.epvdpos.proposals {
		proposals[address] = auth
	}
	return proposals
}

func (e *EAPI) GetSnapshotAtHash(hash common.Hash) (*Snapshot, error) {
	header := e.chain.GetHeaderByHash(hash)
	if header == nil {
		return nil, errUnknownBlock
	}
	return e.epvdpos.snapshot(e.chain, header.Number.Uint64(), header.Hash(), nil)
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
	snap, err := e.epvdpos.snapshot(e.chain, header.Number.Uint64(), header.Hash(), nil)
	if err != nil {
		return nil, err
	}
	return snap.signers(), nil
}

func (e *EAPI) Discard(address common.Address) {
	e.epvdpos.lock.Lock()
	defer e.epvdpos.lock.Unlock()

	delete(e.epvdpos.proposals, address)
}

func (e *EAPI) GetSignersAtHash(hash common.Hash) ([]common.Address, error) {
	header := e.chain.GetHeaderByHash(hash)
	if header == nil {
		return nil, errUnknownBlock
	}
	snap, err := e.epvdpos.snapshot(e.chain, header.Number.Uint64(), header.Hash(), nil)
	if err != nil {
		return nil, err
	}
	return snap.signers(), nil
}
