// Copyright 2015 The go-epvchain Authors
// This file is part of the go-epvchain library.
//
// The go-epvchain library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-epvchain library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-epvchain library. If not, see <http://www.gnu.org/licenses/>.

package epv

import (
	"context"
	"math/big"

	"github.com/epvchain/go-epvchain/act"
	"github.com/epvchain/go-epvchain/public"
	"github.com/epvchain/go-epvchain/public/math"
	"github.com/epvchain/go-epvchain/kernel"
	"github.com/epvchain/go-epvchain/kernel/bloombits"
	"github.com/epvchain/go-epvchain/kernel/state"
	"github.com/epvchain/go-epvchain/kernel/types"
	"github.com/epvchain/go-epvchain/kernel/vm"
	"github.com/epvchain/go-epvchain/epv/downloader"
	"github.com/epvchain/go-epvchain/epv/gasprice"
	"github.com/epvchain/go-epvchain/data"
	"github.com/epvchain/go-epvchain/notice"
	"github.com/epvchain/go-epvchain/content"
	"github.com/epvchain/go-epvchain/remote"
)

// EPVApiBackend implements epvapi.Backend for full nodes
type EPVApiBackend struct {
	epv *EPVchain
	gpo *gasprice.Oracle
}

func (b *EPVApiBackend) ChainConfig() *params.ChainConfig {
	return b.epv.chainConfig
}

func (b *EPVApiBackend) CurrentBlock() *types.Block {
	return b.epv.blockchain.CurrentBlock()
}

func (b *EPVApiBackend) SetHead(number uint64) {
	b.epv.protocolManager.downloader.Cancel()
	b.epv.blockchain.SetHead(number)
}

func (b *EPVApiBackend) HeaderByNumber(ctx context.Context, blockNr rpc.BlockNumber) (*types.Header, error) {
	// Pending block is only known by the miner
	if blockNr == rpc.PendingBlockNumber {
		block := b.epv.miner.PendingBlock()
		return block.Header(), nil
	}
	// Otherwise resolve and return the block
	if blockNr == rpc.LatestBlockNumber {
		return b.epv.blockchain.CurrentBlock().Header(), nil
	}
	return b.epv.blockchain.GetHeaderByNumber(uint64(blockNr)), nil
}

func (b *EPVApiBackend) BlockByNumber(ctx context.Context, blockNr rpc.BlockNumber) (*types.Block, error) {
	// Pending block is only known by the miner
	if blockNr == rpc.PendingBlockNumber {
		block := b.epv.miner.PendingBlock()
		return block, nil
	}
	// Otherwise resolve and return the block
	if blockNr == rpc.LatestBlockNumber {
		return b.epv.blockchain.CurrentBlock(), nil
	}
	return b.epv.blockchain.GetBlockByNumber(uint64(blockNr)), nil
}

func (b *EPVApiBackend) StateAndHeaderByNumber(ctx context.Context, blockNr rpc.BlockNumber) (*state.StateDB, *types.Header, error) {
	// Pending state is only known by the miner
	if blockNr == rpc.PendingBlockNumber {
		block, state := b.epv.miner.Pending()
		return state, block.Header(), nil
	}
	// Otherwise resolve the block number and return its state
	header, err := b.HeaderByNumber(ctx, blockNr)
	if header == nil || err != nil {
		return nil, nil, err
	}
	stateDb, err := b.epv.BlockChain().StateAt(header.Root)
	return stateDb, header, err
}

func (b *EPVApiBackend) GetBlock(ctx context.Context, blockHash common.Hash) (*types.Block, error) {
	return b.epv.blockchain.GetBlockByHash(blockHash), nil
}

func (b *EPVApiBackend) GetReceipts(ctx context.Context, blockHash common.Hash) (types.Receipts, error) {
	return core.GetBlockReceipts(b.epv.chainDb, blockHash, core.GetBlockNumber(b.epv.chainDb, blockHash)), nil
}

func (b *EPVApiBackend) GetTd(blockHash common.Hash) *big.Int {
	return b.epv.blockchain.GetTdByHash(blockHash)
}

