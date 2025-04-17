package main

import (
	"io"
	"net/http"
)

func StreamAd(w http.ResponseWriter, adURL string) error {
	resp, err := http.Get(adURL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	buf := make([]byte, 4096)
	for {
		n, err := resp.Body.Read(buf)
		if n > 0 {
			w.Write(buf[:n])
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
