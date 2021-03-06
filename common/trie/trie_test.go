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

package trie_test

import (
	"encoding/hex"
	"testing"

	"github.com/medibloc/go-medibloc/common/trie"
	"github.com/medibloc/go-medibloc/storage"
	"github.com/stretchr/testify/assert"
)

func TestNewTrie(t *testing.T) {
	type args struct {
		rootHash []byte
	}
	s, err := storage.NewMemoryStorage()
	assert.Nil(t, err)
	tests := []struct {
		args    args
		name    string
		storage storage.Storage
		wantErr bool
	}{
		{
			args{nil},
			"test1",
			s,
			false,
		},
		{
			args{[]byte("root")},
			"test2",
			s,
			true,
		},
	}

	for _, t2 := range tests {
		t.Run(t2.name, func(t3 *testing.T) {
			tr, err := trie.NewTrie(t2.args.rootHash, t2.storage)
			assert.Equal(t3, t2.wantErr, err != nil)
			if err == nil {
				assert.Equal(t3, tr.RootHash(), t2.args.rootHash)
			}
		})
	}
}

func TestTrie_Operations(t *testing.T) {
	s, err := storage.NewMemoryStorage()
	assert.Nil(t, err)

	tr, err := trie.NewTrie(nil, s)
	assert.Nil(t, err)
	assert.NotNil(t, tr)

	key1, _ := hex.DecodeString("2ed1b10c")
	val1 := []byte("leafmedi1")
	err = tr.Put(key1, val1)
	assert.Nil(t, err)
	curRootHash := []byte{0xe5, 0xa2, 0x8b, 0x11, 0x2e, 0x6f, 0x65, 0x91, 0xbd, 0xa4, 0x37, 0xd3, 0x86, 0xa9, 0x19, 0xd9, 0xb3, 0x5d, 0x9d, 0x86, 0x3d, 0x6e, 0xd6, 0x5b, 0x19, 0x3, 0x6c, 0x7, 0x6f, 0x25, 0xd4, 0x85}
	assert.Equal(t, curRootHash, tr.RootHash())

	val2, err := tr.Get(key1)
	assert.Nil(t, err)
	assert.Equal(t, val2, val1)

	err = tr.Delete(key1)
	assert.Nil(t, err)
	assert.Nil(t, tr.RootHash())

	err = tr.Put(key1, val1)
	assert.Nil(t, err)

	key2, _ := hex.DecodeString("2ed1c000")
	val2 = []byte("leafmedi2")
	err = tr.Put(key2, val2)
	assert.Nil(t, err)

	err = tr.Delete(key2)
	assert.Nil(t, err)

	val2, err = tr.Get(key1)
	assert.Nil(t, err)
	assert.Equal(t, val2, val1)
	// TODO make it same topology after put & delete
	// assert.Equal(t, curRootHash, tr.RootHash())

	err = tr.Delete(key1)
	assert.Nil(t, err)
	tr.Get(key2)
	assert.Nil(t, tr.RootHash())
}

type kvs struct {
	key []byte
	val []byte
}

func getKVs() []kvs {
	return []kvs{
		{[]byte("aaabbb"), []byte("val1")},
		{[]byte("aaaccc"), []byte("val2")},
		{[]byte("zzzccc"), []byte("val3")},
	}
}

func TestUpdateExtBranch(t *testing.T) {
	kvs := getKVs()
	s, err := storage.NewMemoryStorage()
	assert.Nil(t, err)

	tr, err := trie.NewTrie(nil, s)
	assert.Nil(t, err)
	assert.NotNil(t, tr)

	for _, kv := range kvs {
		err = tr.Put(kv.key, kv.val)
		assert.Nil(t, err)
	}
	for _, kv := range kvs {
		val, err := tr.Get(kv.key)
		assert.Nil(t, err)
		assert.EqualValues(t, kv.val, val)
	}
}

func TestTrie_Clone(t *testing.T) {
	s, err := storage.NewMemoryStorage()
	assert.Nil(t, err)

	tr1, err := trie.NewTrie(nil, s)
	assert.Nil(t, err)

	tr2, err := tr1.Clone()
	assert.Nil(t, err)

	kvs := getKVs()
	for _, kv := range kvs {
		err = tr1.Put(kv.key, kv.val)
		assert.Nil(t, err)
	}
	assert.Nil(t, tr2.RootHash())
	for _, kv := range kvs {
		err = tr2.Put(kv.key, kv.val)
		assert.Nil(t, err)
	}
	assert.Equal(t, tr1.RootHash(), tr2.RootHash())
}
