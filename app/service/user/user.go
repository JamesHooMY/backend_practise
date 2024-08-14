package user

import (
	"context"
	"errors"

	"go_backend/domain"
	"go_backend/model"
)

var ErrPasswordIncorrect = errors.New("password incorrect")

type userService struct {
	userQryRepo domain.IUserQueryRepo
	userCmdRepo domain.IUserCommandRepo
}

// add database repo here
func NewUserService(userQryRepo domain.IUserQueryRepo, userCmdRepo domain.IUserCommandRepo) domain.IUserService {
	return &userService{
		userQryRepo: userQryRepo,
		userCmdRepo: userCmdRepo,
	}
}

func (s *userService) GetUserByID(ctx context.Context, id uint) (userResp *domain.UserResp, err error) {
	user, err := s.userQryRepo.GetUserByID(ctx, id)
	if err != nil {
		return nil, err
	}

	userResp = &domain.UserResp{
		ID:        user.ID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		DeletedAt: user.DeletedAt.Time,
		Mobile:    user.Mobile,
		Name:      user.Name,
		Age:       user.Age,
	}

	return userResp, nil
}

func (s *userService) GetUserList(ctx context.Context, page, limit int) (userListResp *domain.UserListResp, err error) {
	userList, total, err := s.userQryRepo.GetUserList(ctx, page, limit)
	if err != nil {
		return nil, err
	}

	userListResp = &domain.UserListResp{
		UserList: make([]*domain.UserResp, 0, len(userList)),
		Total:    total,
		Page:     page,
		Limit:    limit,
	}

	for _, user := range userList {
		userListResp.UserList = append(userListResp.UserList, &domain.UserResp{
			ID:        user.ID,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
			DeletedAt: user.DeletedAt.Time,
			Mobile:    user.Mobile,
			Name:      user.Name,
			Age:       user.Age,
		})
	}

	return userListResp, nil
}

func (s *userService) UpdateUser(ctx context.Context, user *model.User) (err error) {
	err = s.userCmdRepo.UpdateUser(ctx, user)
	if err != nil {
		return err
	}

	return nil
}

func (s *userService) DeleteUserByID(ctx context.Context, id uint) (err error) {
	err = s.userCmdRepo.DeleteUser(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
