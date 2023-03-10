package common

import (
	"crypto/tls"
	"crypto/x509"
	"os"

	"google.golang.org/grpc/credentials"
)

func LoadCertificateAndKey(CAcertificate string, certFile string, keyFile string) (credentials.TransportCredentials, error) {
	CAcert, err := os.ReadFile(CAcertificate)
	if err != nil {
		return nil, err
	}

	CAcertPool := x509.NewCertPool()
	if !CAcertPool.AppendCertsFromPEM(CAcert) {
		return nil, err
	}

	hostCertFile, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		return nil, err
	}

	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{hostCertFile},
		RootCAs:      CAcertPool,
		ClientCAs:    CAcertPool,
		ClientAuth:   tls.RequireAndVerifyClientCert,
	}
	return credentials.NewTLS(tlsConfig), nil
}
