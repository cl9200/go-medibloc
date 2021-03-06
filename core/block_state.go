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

package core

import (
	"sort"

	"github.com/gogo/protobuf/proto"
	"github.com/medibloc/go-medibloc/common"
	"github.com/medibloc/go-medibloc/core/pb"
	"github.com/medibloc/go-medibloc/storage"
	"github.com/medibloc/go-medibloc/util"
	"github.com/medibloc/go-medibloc/util/logging"
	"github.com/sirupsen/logrus"
)

type states struct {
	accState           *AccountStateBatch
	txsState           *TrieBatch
	usageState         *TrieBatch
	recordsState       *TrieBatch
	consensusState     ConsensusState
	candidacyState     *TrieBatch
	certificationState *TrieBatch

	reservationQueue *ReservationQueue
	votesCache       *votesCache

	storage storage.Storage
}

func newStates(consensus Consensus, stor storage.Storage) (*states, error) {
	accState, err := NewAccountStateBatch(nil, stor)
	if err != nil {
		return nil, err
	}

	txsState, err := NewTrieBatch(nil, stor)
	if err != nil {
		return nil, err
	}

	usageState, err := NewTrieBatch(nil, stor)
	if err != nil {
		return nil, err
	}

	recordsState, err := NewTrieBatch(nil, stor)
	if err != nil {
		return nil, err
	}

	consensusState, err := consensus.NewConsensusState(nil, stor)
	if err != nil {
		return nil, err
	}

	candidacyState, err := NewTrieBatch(nil, stor)
	if err != nil {
		return nil, err
	}

	certificationState, err := NewTrieBatch(nil, stor)
	if err != nil {
		return nil, err
	}

	reservationQueue := NewEmptyReservationQueue(stor)
	votesCache := newVotesCache()

	return &states{
		accState:           accState,
		txsState:           txsState,
		usageState:         usageState,
		recordsState:       recordsState,
		consensusState:     consensusState,
		candidacyState:     candidacyState,
		certificationState: certificationState,
		reservationQueue:   reservationQueue,
		votesCache:         votesCache,
		storage:            stor,
	}, nil
}

func (st *states) Clone() (*states, error) {
	accState, err := NewAccountStateBatch(st.accState.RootHash(), st.storage)
	if err != nil {
		return nil, err
	}

	txsState, err := NewTrieBatch(st.txsState.RootHash(), st.storage)
	if err != nil {
		return nil, err
	}

	usageState, err := NewTrieBatch(st.usageState.RootHash(), st.storage)
	if err != nil {
		return nil, err
	}

	recordsState, err := NewTrieBatch(st.recordsState.RootHash(), st.storage)
	if err != nil {
		return nil, err
	}

	consensusState, err := st.consensusState.Clone()
	if err != nil {
		return nil, err
	}

	candidacyState, err := NewTrieBatch(st.candidacyState.RootHash(), st.storage)
	if err != nil {
		return nil, err
	}

	certificationState, err := NewTrieBatch(st.certificationState.RootHash(), st.storage)
	if err != nil {
		return nil, err
	}

	reservationQueue, err := LoadReservationQueue(st.storage, st.reservationQueue.Hash())
	if err != nil {
		return nil, err
	}

	votesCache := st.votesCache.Clone()

	return &states{
		accState:           accState,
		txsState:           txsState,
		usageState:         usageState,
		recordsState:       recordsState,
		consensusState:     consensusState,
		candidacyState:     candidacyState,
		certificationState: certificationState,
		reservationQueue:   reservationQueue,
		votesCache:         votesCache,
		storage:            st.storage,
	}, nil
}

func (st *states) BeginBatch() error {
	if err := st.accState.BeginBatch(); err != nil {
		return err
	}
	if err := st.txsState.BeginBatch(); err != nil {
		return err
	}
	if err := st.usageState.BeginBatch(); err != nil {
		return err
	}
	if err := st.recordsState.BeginBatch(); err != nil {
		return err
	}
	if err := st.candidacyState.BeginBatch(); err != nil {
		return err
	}
	if err := st.certificationState.BeginBatch(); err != nil {
		return err
	}
	return st.reservationQueue.BeginBatch()
}

