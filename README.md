# 🎧 Audio SSAI Middleware

Un middleware en Go pour **l'insertion dynamique de publicités audio (SSAI)** dans un flux MP3 (type Icecast), déclenchée automatiquement par détection de bip sonore.

## ⚙️ Fonctionnalités

- 🔊 Détection d’un bip à 1kHz dans le flux audio
- 🎯 Déclenchement automatique d'une requête VAST
- 📥 Récupération et injection d'une publicité audio en ligne
- 🔁 Reprise fluide du flux source après l'insertion

## 📦 Technologies utilisées

- Go 1.22+
- [go-mp3](https://github.com/hajimehoshi/go-mp3) — décodage MP3
- [go-dsp/fft](https://github.com/mjibson/go-dsp) — transformée de Fourier
- HTTP/VAST XML pour récupération de publicités

## 🚀 Lancer le projet

### 1. Cloner le repo

```bash
git clone https://github.com/ton-utilisateur/audio-ssai.git
cd audio-ssai
```

### 2. Installer les dépendances
```bash
go mod tidy
```

### 3. Lancer le serveur
```bash
go run main.go
```

➡️ Accéder ensuite à http://localhost:8080/stream pour consommer le flux avec SSAI.

🧪 Tester la détection de bip
Assurez-vous que le flux source (stream.go) contient un bip clair à 1000 Hz.
Il est possible de générer un fichier de test avec Audacity, par exemple.

## 📄 Architecture
```
.
├── main.go           # Serveur HTTP
├── stream.go         # Traitement du flux et détection de bip
├── fft.go            # Analyse FFT pour repérer le bip
├── vast.go           # Parsing VAST pour trouver l'audio à injecter
├── ad_injector.go    # Lecture et diffusion du flux publicitaire
└── go.mod
```

## 🔧 Personnalisation
Modifier FrequencyToDetect, Threshold ou WindowSize dans stream.go pour affiner la détection.

Adapter le parsing VAST dans vast.go selon les spécificités de l'ad server.

## 📜 Licence
MIT

✨ Projet développé pour intégrer facilement la publicité audio côté serveur, sans impliquer le client final.