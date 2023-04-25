// Copyright (C) 2019-2023 Algorand, Inc.
// This file is part of go-algorand
//
// go-algorand is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as
// published by the Free Software Foundation, either version 3 of the
// License, or (at your option) any later version.
//
// go-algorand is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with go-algorand.  If not, see <https://www.gnu.org/licenses/>.

// Check that devmode is functioning as designed.
package devmode

import (
	"math"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/algorand/go-algorand/crypto"
	"github.com/algorand/go-algorand/data/basics"
	"github.com/algorand/go-algorand/test/framework/fixtures"
	"github.com/algorand/go-algorand/test/partitiontest"
)

func TestDevMode(t *testing.T) {
	partitiontest.PartitionTest(t)

	if testing.Short() {
		t.Skip()
	}

	// Start devmode network, and make sure everything is primed by sending a transaction.
	var fixture fixtures.RestClientFixture
	fixture.SetupNoStart(t, filepath.Join("nettemplates", "DevModeNetwork.json"))
	fixture.Start()
	defer fixture.Shutdown()
	sender, err := fixture.GetRichestAccount()
	require.NoError(t, err)
	key := crypto.GenerateSignatureSecrets(crypto.Seed{})
	receiver := basics.Address(key.SignatureVerifier)
	txn := fixture.SendMoneyAndWait(0, 100000, 1000, sender.Address, receiver.String(), "")
	require.NotNil(t, txn.ConfirmedRound)
	firstRound := *txn.ConfirmedRound + 1
	blk, err := fixture.AlgodClient.Block(*txn.ConfirmedRound)
	require.NoError(t, err)
	seconds, _ := math.Modf(blk.Block["ts"].(float64))
	startTime := time.Unix(int64(seconds), 0)
	blkOffset := uint64(1_000_000)
	err = fixture.AlgodClient.SetBlockTimestampOffset(blkOffset)
	require.NoError(t, err)
	resp, err := fixture.AlgodClient.GetBlockTimestampOffset()
	require.NoError(t, err)
	require.Equal(t, blkOffset, resp.Offset)

	// 2 transactions should be sent within one normal confirmation time.
	for i := uint64(0); i < 2; i++ {
		round := firstRound + i
		txn = fixture.SendMoneyAndWait(round, 100001, 1000, sender.Address, receiver.String(), "")
		// SendMoneyAndWait subtracts 1 from firstValid
		require.Equal(t, round-1, uint64(txn.Txn.Txn.FirstValid))
		newBlk, err := fixture.AlgodClient.Block(round)
		require.NoError(t, err)
		newBlkSeconds, _ := math.Modf(newBlk.Block["ts"].(float64))
		require.GreaterOrEqual(t, time.Unix(int64(newBlkSeconds), 0), startTime.Add(1_000_000*time.Second))
	}

	// Without transactions there should be no rounds even after a normal confirmation time.
	time.Sleep(10 * time.Second)
	status, err := fixture.LibGoalClient.Status()
	require.NoError(t, err)
	require.Equal(t, *txn.ConfirmedRound, status.LastRound, "There should be no rounds without a transaction.")
}
