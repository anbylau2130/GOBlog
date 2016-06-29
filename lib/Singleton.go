package lib

type Singleton struct {
}

var instance *Singleton

func NewSingleton() *Singleton {
	if instance == nil {
		instance = &Singleton{}
	}
	return instance
}
