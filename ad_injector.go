package main

import (
	"io"
	"net/http"
	"time"
)

const BitrateKbps = 128

func StreamAd(w http.ResponseWriter, adURL string) error {
	resp, err := http.Get(adURL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	buf := make([]byte, 4096)
	delay := time.Duration(float64(len(buf)*8) / (BitrateKbps * 1000) * 1e9) // nanosecondes
	for {
		n, err := resp.Body.Read(buf)
		if n > 0 {
			w.Write(buf[:n])
			time.Sleep(delay)
		}
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}
	}
	return nil
}