func (st *states) Commit() error {
	if err := st.accState.Commit(); err != nil {
		return err
	}
	if err := st.txsState.Commit(); err != nil {
		return err
	}
	if err := st.usageState.Commit(); err != nil {
		return err
	}
	if err := st.recordsState.Commit(); err != nil {
		return err
	}
	if err := st.candidacyState.Commit(); err != nil {
		return err
	}
	if err := st.certificationState.Commit(); err != nil {
		return err
	}
	return st.reservationQueue.Commit()
}

func (st *states) AccountsRoot() []byte {
	return st.accState.RootHash()
}

func (st *states) TransactionsRoot() []byte {
	return st.txsState.RootHash()
}

func (st *states) UsageRoot() []byte {
	return st.usageState.RootHash()
}

func (st *states) RecordsRoot() []byte {
	return st.recordsState.RootHash()
}

func (st *states) ConsensusRoot() ([]byte, error) {
	return st.consensusState.RootBytes()
}

func (st *states) CandidacyRoot() []byte {
	return st.candidacyState.RootHash()
}

func (st *states) CertificationRoot() []byte {
	return st.certificationState.RootHash()
}

func (st *states) ReservationQueueHash() []byte {
	return st.reservationQueue.Hash()
}

func (st *states) LoadAccountsRoot(rootHash []byte) error {
	accState, err := NewAccountStateBatch(rootHash, st.storage)
	if err != nil {
		return err
	}
	st.accState = accState
	return nil
}

func (st *states) LoadTransactionsRoot(rootHash []byte) error {
	txsState, err := NewTrieBatch(rootHash, st.storage)
	if err != nil {
		return err
	}
	st.txsState = txsState
	return nil
}

func (st *states) LoadUsageRoot(rootHash []byte) error {
	usageState, err := NewTrieBatch(rootHash, st.storage)
	if err != nil {
		return err
	}
	st.usageState = usageState
	return nil
}

func (st *states) LoadRecordsRoot(rootHash []byte) error {
	recordsState, err := NewTrieBatch(rootHash, st.storage)
	if err != nil {
		return err
	}
	st.recordsState = recordsState
	return nil
}

func (st *states) LoadConsensusRoot(consensus Consensus, rootBytes []byte) error {
	consensusState, err := consensus.LoadConsensusState(rootBytes, st.storage)
	if err != nil {
		return err
	}
	st.consensusState = consensusState
	return nil
}

func (st *states) LoadCandidacyRoot(rootHash []byte) error {
	candidacyState, err := NewTrieBatch(rootHash, st.storage)
	if err != nil {
		return err
	}
	st.candidacyState = candidacyState
	return nil
}

func (st *states) LoadCertificationRoot(rootHash []byte) error {
	certificationState, err := NewTrieBatch(rootHash, st.storage)
	if err != nil {
		return err
	}
	st.certificationState = certificationState
	return nil
}

func (st *states) LoadReservationQueue(hash []byte) error {
	rq, err := LoadReservationQueue(st.storage, hash)
	if err != nil {
		return err
	}
	st.reservationQueue = rq
	return nil
}

func (st *states) ConstructVotesCache() error {
	votes := make(map[common.Address]*util.Uint128)
	votesCache := newVotesCache()

	accIter, err := st.accState.as.accounts.Iterator(nil)
	if err != nil {
		return err
	}
	exist, err := accIter.Next()
	for exist {
		if err != nil {
			return err
		}
		accBytes := accIter.Value()
		acc, err := loadAccount(accBytes)
		if err != nil {
			return err
		}
		votedAddr := common.BytesToAddress(acc.Voted())
		_, err = st.GetCandidate(votedAddr)
		if err != nil {
			exist, err = accIter.Next()
			continue
		}
		_, ok := votes[votedAddr]
		if ok {
			votes[votedAddr], err = votes[votedAddr].Add(acc.Vesting())
			if err != nil {
				return err
			}
		} else {
			votes[votedAddr] = acc.Vesting()
		}
		exist, err = accIter.Next()
	}
	candIter, err := st.candidacyState.Iterator(nil)
	if err != nil {
		return err
	}
	exist, err = candIter.Next()
	for exist {
		if err != nil {
			return err
		}
		candBytes := candIter.Value()
		pbCandidate := new(corepb.Candidate)
		if err := proto.Unmarshal(candBytes, pbCandidate); err != nil {
			return err
		}
		candAddr := common.BytesToAddress(pbCandidate.Address)
		if _, ok := votes[candAddr]; ok {
			votesCache.AddCandidate(candAddr, votes[candAddr])
		}
		exist, err = candIter.Next()
	}
	st.votesCache = votesCache
	return nil
}

