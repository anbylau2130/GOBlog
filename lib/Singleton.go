package lib

type Singleton struct {
	USERSESSION  string
	THEMESESSION string
}

var instance *Singleton

func NewSingleton() *Singleton {
	if instance == nil {
		instance = &Singleton{
			USERSESSION:  "CurrentUser",
			THEMESESSION: "CurrentTheme",
		}
	}
	return instance
}
