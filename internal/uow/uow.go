package uow

import user "github.com/fla-t/go-ing/internal/domain/user"

// UnitOfWorkInterface is a interface that holds all the interfaces to all the repositories
type UnitOfWorkInterface interface {
	Begin() error
	Commit() error
	Rollback() error
	UserRepository() user.RepositoryInterface
}
