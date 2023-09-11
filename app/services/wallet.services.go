package services

import (
	"context"
	"errors"
	"fmt"
	"lido-core/v1/app/models"
	"lido-core/v1/pkg/constants"
	"lido-core/v1/platform/cache"
	"log"
	"math/big"

	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/bson"
)

func BalanceOf(owner string) (float64, error) {
	balanceStr, err := cache.Get(fmt.Sprintf("%s%s", constants.RedisBalanceFormat, owner))
	if err != nil {
		if err == redis.Nil {
			collections := models.WalletCollection()
			var wallet models.Wallet
			if err := collections.FindOne(context.TODO(), bson.M{"address": owner}).Decode(&wallet); err != nil {
				if err := cache.Save(fmt.Sprintf("%s%s", constants.RedisBalanceFormat, owner), "0"); err != nil {
					return 0, err
				}
				return 0, err
			}
			balance, ok := new(big.Float).SetString(wallet.Balance)
			if !ok {
				return 0, errors.New("error convert float number")
			}
			balance.Quo(balance, big.NewFloat(1e18))
			res, _ := balance.Float64()
			if err := cache.Save(fmt.Sprintf("%s%s", constants.RedisBalanceFormat, owner), wallet.Balance); err != nil {
				return 0, err
			}
			return res, nil
		}
		log.Fatal(err)
	}
	balance, ok := new(big.Float).SetString(balanceStr)
	if !ok {
		log.Fatal(balanceStr)
	}
	balance.Quo(balance, big.NewFloat(1e18))
	res, _ := balance.Float64()
	return res, nil
}
