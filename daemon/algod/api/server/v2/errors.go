// Copyright (C) 2019-2022 Algorand, Inc.
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

package v2

var (
	errAppDoesNotExist                         = "application does not exist"
	errAssetDoesNotExist                       = "asset does not exist"
	errAccountAppDoesNotExist                  = "account application info not found"
	errAccountAssetDoesNotExist                = "account asset info not found"
	errBoxDoesNotExist                         = "box not found"
	errFailedLookingUpLedger                   = "failed to retrieve information from the ledger"
	errFailedLookingUpTransactionPool          = "failed to retrieve information from the transaction pool"
	errFailedRetrievingAccountDeltas           = "failed retrieving account deltas"
	errFailedRetrievingNodeStatus              = "failed retrieving node status"
	errFailedRetrievingLatestBlockHeaderStatus = "failed retrieving latests block header"
	errFailedRetrievingSyncRound               = "failed retrieving sync round from ledger"
	errFailedSettingSyncRound                  = "failed to set sync round on the ledger"
	errSyncModeNotEnabled                      = "sync mode must be enabled"
	errFailedParsingFormatOption               = "failed to parse the format option"
	errFailedToParseAddress                    = "failed to parse the address"
	errFailedToParseExclude                    = "failed to parse exclude"
	errFailedToParseTransaction                = "failed to parse transaction"
	errFailedToParseBlock                      = "failed to parse block"
	errFailedToParseCert                       = "failed to parse cert"
	errFailedToParseSourcemap                  = "failed to parse sourcemap"
	errFailedToEncodeResponse                  = "failed to encode response"
	errInternalFailure                         = "internal failure"
	errNoValidTxnSpecified                     = "no valid transaction ID was specified"
	errInvalidHashType                         = "invalid hash type"
	errTransactionNotFound                     = "could not find the transaction in the transaction pool or in the last 1000 confirmed rounds"
	errServiceShuttingDown                     = "operation aborted as server is shutting down"
	errRequestedRoundInUnsupportedRound        = "requested round would reach only after the protocol upgrade which isn't supported"
	errFailedToParseCatchpoint                 = "failed to parse catchpoint"
	errFailedToAbortCatchup                    = "failed to abort catchup : %v"
	errFailedToStartCatchup                    = "failed to start catchup : %v"
	errOperationNotAvailableDuringCatchup      = "operation not available during catchup"
	errRESTPayloadZeroLength                   = "payload was of zero length"
	errRoundGreaterThanTheLatest               = "given round is greater than the latest round"
)
