package fractional_scale

//go:generate go run github.com/raitonoberu/go-wayland/cmd/go-wayland-scanner -pkg fractional_scale -prefix wp -suffix v1 -o fractional_scale.go -i https://gitlab.freedesktop.org/wayland/wayland-protocols/-/raw/$WAYLAND_PROTOCOLS_TAG/staging/fractional-scale/fractional-scale-v1.xml
