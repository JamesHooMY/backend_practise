package user

import (
	"context"
	"errors"

	common "go_backend/app/repo/mysql"
	"go_backend/domain"
	"go_backend/model"

	"gorm.io/gorm"
)

type userQueryRepo struct {
	db *gorm.DB
}

func NewUserQueryRepo(db *gorm.DB) domain.IUserQueryRepo {
	return &userQueryRepo{
		db: db,
	}
}

func (q *userQueryRepo) GetUserByEmail(ctx context.Context, email string) (user *model.User, err error) {
	err = q.db.WithContext(ctx).Model(&model.User{}).Where("email = ?", email).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}

	return user, nil
}

func (q *userQueryRepo) GetUserByID(ctx context.Context, id uint) (user *model.User, err error) {
	err = q.db.WithContext(ctx).Model(&model.User{}).Where("id = ?", id).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}

	return user, nil
}

func (q *userQueryRepo) GetUserList(ctx context.Context, page, limit int) (userList []*model.User, total int64, err error) {
	err = q.db.WithContext(ctx).Model(&model.User{}).
		Scopes(common.Pagination(page, limit)).Find(&userList). // pagination
		Offset(-1).Limit(-1).Count(&total).                     // get total
		Error
	if err != nil {
		return nil, 0, err
	}

	return userList, total, nil
}
