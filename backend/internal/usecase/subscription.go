package usecase

import "cooking/backend/internal/usecase/subscription"

type SubscriptionUseCase struct {
	SubscriptionRepo subscription.Repo
}

func SubscriptionInstance(repo subscription.Repo) *SubscriptionUseCase {
	return &SubscriptionUseCase{SubscriptionRepo: repo}
}
