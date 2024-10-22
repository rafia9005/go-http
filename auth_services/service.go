package main

import "context"

type Service interface {
  Get(context.Context) (*GetFact, error)
}

type GetFactService struct {
  url string
}

func (s *GetFactService) GetCatFact(ctx context.Context) (*GetFact, error) {
  
}
