package lib

type Singleton struct {
	Structs []string
}

var instance *Singleton

func NewSingleton() *Singleton {
	if instance == nil {
		instance = &Singleton{
			Structs: make([]string, 0, 0),
		}
	}
	return instance
}
