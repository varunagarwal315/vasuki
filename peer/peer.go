// Copyright 2014 The Metabase Authors
// This file is part of vasuki p2p library.
//
// The vasuki library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The vasuki library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package peer

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"math/bits"

	"github.com/Metabase-Network/vasuki/crypto/hash"
)

// ID is an identity of nodes, using its public key hash and network address.
type ID struct {
	PublicKey []byte
	Address   string
	Id        []byte
}

// CreateID is a factory function creating ID.
func CreateID(address string, publicKey []byte) ID {
	return ID{Address: address, PublicKey: publicKey, Id: hash.new().HashBytes(publicKey)}
}

// String returns the identity address and public key.
func (id ID) String() string {
	return fmt.Sprintf("ID{Address: %v, Id: %v}", id.Address, id.Id)
}

// Equals determines if two peer IDs are equal to each other based on the contents of their public keys.
func (id ID) Equals(other ID) bool {
	return bytes.Equal(id.Id, other.Id)
}

// Less determines if this peer ID's public key is less than other ID's public key.
func (id ID) Less(other interface{}) bool {
	if other, is := other.(ID); is {
		return bytes.Compare(id.Id, other.Id) == -1
	}
	return false
}

// PublicKeyHex generates a hex-encoded string of public key hash of this given peer ID.
func (id ID) PublicKeyHex() string {
	return hex.EncodeToString(id.PublicKey)
}

// Xor performs XOR (^) over another peer ID's public key.
func (id ID) Xor(other ID) ID {
	result := make([]byte, len(id.PublicKey))

	for i := 0; i < len(id.PublicKey) && i < len(other.PublicKey); i++ {
		result[i] = id.PublicKey[i] ^ other.PublicKey[i]
	}
	return ID{Address: id.Address, PublicKey: result}
}

// XorID performs XOR (^) over another peer ID's public key hash.
func (id ID) XorID(other ID) ID {
	result := make([]byte, len(id.Id))

	for i := 0; i < len(id.Id) && i < len(other.Id); i++ {
		result[i] = id.Id[i] ^ other.Id[i]
	}
	return ID{Address: id.Address, Id: result}
}

// PrefixLen returns the number of prefixed zeros in a peer ID.
func (id ID) PrefixLen() int {
	for i, b := range id.Id {
		if b != 0 {
			return i*8 + bits.LeadingZeros8(uint8(b))
		}
	}
	return len(id.Id)*8 - 1
}
