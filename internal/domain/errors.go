package models

import "errors"

var (
	ForbiddenErr = errors.New("forbiden")
	TokenExpiredErr = errors.New("token expired")
	TokenInvalidErr = errors.New("token invalid")
	NotFoundErr = errors.New("not found")
	ConflictErr =  errors.New("already exist")
)

