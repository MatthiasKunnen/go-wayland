package client

import (
	"errors"
	"fmt"
	"net"
	"os"
	"sync"
)

type Context struct {
	conn          *net.UnixConn
	objects       map[uint32]Proxy
	currentID     uint32
	dispatchMutex sync.Mutex
}

func (ctx *Context) Register(p Proxy) {
	ctx.currentID++
	p.SetID(ctx.currentID)
	p.SetContext(ctx)
	ctx.objects[ctx.currentID] = p
}

func (ctx *Context) Unregister(p Proxy) {
	delete(ctx.objects, p.ID())
}

func (ctx *Context) GetProxy(id uint32) Proxy {
	return ctx.objects[id]
}

func (ctx *Context) Close() error {
	return ctx.conn.Close()
}

// Dispatch reads and processes incoming messages and calls [client.Dispatcher.Dispatch] on the
// respective wayland protocol.
// A Dispatch loop is usually used to handle incoming messages.
// Dispatch blocks if there are no incoming messages.
// Dispatch can be safe to call from another goroutine, if, before any wayland function is called,
// or any wayland object is used, a lock is established using [Context.Lock].
func (ctx *Context) Dispatch() error {
	senderID, opcode, fd, data, err := ctx.ReadMsg() // Blocks if there are no incoming messages
	if err != nil {
		return fmt.Errorf("ctx.Dispatch: unable to read msg: %w", err)
	}

	// Take the dispatch lock. If the lock is already taken, this waits until another goroutine has
	// made their changes and releases the lock. If the lock is free, we take it, dispatch, and
	// release.
	ctx.dispatchMutex.Lock()
	defer ctx.dispatchMutex.Unlock()
	sender, ok := ctx.objects[senderID]
	if ok {
		if sender, ok := sender.(Dispatcher); ok {
			sender.Dispatch(opcode, fd, data)
		} else {
			return fmt.Errorf("ctx.Dispatch: sender doesn't implement Dispatch method (senderID=%d)", senderID)
		}
	} else {
		return fmt.Errorf("ctx.Dispatch: unable find sender (senderID=%d)", senderID)
	}

	return nil
}

// Lock takes a lock that will prevent further dispatching until unlocked.
// This is necessary if Dispatch is called in another goroutine.
func (ctx *Context) Lock() {
	ctx.dispatchMutex.Lock()
}

// Unlock unlocks the dispatch lock.
// This is necessary after [Context.Lock] if [Context.Dispatch] is called in another goroutine.
func (ctx *Context) Unlock() {
	ctx.dispatchMutex.Unlock()
}

func Connect(addr string) (*Display, error) {
	if addr == "" {
		runtimeDir := os.Getenv("XDG_RUNTIME_DIR")
		if runtimeDir == "" {
			return nil, errors.New("env XDG_RUNTIME_DIR not set")
		}
		if addr == "" {
			addr = os.Getenv("WAYLAND_DISPLAY")
		}
		if addr == "" {
			addr = "wayland-0"
		}
		addr = runtimeDir + "/" + addr
	}

	ctx := &Context{
		objects: map[uint32]Proxy{},
	}

	conn, err := net.DialUnix("unix", nil, &net.UnixAddr{Name: addr, Net: "unix"})
	if err != nil {
		return nil, err
	}
	ctx.conn = conn

	return NewDisplay(ctx), nil
}
