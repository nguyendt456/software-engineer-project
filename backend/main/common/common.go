package common

import (
	"crypto/tls"
	"crypto/x509"
	"os"
)

func LoadCertificateAndKey(CAcertificate string, certFile string, keyFile string) *tls.Config {
	CAcert, err := os.ReadFile(CAcertificate)
	if err != nil {
		panic(err)
	}

	CAcertPool := x509.NewCertPool()
	if !CAcertPool.AppendCertsFromPEM(CAcert) {
		panic(err)
	}

	hostCertFile, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		panic(err)
	}

	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{hostCertFile},
		RootCAs:      CAcertPool,
		ClientCAs:    CAcertPool,
		ClientAuth:   tls.RequireAndVerifyClientCert,
	}
	return tlsConfig
}
