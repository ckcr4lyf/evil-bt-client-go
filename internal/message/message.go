package message

import (
	"encoding/binary"
	"encoding/hex"
	"math/rand"
)

//TODO: Seed with torrent hash
func MakeHandshakeMessage(infohash string) []byte {

	buf := make([]byte, 68)

	buf[0] = uint8(19)
	copy(buf[1:20], "BitTorrent protocol")

	infohashBytes, _ := hex.DecodeString(infohash)
	copy(buf[28:48], infohashBytes)

	//TODO: Random peer Id
	peerId := make([]byte, 20)
	rand.Read(peerId)
	copy(buf[48:68], peerId)

	return buf
}

func MakeInterestedMessage() []byte {
	buf := make([]byte, 5)
	binary.BigEndian.PutUint32(buf[0:], 0x01)
	buf[4] = uint8(0x02)
	return buf
}

func MakeRequest(index uint32, offset uint32) []byte {
	buf := make([]byte, 17)
	binary.BigEndian.PutUint32(buf[0:], 0x0D)
	buf[4] = 0x06
	binary.BigEndian.PutUint32(buf[5:], index)
	binary.BigEndian.PutUint32(buf[9:], offset)
	binary.BigEndian.PutUint32(buf[13:], 0x4000)
	return buf
}
