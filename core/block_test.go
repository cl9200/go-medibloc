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

package core_test

import (
	"testing"
	"time"

	"github.com/medibloc/go-medibloc/common"
	"github.com/medibloc/go-medibloc/consensus/dpos"
	"github.com/medibloc/go-medibloc/core"
	"github.com/medibloc/go-medibloc/crypto"
	"github.com/medibloc/go-medibloc/crypto/signature"
	"github.com/medibloc/go-medibloc/crypto/signature/algorithm"
	"github.com/medibloc/go-medibloc/medlet"
	"github.com/medibloc/go-medibloc/util"
	"github.com/medibloc/go-medibloc/util/testutil"
	"github.com/stretchr/testify/assert"
)

func TestNewBlock(t *testing.T) {
	genesis, dynasties, _ := testutil.NewTestGenesisBlock(t)

	coinbase := dynasties[0].Addr
	_, err := core.NewBlock(testutil.ChainID, coinbase, genesis)
	assert.NoError(t, err)
}

func TestSendExecution(t *testing.T) {
	genesis, dynasties, _ := testutil.NewTestGenesisBlock(t)

	coinbase := dynasties[0].Addr
	newBlock, err := core.NewBlock(testutil.ChainID, coinbase, genesis)
	assert.NoError(t, err)

	cases := []struct {
		from                 common.Address
		privKey              signature.PrivateKey
		to                   common.Address
		amount               *util.Uint128
		expectedResultAmount *util.Uint128
	}{
		{
			dynasties[1].Addr,
			dynasties[1].PrivKey,
			dynasties[2].Addr,
			util.NewUint128FromUint(10),
			util.NewUint128FromUint(1000000090),
		},
		{
			dynasties[2].Addr,
			dynasties[2].PrivKey,
			dynasties[0].Addr,
			util.NewUint128FromUint(20),
			util.NewUint128FromUint(999999990),
		},
		{
			dynasties[0].Addr,
			dynasties[0].PrivKey,
			dynasties[1].Addr,
			util.NewUint128FromUint(100),
			util.NewUint128FromUint(999999920),
		},
	}

	txs := make(core.Transactions, len(cases))
	signers := make([]signature.Signature, len(cases))

	for i, c := range cases {
		txs[i], err = core.NewTransaction(testutil.ChainID, c.from, c.to, c.amount, 1, core.TxPayloadBinaryType, []byte{})
		assert.NoError(t, err)

		signers[i], err = crypto.NewSignature(algorithm.SECP256K1)
		assert.NoError(t, err)
		signers[i].InitSign(c.privKey)
		assert.NoError(t, txs[i].SignThis(signers[i]))
	}

	newBlock.SetTransactions(txs)

	newBlock.BeginBatch()
	assert.NoError(t, newBlock.ExecuteAll())
	newBlock.Commit()
	assert.NoError(t, newBlock.Seal())

	coinbaseKey := dynasties[0].PrivKey

	blockSigner, err := crypto.NewSignature(algorithm.SECP256K1)
	assert.NoError(t, err)
	blockSigner.InitSign(coinbaseKey)

	assert.NoError(t, newBlock.SignThis(blockSigner))

	assert.NoError(t, newBlock.VerifyState())

	accStateBatch, err := core.NewAccountStateBatch(newBlock.AccountsRoot(), newBlock.Storage())
	assert.NoError(t, err)
	accState := accStateBatch.AccountState()

	for _, c := range cases {
		acc, err := accState.GetAccount(c.from.Bytes())
		assert.NoError(t, err)

		assert.Zero(t, acc.Balance().Cmp(c.expectedResultAmount))
	}

	assert.NoError(t, newBlock.VerifyIntegrity())
}

func TestSendMoreThanBalance(t *testing.T) {
	genesis, dynasties, _ := testutil.NewTestGenesisBlock(t)
	coinbase := dynasties[0].Addr
	newBlock, err := core.NewBlock(testutil.ChainID, coinbase, genesis)
	assert.NoError(t, err)

	fromKey := dynasties[1].PrivKey
	from, to := dynasties[1].Addr, dynasties[2].Addr

	balance := util.NewUint128FromUint(1000000090)
	sendingAmount, err := balance.Add(util.NewUint128FromUint(1))

	tx, err := core.NewTransaction(testutil.ChainID, from, to, sendingAmount, 1, core.TxPayloadBinaryType, []byte{})
	assert.NoError(t, err)

	signer, err := crypto.NewSignature(algorithm.SECP256K1)
	assert.NoError(t, err)
	signer.InitSign(fromKey)
	assert.NoError(t, tx.SignThis(signer))

	blockState := newBlock.State()

	newBlock.BeginBatch()
	assert.Equal(t, blockState.ExecuteTx(tx), core.ErrBalanceNotEnough)
	newBlock.RollBack()
}

