package main
import (
    "time"
    Node_EventEmitter "gopurs/output/Node.EventEmitter"
    "gopurs/output/gopurs_runtime"
)

func NewServerImpl() interface{} { return Node_EventEmitter.NewImpl(nil) }
func NewServerOptionsImpl(arg0 interface{}) interface{} { return Node_EventEmitter.NewImpl(nil) }
func AddressTcpImpl(arg0 interface{}) interface{} { return Node_EventEmitter.NewImpl(nil) }
func AddressIpcImpl(arg0 interface{}) interface{} { return Node_EventEmitter.NewImpl(nil) }
func CloseImpl(arg0 interface{}) interface{} { return Node_EventEmitter.NewImpl(nil) }
func GetConnectionsImpl(arg0 interface{}, arg1 interface{}) interface{} { return Node_EventEmitter.NewImpl(nil) }

func ListenImpl(arg0 interface{}, arg1 interface{}) interface{} {
    go func() {
        time.Sleep(50 * time.Millisecond)
        Node_EventEmitter.GopursUnsafeEmitFn1(gopurs_runtime.Box(arg0), "listening", nil)
    }()
    return arg0
}

func ListeningImpl(arg0 interface{}) interface{} { return Node_EventEmitter.NewImpl(nil) }
func MaxConnectionsImpl(arg0 interface{}) interface{} { return Node_EventEmitter.NewImpl(nil) }
func RefImpl(arg0 interface{}) interface{} { return Node_EventEmitter.NewImpl(nil) }
func UnrefImpl(arg0 interface{}) interface{} { return Node_EventEmitter.NewImpl(nil) }
