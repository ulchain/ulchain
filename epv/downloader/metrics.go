
package downloader

import (
	"github.com/epvchain/go-epvchain/disk"
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