func (st *states) GetAccount(address common.Address) (Account, error) {
	return st.accState.GetAccount(address.Bytes())
}

func (st *states) AddBalance(address common.Address, amount *util.Uint128) error {
	return st.accState.AddBalance(address.Bytes(), amount)
}

func (st *states) SubBalance(address common.Address, amount *util.Uint128) error {
	return st.accState.SubBalance(address.Bytes(), amount)
}

func (st *states) AddRecord(tx *Transaction, hash []byte, owner common.Address) error {
	record := &corepb.Record{
		Hash:      hash,
		Owner:     tx.from.Bytes(),
		Timestamp: tx.Timestamp(),
	}
	recordBytes, err := proto.Marshal(record)
	if err != nil {
		return err
	}

	if err := st.recordsState.Put(hash, recordBytes); err != nil {
		return err
	}

	return st.accState.AddRecord(tx.from.Bytes(), hash)
}

func (st *states) GetRecord(hash []byte) (*corepb.Record, error) {
	recordBytes, err := st.recordsState.Get(hash)
	if err != nil {
		return nil, err
	}
	pbRecord := new(corepb.Record)
	if err := proto.Unmarshal(recordBytes, pbRecord); err != nil {
		return nil, err
	}
	return pbRecord, nil
}

func (st *states) incrementNonce(address common.Address) error {
	return st.accState.IncrementNonce(address.Bytes())
}

func (st *states) GetTx(txHash []byte) ([]byte, error) {
	return st.txsState.Get(txHash)
}

func (st *states) PutTx(txHash []byte, txBytes []byte) error {
	return st.txsState.Put(txHash, txBytes)
}

func (st *states) updateUsage(tx *Transaction, blockTime int64) error {
	weekSec := int64(604800)

	if tx.Timestamp() < blockTime-weekSec {
		return ErrTooOldTransaction
	}

	payer, err := tx.recoverPayer()
	if err == ErrPayerSignatureNotExist {
		payer = tx.from
	} else if err != nil {
		logging.Console().WithFields(logrus.Fields{
			"err": err,
		}).Warn("Failed to recover payer address.")
		return err
	}

	usageBytes, err := st.usageState.Get(payer.Bytes())
	switch err {
	case nil:
	case ErrNotFound:
		usage := &corepb.Usage{
			Timestamps: []*corepb.TxTimestamp{
				{
					Hash:      tx.Hash(),
					Timestamp: tx.Timestamp(),
				},
			},
		}
		usageBytes, err = proto.Marshal(usage)
		if err != nil {
			logging.Console().WithFields(logrus.Fields{
				"usage": usage,
				"err":   err,
			}).Error("Failed to marshal usage.")
			return err
		}
		return st.usageState.Put(payer.Bytes(), usageBytes)
	default:
		logging.Console().WithFields(logrus.Fields{
			"payer": payer.Hex(),
			"err":   err,
		}).Error("Failed to get usage from trie.")
		return err
	}

	pbUsage := new(corepb.Usage)
	if err := proto.Unmarshal(usageBytes, pbUsage); err != nil {
		logging.Console().WithFields(logrus.Fields{
			"err": err,
			"pb":  pbUsage,
		}).Error("Failed to unmarshal proto.")
		return err
	}

	var idx int
	for idx = range pbUsage.Timestamps {
		if blockTime-weekSec < tx.Timestamp() {
			break
		}
	}
	pbUsage.Timestamps = append(pbUsage.Timestamps[idx:], &corepb.TxTimestamp{Hash: tx.Hash(), Timestamp: tx.Timestamp()})
	sort.Slice(pbUsage.Timestamps, func(i, j int) bool {
		return pbUsage.Timestamps[i].Timestamp < pbUsage.Timestamps[j].Timestamp
	})

	pbBytes, err := proto.Marshal(pbUsage)
	if err != nil {
		logging.Console().WithFields(logrus.Fields{
			"err": err,
			"pb":  pbUsage,
		}).Error("Failed to marshal proto.")
		return err
	}

	return st.usageState.Put(payer.Bytes(), pbBytes)
}

