package main

var Repo *Respository

type Respository struct {
	App *AppConfig
}

func NewRepo(a *AppConfig) *Respository {
	return &Respository{
		App: a,
	}
}

func NewHandlers(r *Respository) {
	Repo = r
}
