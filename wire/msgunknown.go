// Copyright (c) 2014-2015 The btcsuite developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package wire

import (
	"fmt"
	"io"
	"io/ioutil"
)

const (
	// MaxUnknownDataSize is the maximum byte size of a data
	// element to add to the Bloom filter.  It is equal to the
	// maximum element size of a script.
	MaxMessageDataSize = 0x02000000
)

// MsgUnknown implements the Message interface and represents a bitcoin
// unknown message.
type MsgUnknown struct {
	Cmd  string
	Data []byte
}

// BtcDecode decodes r using the bitcoin protocol encoding into the receiver.
// This is part of the Message interface implementation.
func (msg *MsgUnknown) BtcDecode(r io.Reader, pver uint32, enc MessageEncoding) error {
	var err error
	msg.Data, err = ioutil.ReadAll(r)
	return err
}

// BtcEncode encodes the receiver to w using the bitcoin protocol encoding.
// This is part of the Message interface implementation.
func (msg *MsgUnknown) BtcEncode(w io.Writer, pver uint32, enc MessageEncoding) error {
	size := len(msg.Data)
	if size > MaxMessageDataSize {
		str := fmt.Sprintf("unknown size too large for message "+
			"[size %v, max %v]", size, MaxMessageDataSize)
		return messageError("MsgUnknown.BtcEncode", str)
	}

	return WriteVarBytes(w, pver, msg.Data)
}

// Command returns the protocol command string for the message.  This is part
// of the Message interface implementation.
func (msg *MsgUnknown) Command() string {
	return msg.Cmd
}

// MaxPayloadLength returns the maximum length the payload can be for the
// receiver.  This is part of the Message interface implementation.
func (msg *MsgUnknown) MaxPayloadLength(pver uint32) uint32 {
	return uint32(VarIntSerializeSize(MaxMessageDataSize)) +
		MaxMessageDataSize
}

// NewMsgUnknown returns a new bitcoin unknown message that conforms to the
// Message interface.  See MsgUnknown for details.
func NewMsgUnknown(command string, data []byte) *MsgUnknown {
	return &MsgUnknown{
		Cmd:  command,
		Data: data,
	}
}
