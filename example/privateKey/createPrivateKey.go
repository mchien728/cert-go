package main

import (
	"bufio"
	"fmt"
	"os"

	certgo "github.com/Alonza0314/cert-go"
	logger "github.com/Alonza0314/logger-go"
)

var privateKeyPath = "./private_key.pem"

func main() {
	fmt.Print("Enter passphrase for the root certificate private key: ")
	reader := bufio.NewReader(os.Stdin)
	passphrase, _ := reader.ReadString('\n')
	// 去除 passphrase 中的換行符號
	passphrase = passphrase[:len(passphrase)-1]

	logger.Info("CreatePrivateKey", "creating private key")

	if _, err := certgo.CreatePrivateKey(privateKeyPath, passphrase); err != nil {
		logger.Error("SignRootCertificate", "failed to sign certificate: "+err.Error())
		return
	}

	logger.Info("CreatePrivateKey", "private key created, you can see the private key in "+privateKeyPath)
}
