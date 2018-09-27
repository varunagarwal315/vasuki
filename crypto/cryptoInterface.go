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

package crypto

import "math/big"

type cryptoInterface interface {
	generateKey() ([]byte, []byte, error)
	PrivateKeySize() int
	PrivateToPublic(privateKey []byte) ([]byte, error)
	PublicKeySize() int
	Sign(privateKey []byte, message []byte) []byte
	RandomKeyPair() *KeyPair
	Verify(publicKey []byte, message []byte, signature []byte) bool
}

type hashInterface interface {
	HashBytes(b []byte) []byte
}

// HashBytes returns a hash of a big integer given a hash policy.
func HashBytes(hp hashInterface, s *big.Int) *big.Int {
	return s.SetBytes(hp.HashBytes(s.Bytes()))
}