func (b *EPVApiBackend) GetEVM(ctx context.Context, msg core.Message, state *state.StateDB, header *types.Header, vmCfg vm.Config) (*vm.EVM, func() error, error) {
	state.SetBalance(msg.From(), math.MaxBig256)
	vmError := func() error { return nil }

	context := core.NewEVMContext(msg, header, b.epv.BlockChain(), nil)
	return vm.NewEVM(context, state, b.epv.chainConfig, vmCfg), vmError, nil
}

func (b *EPVApiBackend) SubscribeRemovedLogsEvent(ch chan<- core.RemovedLogsEvent) event.Subscription {
	return b.epv.BlockChain().SubscribeRemovedLogsEvent(ch)
}

func (b *EPVApiBackend) SubscribeChainEvent(ch chan<- core.ChainEvent) event.Subscription {
	return b.epv.BlockChain().SubscribeChainEvent(ch)
}

func (b *EPVApiBackend) SubscribeChainHeadEvent(ch chan<- core.ChainHeadEvent) event.Subscription {
	return b.epv.BlockChain().SubscribeChainHeadEvent(ch)
}

func (b *EPVApiBackend) SubscribeChainSideEvent(ch chan<- core.ChainSideEvent) event.Subscription {
	return b.epv.BlockChain().SubscribeChainSideEvent(ch)
}

func (b *EPVApiBackend) SubscribeLogsEvent(ch chan<- []*types.Log) event.Subscription {
	return b.epv.BlockChain().SubscribeLogsEvent(ch)
}

func (b *EPVApiBackend) SendTx(ctx context.Context, signedTx *types.Transaction) error {
	return b.epv.txPool.AddLocal(signedTx)
}

func (b *EPVApiBackend) GetPoolTransactions() (types.Transactions, error) {
	pending, err := b.epv.txPool.Pending()
	if err != nil {
		return nil, err
	}
	var txs types.Transactions
	for _, batch := range pending {
		txs = append(txs, batch...)
	}
	return txs, nil
}

func (b *EPVApiBackend) GetPoolTransaction(hash common.Hash) *types.Transaction {
	return b.epv.txPool.Get(hash)
}

func (b *EPVApiBackend) GetPoolNonce(ctx context.Context, addr common.Address) (uint64, error) {
	return b.epv.txPool.State().GetNonce(addr), nil
}

func (b *EPVApiBackend) Stats() (pending int, queued int) {
	return b.epv.txPool.Stats()
}

func (b *EPVApiBackend) TxPoolContent() (map[common.Address]types.Transactions, map[common.Address]types.Transactions) {
	return b.epv.TxPool().Content()
}

func (b *EPVApiBackend) SubscribeTxPreEvent(ch chan<- core.TxPreEvent) event.Subscription {
	return b.epv.TxPool().SubscribeTxPreEvent(ch)
}

func (b *EPVApiBackend) Downloader() *downloader.Downloader {
	return b.epv.Downloader()
}

func (b *EPVApiBackend) ProtocolVersion() int {
	return b.epv.EPVVersion()
}

func (b *EPVApiBackend) SuggestPrice(ctx context.Context) (*big.Int, error) {
	return b.gpo.SuggestPrice(ctx)
}

func (b *EPVApiBackend) ChainDb() epvdb.Database {
	return b.epv.ChainDb()
}

func (b *EPVApiBackend) EventMux() *event.TypeMux {
	return b.epv.EventMux()
}

func (b *EPVApiBackend) AccountManager() *accounts.Manager {
	return b.epv.AccountManager()
}

func (b *EPVApiBackend) BloomStatus() (uint64, uint64) {
	sections, _, _ := b.epv.bloomIndexer.Sections()
	return params.BloomBitsBlocks, sections
}

func (b *EPVApiBackend) ServiceFilter(ctx context.Context, session *bloombits.MatcherSession) {
	for i := 0; i < bloomFilterThreads; i++ {
		go session.Multiplex(bloomRetrievalBatch, bloomRetrievalWait, b.epv.bloomRequests)
	}
}
