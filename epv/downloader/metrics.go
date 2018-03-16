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

// Contains the metrics collected by the downloader.

package downloader

import (
	"github.com/epvchain/go-epvchain/metrics"
)

var (
	headerInMeter      = metrics.NewMeter("epv/downloader/headers/in")
	headerReqTimer     = metrics.NewTimer("epv/downloader/headers/req")
	headerDropMeter    = metrics.NewMeter("epv/downloader/headers/drop")
	headerTimeoutMeter = metrics.NewMeter("epv/downloader/headers/timeout")

	bodyInMeter      = metrics.NewMeter("epv/downloader/bodies/in")
	bodyReqTimer     = metrics.NewTimer("epv/downloader/bodies/req")
	bodyDropMeter    = metrics.NewMeter("epv/downloader/bodies/drop")
	bodyTimeoutMeter = metrics.NewMeter("epv/downloader/bodies/timeout")

	receiptInMeter      = metrics.NewMeter("epv/downloader/receipts/in")
	receiptReqTimer     = metrics.NewTimer("epv/downloader/receipts/req")
	receiptDropMeter    = metrics.NewMeter("epv/downloader/receipts/drop")
	receiptTimeoutMeter = metrics.NewMeter("epv/downloader/receipts/timeout")

	stateInMeter   = metrics.NewMeter("epv/downloader/states/in")
	stateDropMeter = metrics.NewMeter("epv/downloader/states/drop")
)
