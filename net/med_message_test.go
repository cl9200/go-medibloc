package net

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type msgFields struct {
	chainID     uint32
	reserved    []byte
	version     byte
	messageName string
	data        []byte
}

func TestNewMedMessageWithLimits(t *testing.T) {
	tests := []struct {
		name          string
		field         msgFields
		expectedError error
	}{
		{
			"ExceedMaxDataLength",
			msgFields(msgFields{
				100,
				[]byte{0x0, 0x0, 0x0},
				0,
				"messageName",
				make([]byte, MaxMedMessageDataLength+1),
			}),
			ErrExceedMaxDataLength,
		},
		{
			"ExceedMaxMessageNameLength",
			msgFields(msgFields{
				100,
				[]byte{0x0, 0x0, 0x0},
				0,
				string(make([]byte, MaxMedMessageNameLength+1)),
				[]byte(""),
			}),
			ErrExceedMaxMessageNameLength,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			msg, err := NewMedMessage(
				tt.field.chainID,
				tt.field.reserved,
				tt.field.version,
				tt.field.messageName,
				tt.field.data,
			)
			assert.Nil(t, msg)
			assert.Equal(t, tt.expectedError, err)
		})
	}
}

func TestNewMedMessage(t *testing.T) {
	tests := []struct {
		name            string
		field           msgFields
		expectedContent []byte
	}{
		{
			"NewMedMessage1",
			msgFields(msgFields{
				100,
				[]byte{0x0, 0x0, 0x0},
				0,
				"hello",
				[]byte("hello"),
			}),
			[]byte{
				0x4e, 0x45, 0x42, 0x31,
				0x0, 0x0, 0x0, 0x64,
				0x0, 0x0, 0x0, 0x0,
				0x68, 0x65, 0x6c, 0x6c,
				0x6f, 0x0, 0x0, 0x0,
				0x0, 0x0, 0x0, 0x0,
				0x0, 0x0, 0x0, 0x5,
				0x36, 0x10, 0xa6, 0x86,
				0xef, 0x78, 0xba, 0x4e,
				0x68, 0x65, 0x6c, 0x6c,
				0x6f},
		},
		{
			"NewMedMessage2",
			msgFields(msgFields{
				1001,
				[]byte{0x0, 0x0, 0x1},
				1,
				"ok",
				[]byte("ok"),
			}),
			[]byte{
				0x4e, 0x45, 0x42, 0x31,
				0x0, 0x0, 0x3, 0xe9,
				0x0, 0x0, 0x1, 0x1,
				0x6f, 0x6b, 0x0, 0x0,
				0x0, 0x0, 0x0, 0x0,
				0x0, 0x0, 0x0, 0x0,
				0x0, 0x0, 0x0, 0x2,
				0x79, 0xdc, 0xdd, 0x47,
				0x5c, 0xdd, 0xb5, 0xe0,
				0x6f, 0x6b},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			msg, err := NewMedMessage(
				tt.field.chainID,
				tt.field.reserved,
				tt.field.version,
				tt.field.messageName,
				tt.field.data,
			)
			assert.Nil(t, err)
			assert.Equal(t, tt.expectedContent, msg.content)
			messageName := msg.MessageName()
			assert.Equal(t, tt.field.messageName, messageName)
		})
	}
}

