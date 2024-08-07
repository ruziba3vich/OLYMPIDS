package storage

import (
	"medal-service/repositroy"
)

type MainStorageImpl struct {
	// redisImpl    repositroy.MedalRepository
	postgresImpl repositroy.MedalRepository
}

func NewMainStorageImpl(postgresImpl repositroy.MedalRepository) *MainStorageImpl {
	return &MainStorageImpl{
		postgresImpl: postgresImpl,
	}
}

func (impl *MainStorageImpl) PostgresImpl() repositroy.MedalRepository {
	return impl.postgresImpl
}

// func (impl *MainStorageImpl) RedisImpl() repositroy.MedalRepository {
// 	return impl.redisImpl
// }
