package main

import (
	"encoding/xml"
	"io"
	"net/http"
)

type VAST struct {
	Ad struct {
		InLine struct {
			Creatives struct {
				Creative struct {
					Linear struct {
						MediaFiles struct {
							MediaFile struct {
								URL string `xml:",chardata"`
							} `xml:"MediaFile"`
						} `xml:"MediaFiles"`
					} `xml:"Linear"`
				} `xml:"Creative"`
			} `xml:"Creatives"`
		} `xml:"InLine"`
	} `xml:"Ad"`
}

func GetAdAudioURL(vastURL string) (string, error) {
	resp, err := http.Get(vastURL)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	var vast VAST
	if err := xml.Unmarshal(body, &vast); err != nil {
		return "", err
	}

	return vast.Ad.InLine.Creatives.Creative.Linear.MediaFiles.MediaFile.URL, nil
}
