// Generated by go-wayland-scanner
// https://github.com/rajveermalviya/go-wayland/cmd/go-wayland-scanner
// XML file : https://gitlab.freedesktop.org/wayland/wayland-protocols/-/raw/1.31/staging/xwayland-shell/xwayland-shell-v1.xml?ref_type=tags
//
// xwayland_shell_v1 Protocol Copyright:
//
// Copyright © 2022 Joshua Ashton
//
// Permission is hereby granted, free of charge, to any person obtaining a
// copy of this software and associated documentation files (the "Software"),
// to deal in the Software without restriction, including without limitation
// the rights to use, copy, modify, merge, publish, distribute, sublicense,
// and/or sell copies of the Software, and to permit persons to whom the
// Software is furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice (including the next
// paragraph) shall be included in all copies or substantial portions of the
// Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.  IN NO EVENT SHALL
// THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
// FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER
// DEALINGS IN THE SOFTWARE.

package xwayland_shell

import "github.com/rajveermalviya/go-wayland/wayland/client"

// XwaylandShell : context object for Xwayland shell
//
// xwayland_shell_v1 is a singleton global object that
// provides the ability to create a xwayland_surface_v1 object
// for a given wl_surface.
//
// This interface is intended to be bound by the Xwayland server.
//
// A compositor must not allow clients other than Xwayland to
// bind to this interface. A compositor should hide this global
// from other clients' wl_registry.
// A client the compositor does not consider to be an Xwayland
// server attempting to bind this interface will result in
// an implementation-defined error.
//
// An Xwayland server that has bound this interface must not
// set the `WL_SURFACE_ID` atom on a window.
type XwaylandShell struct {
	client.BaseProxy
}

// NewXwaylandShell : context object for Xwayland shell
//
// xwayland_shell_v1 is a singleton global object that
// provides the ability to create a xwayland_surface_v1 object
// for a given wl_surface.
//
// This interface is intended to be bound by the Xwayland server.
//
// A compositor must not allow clients other than Xwayland to
// bind to this interface. A compositor should hide this global
// from other clients' wl_registry.
// A client the compositor does not consider to be an Xwayland
// server attempting to bind this interface will result in
// an implementation-defined error.
//
// An Xwayland server that has bound this interface must not
// set the `WL_SURFACE_ID` atom on a window.
func NewXwaylandShell(ctx *client.Context) *XwaylandShell {
	xwaylandShellV1 := &XwaylandShell{}
	ctx.Register(xwaylandShellV1)
	return xwaylandShellV1
}

// Destroy : destroy the Xwayland shell object
//
// Destroy the xwayland_shell_v1 object.
//
// The child objects created via this interface are unaffected.
func (i *XwaylandShell) Destroy() error {
	defer i.Context().Unregister(i)
	const opcode = 0
	const _reqBufLen = 8
	var _reqBuf [_reqBufLen]byte
	l := 0
	client.PutUint32(_reqBuf[l:4], i.ID())
	l += 4
	client.PutUint32(_reqBuf[l:l+4], uint32(_reqBufLen<<16|opcode&0x0000ffff))
	l += 4
	err := i.Context().WriteMsg(_reqBuf[:], nil)
	return err
}

// GetXwaylandSurface : assign the xwayland_surface surface role
//
// Create an xwayland_surface_v1 interface for a given wl_surface
// object and gives it the xwayland_surface role.
//
// It is illegal to create an xwayland_surface_v1 for a wl_surface
// which already has an assigned role and this will result in the
// `role` protocol error.
//
// See the documentation of xwayland_surface_v1 for more details
// about what an xwayland_surface_v1 is and how it is used.
func (i *XwaylandShell) GetXwaylandSurface(surface *client.Surface) (*XwaylandSurface, error) {
	id := NewXwaylandSurface(i.Context())
	const opcode = 1
	const _reqBufLen = 8 + 4 + 4
	var _reqBuf [_reqBufLen]byte
	l := 0
	client.PutUint32(_reqBuf[l:4], i.ID())
	l += 4
	client.PutUint32(_reqBuf[l:l+4], uint32(_reqBufLen<<16|opcode&0x0000ffff))
	l += 4
	client.PutUint32(_reqBuf[l:l+4], id.ID())
	l += 4
	client.PutUint32(_reqBuf[l:l+4], surface.ID())
	l += 4
	err := i.Context().WriteMsg(_reqBuf[:], nil)
	return id, err
}

type XwaylandShellError uint32

// XwaylandShellError :
const (
	// XwaylandShellErrorRole : given wl_surface has another role
	XwaylandShellErrorRole XwaylandShellError = 0
)

func (e XwaylandShellError) Name() string {
	switch e {
	case XwaylandShellErrorRole:
		return "role"
	default:
		return ""
	}
}

func (e XwaylandShellError) Value() string {
	switch e {
	case XwaylandShellErrorRole:
		return "0"
	default:
		return ""
	}
}

