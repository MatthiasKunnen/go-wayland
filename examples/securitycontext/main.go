package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"

	"github.com/MatthiasKunnen/go-wayland/wayland/client"
	security_context "github.com/MatthiasKunnen/go-wayland/wayland/staging/security-context-v1"
)

const socketName = "wayland-restricted-0"

func main() {
	display, err := client.Connect("")
	if err != nil {
		panic(err)
	}

	runtimeDir := os.Getenv("XDG_RUNTIME_DIR")
	if runtimeDir == "" {
		log.Fatal("env XDG_RUNTIME_DIR not set")
	}

	socketPath := filepath.Join(runtimeDir, socketName)

	display.SetErrorHandler(func(dee client.DisplayErrorEvent) {
		log.Fatalf("display error event: %v", dee)
	})

	registry, err := display.GetRegistry()
	if err != nil {
		log.Fatalf("error getting registry: %v", err)
	}

	scmCh := make(chan *security_context.WpSecurityContextManager, 1)

	registry.SetGlobalHandler(func(rge client.RegistryGlobalEvent) {
		switch rge.Interface {
		case security_context.WpSecurityContextManagerInterfaceName:
			scm := security_context.NewWpSecurityContextManager(display.Context())
			err := registry.Bind(rge.Name, rge.Interface, rge.Version, scm)
			if err != nil {
				log.Fatalf("unable to bind scm interface: %v", err)
			}
			scmCh <- scm
		}
	})

	// Wait for interface to register
	display.Roundtrip()
	// Wait for handler events
	display.Roundtrip()

	scm := <-scmCh
	close(scmCh)

	var closeSocket func() error

	defer func() {
		if closeSocket != nil {
			log.Println("closing socket")
			closeSocket()
		}
	}()

	// Ignore how ugly this is >.<

	toggleSocket := func() {
		var err error
		if closeSocket == nil {
			closeSocket, err = openSocket(scm, socketPath)
		} else {
			log.Println("closing socket, press enter to open it")
			err = closeSocket()
			closeSocket = nil
		}
		if err != nil {
			log.Fatal(err)
		}
	}

	toggleSocket()

	ctx, _ := signal.NotifyContext(context.Background(), os.Interrupt)

	go func() {
		<-ctx.Done()
		if closeSocket != nil {
			log.Println("closing socket")
			closeSocket()
		}
		os.Exit(1)
	}()

	sc := bufio.NewScanner(os.Stdin)
	for ctx.Err() == nil && sc.Scan() {
		toggleSocket()
	}
}

func openSocket(scm *security_context.WpSecurityContextManager, socketPath string) (closeSocket func() error, err error) {
	log.Println("opening socket, press enter to close it")

	pipeR, pipeW, err := os.Pipe()
	if err != nil {
		return nil, fmt.Errorf("error opening pipe: %w", err)
	}

	closeFD := int(pipeR.Fd())

	sockFD, err := syscall.Socket(syscall.AF_UNIX, syscall.SOCK_STREAM, 0)
	if err != nil {
		return nil, fmt.Errorf("error creating socket: %w", err)
	}

	_ = os.Remove(socketPath)

	err = syscall.Bind(sockFD, &syscall.SockaddrUnix{
		Name: socketPath,
	})
	if err != nil {
		return nil, fmt.Errorf("error calling bind on socket: %w", err)
	}

	if err := syscall.Listen(sockFD, 0); err != nil {
		return nil, fmt.Errorf("error calling listen on socket: %w", err)
	}

	sctx, err := scm.CreateListener(sockFD, closeFD)
	if err != nil {
		return nil, fmt.Errorf("error creating security context: %w", err)
	}

	if err := sctx.SetSandboxEngine("com.example.sandbox"); err != nil {
		return nil, fmt.Errorf("error setting sandbox engine on security context: %w", err)
	}

	if err := sctx.SetAppId("com.example.sandboxed-app"); err != nil {
		return nil, fmt.Errorf("error setting app ID on security context: %w", err)
	}

	if err := sctx.SetInstanceId("instance-1"); err != nil {
		return nil, fmt.Errorf("error setting instance ID on security context: %w", err)
	}

	if err := sctx.Commit(); err != nil {
		return nil, fmt.Errorf("error committing security context: %w", err)
	}

	if err := syscall.Close(sockFD); err != nil {
		return nil, fmt.Errorf("error closing socket from companion program: %w", err)
	}

	return func() error {
		_ = pipeR.Close()
		return pipeW.Close()
	}, nil
}
