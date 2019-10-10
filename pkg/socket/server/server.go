package server

import (
	"bytes"
	"encoding/binary"
	"encoding/gob"
	"go.uber.org/zap"
	"io"
	"net"
	"os"
	"time"
	. "xtc/sofa/log"
	"xtc/sofa/model"
	"xtc/sofa/pkg/store"
)

const (
	UNIX_SOCK_PIPE_PATH = "/tmp/sofa.sock" // socket file path
)

func Start() {
	// Remove socket file
	os.Remove(UNIX_SOCK_PIPE_PATH)
	// Get unix socket address based on file path
	uaddr, err := net.ResolveUnixAddr("unix", UNIX_SOCK_PIPE_PATH)
	if err != nil {
		Logger.Error("resolve unix addr error", zap.Error(err))
		os.Exit(1)
	}

	// Listen on the address
	unixListener, err := net.ListenUnix("unix", uaddr)
	if err != nil {
		Logger.Error("listen on the address error", zap.Error(err))
		os.Exit(1)
	}

	// Close listener when close this function, you can also emit it because this function will not terminate gracefully
	defer unixListener.Close()

	Logger.Info("server start successed")

	// Monitor request and process
	for {
		uconn, err := unixListener.AcceptUnix()
		if err != nil {
			Logger.Error("accepts the request error", zap.Error(err))
			continue
		}

		// Handle request
		go handleConnection(uconn)
	}
}

/*******************************************************
* Handle connection and request
* conn: conn handler
*******************************************************/
func handleConnection(conn *net.UnixConn) {
	// Close connection when finish handling
	defer func() {
		conn.Close()
	}()

	// Read data and return response
	data, err := parseRequest(conn)

	if err != nil {
		Logger.Error("read the request data error", zap.Error(err))
		return
	}

	// 反序列化
	var call model.Call
	decoder := gob.NewDecoder(bytes.NewReader(data))
	decoder.Decode(&call)

	// 存储到 redis 交给 logstash 处理
	store.Save(&call)

	// Send back response
	sendResponse(conn, []byte(time.Now().String()))

}

/*******************************************************
* Parse request of unix socket
* conn: conn handler
*******************************************************/
func parseRequest(conn *net.UnixConn) ([]byte, error) {
	var reqLen uint32
	lenBytes := make([]byte, 4)
	if _, err := io.ReadFull(conn, lenBytes); err != nil {
		return nil, err
	}

	lenBuf := bytes.NewBuffer(lenBytes)
	if err := binary.Read(lenBuf, binary.BigEndian, &reqLen); err != nil {
		return nil, err
	}

	reqBytes := make([]byte, reqLen)
	_, err := io.ReadFull(conn, reqBytes)

	if err != nil {
		return nil, err
	}

	return reqBytes, nil
}

/*******************************************************
* Send response to client
* conn: conn handler
*******************************************************/
func sendResponse(conn *net.UnixConn, data []byte) {
	buf := new(bytes.Buffer)
	msglen := uint32(len(data))

	binary.Write(buf, binary.BigEndian, &msglen)
	data = append(buf.Bytes(), data...)

	conn.Write(data)

}
