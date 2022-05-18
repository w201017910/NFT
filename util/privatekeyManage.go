package util

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

var base = "D:\\桌面\\NFT\\"

func BindOptsByKeystore(keyfile string, passwd string) (*bind.TransactOpts, error) {
	reader, _ := os.Open(keyfile)
	opts, err := bind.NewTransactor(reader, passwd)
	if err != nil {
		log.Fatal("NewTransactor", err)
		return nil, err
	}
	return opts, nil
}
func NewAccounts(passwd string) (publicKey string, keystoreName string, mnemonics string, err error) {
	mnemonic := Create_mneonic()
	key := keystore.NewKeyStore("./keyStoreFile", keystore.StandardScryptN, keystore.StandardScryptP)
	account, err := key.NewAccount(passwd)
	address := account.Address.String()
	path := account.URL.Path
	path, _ = filepath.Rel(base, path)
	path = "./" + path
	key.Wallets()
	if err != nil {
		return "", "", "", err
	}
	return address, path, mnemonic, nil
}
func ImportAccountByPrivateKey(privateKey string, passwd string) (keystoreName string, mnemonics string, err string) {
	mnemonic := Create_mneonic()
	key := keystore.NewKeyStore("./keyStoreFile", keystore.StandardScryptN, keystore.StandardScryptP)
	privateKey_, _ := crypto.HexToECDSA(privateKey)
	account, _ := key.ImportECDSA(privateKey_, passwd)
	path := account.URL.Path
	path, _ = filepath.Rel(base, path)
	path = "./" + path
	if path == "" {
		return "", "", "该账户已存在"
	}
	return path, mnemonic, ""
	//pk, _ := crypto.HexToECDSA(privateKey)
	//path := "keyStoreFile"
	//hdKeyStore := NewHDkeyStore(path, pk)
	//fmt.Println("hdKeyStore.keysDirPath: ", hdKeyStore.keysDirPath)
	//err := hdKeyStore.StoreKey(hdKeyStore.keysDirPath, &hdKeyStore.Key, passwd)
	//if err != nil {
	//	log.Fatal("StoreKey: ", err)
	//}
	//fmt.Println(NewAccount(hdKeyStore.keysDirPath, passwd))
}
func ImportAccountByKetstoreFile(keyfile string, passwd string) (address string, mnemonics string, err error) {
	mnemonic := Create_mneonic()
	file, err := ioutil.ReadFile(keyfile)
	if err != nil {
		log.Fatal("readKeystore: ", err)
		return "", "", err
	}
	key, err := keystore.DecryptKey(file, passwd)
	if err != nil {
		log.Fatal("DecryptKey", err)
		return "", "", err
	}
	publicKey := crypto.PubkeyToAddress(key.PrivateKey.PublicKey).String()
	//fmt.Println(hexutil.Encode(crypto.FromECDSA(key.PrivateKey)))
	return publicKey, mnemonic, nil
}
func ExportPrivateKey(keyfile string, passwd string) (string, error) {
	file, err := ioutil.ReadFile(keyfile)
	if err != nil {
		log.Fatal("readKeystore", err)
		return "", err
	}
	key, err := keystore.DecryptKey(file, passwd)
	if err != nil {
		log.Fatal("DecryptKey", err)
		return "", err
	}
	return hexutil.Encode(crypto.FromECDSA(key.PrivateKey)), nil
}
