export WAYLAND_TAG ?= 1.23
export WAYLAND_PROTOCOLS_TAG ?= 1.38
export MESA_TAG ?= 24.2

gen:
	@go generate -x ./...
