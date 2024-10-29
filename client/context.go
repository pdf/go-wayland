package client

import (
	"fmt"
	"net"
	"os"
	"reflect"
)

type Context struct {
	conn      *net.UnixConn
	objects   map[uint32]Proxy
	currentID uint32
}

func Connect(addr string) (*Display, error) {
	if addr == "" {
		addr = getAddress()
	}

	conn, err := net.DialUnix("unix", nil, &net.UnixAddr{Name: addr, Net: "unix"})
	if err != nil {
		return nil, err
	}

	ctx := &Context{
		conn:    conn,
		objects: map[uint32]Proxy{},
	}
	return NewDisplay(ctx), nil
}

func (ctx *Context) Register(p Proxy) {
	ctx.currentID++
	ctx.register(ctx.currentID, p)
}

func (ctx *Context) register(id uint32, p Proxy) {
	p.SetID(id)
	p.SetContext(ctx)
	ctx.objects[id] = p
}

func (ctx *Context) Unregister(p Proxy) {
	delete(ctx.objects, p.ID())
}

func (ctx *Context) Get(id uint32) Proxy {
	return ctx.objects[id]
}

// GetOrRegister will register proxy of type p if it's not registered
// Example: ctx.GetOrRegister(id, (*Object)(nil)).(*Object)
func (ctx *Context) GetOrRegister(id uint32, p Proxy) Proxy {
	if p, ok := ctx.objects[id]; ok {
		return p
	}

	proxy := reflect.New(reflect.TypeOf(p).Elem()).Interface().(Proxy)
	ctx.register(id, proxy)
	return proxy
}

func (ctx *Context) Close() error {
	return ctx.conn.Close()
}

func (ctx *Context) Dispatch() error {
	senderID, opcode, fd, data, err := ctx.ReadMsg()
	if err != nil {
		return fmt.Errorf("unable to read msg: %w", err)
	}

	sender, ok := ctx.objects[senderID]
	if ok {
		if sender, ok := sender.(Dispatcher); ok {
			sender.Dispatch(opcode, fd, data)
		} else {
			return fmt.Errorf("sender doesn't implement Dispatch method (senderID=%d)", senderID)
		}
	} else {
		return fmt.Errorf("unable find sender (senderID=%d)", senderID)
	}

	return nil
}

func getAddress() string {
	dir := os.Getenv("XDG_RUNTIME_DIR")
	if dir == "" {
		dir = "/run/user/1000"
	}
	display := os.Getenv("WAYLAND_DISPLAY")
	if display == "" {
		display = "wayland-0"
	}
	return dir + "/" + display
}