func TestExecuteOnParentBlock(t *testing.T) {
	genesis, dynasties, _ := testutil.NewTestGenesisBlock(t)
	coinbase := dynasties[0].Addr
	firstBlock, err := core.NewBlock(testutil.ChainID, coinbase, genesis)
	assert.NoError(t, err)

	cases := []struct {
		from    common.Address
		privKey signature.PrivateKey
		to      common.Address
		amount  *util.Uint128
	}{
		{
			dynasties[1].Addr,
			dynasties[1].PrivKey,
			dynasties[2].Addr,
			util.NewUint128FromUint(1),
		},
		{
			dynasties[2].Addr,
			dynasties[2].PrivKey,
			dynasties[1].Addr,
			util.NewUint128FromUint(1000000001),
		},
	}

	txs := make(core.Transactions, len(cases))
	signers := make([]signature.Signature, len(cases))

	for i, c := range cases {
		txs[i], err = core.NewTransaction(testutil.ChainID, c.from, c.to, c.amount, 1, core.TxPayloadBinaryType, []byte{})
		assert.NoError(t, err)

		signers[i], err = crypto.NewSignature(algorithm.SECP256K1)
		assert.NoError(t, err)
		signers[i].InitSign(c.privKey)
		assert.NoError(t, txs[i].SignThis(signers[i]))
	}

	firstBlock.BeginBatch()
	assert.NoError(t, firstBlock.State().TransitionDynasty(firstBlock.Timestamp()))
	assert.NoError(t, firstBlock.ExecuteTransaction(txs[0]))
	assert.NoError(t, firstBlock.AcceptTransaction(txs[0]))
	firstBlock.Commit()

	assert.NoError(t, firstBlock.Seal())

	coinbaseKey := dynasties[0].PrivKey

	blockSigner, err := crypto.NewSignature(algorithm.SECP256K1)
	assert.NoError(t, err)
	blockSigner.InitSign(coinbaseKey)

	assert.NoError(t, firstBlock.SignThis(blockSigner))
	assert.NoError(t, firstBlock.VerifyState())

	secondBlock, err := core.NewBlock(testutil.ChainID, coinbase, firstBlock)
	assert.NoError(t, err)

	nextBlockTime := (firstBlock.Timestamp()/int64(dpos.DynastyInterval/time.Second) + 1) * int64(dpos.DynastyInterval/time.Second)
	deadline, err := dpos.CheckDeadline(firstBlock, time.Unix(nextBlockTime, 0))
	assert.NoError(t, err)
	secondBlock.BeginBatch()
	secondBlock.SetTimestamp(deadline.Unix())
	assert.NoError(t, secondBlock.State().TransitionDynasty(secondBlock.Timestamp()))
	assert.NoError(t, secondBlock.ExecuteTransaction(txs[1]))
	assert.NoError(t, secondBlock.AcceptTransaction(txs[1]))
	secondBlock.Commit()

	assert.NoError(t, secondBlock.Seal())

	assert.NoError(t, secondBlock.SignThis(blockSigner))
	assert.NoError(t, secondBlock.VerifyState())
	assert.Error(t, secondBlock.ExecuteAll())
	_, err = secondBlock.BlockData.ExecuteOnParentBlock(firstBlock)
	assert.NoError(t, err)
}

func TestGetExecutedBlock(t *testing.T) {
	genesis, dynasties, _ := testutil.NewTestGenesisBlock(t)
	coinbase := dynasties[0].Addr
	newBlock, err := core.NewBlock(testutil.ChainID, coinbase, genesis)
	assert.NoError(t, err)

	cases := []struct {
		from    common.Address
		privKey signature.PrivateKey
		to      common.Address
		amount  *util.Uint128
	}{
		{
			dynasties[1].Addr,
			dynasties[1].PrivKey,
			dynasties[0].Addr,
			util.NewUint128FromUint(1),
		},
	}

	txs := make(core.Transactions, len(cases))
	signers := make([]signature.Signature, len(cases))

	for i, c := range cases {
		txs[i], err = core.NewTransaction(testutil.ChainID, c.from, c.to, c.amount, 1, core.TxPayloadBinaryType, []byte{})
		assert.NoError(t, err)

		signers[i], err = crypto.NewSignature(algorithm.SECP256K1)
		assert.NoError(t, err)
		signers[i].InitSign(c.privKey)
		assert.NoError(t, txs[i].SignThis(signers[i]))
	}

	newBlock.BeginBatch()
	assert.NoError(t, newBlock.State().TransitionDynasty(newBlock.Timestamp()))
	assert.NoError(t, newBlock.ExecuteTransaction(txs[0]))
	assert.NoError(t, newBlock.AcceptTransaction(txs[0]))
	newBlock.Commit()

	assert.NoError(t, newBlock.Seal())

	coinbaseKey := dynasties[0].PrivKey

	blockSigner, err := crypto.NewSignature(algorithm.SECP256K1)
	assert.NoError(t, err)
	blockSigner.InitSign(coinbaseKey)

	assert.NoError(t, newBlock.SignThis(blockSigner))
	assert.NoError(t, newBlock.VerifyState())

	bd := newBlock.GetBlockData()

	cfg := medlet.DefaultConfig()
	cfg.Chain.BlockCacheSize = 1
	cfg.Chain.Coinbase = "02fc22ea22d02fc2469f5ec8fab44bc3de42dda2bf9ebc0c0055a9eb7df579056c"
	cfg.Chain.Miner = "02fc22ea22d02fc2469f5ec8fab44bc3de42dda2bf9ebc0c0055a9eb7df579056c"

	consensus, err := dpos.New(cfg)
	assert.NoError(t, err)
	executedBlock, err := bd.GetExecutedBlock(consensus, newBlock.Storage())
	assert.NoError(t, err)
	assert.NoError(t, executedBlock.VerifyState())
}

