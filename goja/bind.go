package goja

import (
	"astraljs"
	"github.com/dop251/goja"
)

func Bind(vm *goja.Runtime, astral *astraljs.AppHostFlatAdapter) (err error) {
	var a = adapter{astral: astral, vm: vm, queue: make(chan func(), 1024)}
	if err = vm.Set(astraljs.Log, a.Log); err != nil {
		return
	}
	if err = vm.Set(astraljs.Sleep, a.Sleep); err != nil {
		return
	}
	if err = vm.Set(astraljs.ServiceRegister, a.ServiceRegister); err != nil {
		return
	}
	if err = vm.Set(astraljs.ServiceClose, a.ServiceClose); err != nil {
		return
	}
	if err = vm.Set(astraljs.ConnAccept, a.ConnAccept); err != nil {
		return
	}
	if err = vm.Set(astraljs.ConnClose, a.ConnClose); err != nil {
		return
	}
	if err = vm.Set(astraljs.ConnWrite, a.ConnWrite); err != nil {
		return
	}
	if err = vm.Set(astraljs.ConnRead, a.ConnRead); err != nil {
		return
	}
	if err = vm.Set(astraljs.Query, a.Query); err != nil {
		return
	}
	if err = vm.Set(astraljs.QueryName, a.QueryName); err != nil {
		return
	}
	if err = vm.Set(astraljs.GetNodeInfo, a.NodeInfo); err != nil {
		return
	}
	if err = vm.Set(astraljs.Resolve, a.Resolve); err != nil {
		return
	}
	go func() {
		for f := range a.queue {
			f()
		}
	}()
	return
}

type adapter struct {
	astral *astraljs.AppHostFlatAdapter
	vm     *goja.Runtime
	queue  chan func()
}

func (a *adapter) Log(arg ...any) {
	a.astral.LogArr(arg)
}

func (a *adapter) Sleep(millis int64) *goja.Promise {
	promise, resolve, _ := a.vm.NewPromise()
	go func() {
		a.astral.Sleep(millis)
		a.queue <- func() {
			resolve(goja.Undefined())
		}
	}()
	return promise
}

func (a *adapter) ServiceRegister(port string) *goja.Promise {
	promise, resolve, reject := a.vm.NewPromise()
	go func() {
		err := a.astral.ServiceRegister(port)
		a.queue <- func() {
			if err != nil {
				reject(err)
			} else {
				resolve(goja.Undefined())
			}
		}
	}()
	return promise
}

func (a *adapter) ServiceClose(port string) *goja.Promise {
	promise, resolve, reject := a.vm.NewPromise()
	go func() {
		err := a.astral.ServiceClose(port)
		a.queue <- func() {
			if err != nil {
				reject(err)
			} else {
				resolve(goja.Undefined())
			}
		}
	}()
	return promise
}

func (a *adapter) ConnAccept(port string) *goja.Promise {
	promise, resolve, reject := a.vm.NewPromise()
	go func() {
		conn, err := a.astral.ConnAccept(port)
		a.queue <- func() {
			if err != nil {
				reject(err)
			} else {
				resolve(conn)
			}
		}
	}()
	return promise
}

func (a *adapter) ConnClose(id string) *goja.Promise {
	promise, resolve, reject := a.vm.NewPromise()
	go func() {
		err := a.astral.ConnClose(id)
		a.queue <- func() {
			if err != nil {
				reject(err)
			} else {
				resolve(goja.Undefined())
			}
		}
	}()
	return promise
}

func (a *adapter) ConnWrite(id string, data string) *goja.Promise {
	promise, resolve, reject := a.vm.NewPromise()
	go func() {
		err := a.astral.ConnWrite(id, data)
		a.queue <- func() {
			if err != nil {
				reject(err)
			} else {
				resolve(goja.Undefined())
			}
		}
	}()
	return promise
}

func (a *adapter) ConnRead(id string) *goja.Promise {
	promise, resolve, reject := a.vm.NewPromise()
	go func() {
		val, err := a.astral.ConnRead(id)
		a.queue <- func() {
			if err != nil {
				reject(err)
			} else {
				resolve(val)
			}
		}
	}()
	return promise
}

func (a *adapter) Query(identity string, query string) *goja.Promise {
	promise, resolve, reject := a.vm.NewPromise()
	go func() {
		val, err := a.astral.Query(identity, query)
		a.queue <- func() {
			if err != nil {
				reject(err)
			} else {
				resolve(val)
			}
		}
	}()
	return promise
}

func (a *adapter) QueryName(name string, query string) *goja.Promise {
	promise, resolve, reject := a.vm.NewPromise()
	go func() {
		val, err := a.astral.Query(name, query)
		a.queue <- func() {
			if err != nil {
				reject(err)
			} else {
				resolve(val)
			}
		}
	}()
	return promise
}

func (a *adapter) Resolve(name string) *goja.Promise {
	promise, resolve, reject := a.vm.NewPromise()
	go func() {
		val, err := a.astral.Resolve(name)
		a.queue <- func() {
			if err != nil {
				reject(err)
			} else {
				resolve(val)
			}
		}
	}()
	return promise
}

func (a *adapter) NodeInfo(identity string) *goja.Promise {
	promise, resolve, reject := a.vm.NewPromise()
	go func() {
		val, err := a.astral.NodeInfo(identity)
		a.queue <- func() {
			if err != nil {
				reject(err)
			} else {
				resolve(val)
			}
		}
	}()
	return promise
}
