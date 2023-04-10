package services

import (
	"context"
	"strings"
)

type meatCounterServer struct {
}

func NewMeatCounterServer() MeatCounterServer {
	return meatCounterServer{}
}

func (meatCounterServer) mustEmbedUnimplementedMeatCounterServer() {}

func (s meatCounterServer) GetMeatSummary(ctx context.Context, in *MeatRequest) (*MeatSummary, error) {
	meats := strings.Split(in.Input, " ")
	meatCounter := make(map[string]int32)

	for _, meat := range meats {
		if strings.ContainsAny(meat, ",") {
			meat = strings.ReplaceAll(meat, ",", "")
		}
		if strings.ContainsAny(meat, ".") {
			meat = strings.ReplaceAll(meat, ".", "")
		}
		if meat != "" {
			meat = strings.ToLower(meat)
			meatCounter[meat]++
		}
	}

	return &MeatSummary{Beef: meatCounter}, nil
}
