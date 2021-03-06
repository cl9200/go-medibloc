// Copyright (C) 2018  MediBloc
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>

package dpos

import (
	"errors"
	"time"
)

// Consensus properties.
const (
	BlockInterval   = 15 * time.Second
	DynastyInterval = 210 * BlockInterval
	DynastySize     = 21
	MinMintDuration = 2 * time.Second

	miningTickInterval = time.Second
)

// Error types of dpos package.
var (
	ErrInvalidBlockInterval         = errors.New("invalid block interval")
	ErrInvalidBlockProposer         = errors.New("invalid block proposer")
	ErrInvalidBlockForgeTime        = errors.New("invalid time to forge block")
	ErrFoundNilProposer             = errors.New("found a nil proposer")
	ErrInvalidProtoToConsensusState = errors.New("protobuf message cannot be converted into ConsensusState")
	ErrBlockMintedInNextSlot        = errors.New("cannot mint block now, there is a block minted in current slot")
	ErrWaitingBlockInLastSlot       = errors.New("cannot mint block now, waiting for last block")
	ErrInvalidDynastySize           = errors.New("invalid dynasty size")
)
