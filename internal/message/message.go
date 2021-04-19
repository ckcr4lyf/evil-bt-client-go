package message

//TODO: Seed with torrent hash
func MakeHandshakeMessage() []byte {

	buf := make([]byte, 68)

	buf[0] = uint8(19)
	copy(buf[1:20], "BitTorrent protocol")

	return buf

	//TODO: Random peer Id
}
