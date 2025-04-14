package main

import (
	"bufio"
	"fmt"
	"os"

	certgo "github.com/Alonza0314/cert-go"
	logger "github.com/Alonza0314/logger-go"
)

var signCertYmlPath = "./signCertCfg.yml"

func main() {
	fmt.Print("Enter passphrase for the root certificate private key: ")
	reader := bufio.NewReader(os.Stdin)
	passphrase, _ := reader.ReadString('\n')
	// 去除 passphrase 中的換行符號
	passphrase = passphrase[:len(passphrase)-1]

	logger.Info("SignRootCertificate", "signing root certificate")

	if _, err := certgo.SignRootCertificate(signCertYmlPath, passphrase); err != nil {
		return
	}

	logger.Info("SignRootCertificate", "root certificate signed, you can see the root certificate in ./root_cert.pem")
}
