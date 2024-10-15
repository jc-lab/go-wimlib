package binding

import (
	"sync"
	"unsafe"
)

type retainerImpl struct {
	mutex   sync.Mutex
	objects map[unsafe.Pointer]interface{}
}

var retainer retainerImpl

func (r *retainerImpl) init() {
	if r.objects == nil {
		r.objects = make(map[unsafe.Pointer]interface{})
	}
}

func (r *retainerImpl) Keep(pointer unsafe.Pointer, obj interface{}) {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	r.init()
	r.objects[pointer] = obj
}

func (r *retainerImpl) Remove(pointer unsafe.Pointer) {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	delete(r.objects, pointer)
}
