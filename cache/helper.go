package cache

import (
	"fmt"
	"os/exec"
	"time"
)

var (
	cmd *exec.Cmd
)

func SetupMemcached() error {
	if cmd != nil {
		KillMemcached()
	}

	sock := fmt.Sprintf("/tmp/test-memcached-1.sock")

	cmd = exec.Command("memcached", "-s", sock)
	if err := cmd.Start(); err != nil {
		return err
	}

	// Wait a bit for the socket to appear.
	for i := 0; i < 10; i++ {
		time.Sleep(time.Duration(25*i) * time.Millisecond)
	}

	Init(sock)
	return nil
}

func KillMemcached() {
	cmd.Process.Kill()
	cmd = nil
}
