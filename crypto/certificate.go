package crypto

import (
	"bytes"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"errors"
)

// function to get user identity
func ParseCert(certificateBytes []byte) (pkix.Name, error) {
	if certificateBytes == nil {
		return pkix.Name{}, errors.New("Nil certificate bytes passed")
	}

	certStart := bytes.IndexAny(certificateBytes, "----BEGIN CERTIFICATE-----")
	if certStart == -1 {
		//logger.Debug("No certificate found")
		return pkix.Name{}, errors.New("No Certificate found")
	}
	certText := certificateBytes[certStart:]
	block, _ := pem.Decode(certText)
	if block == nil {
		//logger.Debug("Error received on pem.Decode of certificate",  certText)
		return pkix.Name{}, errors.New("Error received on pem.Decode of certificate text")
	}

	ucert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		//logger.Debug("Error received on ParseCertificate", err)
		return pkix.Name{}, errors.New("Error received on ParseCertificate")
	}
	return ucert.Subject, nil
}
