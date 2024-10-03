package main

type IRepo interface {
	Demo1() string
	Demo2() string
}

type Repo struct{}

func NewRepo() IRepo {
	return &Repo{}
}

func (r *Repo) Demo1() string {
	return "Demo1"
}

func (r *Repo) Demo2() string {
	return "Demo2"
}
