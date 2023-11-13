package loader

import (
	"sync"
	"time"

	"github.com/graph-gophers/dataloader"
)

type loaderWithLastAccess struct {
	loader     *dataloader.Loader
	lastAccess int64
}

func getOrSetFromLoaderMaps(
	mutex *sync.Mutex,
	loaderMaps map[string]*loaderWithLastAccess,
	key string,
	fn func() *dataloader.Loader,
) *loaderWithLastAccess {
	mutex.Lock()
	loaderWrapper, ok := loaderMaps[key]
	if !ok {
		loaderWrapper = &loaderWithLastAccess{
			loader: fn(),
		}

		loaderMaps[key] = loaderWrapper
	}
	loaderWrapper.lastAccess = time.Now().Unix()
	mutex.Unlock()

	return loaderWrapper
}

func autoRemoveExpiredLoaderMaps(
	mutex *sync.Mutex,
	loaderMaps map[string]*loaderWithLastAccess,
	delay time.Duration,
	maxTtlInSecond int64,
	closeSig <-chan struct{},
) {
	go func() {
		ticker := time.NewTicker(delay)
		defer ticker.Stop()

		for {
			select {
			case now := <-ticker.C:
				mutex.Lock()
				for key, val := range loaderMaps {
					if now.Unix()-val.lastAccess > maxTtlInSecond {
						delete(loaderMaps, key)
					}
				}
				mutex.Unlock()
			case <-closeSig:
				return
			}
		}
	}()
}

func NewDataloader(batchFn dataloader.BatchFunc) dataloader.Loader {
	return *dataloader.NewBatchedLoader(
		batchFn,
	)
}

func NewDataLoaderP(batchFn dataloader.BatchFunc) *dataloader.Loader {
	loader := NewDataloader(batchFn)
	return &loader
}
