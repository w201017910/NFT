package util

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"io/ioutil"
	"os"
	"path/filepath"
)

var base = "D:\\桌面\\NFT\\"

func BindOptsByKeystore(keyfile string, passwd string) (*bind.TransactOpts, error) {
	reader, _ := os.Open(keyfile)
	opts, err := bind.NewTransactor(reader, passwd)
	if err != nil {
		fmt.Println("NewTransactor", err)
		return nil, err
	}
	return opts, nil
}
func NewAccounts(passwd string) (publicKey string, keystoreName string, err error) {
	key := keystore.NewKeyStore("./keyStoreFile", keystore.StandardScryptN, keystore.StandardScryptP)
	account, err := key.NewAccount(passwd)
	address := account.Address.String()
	path := account.URL.Path
	path, _ = filepath.Rel(base, path)
	path = "./" + path
	key.Wallets()
	if err != nil {
		return "", "", err
	}
	return address, path, nil
}
func ImportAccountByPrivateKey(privateKey string, passwd string) (keystoreName string, err string) {
	key := keystore.NewKeyStore("./keyStoreFile", keystore.StandardScryptN, keystore.StandardScryptP)
	privateKey_, _ := crypto.HexToECDSA(privateKey)
	account, _ := key.ImportECDSA(privateKey_, passwd)
	path := account.URL.Path
	path, _ = filepath.Rel(base, path)
	path = "./" + path
	if path == "" {
		return "", "该账户已存在"
	}
	return path, ""
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
func ImportAccountByKetstoreFile(keyfile string, passwd string) (address string, err error) {
	file, fileErr := ioutil.ReadFile(keyfile)
	if fileErr != nil {
		fmt.Println("readKeystore: ", fileErr)
		return "", err
	}
	key, DecryptKeyErr := keystore.DecryptKey(file, passwd)
	if DecryptKeyErr != nil {
		fmt.Println("DecryptKeyErr : ", DecryptKeyErr)
		return "", DecryptKeyErr
	}
	publicKey := crypto.PubkeyToAddress(key.PrivateKey.PublicKey).String()
	//fmt.Println(hexutil.Encode(crypto.FromECDSA(key.PrivateKey)))
	return publicKey, nil
}
func ExportPrivateKey(keyfile string, passwd string) (string, error) {
	file, err := ioutil.ReadFile(keyfile)
	if err != nil {
		fmt.Println("readKeystore", err)
		return "", err
	}
	key, err := keystore.DecryptKey(file, passwd)
	if err != nil {
		fmt.Println("DecryptKey", err)
		return "", err
	}
	return hexutil.Encode(crypto.FromECDSA(key.PrivateKey)), nil
}
