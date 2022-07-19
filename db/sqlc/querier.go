// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0

package db

import (
	"context"
)

type Querier interface {
	CreateKategori(ctx context.Context, arg CreateKategoriParams) (Kategori, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	GetMobilJoinMany(ctx context.Context, arg GetMobilJoinManyParams) ([]GetMobilJoinManyRow, error)
	GetUser(ctx context.Context, username string) (User, error)
	ListKategori(ctx context.Context, arg ListKategoriParams) ([]Kategori, error)
}

var _ Querier = (*Queries)(nil)
