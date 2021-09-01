package utils

type Bindable interface {
	BindKey() string
}

type containerDJ struct {
	m map[string]interface{}
}

var Container *containerDJ

func init() {
	Container = &containerDJ{m: make(map[string]interface{})}
}

func (c *containerDJ) Bind(key string, instance interface{}) {
	c.m[key] = instance
}

func (c *containerDJ) BindObject(bindable Bindable) {
	c.m[bindable.BindKey()] = bindable
}

func (c *containerDJ) Lookup(key string) (interface{}, bool) {
	v, ok := c.m[key]
	return v, ok
}

func (c *containerDJ) Get(key string) interface{} {
	v := c.m[key]
	return v
}

func (c *containerDJ) GetFromBindAble(bindable Bindable) interface{} {
	v := c.m[bindable.BindKey()]
	return v
}