func (st *states) GetUsage(addr common.Address) ([]*corepb.TxTimestamp, error) {
	usageBytes, err := st.usageState.Get(addr.Bytes())
	switch err {
	case nil:
	case ErrNotFound:
		return []*corepb.TxTimestamp{}, nil
	default:
		return nil, err
	}

	pbUsage := new(corepb.Usage)
	if err := proto.Unmarshal(usageBytes, pbUsage); err != nil {
		return nil, err
	}
	return pbUsage.Timestamps, nil
}

// Dynasty returns members belonging to the current dynasty.
func (st *states) Dynasty() ([]*common.Address, error) {
	return st.consensusState.Dynasty()
}

// DynastySize returns size of dynasties.
func (st *states) DynastySize() int {
	return st.consensusState.DynastySize()
}

// SetDynasty sets dynasty members.
func (st *states) SetDynasty(miners []*common.Address, dynastySize int, startTime int64) error {
	return st.consensusState.InitDynasty(miners, dynastySize, startTime)
}

// Proposer returns address of block proposer set in consensus state
func (st *states) Proposer() common.Address {
	return st.consensusState.Proposer()
}

// TransitionDynasty transitions dynasty to a new one that is correct for the given time
func (st *states) TransitionDynasty(now int64) error {
	if st.consensusState.Timestamp() == GenesisTimestamp {
		cs, err := st.consensusState.GetNextStateAfterGenesis(now)
		if err != nil {
			return err
		}
		st.consensusState = cs
		return nil
	}
	cs, err := st.consensusState.GetNextStateAfter(now - st.consensusState.Timestamp())
	if err == nil {
		st.consensusState = cs
		return nil
	}
	if err != nil && err != ErrDynastyExpired {
		return err
	}
	var miners []*common.Address
	minerNum := st.consensusState.DynastySize()
	if len(st.votesCache.candidates) < minerNum {
		minerNum = len(st.votesCache.candidates)
	}
	for _, candidate := range st.votesCache.candidates {
		if candidate.candidacy {
			miners = append(miners, &candidate.address)
			if len(miners) == minerNum {
				break
			}
		}
	}
	if err := st.consensusState.InitDynasty(miners, minerNum, now); err != nil {
		return err
	}
	return nil
}

func (st *states) GetCandidate(address common.Address) (*corepb.Candidate, error) {
	candidateBytes, err := st.candidacyState.Get(address.Bytes())
	if err != nil {
		return nil, err
	}
	pbCandidate := new(corepb.Candidate)
	if err := proto.Unmarshal(candidateBytes, pbCandidate); err != nil {
		return nil, err
	}
	return pbCandidate, nil
}

// AddCandidate makes an address candidate
func (st *states) AddCandidate(address common.Address, collateral *util.Uint128) error {
	_, err := st.GetCandidate(address)
	if err != nil && err != ErrNotFound {
		return err
	}
	if err == nil {
		return ErrAlreadyInCandidacy
	}
	if err := st.SubBalance(address, collateral); err != nil {
		return err
	}
	collateralBytes, err := collateral.ToFixedSizeByteSlice()
	if err != nil {
		return err
	}
	pbCandidate := &corepb.Candidate{
		Address:    address.Bytes(),
		Collateral: collateralBytes,
	}
	candidateBytes, err := proto.Marshal(pbCandidate)
	if err != nil {
		return err
	}
	if err := st.candidacyState.Put(address.Bytes(), candidateBytes); err != nil {
		return err
	}
	if _, _, err := st.votesCache.GetCandidate(address); err == ErrCandidateNotFound {
		st.votesCache.AddCandidate(address, util.Uint128Zero())
		return nil
	}
	return st.votesCache.SetCandidacy(address, true)
}

// QuitCandidacy makes an account quit from candidacy
func (st *states) QuitCandidacy(address common.Address) error {
	candidate, err := st.GetCandidate(address)
	if err != nil {
		return err
	}
	collateral, err := util.NewUint128FromFixedSizeByteSlice(candidate.Collateral)
	if err != nil {
		return err
	}
	if err := st.AddBalance(address, collateral); err != nil {
		return err
	}
	if err := st.candidacyState.Delete(address.Bytes()); err != nil {
		return err
	}
	return st.votesCache.SetCandidacy(address, false)
}

// GetReservedTasks returns reserved tasks in reservation queue
func (st *states) GetReservedTasks() []*ReservedTask {
	return st.reservationQueue.Tasks()
}

