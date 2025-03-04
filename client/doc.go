// Package client is Go port of wayland-client library
// for writing pure Go GUI software for wayland supported
// platforms.
package client

//go:generate go run github.com/pdf/go-wayland/cmd/go-wayland-scanner -pkg client -prefix wl -o client.go -i https://gitlab.freedesktop.org/wayland/wayland/-/raw/$WAYLAND_TAG/protocol/wayland.xml
