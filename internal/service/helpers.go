package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/tredoc/go-crud-api/internal/cache"
	"github.com/tredoc/go-crud-api/pkg/log"
	"strings"
	"time"
)

func setToCache(fn func(key string, value interface{}, expiration time.Duration) error, key string, entity any, expiration time.Duration) {
	str, err := json.Marshal(entity)
	if err != nil {
		log.Error(fmt.Sprintf("cache error on converting struct %s to string: %s", key, err.Error()))
	}

	err = fn(key, str, expiration)
	if err != nil {
		log.Error("cache error on save: " + err.Error())
	}
}

func getFromCache(fn func(key string) (string, error), key string, entity any) error {
	authorCached, err := fn(key)
	if err == nil {
		decodeErr := json.NewDecoder(strings.NewReader(authorCached)).Decode(entity)
		if decodeErr != nil {
			log.Error(fmt.Sprintf("cache error on decoding %s from cache string: %s", key, err.Error()))
		}
		return nil
	}

	if err != nil && !errors.Is(err, cache.ErrNotFound) {
		log.Error("can't get author from cache: ", err.Error())
	}
	return err
}
