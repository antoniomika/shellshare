// Package share handles the shellshare logic
package share

import (
	"bufio"
	"io"
	"log"
	"net"
	"os"
	"os/exec"
	"os/signal"
	"strings"
	"sync"
	"syscall"
	"time"

	"github.com/creack/pty"
	"github.com/google/shlex"
	"github.com/spf13/viper"
	"golang.org/x/crypto/ssh/terminal"
)

var once sync.Once

// Start initializes the shell share command.
func Start() {
	splitCmd, err := shlex.Split(viper.GetString("command"))
	if err != nil {
		log.Fatalln("error splitting command:", err)
	}

	cmd := exec.Command(splitCmd[0], splitCmd[1:]...)

	conn, err := net.Dial("tcp", viper.GetString("address"))
	if err != nil {
		log.Fatalln("error connecting to socket:", err)
	}

	firstLine, err := bufio.NewReader(conn).ReadBytes('\n')
	if err != nil {
		log.Fatalln("error reading first line:", err)
	}

	log.Println(strings.TrimSpace(string(firstLine)))

	<-time.After(viper.GetDuration("delay"))

	size, err := pty.GetsizeFull(os.Stdout)
	if err != nil {
		log.Fatalln("error getting terminal size:", err)
	}

	ptmx, err := pty.StartWithSize(cmd, size)
	if err != nil {
		log.Fatalln("error starting with pty:", err)
	}

	defer ptmx.Close()

	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGWINCH)
	go func() {
		for range ch {
			if err := pty.InheritSize(os.Stdin, ptmx); err != nil {
				log.Fatalln("error resizing pty:", err)
			}
		}
	}()

	oldState, err := terminal.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		log.Fatalln("error making terminal raw:", err)
	}

	defer func() {
		err := terminal.Restore(int(os.Stdin.Fd()), oldState)
		if err != nil {
			log.Fatalln("error returning terminal:", err)
		}
	}()

	stop := make(chan bool)

	go copyStop(ptmx, os.Stdin, stop)

	if viper.GetBool("remote") {
		go copyStop(ptmx, conn, stop)
	}

	go copyStop(io.MultiWriter(os.Stdout, conn), ptmx, stop)

	<-stop
}

func copyStop(dst io.Writer, src io.Reader, stop chan bool) {
	_, err := io.Copy(dst, src)
	if err != nil {
		log.Println("error copying to destination:", err)
	}

	once.Do(func() {
		close(stop)
	})
}
