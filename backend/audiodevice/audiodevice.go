package audiodevice

type AudioDevice struct {
	ID       string
	Name     string
	Selected bool
}

type AudioDevices []AudioDevice

func (a AudioDevices) Len() int           { return len(a) }
func (a AudioDevices) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a AudioDevices) Less(i, j int) bool { return a[i].Name < a[j].Name }
