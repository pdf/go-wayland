package idle_inhibit

//go:generate go-wayland-scanner -pkg idle_inhibit -prefix zwp -suffix v1 -o idle_inhibit.go -i https://raw.githubusercontent.com/wayland-project/wayland-protocols/1.22/unstable/idle-inhibit/idle-inhibit-unstable-v1.xml
