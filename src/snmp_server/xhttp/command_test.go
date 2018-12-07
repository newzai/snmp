package xhttp

import "testing"

func TestCommand(t *testing.T) {

	command, err := getCommand(1)
	if err != nil {
		t.Error(err)
	} else if command != "ps" {
		t.Error("command not ps")
	}

	command3, err := getCommand(3)
	if err != nil {
		t.Error(err)
	} else if command3 != "ping -c 4 127.0.0.1" {
		t.Error("command 3 not ping -c 4 127.0.0.1")
	}
}
