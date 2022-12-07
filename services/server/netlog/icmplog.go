package netlog

import (
	"net"

	"github.com/virzz/virzz/logger"
)

func ICMP(ip string) {
	const IcmpLen = 8
	var msg [512]byte
	msg[0] = 8 // type
	msg[1] = 0 // code
	msg[2] = 0 // checkSum -> 2 byte
	msg[3] = 0
	msg[4] = 0  // identifier[0]
	msg[5] = 13 // identifier[1]
	msg[6] = 0  // sequence[0]
	msg[7] = 37 // sequence[1]
	// 检验和
	check := checkSum(msg[:IcmpLen])
	msg[2] = byte(check >> 8)
	msg[3] = byte(check & 255)

	remoteAddr, err := net.ResolveIPAddr("ip", ip)
	if err != nil {
		logger.Error("error: ", err)
		return
	}
	// 拨号
	conn, err := net.DialIP("ip:icmp", nil, remoteAddr)
	if err != nil {
		logger.Error("error: ", err)
		return
	}
	// 发送数据
	if _, err := conn.Write(msg[:IcmpLen]); err != nil {
		logger.Error("send data error: ", err)
		return
	}
	// 读取返回的数据
	size, err := conn.Read(msg[:])
	if err != nil {
		logger.Error("error: ", err)
		return
	}
	logger.Error(msg[20:size])
}

func checkSum(msg []byte) uint16 {
	sum := 0
	for n := 0; n < len(msg); n += 2 {
		sum += int(msg[n])*256 + int(msg[n+1])
	}
	sum = (sum >> 16) + sum&0xffff
	sum += sum >> 16
	return uint16(^sum)
}
