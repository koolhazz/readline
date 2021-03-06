package readline

import (
	"container/list"
	"fmt"
	"os"
	"syscall"
	"time"
	"unicode/utf8"
	"unsafe"

	"golang.org/x/crypto/ssh/terminal"
)

// IsTerminal returns true if the given file descriptor is a terminal.
func IsTerminal(fd int) bool {
	return terminal.IsTerminal(fd)
}

func MakeRaw(fd int) (*terminal.State, error) {
	return terminal.MakeRaw(fd)
}

func Restore(fd int, state *terminal.State) error {
	return terminal.Restore(fd, state)
}

func IsPrintable(key rune) bool {
	isInSurrogateArea := key >= 0xd800 && key <= 0xdbff
	return key >= 32 && !isInSurrogateArea
}

func escapeExKey(r rune) rune {
	switch r {
	case 'D':
		r = CharBackward
	case 'C':
		r = CharForward
	case 'A':
		r = CharPrev
	case 'B':
		r = CharNext
	}
	return r
}

func escapeKey(r rune) rune {
	switch r {
	case 'b':
		r = MetaPrev
	case 'f':
		r = MetaNext
	case 'd':
		r = MetaDelete
	case CharTranspose:
		r = MetaTranspose
	case CharBackspace:
		r = MetaBackspace
	case CharEsc:

	}
	return r
}

func Debug(o ...interface{}) {
	f, _ := os.OpenFile("debug.tmp", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	fmt.Fprintln(f, o...)
	f.Close()
}

type winsize struct {
	Row    uint16
	Col    uint16
	Xpixel uint16
	Ypixel uint16
}

func getWidth() int {
	ws := &winsize{}
	retCode, _, errno := syscall.Syscall(syscall.SYS_IOCTL,
		uintptr(syscall.Stdin),
		uintptr(syscall.TIOCGWINSZ),
		uintptr(unsafe.Pointer(ws)))

	if int(retCode) == -1 {
		panic(errno)
	}
	return int(ws.Col)
}

func debugList(l *list.List) {
	idx := 0
	for e := l.Front(); e != nil; e = e.Next() {
		Debug(idx, fmt.Sprintf("%+v", e.Value))
		idx++
	}
}

func equalRunes(a, b []rune) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func sleep(n int) {
	Debug(n)
	time.Sleep(2000 * time.Millisecond)
}

func LineCount(w int) int {
	screenWidth := getWidth()
	r := w / screenWidth
	if w%screenWidth != 0 {
		r++
	}
	return r
}

func RunesWidth(r []rune) (length int) {
	for i := 0; i < len(r); i++ {
		if utf8.RuneLen(r[i]) > 3 {
			length += 2
		} else {
			length += 1
		}
	}
	return
}

func RunesIndexBck(r, sub []rune) int {
	for i := len(r) - len(sub); i >= 0; i-- {
		found := true
		for j := 0; j < len(sub); j++ {
			if r[i+j] != sub[j] {
				found = false
				break
			}
		}
		if found {
			return i
		}
	}
	return -1
}

func RunesIndex(r, sub []rune) int {
	for i := 0; i < len(r); i++ {
		found := true
		if len(r[i:]) < len(sub) {
			return -1
		}
		for j := 0; j < len(sub); j++ {
			if r[i+j] != sub[j] {
				found = false
				break
			}
		}
		if found {
			return i
		}
	}
	return -1
}

func IsWordBreak(i rune) bool {
	if i >= 'a' && i <= 'z' {
		return false
	}
	if i >= 'A' && i <= 'Z' {
		return false
	}
	return true
}
