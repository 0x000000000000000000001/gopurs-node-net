package main

import (
	"fmt"
	"gopurs/output/Node.EventEmitter"
	"gopurs/output/gopurs_runtime"
	"net"
)

func NewServerImpl() interface{} { 
    return Node_EventEmitter.NewImpl(nil) 
}

func NewServerOptionsImpl(arg0 interface{}) interface{} { 
    return Node_EventEmitter.NewImpl(nil) 
}

func ListenImpl(arg0 interface{}, arg1 interface{}) interface{} {
	s := gopurs_runtime.Unbox[*Node_EventEmitter.EventEmitter](arg0)
	
	options := gopurs_runtime.RecordToMap(arg1.(gopurs_runtime.Value))
	port := gopurs_runtime.Unbox[int64](options["port"])
	
	hostVal, hasHost := options["host"]
	var host string
	if hasHost {
		host = gopurs_runtime.Unbox[string](hostVal)
	} else {
		host = "localhost"
	}
	
	address := fmt.Sprintf("%s:%d", host, port)

	go func() {
		listener, err := net.Listen("tcp", address)
		if err != nil {
			Node_EventEmitter.GopursUnsafeEmitFn2(gopurs_runtime.Box(s), "error", gopurs_runtime.Box(err.Error()), nil)
			return
		}
		
		s.Any = listener
		
		Node_EventEmitter.GopursUnsafeEmitFn1(gopurs_runtime.Box(s), "listening", nil)
		
		for {
			conn, err := listener.Accept()
			if err != nil {
				Node_EventEmitter.GopursUnsafeEmitFn1(gopurs_runtime.Box(s), "close", nil)
				break
			}
			
			// We need to create a new Socket object and emit "connection"
			socketObj := Node_EventEmitter.NewImpl(nil).(*Node_EventEmitter.EventEmitter)
			socketObj.Any = conn
			
			// Emit connection event with the socket
			Node_EventEmitter.GopursUnsafeEmitFn2(gopurs_runtime.Box(s), "connection", gopurs_runtime.Box(socketObj), nil)
			
			// start read loop for the accepted socket
			go func(c net.Conn, sock *Node_EventEmitter.EventEmitter) {
				if sock.Any == nil {
					return // HTTP Server stole the socket, don't read from it!
				}
				buf := make([]byte, 8192)
				for {
					n, err := c.Read(buf)
					if n > 0 {
						data := make([]byte, n)
						copy(data, buf[:n])
						Node_EventEmitter.GopursUnsafeEmitFn2(gopurs_runtime.Box(sock), "data", gopurs_runtime.Box(data), nil)
					}
					if err != nil {
						Node_EventEmitter.GopursUnsafeEmitFn1(gopurs_runtime.Box(sock), "end", nil)
						Node_EventEmitter.GopursUnsafeEmitFn2(gopurs_runtime.Box(sock), "close", gopurs_runtime.Box(false), nil)
						break
					}
				}
			}(conn, socketObj)
		}
	}()
	return nil
}

func CloseImpl(arg0 interface{}) interface{} {
    s := gopurs_runtime.Unbox[*Node_EventEmitter.EventEmitter](arg0)
    if listener, ok := s.Any.(net.Listener); ok {
        listener.Close()
    }
    return nil
}

func AddressTcpImpl(arg0 interface{}) interface{} {
    s := gopurs_runtime.Unbox[*Node_EventEmitter.EventEmitter](arg0)
    if listener, ok := s.Any.(net.Listener); ok {
        if addr, ok2 := listener.Addr().(*net.TCPAddr); ok2 {
            rec := make(map[string]gopurs_runtime.Value)
            rec["port"] = gopurs_runtime.Box(int64(addr.Port))
            rec["family"] = gopurs_runtime.Box("IPv4")
            rec["address"] = gopurs_runtime.Box(addr.IP.String())
            return rec
        }
    }
    return nil 
}

func AddressIpcImpl(arg0 interface{}) interface{} { return nil }
func GetConnectionsImpl(arg0 interface{}, arg1 interface{}) interface{} { return nil }
func ListeningImpl(arg0 interface{}) interface{} { return true }
func MaxConnectionsImpl(arg0 interface{}) interface{} { return 0 }
func RefImpl(arg0 interface{}) interface{} { return nil }
func UnrefImpl(arg0 interface{}) interface{} { return nil }
