package main

import (
	"crypto/tls"
	"net"
	"net/http"
	"net/url"

	utls "github.com/refraction-networking/utls"
)

func buildClient(proxyStr string, identifier string) *http.Client {

	var transport *http.Transport

	if proxyStr != "" {
		proxyURL, _ := url.Parse(proxyStr)
		transport = &http.Transport{
			Proxy: http.ProxyURL(proxyURL),
		}
	} else {
		transport = &http.Transport{}
	}

	// TLS Fingerprint selection
	switch identifier {
	case "chrome":
		transport.DialTLS = dialTLSWithFingerprint(utls.HelloChrome_Auto)
	case "firefox":
		transport.DialTLS = dialTLSWithFingerprint(utls.HelloFirefox_Auto)
	case "safari":
		transport.DialTLS = dialTLSWithFingerprint(utls.HelloSafari_Auto)
	case "safari_ios":
		transport.DialTLS = dialTLSWithFingerprint(utls.HelloIOS_Auto)
	case "okhttp":
		transport.DialTLS = dialTLSWithFingerprint(utls.HelloAndroid_11_OkHttp)
	default:
		transport.TLSClientConfig = &tls.Config{}
	}

	return &http.Client{Transport: transport}
}

func dialTLSWithFingerprint(profile utls.ClientHelloID) func(string, string) (net.Conn, error) {

	return func(network, addr string) (net.Conn, error) {

		conn, err := net.Dial(network, addr)
		if err != nil {
			return nil, err
		}

		config := &utls.Config{
			ServerName: "",
		}

		uconn := utls.UClient(conn, config, profile)
		err = uconn.Handshake()

		return uconn, err
	}
}
