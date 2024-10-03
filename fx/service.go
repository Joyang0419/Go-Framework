package main

type IService interface {
	Demo1() string
	Demo2() string
}

type Service struct {
	repo IRepo
}

func NewService(repo IRepo) IService {
	return &Service{repo: repo}
}

func (s *Service) Demo1() string {
	return s.repo.Demo1()
}

func (s *Service) Demo2() string {
	return s.repo.Demo2()
}
