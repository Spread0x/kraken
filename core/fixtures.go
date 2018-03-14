package core

import (
	"bytes"

	"code.uber.internal/infra/kraken/utils/randutil"
)

// BlobFixture joins all information associated with a blob for testing convenience.
type BlobFixture struct {
	Content  []byte
	Digest   Digest
	MetaInfo *MetaInfo
}

// CustomBlobFixture creates a BlobFixture with custom fields.
func CustomBlobFixture(content []byte, digest Digest, mi *MetaInfo) *BlobFixture {
	return &BlobFixture{content, digest, mi}
}

// SizedBlobFixture creates a randomly generated BlobFixture of given size with given piece lengths.
func SizedBlobFixture(size uint64, pieceLength uint64) *BlobFixture {
	b := randutil.Text(size)
	d, err := NewDigester().FromBytes(b)
	if err != nil {
		panic(err)
	}
	mi, err := NewMetaInfoFromBlob(d.Hex(), bytes.NewReader(b), int64(pieceLength))
	if err != nil {
		panic(err)
	}
	return &BlobFixture{
		Content:  b,
		Digest:   d,
		MetaInfo: mi,
	}
}

// NewBlobFixture creates a randomly generated BlobFixture.
func NewBlobFixture() *BlobFixture {
	return SizedBlobFixture(256, 8)
}

// PeerIDFixture returns a randomly generated PeerID.
func PeerIDFixture() PeerID {
	p, err := RandomPeerID()
	if err != nil {
		panic(err)
	}
	return p
}

// PeerInfoFixture returns a randomly generated PeerInfo.
func PeerInfoFixture() *PeerInfo {
	return NewPeerInfo(PeerIDFixture(), randutil.IP(), randutil.Port(), false, false)
}

// OriginPeerInfoFixture returns a randomly generated PeerInfo for an origin.
func OriginPeerInfoFixture() *PeerInfo {
	return NewPeerInfo(PeerIDFixture(), randutil.IP(), randutil.Port(), true, true)
}

// MetaInfoFixture returns a randomly generated MetaInfo.
func MetaInfoFixture() *MetaInfo {
	return NewBlobFixture().MetaInfo
}

// InfoHashFixture returns a randomly generated InfoHash.
func InfoHashFixture() InfoHash {
	return MetaInfoFixture().InfoHash
}

// DigestFixture returns a random Digest.
func DigestFixture() Digest {
	return NewBlobFixture().Digest
}

// PeerContextFixture returns a randomly generated PeerContext.
func PeerContextFixture() PeerContext {
	return PeerContext{
		IP:     randutil.IP(),
		Port:   randutil.Port(),
		PeerID: PeerIDFixture(),
		Zone:   "sjc1",
	}
}

// OriginContextFixture returns a randomly generated origin PeerContext.
func OriginContextFixture() PeerContext {
	octx := PeerContextFixture()
	octx.Origin = true
	return octx
}