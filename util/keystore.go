package util

import (
	"crypto/ecdsa"
	"crypto/rand"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/google/uuid"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

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
func change(uu []byte) [16]byte {
	var uu1 [16]byte
	for i, v := range uu {
		uu1[i] = v
	}
	return uu1
}
func NewHDkeyStore(path string, privateKey *ecdsa.PrivateKey) *HDkeyStore {
	uuid := []byte(NewRandom())

	key := keystore.Key{
		Id:         change(uuid),
		Address:    crypto.PubkeyToAddress(privateKey.PublicKey),
		PrivateKey: privateKey,
	}
	return &HDkeyStore{
		keysDirPath: path,
		scryptN:     keystore.LightScryptN,
		scryptP:     keystore.LightScryptP,
		Key:         key,
	}
}
func (ks HDkeyStore) StoreKey(filename string, key *keystore.Key, auth string) error {
	keyjson, err := keystore.EncryptKey(key, auth, ks.scryptN, ks.scryptP)
	if err != nil {
		return err
	}
	return WriteKeyFile(filename, keyjson)
}
func WriteKeyFile(file string, content []byte) error {
	const dirPerm = 0700
	if err := os.MkdirAll(filepath.Dir(file), dirPerm); err != nil {
		return err
	}
	f, err := ioutil.TempFile(filepath.Dir(file), "."+filepath.Base(file)+".tmp")
	if err != nil {
		return err
	}
	if _, err := f.Write(content); err != nil {
		f.Close()
		os.Remove(f.Name())
		return err
	}
	f.Close()
	return os.Rename(f.Name(), file)
}
func (ks HDkeyStore) JoinPath(filename string) string {
	if filepath.IsAbs(filename) {
		return filename
	}
	return filepath.Join(ks.keysDirPath, filename)
}
func (ks *HDkeyStore) GetKey(addr common.Address, filename string, auth string) (*keystore.Key, error) {
	keyjson, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	key, err := keystore.DecryptKey(keyjson, auth)
	if err != nil {
		return nil, err
	}
	if key.Address != addr {
		return nil, fmt.Errorf("key content mismatch: have account %x,want %x", key.Address, addr)
	}
	ks.Key = *key
	return key, nil
}

const defaultDerivationPath = "m/44'/60'/0'/0/1"

type HDWallet struct {
	Address    common.Address
	HDKeyStore *HDkeyStore
}

func NewWallet(keypath string) (*HDWallet, string, string, error) {
	mn := Create_mneonic()
	//if err!=nil {
	//	fmt.Println("Failed to NewAccount",err)
	//}
	privateKey, publicKey := DeriveAddressFromMnemonic(mn)
	address := crypto.PubkeyToAddress(*publicKey)
	hdks := NewHDkeyStore(keypath, privateKey)

	return &HDWallet{address, hdks}, address.String(), mn, nil
}
func (w HDWallet) StoreKey(pass string) error {
	filename := w.HDKeyStore.JoinPath(w.Address.Hex())
	return w.HDKeyStore.StoreKey(filename, &w.HDKeyStore.Key, pass)
}
func NewAccount(filepath string, pass string) (path string, address string, mneonic string) {
	w, address, mn, err := NewWallet(filepath)
	if err != nil {
		fmt.Println(err)
	}
	w.StoreKey(pass)
	return "./keystore" + address, address, mn
}
