package util

import (
	"fmt"
	shell "github.com/ipfs/go-ipfs-api"
	"os"
)

func UpLoad(fileName string) (imageCid string) {
	r, err := os.Open(fileName)
	defer r.Close()
	//ipfsPort := "localhost:5001"
	ipfsPort := "175.178.215.53:5001"
	sh := shell.NewShell(ipfsPort)
	cid, err := sh.Add(r)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s", err)
		os.Exit(1)
	}
	fmt.Printf("added %s", cid)
	return cid
}

func DelImage(fileName string) (status error) {
	err := os.Remove(fileName)
	if err != nil {
		fmt.Println(err)
	}
	return err
}
