package linux_dmabuf

//go:generate go run github.com/pdf/go-wayland/cmd/go-wayland-scanner -pkg linux_dmabuf -prefix zwp --suffix v1 -o linux_dmabuf.go -i https://gitlab.freedesktop.org/wayland/wayland-protocols/-/raw/$WAYLAND_PROTOCOLS_TAG/stable/linux-dmabuf/linux-dmabuf-v1.xml
