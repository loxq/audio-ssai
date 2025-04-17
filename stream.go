package main

import (
	"encoding/binary"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/hajimehoshi/go-mp3"
)

const (
	FrequencyToDetect = 1000.0
	SampleRate        = 44100
	WindowSize        = 2048
	Threshold         = 0.2
	SourceURL         = "http://monserveur-icecast:8000/live.mp3"
	VASTURL           = "https://ads.myadserver.com/vast.xml"
)

func streamHandler(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get(SourceURL)
	if err != nil {
		http.Error(w, "Erreur source", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	decoder, err := mp3.NewDecoder(resp.Body)
	if err != nil {
		http.Error(w, "Erreur de décodage MP3", http.StatusInternalServerError)
		return
	}

	pcmBuf := make([]byte, 2*WindowSize)
	floatBuf := make([]float64, WindowSize)

	w.Header().Set("Content-Type", "audio/mpeg")

	for {
		_, err := io.ReadFull(decoder, pcmBuf)
		if err != nil {
			break
		}

		for i := 0; i < WindowSize; i++ {
			floatBuf[i] = float64(int16(binary.LittleEndian.Uint16(pcmBuf[i*2 : i*2+2])))
		}

		if containsBeep(floatBuf) {
			fmt.Println("🔊 Bip détecté ! Insertion de pub...")
			adURL, err := GetAdAudioURL(VASTURL)
			if err == nil {
				StreamAd(w, adURL)
			}
			fmt.Println("✅ Fin de pub. Reprise du flux.")
			continue
		}

		w.Write(pcmBuf)
		time.Sleep(time.Duration(len(pcmBuf)*8*1e6/(SampleRate*16)) * time.Microsecond) // approx 128kbps pacing
	}
}