func TestExecuteReservedTasks(t *testing.T) {
	genesis, dynasties, _ := testutil.NewTestGenesisBlock(t)
	from := dynasties[0].Addr
	vestTx, err := core.NewTransaction(
		testutil.ChainID,
		from,
		common.Address{},
		util.NewUint128FromUint(333), 1,
		core.TxOperationVest, []byte{},
	)
	withdrawTx, err := core.NewTransaction(
		testutil.ChainID,
		from,
		common.Address{},
		util.NewUint128FromUint(333), 2,
		core.TxOperationWithdrawVesting, []byte{})
	assert.NoError(t, err)
	withdrawTx.SetTimestamp(int64(1000))

	privKey := dynasties[0].PrivKey
	assert.NoError(t, err)
	sig, err := crypto.NewSignature(algorithm.SECP256K1)
	assert.NoError(t, err)
	sig.InitSign(privKey)
	assert.NoError(t, vestTx.SignThis(sig))
	assert.NoError(t, withdrawTx.SignThis(sig))

	coinbase := from
	newBlock, err := core.NewBlock(testutil.ChainID, coinbase, genesis)
	assert.NoError(t, err)
	newBlock.SetTimestamp(int64(1000))

	newBlock.BeginBatch()
	assert.NoError(t, newBlock.ExecuteTransaction(vestTx))
	assert.NoError(t, newBlock.AcceptTransaction(vestTx))
	assert.NoError(t, newBlock.ExecuteTransaction(withdrawTx))
	assert.NoError(t, newBlock.AcceptTransaction(withdrawTx))
	assert.NoError(t, newBlock.ExecuteReservedTasks())
	newBlock.Commit()

	state := newBlock.State()

	acc, err := state.GetAccount(from)
	assert.NoError(t, err)
	assert.Equal(t, acc.Vesting(), util.NewUint128FromUint(uint64(333)))
	assert.Equal(t, acc.Balance(), util.NewUint128FromUint(uint64(1000000000-333)))
	tasks := state.GetReservedTasks()
	assert.Equal(t, 3, len(tasks))
	for i := 0; i < len(tasks); i++ {
		assert.Equal(t, core.RtWithdrawType, tasks[i].TaskType())
		assert.Equal(t, from, tasks[i].From())
		assert.Equal(t, withdrawTx.Timestamp()+int64(i+1)*core.RtWithdrawInterval, tasks[i].Timestamp())
	}

	newBlock.SetTimestamp(newBlock.Timestamp() + int64(2)*core.RtWithdrawInterval)
	newBlock.BeginBatch()
	assert.NoError(t, newBlock.ExecuteReservedTasks())
	newBlock.Commit()

	acc, err = state.GetAccount(from)
	assert.NoError(t, err)
	assert.Equal(t, acc.Vesting(), util.NewUint128FromUint(uint64(111)))
	assert.Equal(t, acc.Balance(), util.NewUint128FromUint(uint64(1000000000-111)))
	tasks = state.GetReservedTasks()
	assert.Equal(t, 1, len(tasks))
	assert.Equal(t, core.RtWithdrawType, tasks[0].TaskType())
	assert.Equal(t, from, tasks[0].From())
	assert.Equal(t, withdrawTx.Timestamp()+int64(3)*core.RtWithdrawInterval, tasks[0].Timestamp())
}

func TestBlock_VerifyState(t *testing.T) {
	genesis, dynasties, _ := testutil.NewTestGenesisBlock(t)
	wrongGenesis, _, _ := testutil.NewTestGenesisBlock(t)
	from, to := dynasties[0], dynasties[1]

	tx, err := core.NewTransaction(testutil.ChainID, from.Addr, to.Addr, util.NewUint128FromUint(100), 1, core.TxPayloadBinaryType, []byte{})
	assert.NoError(t, err)
	txSigner, err := crypto.NewSignature(algorithm.SECP256K1)
	assert.NoError(t, err)
	txSigner.InitSign(from.PrivKey)
	err = tx.SignThis(txSigner)
	assert.NoError(t, err)

	block, err := core.NewBlock(testutil.ChainID, from.Addr, genesis)
	assert.NoError(t, err)
	assert.NoError(t, block.SetTransactions(core.Transactions{tx}))
	assert.NoError(t, block.State().TransitionDynasty(block.Timestamp()))
	assert.NoError(t, block.ExecuteAll())
	assert.NoError(t, block.Seal())

	bd := block.GetBlockData()
	block, err = bd.ExecuteOnParentBlock(genesis)
	assert.NoError(t, err)
	assert.NoError(t, block.VerifyState())

	bd = block.GetBlockData()
	_, err = bd.ExecuteOnParentBlock(wrongGenesis)
	assert.Error(t, err)
}
