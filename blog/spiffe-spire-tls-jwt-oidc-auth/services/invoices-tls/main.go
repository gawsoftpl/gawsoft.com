package main

import (
	"crypto/x509"
	"crypto/tls"
	"log"
	"net/http"
	"os"
)

// checkSpiffeId checks if the client certificate contains the correct SPIFFE URI in its SAN
func checkSpiffeId(cert *x509.Certificate, spiffeURI string) bool {
	// Iterate through the Subject Alternative Names (SAN) in the certificate
	for _, san := range cert.URIs {
		if san.String() == spiffeURI {
			return true
		}
	}
	return false
}

// handler processes incoming requests and validates the client certificate
func handler(w http.ResponseWriter, r *http.Request) {
	// Get the client's certificate
	clientCert := r.TLS.PeerCertificates[0]

	// Get the allowed SPIFFE URI from the environment variable
	allowSpiffeURI := os.Getenv("ALLOW_SPIFFE_URI")
	if allowSpiffeURI == "" {
		// If ALLOW_SPIFFE_URI is not set, return a server error
		http.Error(w, "ALLOW_SPIFFE_URI environment variable not set", http.StatusInternalServerError)
		return
	}

	// Check if the client's certificate contains the correct SPIFFE URI
	if checkSpiffeId(clientCert, allowSpiffeURI) {
		// If valid, grant access
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Access granted. Client certificate is valid."))
	} else {
		// If invalid, deny access
		http.Error(w, "Access denied. SPIFFE URI is invalid.", http.StatusForbidden)
	}
}

func main() {
	// Read paths to certificates and CA certificate from environment variables
	certFile := os.Getenv("CERT_FILE")    // Server certificate path
	keyFile := os.Getenv("KEY_FILE")      // Server private key path
	caCertFile := os.Getenv("CA_CERT_FILE") // CA certificate path

	// Check if any of the necessary environment variables are missing
	if certFile == "" || keyFile == "" || caCertFile == "" {
		log.Fatalf("Missing necessary environment variables (CERT_FILE, KEY_FILE, CA_CERT_FILE)")
	}

	// Read the CA certificate
	caCert, err := os.ReadFile(caCertFile)
	if err != nil {
		log.Fatalf("Error loading CA certificate: %v", err)
	}

	// Create a new certificate pool for the CA
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	// Configure TLS settings
	tlsConfig := &tls.Config{
		ClientCAs:    caCertPool,        // Use the CA certificate pool
		ClientAuth:   tls.RequireAndVerifyClientCert, // Require client certificates
		InsecureSkipVerify: false, // Ensure the server's certificate is validated
	}

	// Create the HTTPS server
	server := &http.Server{
		Addr:      ":443",  // Listen on port 443
		TLSConfig: tlsConfig,
		Handler:   http.HandlerFunc(handler),
	}

	// Start the server with TLS
	log.Println("HTTPS server running on port 443")
	err = server.ListenAndServeTLS(certFile, keyFile)
	if err != nil {
		log.Fatalf("Server error: %v", err)
	}
}
