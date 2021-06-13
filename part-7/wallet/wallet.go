package wallet

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"

	"github.com/golang-blockchain/part-7/utils"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/ripemd160"
)

type Wallet struct {
	PrivateKey ecdsa.PrivateKey
	PublicKey  []byte
}

func (w Wallet) Address() []byte {
	pubHash := PublicKeyHash(w.PublicKey)

	versionedHash := append([]byte{utils.Version}, pubHash...)
	checkSum := CheckSum(versionedHash)

	fullHash := append(versionedHash, checkSum...)
	address := Base58Encode(fullHash)

	log.Infof("PubKey: %x\n", w.PublicKey)
	log.Infof("PubHash: %x\n", pubHash)
	log.Infof("Address: %x\n", address)

	return address
}

func ValidateAddress(address string) bool {
	pubKeyHash := Base58Decode([]byte(address))
	actualCheckSum := pubKeyHash[len(pubKeyHash)-utils.ChecksumLength:]
	version := pubKeyHash[0]
	pubKeyHash = pubKeyHash[1 : len(pubKeyHash)-utils.ChecksumLength]
	targetCheckSum := CheckSum(append([]byte{version}, pubKeyHash...))

	return bytes.Compare(actualCheckSum, targetCheckSum) == 0
}

func NewKeyPair() (ecdsa.PrivateKey, []byte) {
	curve := elliptic.P256()
	private, err := ecdsa.GenerateKey(curve, rand.Reader)
	utils.Handle(err)

	pub := append(private.PublicKey.X.Bytes(), private.PublicKey.Y.Bytes()...)
	return *private, pub
}

func MakeWallet() *Wallet {
	private, public := NewKeyPair()
	wallet := Wallet{private, public}
	return &wallet
}

func PublicKeyHash(pubKey []byte) []byte {
	pubHash := sha256.Sum256(pubKey)

	hasher := ripemd160.New()
	_, err := hasher.Write(pubHash[:])
	utils.Handle(err)

	publicRipMD := hasher.Sum(nil)

	return publicRipMD
}

func CheckSum(payload []byte) []byte {
	firstHash := sha256.Sum256(payload)
	secondHash := sha256.Sum256(firstHash[:])

	return secondHash[:utils.ChecksumLength]
}
