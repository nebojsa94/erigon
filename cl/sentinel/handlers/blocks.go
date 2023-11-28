/*
   Copyright 2022 Erigon-Lightclient contributors
   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at
       http://www.apache.org/licenses/LICENSE-2.0
   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

package handlers

import (
	"github.com/nebojsa94/erigon/cl/sentinel/communication/ssz_snappy"
	"github.com/ledgerwatch/log/v3"
	"github.com/libp2p/go-libp2p/core/network"
)

// func (c *ConsensusHandlers) blocksByRangeHandlerPROTODONOTTOUCH69(stream network.Stream) error {
// 	log.Trace("Got block by range handler call")
// 	return ssz_snappy.EncodeAndWrite(stream, &emptyString{}, ResourceUnavaiablePrefix)
// }

func (c *ConsensusHandlers) blocksByRangeHandler(stream network.Stream) error {
	log.Trace("Got block by range handler call")
	return ssz_snappy.EncodeAndWrite(stream, &emptyString{}, ResourceUnavaiablePrefix)
}

func (c *ConsensusHandlers) beaconBlocksByRootHandler(stream network.Stream) error {
	log.Trace("Got beacon block by root handler call")
	return ssz_snappy.EncodeAndWrite(stream, &emptyString{}, ResourceUnavaiablePrefix)
}

type emptyString struct{}

func (e *emptyString) EncodeSSZ(xs []byte) ([]byte, error) {
	return append(xs, 0), nil
}

func (e *emptyString) EncodingSizeSSZ() int {
	return 1
}