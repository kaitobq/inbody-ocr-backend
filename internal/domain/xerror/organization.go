package xerror

import "errors"

var (
	ErrUserIsNotAdmin        = errors.New("user is not admin")
	ErrCannotUpdateOwnerRole = errors.New("cannot update owner role")
	ErrCannotDeleteOwner     = errors.New("cannot delete owner")
)
