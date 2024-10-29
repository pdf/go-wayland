module github.com/raitonoberu/go-wayland/examples/imageviewer

go 1.23.1

require (
	github.com/nfnt/resize v0.0.0-20180221191011-83c6a9932646
	github.com/raitonoberu/go-wayland v0.0.0-20230130181619-0ad78d1310b2
	golang.org/x/image v0.21.0
	golang.org/x/sys v0.26.0
)

replace github.com/raitonoberu/go-wayland => ../../
