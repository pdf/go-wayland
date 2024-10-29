package ext_session_lock

//go:generate go run github.com/raitonoberu/go-wayland/cmd/go-wayland-scanner -pkg ext_session_lock -suffix v1 -o ext_session_lock.go -i https://gitlab.freedesktop.org/wayland/wayland-protocols/-/raw/$WAYLAND_PROTOCOLS_TAG/staging/ext-session-lock/ext-session-lock-v1.xml
