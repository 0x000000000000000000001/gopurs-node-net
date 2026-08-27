package main

import (
	"fmt"
	"gopurs/output/gopurs_runtime"
	"net"
)

func NewImpl(arg0 interface{}) interface{} {
	e := Node_EventEmitter_NewImpl(nil).(*EventEmitter)
	return e
}

func ConnectTcpImpl(arg0 interface{}, arg1 interface{}) interface{} {
	s := gopurs_runtime.Unbox[*EventEmitter](arg0)
	
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
		conn, err := net.Dial("tcp", address)
		if err != nil {
			Node_EventEmitter_GopursUnsafeEmitFn2(gopurs_runtime.Box(s), "error", gopurs_runtime.Box(err.Error()), nil)
			return
		}
		
		s.Any = conn // Set io.Writer / io.Closer for Node.Stream
		
		Node_EventEmitter_GopursUnsafeEmitFn1(gopurs_runtime.Box(s), "connect", nil)
		Node_EventEmitter_GopursUnsafeEmitFn1(gopurs_runtime.Box(s), "ready", nil)
		
		buf := make([]byte, 8192)
		for {
			n, err := conn.Read(buf)
			if n > 0 {
				data := make([]byte, n)
				copy(data, buf[:n])
				Node_EventEmitter_GopursUnsafeEmitFn2(gopurs_runtime.Box(s), "data", gopurs_runtime.Box(data), nil)
			}
			if err != nil {
				Node_EventEmitter_GopursUnsafeEmitFn1(gopurs_runtime.Box(s), "end", nil)
				Node_EventEmitter_GopursUnsafeEmitFn2(gopurs_runtime.Box(s), "close", gopurs_runtime.Box(false), nil)
				break
			}
		}
	}()
	return arg0
}

func CreateConnectionImpl(arg0 interface{}) interface{} {
	s := Node_Net_Socket_NewImpl(nil)
	return Node_Net_Socket_ConnectTcpImpl(s, arg0)
}

func AddressImpl(arg0 interface{}) interface{} { return Node_EventEmitter_NewImpl(nil) }
func BytesReadImpl(arg0 interface{}) interface{} { return Node_EventEmitter_NewImpl(nil) }
func BytesWrittenImpl(arg0 interface{}) interface{} { return Node_EventEmitter_NewImpl(nil) }
func ConnectIpcImpl(arg0 interface{}, arg1 interface{}) interface{} { return arg0 }
func ConnectingImpl(arg0 interface{}) interface{} { return false }
func DestroySoonImpl(arg0 interface{}) interface{} { 
    if s, ok := gopurs_runtime.Unbox[*EventEmitter](arg0).Any.(net.Conn); ok {
        s.Close()
    }
    return nil 
}
func LocalAddressImpl(arg0 interface{}) interface{} { return "" }
func LocalPortImpl(arg0 interface{}) interface{} { return 0 }
func LocalFamilyImpl(arg0 interface{}) interface{} { return "" }
func PendingImpl(arg0 interface{}) interface{} { return false }
func RefImpl(arg0 interface{}) interface{} { return nil }
func RemoteAddressImpl(arg0 interface{}) interface{} { return "" }
func RemotePortImpl(arg0 interface{}) interface{} { return 0 }
func RemoteFamilyImpl(arg0 interface{}) interface{} { return "" }
func ResetAndDestroyImpl(arg0 interface{}) interface{} { return nil }
func SetKeepAliveImpl(arg0 interface{}) interface{} { return nil }
func SetKeepAliveBooleanImpl(arg0 interface{}, arg1 interface{}) interface{} { return nil }
func SetKeepAliveInitialDelayImpl(arg0 interface{}, arg1 interface{}) interface{} { return nil }
func SetKeepAliveAllImpl(arg0 interface{}, arg1 interface{}, arg2 interface{}) interface{} { return nil }
func SetNoDelayImpl(arg0 interface{}) interface{} { return nil }
func SetNoDelayBooleanImpl(arg0 interface{}, arg1 interface{}) interface{} { return nil }
func SetTimeoutImpl(arg0 interface{}, arg1 interface{}) interface{} { return nil }
func TimeoutImpl(arg0 interface{}) interface{} { return nil }
func UnrefImpl(arg0 interface{}) interface{} { return nil }
func ReadyStateImpl(arg0 interface{}) interface{} { return "open" }
