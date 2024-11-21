package postges

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"

	"comics/storage"
)

type Store struct {
	db         *pgxpool.Pool
	user       storage.UserRepoI
	role       storage.RoleRepoI
	permission storage.PermissionRepoI
	userRole   storage.UserRoleRepoI
	rolePermission storage.RolePermissionRepoI
	comics storage.ComicRepoI
	comicsReview storage.ComicReviewRepoI
	comicsPages storage.ComicPageRepoI
	order storage.OrderRepoI
	orderItem storage.OrderItemRepoI
}

func NewPostgres(psqlConnString string) storage.StorageRepoI {
	config, err := pgxpool.ParseConfig(psqlConnString)
	if err != nil {
		log.Panicf("Unable to parse connection string: %v", err)
	}

	pool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		log.Panicf("Unable to connect to the database: %v", err)
	}

	return &Store{
		db: pool,
	}
}

func (s *Store) CloseDB() {
	s.db.Close()
}

func (s *Store) User() storage.UserRepoI {
	if s.user == nil {
		s.user = &userRepo{
			db: s.db,
		}
	}
	return s.user
}

func (s *Store) Role() storage.RoleRepoI {
	if s.role == nil {
		s.role = &roleRepo{
			db: s.db,
		}
	}
	return s.role
}

func (s *Store) Permission() storage.PermissionRepoI {
	if s.permission == nil {
		s.permission = &permissionRepo{
			db: s.db,
		}
	}
	return s.permission
}

func (s *Store) UserRole() storage.UserRoleRepoI {
	if s.userRole == nil {
		s.userRole = &userRoleRepo{
			db: s.db,
		}
	}
	return s.userRole
}

func (s *Store) RolePermission() storage.RolePermissionRepoI {
	if s.rolePermission == nil {
		s.rolePermission = &rolePermissionRepo{
			db: s.db,
		}
	}
	return s.rolePermission
}

func (s *Store) Comics() storage.ComicRepoI {
	if s.comics == nil {
		s.comics = &comicsRepo{
			db: s.db,
		}
	}
	return s.comics
}

func (s *Store) ComicReview() storage.ComicReviewRepoI {
	if s.comicsReview == nil {
		s.comicsReview = &comicsReviewRepo{
			db: s.db,
		}
	}
	return s.comicsReview
}

func (s *Store) ComicsPages() storage.ComicPageRepoI {
	if s.comicsPages == nil {
		s.comicsPages = &comicsPagesRepo{
			db: s.db,
		}
	}
	return s.comicsPages
}

func (s *Store) Order() storage.OrderRepoI {
	if s.order == nil {
		s.order = &orderRepo{
			db: s.db,
		}
	}
	return s.order
}

func (s *Store) OrderItem() storage.OrderItemRepoI {
	if s.orderItem == nil {
		s.orderItem = &orderItemRepo{
			db: s.db,
		}
	}
	return s.orderItem
}