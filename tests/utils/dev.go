// Copyright 2020 ChainSafe Systems (ON) Corp.
// This file is part of gossamer.
//
// The gossamer library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The gossamer library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the gossamer library. If not, see <http://www.gnu.org/licenses/>.

package utils

import (
	"encoding/binary"
	"strconv"
	"testing"
	"time"

	"github.com/ChainSafe/gossamer/lib/common"
	"github.com/stretchr/testify/require"
)

// PauseBABE calls the endpoint dev_control with the params ["babe", "stop"]
func PauseBABE(t *testing.T, node *Node) error {
	_, err := PostRPC(DevControl, NewEndpoint(node.RPCPort), "[\"babe\", \"stop\"]")
	return err
}

// SlotDuration Calls dev endpoint for slot duration
func SlotDuration(t *testing.T, node *Node) time.Duration {
	slotDuration, err := PostRPC("dev_slotDuration", NewEndpoint(node.RPCPort), "[]")

	if err != nil {
		require.NoError(t, err)
	}

	slotDurationDecoded := new(string)
	err = DecodeRPC(t, slotDuration, slotDurationDecoded)
	require.NoError(t, err)

	slotDurationParsed := binary.LittleEndian.Uint64(common.MustHexToBytes(*slotDurationDecoded))
	duration, err := time.ParseDuration(strconv.Itoa(int(slotDurationParsed)) + "ms")
	require.NoError(t, err)
	return duration
}

// EpochLength Calls dev endpoint for epoch length
func EpochLength(t *testing.T, node *Node) uint64 {
	epochLength, err := PostRPC("dev_epochLength", NewEndpoint(node.RPCPort), "[]")

	if err != nil {
		require.NoError(t, err)
	}

	epochLengthDecoded := new(string)
	err = DecodeRPC(t, epochLength, epochLengthDecoded)
	require.NoError(t, err)

	epochLengthParsed := binary.LittleEndian.Uint64(common.MustHexToBytes(*epochLengthDecoded))
	return epochLengthParsed
}
