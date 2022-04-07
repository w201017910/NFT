package util

import (
	"crypto/ecdsa"
	"crypto/rand"
	"errors"
	"fmt"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil/hdkeychain"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/google/uuid"
	"github.com/tyler-smith/go-bip39"
	"io"
	"log"
)

func Create_mneonic() string { //生成助记词
	b, err := bip39.NewEntropy(128)
	if err != nil {
		log.Panic(err)
	}
	nm, err := bip39.NewMnemonic(b)
	fmt.Println(nm)
	return nm
}
func DeriveAddressFromMnemonic(nm string) (string, string) { //输入助记词生成账户和私钥
	path, err := accounts.ParseDerivationPath("m/44'/60'/0'/0/1")
	if err != nil {
		panic(err)
	}
	seed, err := bip39.NewSeedWithErrorChecking(nm, "")
	masterKey, err := hdkeychain.NewMaster(seed, &chaincfg.MainNetParams)
	if err != nil {
		fmt.Println(err)
		return "", ""
	}
	privateKey, err := DerivePrivateKey(path, masterKey)
	publicKey, err := DerivePublicKey(privateKey)
	address := crypto.PubkeyToAddress(*publicKey)

	return hexutil.Encode(crypto.FromECDSA(privateKey)), address.Hex()
}
func DerivePrivateKey(path accounts.DerivationPath, masterKey *hdkeychain.ExtendedKey) (*ecdsa.PrivateKey, error) {
	var err error
	key := masterKey
	for _, n := range path {
		key, err = key.Child(n)
		if err != nil {
			return nil, err
		}
	}
	privateKey, err := key.ECPrivKey()
	privateKeyECDSA := privateKey.ToECDSA()
	if err != nil {
		return nil, err
	}
	return privateKeyECDSA, nil
}
func DerivePublicKey(privateKey *ecdsa.PrivateKey) (*ecdsa.PublicKey, error) {
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, errors.New("failed")
	}
	return publicKeyECDSA, nil
}

type HDkeyStore struct {
	keysDirPath string
	scryptN     int
	scryptP     int
	Key         keystore.Key
}
type Key struct {
	Id         uuid.UUID
	Address    common.Address
	PrivateKey *ecdsa.PrivateKey
}
type UUID []byte

var rander = rand.Reader

func NewRandom() UUID {
	uuid := make([]byte, 16)
	io.ReadFull(rand.Reader, uuid)
	uuid[6] = (uuid[6] & 0x0f) | 0x40
	uuid[8] = (uuid[8] & 0x3f) | 0x80
	return uuid
}