func TestMedMessage_VerifyHeader(t *testing.T) {
	tests := []struct {
		name          string
		content       []byte
		expectedError error
	}{
		{
			"ValidMedMessage1",
			[]byte{
				0x4e, 0x45, 0x42, 0x31,
				0x0, 0x0, 0x0, 0x64,
				0x0, 0x0, 0x0, 0x0,
				0x68, 0x65, 0x6c, 0x6c,
				0x6f, 0x0, 0x0, 0x0,
				0x0, 0x0, 0x0, 0x0,
				0x0, 0x0, 0x0, 0x5,
				0x36, 0x10, 0xa6, 0x86,
				0xef, 0x78, 0xba, 0x4e,
				0x68, 0x65, 0x6c, 0x6c,
				0x6f},
			nil,
		},
		{
			"InvalidMagicNumberMedMessage1",
			[]byte{
				0x4e, 0x45, 0x42, 0x30,
				0x0, 0x0, 0x0, 0x64,
				0x0, 0x0, 0x0, 0x0,
				0x68, 0x65, 0x6c, 0x6c,
				0x6f, 0x0, 0x0, 0x0,
				0x0, 0x0, 0x0, 0x0,
				0x0, 0x0, 0x0, 0x5,
				0x36, 0x10, 0xa6, 0x86,
				0xef, 0x78, 0xba, 0x4e,
				0x68, 0x65, 0x6c, 0x6c,
				0x6f},
			ErrInvalidMagicNumber,
		},
		{
			"InvalidHeaderCheckSumMedMessage1",
			[]byte{
				0x4e, 0x45, 0x42, 0x31,
				0x0, 0x0, 0x0, 0x64,
				0x0, 0x0, 0x0, 0x0,
				0x68, 0x65, 0x6c, 0x6c,
				0x6f, 0x0, 0x0, 0x0,
				0x0, 0x0, 0x0, 0x0,
				0x0, 0x0, 0x0, 0x5,
				0x36, 0x10, 0xa6, 0x86,
				0xef, 0x78, 0xba, 0x4f,
				0x68, 0x65, 0x6c, 0x6c,
				0x6f},
			ErrInvalidHeaderCheckSum,
		},
		{
			"ExceedMaxDataMedMessage1",
			[]byte{
				0x4e, 0x45, 0x42, 0x31,
				0x0, 0x0, 0x0, 0x64,
				0x0, 0x0, 0x0, 0x0,
				0x68, 0x65, 0x6c, 0x6c,
				0x6f, 0x0, 0x0, 0x0,
				0x0, 0x0, 0x0, 0x0,
				0x20, 0x00, 0x00, 0x01,
				0x36, 0x10, 0xa6, 0x86,
				0xe3, 0x8e, 0x7e, 0xd8,
				0x68, 0x65, 0x6c, 0x6c,
				0x6f},
			ErrExceedMaxDataLength,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			msg := &MedMessage{
				content: make([]byte, len(tt.content)),
			}
			copy(msg.content, tt.content)
			err := msg.VerifyHeader()
			assert.Equal(t, tt.expectedError, err)
		})
	}
}

func TestMedMessage_VerifyData(t *testing.T) {
	tests := []struct {
		name          string
		content       []byte
		expectedError error
	}{
		{
			"ValidMedMessage1",
			[]byte{
				0x4e, 0x45, 0x42, 0x31,
				0x0, 0x0, 0x0, 0x64,
				0x0, 0x0, 0x0, 0x0,
				0x68, 0x65, 0x6c, 0x6c,
				0x6f, 0x0, 0x0, 0x0,
				0x0, 0x0, 0x0, 0x0,
				0x0, 0x0, 0x0, 0x5,
				0x36, 0x10, 0xa6, 0x86,
				0xef, 0x78, 0xba, 0x4e,
				0x68, 0x65, 0x6c, 0x6c,
				0x6f},
			nil,
		},
		{
			"InvalidDataCheckSumMedMessage1",
			[]byte{
				0x4e, 0x45, 0x42, 0x31,
				0x0, 0x0, 0x0, 0x64,
				0x0, 0x0, 0x0, 0x0,
				0x68, 0x65, 0x6c, 0x6c,
				0x6f, 0x0, 0x0, 0x0,
				0x0, 0x0, 0x0, 0x0,
				0x0, 0x0, 0x0, 0x5,
				0x36, 0x10, 0xa6, 0x87,
				0x98, 0x7f, 0x8a, 0xd8,
				0x68, 0x65, 0x6c, 0x6c,
				0x6f},
			ErrInvalidDataCheckSum,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			msg := &MedMessage{
				content: make([]byte, len(tt.content)),
			}
			copy(msg.content, tt.content)
			err := msg.VerifyData()
			assert.Equal(t, tt.expectedError, err)
		})
	}
}