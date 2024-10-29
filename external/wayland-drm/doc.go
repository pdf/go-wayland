package wayland_drm

//go:generate go run github.com/raitonoberu/go-wayland/cmd/go-wayland-scanner -pkg wayland_drm -prefix wl -o wayland_drm.go -i https://gitlab.freedesktop.org/mesa/mesa/-/raw/$MESA_TAG/src/egl/wayland/wayland-drm/wayland-drm.xml
