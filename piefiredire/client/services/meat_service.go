package services

import (
	"context"
	"fmt"
)

type MeatCounterService interface {
	GetMeatSummary(input string) error
}

type meatCounterService struct {
	meatCounterClient MeatCounterClient
}

func NewMeatCounterService(meatCounterClient MeatCounterClient) MeatCounterService {
	return meatCounterService{meatCounterClient}
}

func (base meatCounterService) GetMeatSummary(input string) error {
	req := MeatRequest{
		Input: input,
	}

	res, err := base.meatCounterClient.GetMeatSummary(context.Background(), &req)
	if err != nil {
		return err
	}

	fmt.Printf("Service : Hello\n")
	fmt.Printf("Request : %v\n", req.Input)
	fmt.Printf("Response: %v\n", res.Beef)
	return nil
}
