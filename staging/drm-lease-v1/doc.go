package drm_lease

//go:generate go run github.com/pdf/go-wayland/cmd/go-wayland-scanner -pkg drm_lease -prefix wp -suffix v1 -o drm_lease.go -i https://gitlab.freedesktop.org/wayland/wayland-protocols/-/raw/$WAYLAND_PROTOCOLS_TAG/staging/drm-lease/drm-lease-v1.xml