// AddReservedTask adds a reserved task in reservation queue
func (st *states) AddReservedTask(task *ReservedTask) error {
	return st.reservationQueue.AddTask(task)
}

// PopReservedTask pops reserved tasks which should be processed before 'before'
func (st *states) PopReservedTasks(before int64) []*ReservedTask {
	return st.reservationQueue.PopTasksBefore(before)
}

func (st *states) PeekHeadReservedTask() *ReservedTask {
	return st.reservationQueue.Peek()
}

// WithdrawVesting makes multiple reserved tasks for withdraw a certain amount of vesting
func (st *states) WithdrawVesting(address common.Address, amount *util.Uint128, blockTime int64) error {
	acc, err := st.GetAccount(address)
	if err != nil {
		return err
	}
	if amount.Cmp(acc.Vesting()) > 0 {
		return ErrVestingNotEnough
	}
	splitAmount, err := amount.Div(util.NewUint128FromUint(RtWithdrawNum))
	if err != nil {
		return err
	}
	amountLeft := amount.DeepCopy()
	payload := new(RtWithdraw)
	for i := 0; i < RtWithdrawNum; i++ {
		if amountLeft.Cmp(splitAmount) <= 0 {
			payload, err = NewRtWithdraw(amountLeft)
			if err != nil {
				return err
			}
		} else {
			payload, err = NewRtWithdraw(splitAmount)
			if err != nil {
				return err
			}
		}
		task := NewReservedTask(RtWithdrawType, address, payload, blockTime+int64(i+1)*RtWithdrawInterval)
		if err := st.AddReservedTask(task); err != nil {
			return err
		}
		amountLeft, _ = amountLeft.Sub(splitAmount)
	}
	return nil
}

func (st *states) Vest(address common.Address, amount *util.Uint128) error {
	if err := st.accState.SubBalance(address.Bytes(), amount); err != nil {
		return err
	}
	if err := st.accState.AddVesting(address.Bytes(), amount); err != nil {
		return err
	}
	voted, err := st.GetVoted(address)
	if err == ErrNotVotedYet {
		return nil
	}
	if err != nil {
		return err
	}
	return st.votesCache.AddVotesPower(voted, amount)
}

func (st *states) SubVesting(address common.Address, amount *util.Uint128) error {
	acc, err := st.GetAccount(address)
	if err != nil {
		return err
	}
	voted := common.BytesToAddress(acc.Voted())
	if voted != (common.Address{}) {
		if err := st.votesCache.SubtractVotesPower(voted, amount); err != nil {
			return err
		}
	}
	return st.accState.SubVesting(address.Bytes(), amount)
}

func (st *states) Vote(address common.Address, voted common.Address) error {
	if _, err := st.GetCandidate(voted); err != nil {
		return err
	}
	acc, err := st.GetAccount(address)
	if err != nil {
		return err
	}
	oldVoted := common.BytesToAddress(acc.Voted())
	if oldVoted == voted {
		return ErrVoteDuplicate
	}
	if oldVoted != (common.Address{}) {
		if err := st.votesCache.SubtractVotesPower(oldVoted, acc.Vesting()); err != nil {
			return err
		}
	}
	if err := st.accState.SetVoted(address.Bytes(), voted.Bytes()); err != nil {
		return err
	}
	return st.votesCache.AddVotesPower(voted, acc.Vesting())
}

func (st *states) GetVoted(address common.Address) (common.Address, error) {
	votedBytes, err := st.accState.GetVoted(address.Bytes())
	if err != nil {
		return common.Address{}, err
	}
	return common.BytesToAddress(votedBytes), nil
}

func (st *states) AddCertification(hash []byte,
	issuerAddr common.Address, certifiedAddr common.Address,
	issueTime int64, expirationTime int64) error {
	if err := st.accState.AddCertReceived(certifiedAddr.Bytes(), hash); err != nil {
		return err
	}
	if err := st.accState.AddCertIssued(issuerAddr.Bytes(), hash); err != nil {
		return err
	}
	pbCertification := &corepb.Certification{
		CertificateHash: hash,
		Issuer:          issuerAddr.Bytes(),
		Certified:       certifiedAddr.Bytes(),
		IssueTime:       issueTime,
		ExpirationTime:  expirationTime,
		RevocationTime:  int64(0),
	}
	certificationBytes, err := proto.Marshal(pbCertification)
	if err != nil {
		return err
	}
	if err := st.certificationState.Put(hash, certificationBytes); err != nil {
		return err
	}
	return nil
}

