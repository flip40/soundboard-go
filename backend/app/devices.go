package app

import (
	"sort"

	"github.com/flip40/soundboard-go/backend/audiodevice"
	"github.com/gen2brain/malgo"
)

func (b *App) GetPlaybackDeviceInfo() []audiodevice.AudioDevice {
	audioDevices, err := b.audioCtx.Devices(malgo.Playback)
	if err != nil {
		panic(err)
	}

	var wrappedDevices audiodevice.AudioDevices
	for _, device := range audioDevices {
		wrappedDevices = append(wrappedDevices, audiodevice.AudioDevice{
			ID:       device.ID.String(),
			Name:     device.Name(),
			Selected: func() bool { return device.ID.String() == b.selectedDeviceID }(),
		})
	}

	sort.Sort(wrappedDevices)

	return wrappedDevices
}

func (b *App) SetPlaybackDevice(id string) {
	b.selectedDeviceID = id
}
