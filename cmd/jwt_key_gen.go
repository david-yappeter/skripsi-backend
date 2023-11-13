//go:build tools

package cmd

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"

	"myapp/global"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(newJwtKeyGenCommand())
}

func newJwtKeyGenCommand() *cobra.Command {
	var isForce bool

	cmd := &cobra.Command{
		Use:   "jwt-key-gen",
		Short: "Generate rsa key pair for jwt",
		Long:  "This command is used to generate rsa key pair for jwt",
		Run: func(_ *cobra.Command, _ []string) {
			isFileExists := false
			if _, err := os.Stat(global.GetJwtPrivateKeyFilePath()); !errors.Is(err, os.ErrNotExist) {
				isFileExists = true
			}

			folderPath := filepath.Dir(global.GetJwtPrivateKeyFilePath())
			os.MkdirAll(folderPath, os.ModePerm)

			if isForce || !isFileExists {
				bitSize := 4096

				// Generate RSA key.
				key, err := rsa.GenerateKey(rand.Reader, bitSize)
				if err != nil {
					panic(err)
				}

				// Extract public component.
				pub := key.Public()

				// Encode private key to PKCS#1 ASN.1 PEM.
				keyPEM := pem.EncodeToMemory(
					&pem.Block{
						Type:  "RSA PRIVATE KEY",
						Bytes: x509.MarshalPKCS1PrivateKey(key),
					},
				)

				// Encode public key to PKCS#1 ASN.1 PEM.
				pubPEM := pem.EncodeToMemory(
					&pem.Block{
						Type:  "RSA PUBLIC KEY",
						Bytes: x509.MarshalPKCS1PublicKey(pub.(*rsa.PublicKey)),
					},
				)

				// Write private key to file.
				if err := ioutil.WriteFile(global.GetJwtPrivateKeyFilePath(), keyPEM, 0700); err != nil {
					panic(err)
				}

				// Write public key to file.
				if err := ioutil.WriteFile(global.GetJwtPublicKeyFilePath(), pubPEM, 0755); err != nil {
					panic(err)
				}
			}
		},
	}
	cmd.Flags().BoolVarP(&isForce, "force", "f", false, "force recreate rsa key pair for jwt")

	return cmd
}
