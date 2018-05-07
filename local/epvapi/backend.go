
package epvapi

import (
	"context"
	"math/big"

	"github.com/epvchain/go-epvchain/act"
	"github.com/epvchain/go-epvchain/public"
	"github.com/epvchain/go-epvchain/kernel"
	"github.com/epvchain/go-epvchain/kernel/state"
	"github.com/epvchain/go-epvchain/kernel/types"
	"github.com/epvchain/go-epvchain/kernel/vm"
	"github.com/epvchain/go-epvchain/epv/downloader"
	"github.com/epvchain/go-epvchain/data"
	"github.com/epvchain/go-epvchain/notice"
	"github.com/epvchain/go-epvchain/content"
	"github.com/epvchain/go-epvchain/remote"
)

type Backend interface {

	Downloader() *downloader.Downloader
	ProtocolVersion() int
	SuggestPrice(ctx context.Context) (*big.Int, error)
	ChainDb() epvdb.Database
	EventMux() *event.TypeMux
	AccountManager() *accounts.Manager

	SetHead(number uint64)
	HeaderByNumber(ctx context.Context, blockNr rpc.BlockNumber) (*types.Header, error)
	BlockByNumber(ctx context.Context, blockNr rpc.BlockNumber) (*types.Block, error)
	StateAndHeaderByNumber(ctx context.Context, blockNr rpc.BlockNumber) (*state.StateDB, *types.Header, error)
	GetBlock(ctx context.Context, blockHash common.Hash) (*types.Block, error)
	GetReceipts(ctx context.Context, blockHash common.Hash) (types.Receipts, error)
	GetTd(blockHash common.Hash) *big.Int
	GetEVM(ctx context.Context, msg core.Message, state *state.StateDB, header *types.Header, vmCfg vm.Config) (*vm.EVM, func() error, error)
	SubscribeChainEvent(ch chan<- core.ChainEvent) event.Subscription
	SubscribeChainHeadEvent(ch chan<- core.ChainHeadEvent) event.Subscription
	SubscribeChainSideEvent(ch chan<- core.ChainSideEvent) event.Subscription

	SendTx(ctx context.Context, signedTx *types.Transaction) error
	GetPoolTransactions() (types.Transactions, error)
	GetPoolTransaction(txHash common.Hash) *types.Transaction
	GetPoolNonce(ctx context.Context, addr common.Address) (uint64, error)
	Stats() (pending int, queued int)
	TxPoolContent() (map[common.Address]types.Transactions, map[common.Address]types.Transactions)
	SubscribeTxPreEvent(chan<- core.TxPreEvent) event.Subscription

	ChainConfig() *params.ChainConfig
	CurrentBlock() *types.Block
}

func GetAPIs(apiBackend Backend) []rpc.API {
	nonceLock := new(AddrLocker)
	return []rpc.API{
		{
			Namespace: "epv",
			Version:   "1.0",
			Service:   NewPublicEPVchainAPI(apiBackend),
			Public:    true,
		}, {
			Namespace: "epv",
			Version:   "1.0",
			Service:   NewPublicBlockChainAPI(apiBackend),
			Public:    true,
		}, {
			Namespace: "epv",
			Version:   "1.0",
			Service:   NewPublicTransactionPoolAPI(apiBackend, nonceLock),
			Public:    true,
		}, {
			Namespace: "txpool",
			Version:   "1.0",
			Service:   NewPublicTxPoolAPI(apiBackend),
			Public:    true,
		}, {
			Namespace: "debug",
			Version:   "1.0",
			Service:   NewPublicDebugAPI(apiBackend),
			Public:    true,
		}, {
			Namespace: "debug",
			Version:   "1.0",
			Service:   NewPrivateDebugAPI(apiBackend),
		}, {
			Namespace: "epv",
			Version:   "1.0",
			Service:   NewPublicAccountAPI(apiBackend.AccountManager()),
			Public:    true,
		}, {
			Namespace: "personal",
			Version:   "1.0",
			Service:   NewPrivateAccountAPI(apiBackend, nonceLock),
			Public:    false,
		},
	}
}
