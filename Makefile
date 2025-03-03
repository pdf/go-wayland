export WAYLAND_TAG ?= 1.23.1
export WAYLAND_PROTOCOLS_TAG ?= 1.41
export MESA_TAG ?= 25.0

gen:
	@go generate -x ./...