func (st *states) RevokeCertification(hash []byte, revoker common.Address, revokeTime int64) error {
	certificationBytes, err := st.certificationState.Get(hash)
	if err != nil {
		return nil
	}
	pbCertification := new(corepb.Certification)
	if err := proto.Unmarshal(certificationBytes, pbCertification); err != nil {
		return err
	}
	if common.BytesToAddress(pbCertification.Issuer) != revoker {
		return ErrInvalidCertificationRevoker
	}
	if pbCertification.RevocationTime > int64(0) {
		return ErrCertAlreadyRevoked
	}
	pbCertification.RevocationTime = revokeTime
	modifiedBytes, err := proto.Marshal(pbCertification)
	if err != nil {
		return err
	}
	if err := st.certificationState.Put(hash, modifiedBytes); err != nil {
		return err
	}
	return nil
}

func (st *states) GetCertification(hash []byte) (*corepb.Certification, error) {
	certificationBytes, err := st.certificationState.Get(hash)
	if err != nil {
		return nil, err
	}
	pbCertification := new(corepb.Certification)
	if err := proto.Unmarshal(certificationBytes, pbCertification); err != nil {
		return nil, err
	}
	return pbCertification, nil
}

// BlockState possesses every states a block should have
type BlockState struct {
	*states
	snapshot *states
}

// NewBlockState creates a new block state
func NewBlockState(consensus Consensus, stor storage.Storage) (*BlockState, error) {
	states, err := newStates(consensus, stor)
	if err != nil {
		return nil, err
	}
	return &BlockState{
		states:   states,
		snapshot: nil,
	}, nil
}

// Clone clones block state
func (bs *BlockState) Clone() (*BlockState, error) {
	states, err := bs.states.Clone()
	if err != nil {
		return nil, err
	}
	return &BlockState{
		states:   states,
		snapshot: nil,
	}, nil
}

// BeginBatch begins batch
func (bs *BlockState) BeginBatch() error {
	snapshot, err := bs.states.Clone()
	if err != nil {
		return err
	}
	if err := bs.states.BeginBatch(); err != nil {
		return err
	}
	bs.snapshot = snapshot
	return nil
}

// RollBack rolls back batch
func (bs *BlockState) RollBack() error {
	bs.states = bs.snapshot
	bs.snapshot = nil
	return nil
}

// Commit saves batch updates
func (bs *BlockState) Commit() error {
	if err := bs.states.Commit(); err != nil {
		return err
	}
	bs.snapshot = nil
	return nil
}

// ExecuteTx and update internal states
func (bs *BlockState) ExecuteTx(tx *Transaction) error {
	return tx.ExecuteOnState(bs)
}

// AcceptTransaction and update internal txsStates
func (bs *BlockState) AcceptTransaction(tx *Transaction, blockTime int64) error {
	pbTx, err := tx.ToProto()
	if err != nil {
		logging.Console().WithFields(logrus.Fields{
			"err": err,
			"tx":  tx,
		}).Error("Failed to convert a transaction to proto.")
		return err
	}

	txBytes, err := proto.Marshal(pbTx)
	if err != nil {
		logging.Console().WithFields(logrus.Fields{
			"err": err,
			"pb":  pbTx,
		}).Error("Failed to marshal proto.")
		return err
	}

	if err := bs.PutTx(tx.hash, txBytes); err != nil {
		logging.Console().WithFields(logrus.Fields{
			"err": err,
			"tx":  tx,
		}).Error("Failed to put a transaction to block state.")
		return err
	}

	if err := bs.updateUsage(tx, blockTime); err != nil {
		logging.Console().WithFields(logrus.Fields{
			"err":       err,
			"tx":        tx,
			"blockTime": blockTime,
		}).Error("Failed to update usage.")
		return err
	}

	return bs.incrementNonce(tx.from)
}

func (bs *BlockState) checkNonce(tx *Transaction) error {
	fromAcc, err := bs.GetAccount(tx.from)
	if err != nil {
		return err
	}

	expectedNonce := fromAcc.Nonce() + 1
	if tx.nonce > expectedNonce {
		return ErrLargeTransactionNonce
	} else if tx.nonce < expectedNonce {
		return ErrSmallTransactionNonce
	}
	return nil
}
