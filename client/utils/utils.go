package utils

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

// --------------------------

func ReadLine(message string) string {
	var name string
	fmt.Print("\r" + message)
	fmt.Scanln(&name)
	return name
}

func GetUrl(host string, port int, params []string, secure bool) string {
	const (
		webSocketProtocol       = "ws://"
		webSocketSecureProtocol = "wss://"
		socketioUrl             = "/socket.io/?EIO=3&transport=websocket"
	)

	var prefix string
	if secure {
		prefix = webSocketSecureProtocol
	} else {
		prefix = webSocketProtocol
	}

	_url, err := url.Parse(prefix + host + ":" + strconv.Itoa(port) + socketioUrl)
	if err != nil {
		fmt.Println("We unable to parse given url: ", _url)
	}

	if len(params) > 0 {
		_uval := _url.Query()
		for _, element := range params {
			s := strings.Split(element, "=")
			_uval.Add(s[0], s[1])
		}
		_url.RawQuery = _uval.Encode()
	}

	return _url.String()
}
