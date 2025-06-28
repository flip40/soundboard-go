# Sound Board

Simple Sound Board written in Go + Angular using Wails framework.

## Prerequisites

* Wails CLI v2.10.1 (https://wails.io/docs/gettingstarted/installation)
* Go 1.22.0 (https://go.dev/doc/install)
* Node 22.16.0 and `npm` (https://nodejs.org/en/download)
* Angular CLI 19.2.13 (install with `npm i -g @angular/cli@19.2.13`)

# Supported Sound Formats
* WAV
* MP3

# How to use

Plays sound to any playback device on your system. This can be your microphone directly or to another device physical or virtual playback device.

I have three suggested methods of how to use the soundboard, with varying levels of complexity, along with different pros and cons.

## 1. Play Audio through Microphone Device
* Required tools:
  * Just the soundboard and a microphone!
* Pros: Very easy setup; people can hear you and the soundboard.
* Cons: Audio is combined, so you can't split who hears the soundboard and who hears the mic; background noise suppression and other common "enhancements" to voice chat **WILL** garble your soundboard noises. You'll have to turn them off, depending on where you play this.

## 2. Play Audio through VB Audio Cable
* Required tools:
  * [VB-CABLE Virtual Audio Device](https://vb-audio.com/Cable/)
* Pros: You can play soundboard sounds on one device, while keeping your microphone purely for your voice.
* Cons: Your microphone audio won't be on the same channel as the soundboard.

## 3. Play Audio with VB Voicemeeter
* Required Tools:
  * [VOICEMEETER Virtual Audio Mixer](https://vb-audio.com/Voicemeeter/)
* Pros: Voicemeeter is a Virtual Audio Mixer that lets you direct and mix input and output audio however you'd like. If you want to have both soundboard and microphone audio output on the same device for your in-game, while only having microphone audio in your external voice app (e.g. Discord), you can do it with Voicemeeter. This is great for more advanced use cases (specific control of audio output, streaming setups, etc).
* Cons: Voicemeeter essentially replaces any built-in audio mixing for Windows. I've had some audio quirks in the past, so this is more advanced to deal with if you don't already want to use a Virtual Audio Mixer like Voicemeeter and don't know how to use it.