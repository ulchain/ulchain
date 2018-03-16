// Copyright 2016 The go-epvchain Authors
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

package epvclient

import "github.com/epvchain/go-epvchain"

// Verify that Client implements the epvchain interfaces.
var (
	_ = epvchain.ChainReader(&Client{})
	_ = epvchain.TransactionReader(&Client{})
	_ = epvchain.ChainStateReader(&Client{})
	_ = epvchain.ChainSyncReader(&Client{})
	_ = epvchain.ContractCaller(&Client{})
	_ = epvchain.GasEstimator(&Client{})
	_ = epvchain.GasPricer(&Client{})
	_ = epvchain.LogFilterer(&Client{})
	_ = epvchain.PendingStateReader(&Client{})
	// _ = epvchain.PendingStateEventer(&Client{})
	_ = epvchain.PendingContractCaller(&Client{})
)
