package unixpath

import (
	"log"
	"strings"
)

// DEPRECATED, found out about path/filepath lib
type Unixpath struct {
	Path string
}

func (up Unixpath) File() string {

	if up.Path == "" {
		log.Println("Incorrect usage of Unixpath.File, input is empty")
		return string(up.Path)
	}
	if condition := strings.LastIndex(string(up.Path), "/"); condition == -1 {
		log.Println("Unnecessarily usage of Unixpath.File, input doesn't contain '/'")
		return string(up.Path)
	}
	return string(up.Path[strings.LastIndex(string(up.Path), "/")+1:])
}

const (
	WITH_SLASH = iota
	WITHOUT_SLASH
)

func (up Unixpath) Directory(slash int) string {
	if up.Path == "" {
		log.Println("Incorrect usage of Unixpath.Directory, input is empty")
		return string(up.Path)
	}
	if condition := strings.LastIndex(string(up.Path), "/"); condition == -1 {
		log.Println("Incorrect usage of Unixpath.Directory , input doesn't contain '/', and consequently doesn't contain directory")
		return string(up.Path)
	}
	if slash == WITHOUT_SLASH {
		return string(up.Path[:strings.LastIndex(string(up.Path), "/")])
	}
	return string(up.Path[:strings.LastIndex(string(up.Path), "/")+1])

}
func File(input string) string {
	if input == "" {
		log.Println("Incorrect usage of Unixpath.File, input is empty")
		return string(input)
	}
	if condition := strings.LastIndex(string(input), "/"); condition == -1 {
		log.Println("Unnecessarily usage of Unixpath.File, input doesn't contain '/'")
		return string(input)
	}
	return string(input[strings.LastIndex(string(input), "/")+1:])
}

func Directory(input string, slash int) string {
	if input == "" {
		log.Println("Incorrect usage of Unixpath.Directory, input is empty")
		return string(input)
	}
	if condition := strings.LastIndex(string(input), "/"); condition == -1 {
		log.Println("Incorrect usage of Unixpath.Directory , input doesn't contain '/', and consequently doesn't contain directory")
		return string(input)
	}
	if slash == WITHOUT_SLASH {
		return string(input[:strings.LastIndex(string(input), "/")])
	}
	return string(input[:strings.LastIndex(string(input), "/")+1])

}
