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

	"github.com/medibloc/go-medibloc/core"
	"github.com/medibloc/go-medibloc/util/byteutils"
	"github.com/medibloc/go-medibloc/util/testutil"
	"github.com/stretchr/testify/assert"
)

func TestTransactionManager(t *testing.T) {
	mgrs, closeFn := testutil.NewTestTransactionManagers(t, 2)
	defer closeFn()

	tx := testutil.NewRandomSignedTransaction(t)

	mgrs[0].Broadcast(tx)
	var actual *core.Transaction
	for actual == nil {
		actual = mgrs[1].Pop()
		time.Sleep(time.Millisecond)
	}
	assert.EqualValues(t, tx.Hash(), actual.Hash())

	tx = testutil.NewRandomSignedTransaction(t)
	mgrs[1].Relay(tx)
	actual = nil
	for actual == nil {
		actual = mgrs[0].Pop()
		time.Sleep(time.Millisecond)
	}
	assert.EqualValues(t, tx.Hash(), actual.Hash())
}

func TestTransactionManagerAbnormalTx(t *testing.T) {
	mgrs, closeFn := testutil.NewTestTransactionManagers(t, 2)
	defer closeFn()

	sender, receiver := mgrs[0], mgrs[1]

	// No signature
	noSign := testutil.NewRandomTransaction(t)
	expectTxFiltered(t, sender, receiver, noSign)

	// Invalid signature
	from, to := testutil.NewPrivateKey(t), testutil.NewPrivateKey(t)
	invalidSign := testutil.NewTransaction(t, from, to, 10)
	testutil.SignTx(t, invalidSign, to)
	expectTxFiltered(t, sender, receiver, invalidSign)
}

func TestTransactionManagerDupTxFromNet(t *testing.T) {
	mgrs, closeFn := testutil.NewTestTransactionManagers(t, 2)
	defer closeFn()

	sender, receiver := mgrs[0], mgrs[1]

	dup := testutil.NewRandomSignedTransaction(t)
	sender.Broadcast(dup)
	sender.Broadcast(dup)
	time.Sleep(100 * time.Millisecond)

	normal := testutil.NewRandomSignedTransaction(t)
	sender.Broadcast(normal)

	var count int
	for {
		recv := receiver.Pop()
		if recv != nil && byteutils.Equal(recv.Hash(), normal.Hash()) {
			break
		}
		if recv != nil {
			count++
		}
		time.Sleep(time.Millisecond)
	}
	assert.Equal(t, 1, count)
}

func TestTransactionManagerDupTxPush(t *testing.T) {
	mgrs, closeFn := testutil.NewTestTransactionManagers(t, 2)
	defer closeFn()

	dup := testutil.NewRandomSignedTransaction(t)
	err := mgrs[0].Push(dup)
	assert.NoError(t, err)
	err = mgrs[0].Push(dup)
	assert.EqualValues(t, core.ErrDuplicatedTransaction, err)

	actual := mgrs[0].Pop()
	assert.EqualValues(t, dup, actual)
	actual = mgrs[0].Pop()
	assert.Nil(t, actual)
}

func expectTxFiltered(t *testing.T, sender, receiver *core.TransactionManager, abnormal *core.Transaction) {
	sender.Broadcast(abnormal)

	time.Sleep(100 * time.Millisecond)

	normal := testutil.NewRandomSignedTransaction(t)
	sender.Broadcast(normal)

	var recv *core.Transaction
	for recv == nil {
		recv = receiver.Pop()
		time.Sleep(time.Millisecond)
	}
	assert.EqualValues(t, normal.Hash(), recv.Hash())
}