func (e XwaylandShellError) String() string {
	return e.Name() + "=" + e.Value()
}

// XwaylandSurface : interface for associating Xwayland windows to wl_surfaces
//
// An Xwayland surface is a surface managed by an Xwayland server.
// It is used for associating surfaces to Xwayland windows.
//
// The Xwayland server associated with actions in this interface is
// determined by the Wayland client making the request.
//
// The client must call wl_surface.commit on the corresponding wl_surface
// for the xwayland_surface_v1 state to take effect.
type XwaylandSurface struct {
	client.BaseProxy
}

// NewXwaylandSurface : interface for associating Xwayland windows to wl_surfaces
//
// An Xwayland surface is a surface managed by an Xwayland server.
// It is used for associating surfaces to Xwayland windows.
//
// The Xwayland server associated with actions in this interface is
// determined by the Wayland client making the request.
//
// The client must call wl_surface.commit on the corresponding wl_surface
// for the xwayland_surface_v1 state to take effect.
func NewXwaylandSurface(ctx *client.Context) *XwaylandSurface {
	xwaylandSurfaceV1 := &XwaylandSurface{}
	ctx.Register(xwaylandSurfaceV1)
	return xwaylandSurfaceV1
}

// SetSerial : associates a Xwayland window to a wl_surface
//
// Associates an Xwayland window to a wl_surface.
// The association state is double-buffered and will be applied at
// the time wl_surface.commit of the corresponding wl_surface is called.
//
// The `serial_lo` and `serial_hi` parameters specify a non-zero
// monotonic serial number which is entirely unique and provided by the
// Xwayland server equal to the serial value provided by a client message
// with a message type of the `WL_SURFACE_SERIAL` atom on the X11 window
// for this surface to be associated to.
//
// The serial value in the `WL_SURFACE_SERIAL` client message is specified
// as having the lo-bits specified in `l[0]` and the hi-bits specified
// in `l[1]`.
//
// If the serial value provided by `serial_lo` and `serial_hi` is not
// valid, the `invalid_serial` protocol error will be raised.
//
// An X11 window may be associated with multiple surfaces throughout its
// lifespan. (eg. unmapping and remapping a window).
//
// For each wl_surface, this state must not be committed more than once,
// otherwise the `already_associated` protocol error will be raised.
//
//	serialLo: The lower 32-bits of the serial number associated with the X11 window
//	serialHi: The upper 32-bits of the serial number associated with the X11 window
func (i *XwaylandSurface) SetSerial(serialLo, serialHi uint32) error {
	const opcode = 0
	const _reqBufLen = 8 + 4 + 4
	var _reqBuf [_reqBufLen]byte
	l := 0
	client.PutUint32(_reqBuf[l:4], i.ID())
	l += 4
	client.PutUint32(_reqBuf[l:l+4], uint32(_reqBufLen<<16|opcode&0x0000ffff))
	l += 4
	client.PutUint32(_reqBuf[l:l+4], uint32(serialLo))
	l += 4
	client.PutUint32(_reqBuf[l:l+4], uint32(serialHi))
	l += 4
	err := i.Context().WriteMsg(_reqBuf[:], nil)
	return err
}

// Destroy : destroy the Xwayland surface object
//
// Destroy the xwayland_surface_v1 object.
//
// Any already existing associations are unaffected by this action.
func (i *XwaylandSurface) Destroy() error {
	defer i.Context().Unregister(i)
	const opcode = 1
	const _reqBufLen = 8
	var _reqBuf [_reqBufLen]byte
	l := 0
	client.PutUint32(_reqBuf[l:4], i.ID())
	l += 4
	client.PutUint32(_reqBuf[l:l+4], uint32(_reqBufLen<<16|opcode&0x0000ffff))
	l += 4
	err := i.Context().WriteMsg(_reqBuf[:], nil)
	return err
}

type XwaylandSurfaceError uint32

// XwaylandSurfaceError :
const (
	// XwaylandSurfaceErrorAlreadyAssociated : given wl_surface is already associated with an X11 window
	XwaylandSurfaceErrorAlreadyAssociated XwaylandSurfaceError = 0
	// XwaylandSurfaceErrorInvalidSerial : serial was not valid
	XwaylandSurfaceErrorInvalidSerial XwaylandSurfaceError = 1
)

func (e XwaylandSurfaceError) Name() string {
	switch e {
	case XwaylandSurfaceErrorAlreadyAssociated:
		return "already_associated"
	case XwaylandSurfaceErrorInvalidSerial:
		return "invalid_serial"
	default:
		return ""
	}
}

func (e XwaylandSurfaceError) Value() string {
	switch e {
	case XwaylandSurfaceErrorAlreadyAssociated:
		return "0"
	case XwaylandSurfaceErrorInvalidSerial:
		return "1"
	default:
		return ""
	}
}

func (e XwaylandSurfaceError) String() string {
	return e.Name() + "=" + e.Value()
}
