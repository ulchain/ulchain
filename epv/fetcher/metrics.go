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

// Contains the metrics collected by the fetcher.

package fetcher

import (
	"github.com/epvchain/go-epvchain/metrics"
)

var (
	propAnnounceInMeter   = metrics.NewMeter("epv/fetcher/prop/announces/in")
	propAnnounceOutTimer  = metrics.NewTimer("epv/fetcher/prop/announces/out")
	propAnnounceDropMeter = metrics.NewMeter("epv/fetcher/prop/announces/drop")
	propAnnounceDOSMeter  = metrics.NewMeter("epv/fetcher/prop/announces/dos")

	propBroadcastInMeter   = metrics.NewMeter("epv/fetcher/prop/broadcasts/in")
	propBroadcastOutTimer  = metrics.NewTimer("epv/fetcher/prop/broadcasts/out")
	propBroadcastDropMeter = metrics.NewMeter("epv/fetcher/prop/broadcasts/drop")
	propBroadcastDOSMeter  = metrics.NewMeter("epv/fetcher/prop/broadcasts/dos")

	headerFetchMeter = metrics.NewMeter("epv/fetcher/fetch/headers")
	bodyFetchMeter   = metrics.NewMeter("epv/fetcher/fetch/bodies")

	headerFilterInMeter  = metrics.NewMeter("epv/fetcher/filter/headers/in")
	headerFilterOutMeter = metrics.NewMeter("epv/fetcher/filter/headers/out")
	bodyFilterInMeter    = metrics.NewMeter("epv/fetcher/filter/bodies/in")
	bodyFilterOutMeter   = metrics.NewMeter("epv/fetcher/filter/bodies/out")
)
