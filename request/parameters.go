package request

import (
	"errors"
	"net/url"
	"sync"
)

type Parameters interface {
	Set(string, string) error
	SetIfNotEmpty(string, string) error
	SetMap(map[string]string) error
	SetArr([]string) error
	Get(string) string
	Remove(string)
	Has(string) bool
	Clear()
	Clone() Parameters
	BuildURLValues() *url.Values
	BuildURLValuesEncode() string
}

type BaseParameters struct {
	mtx             sync.RWMutex
	data            map[string]string
	urlValues       *url.Values // cache
	urlValuesEncode string      // cache
}

func NewBaseRequestParameters(m map[string]string) *BaseParameters {
	p := &BaseParameters{data: make(map[string]string), urlValues: nil}
	err := p.SetMap(m)
	if err != nil {
		panic(err)
	}
	return p
}

func NewBaseRequestParametersArr(arr []string) *BaseParameters {
	p := &BaseParameters{data: make(map[string]string), urlValues: nil}
	err := p.SetArr(arr)
	if err != nil {
		panic(err)
	}
	return p
}

func (p *BaseParameters) Set(key, value string) error {
	if key == "" {
		return errors.New("BaseParameters.Set(): empty key")
	}

	p.mtx.Lock()
	p.data[key] = value
	p.resetCache()
	p.mtx.Unlock()

	return nil
}

func (p *BaseParameters) SetIfNotEmpty(key, value string) error {
	if value != "" {
		return p.Set(key, value)
	}

	return nil
}

func (p *BaseParameters) SetMap(m map[string]string) error {
	if m == nil {
		return errors.New("BaseParameters.SetMap(): map is nil")
	}

	p.mtx.Lock()
	defer p.mtx.Unlock()

	for k, v := range m {
		p.data[k] = v
	}

	p.resetCache()

	return nil
}

// SetArr [key, value, key, value...]
func (p *BaseParameters) SetArr(pairs []string) error {
	if len(pairs) == 0 || len(pairs)%2 != 0 {
		return errors.New("BaseParameters.SetArr(): pairs length must be even")
	}

	p.mtx.Lock()
	defer p.mtx.Unlock()

	for i := 0; i < len(pairs)-1; i = i + 2 {
		p.data[pairs[i]] = pairs[i+1]
	}

	p.resetCache()

	return nil
}

func (p *BaseParameters) Get(key string) string {
	p.mtx.RLock()

	value, ok := p.data[key]
	if ok {
		p.mtx.RUnlock()
		return value
	}

	p.mtx.RUnlock()

	return ""
}

func (p *BaseParameters) Remove(key string) {
	p.mtx.Lock()
	delete(p.data, key)
	p.resetCache()
	p.mtx.Unlock()
}

func (p *BaseParameters) Has(key string) bool {
	p.mtx.RLock()
	_, ok := p.data[key]
	p.mtx.RUnlock()
	return ok
}

func (p *BaseParameters) Clear() {
	p.mtx.Lock()
	clear(p.data)
	p.resetCache()
	p.mtx.Unlock()
}

func (p *BaseParameters) Clone() Parameters {
	p.mtx.RLock()
	clone := NewBaseRequestParameters(p.data)
	p.mtx.RUnlock()
	return clone
}

func (p *BaseParameters) BuildURLValues() *url.Values {
	p.mtx.RLock()
	if p.urlValues != nil {
		defer p.mtx.RUnlock()
		return p.urlValues
	}
	p.mtx.RUnlock()

	p.mtx.Lock()
	defer p.mtx.Unlock()

	p.urlValues = &url.Values{}
	for key, value := range p.data {
		if value != "" {
			p.urlValues.Add(key, value)
		}
	}

	return p.urlValues
}

func (p *BaseParameters) BuildURLValuesEncode() string {
	p.mtx.RLock()
	if p.urlValuesEncode != "" {
		defer p.mtx.RUnlock()
		return p.urlValuesEncode
	}
	p.mtx.RUnlock()

	p.mtx.Lock()
	defer p.mtx.Unlock()

	if p.urlValues == nil {
		p.urlValues = &url.Values{}
		for key, value := range p.data {
			if value != "" {
				p.urlValues.Add(key, value)
			}
		}
	}

	p.urlValuesEncode = p.urlValues.Encode()

	return p.urlValuesEncode
}

func (p *BaseParameters) resetCache() {
	p.urlValues = nil
	p.urlValuesEncode = ""
}
