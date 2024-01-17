package handler

import (
	"context"
	"crypto/sha512"
	"fmt"
	"github.com/anaskhan96/go-password-encoder"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
	"mvshop_srvs/user_srv/global"
	"mvshop_srvs/user_srv/model"
	"mvshop_srvs/user_srv/proto"
	"strings"
	"time"
)

func ModelToResponse(user model.User) __proto.UserInfoResponse {
	// 在grpc 的message 中字段有默认值 你不能随便赋值nil进去，容易出错
	userInfoRsp := __proto.UserInfoResponse{
		Id:       user.ID,
		PassWord: user.Password,
		NickName: user.NickName,
		Gender:   user.Gender,
		Mobile:   user.Mobile,
		Role:     int32(user.Role),
	}
	if user.Birthday != nil {
		userInfoRsp.BirthDay = uint64(user.Birthday.Unix())
	}
	return userInfoRsp
}

func Paginate(page, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page == 0 {
			page = 1
		}
		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}
		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}

type UserServer struct {
}

// GetUserList(context.Context, *PageInfo) (*UserListResponse, error)
func (s *UserServer) GetUserList(ctx context.Context, req *__proto.PageInfo) (*__proto.UserListResponse, error) {
	// 获取用户列表
	// 获取全局gin连接
	var users []model.User
	result := global.DB.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	rsp := &__proto.UserListResponse{}
	rsp.Total = int32(result.RowsAffected)

	global.DB.Scopes(Paginate(int(req.Pn), int(req.PSize))).Find(&users)

	for _, user := range users {
		userInfoRsp := ModelToResponse(user)
		rsp.Data = append(rsp.Data, &userInfoRsp)
	}
	return rsp, nil
}

// 通过手机号码查询用户
func (s *UserServer) GetUserByMobile(ctx context.Context, req *__proto.MobileRequest) (*__proto.UserInfoResponse, error) {
	var user model.User
	result := global.DB.Where(&model.User{Mobile: req.Mobile}).First(&user)
	if result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "用户不存在")
	}

	userInfoRsp := ModelToResponse(user)

	return &userInfoRsp, nil
}

// 通过ID查询用户
func (s *UserServer) GetUserById(ctx context.Context, req *__proto.IdRequest) (*__proto.UserInfoResponse, error) {
	var user model.User
	result := global.DB.First(&user, req.Id)
	if result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "用户不存在")
	}

	userInfoRsp := ModelToResponse(user)

	return &userInfoRsp, nil
}

func (s *UserServer) CreateUser(ctx context.Context, req *__proto.CreateUserInfo) (*__proto.UserInfoResponse, error) {
	var user model.User
	result := global.DB.Where(&model.User{Mobile: req.Mobile}).First(&user)
	if result.RowsAffected == 1 {
		return nil, status.Errorf(codes.AlreadyExists, "用户已存在")
	}
	user.Mobile = req.Mobile
	user.NickName = req.NickName
	// 密码加密
	options := &password.Options{16, 100, 32, sha512.New}
	salt, encodedPwd := password.Encode(req.PassWord, options)
	user.Password = fmt.Sprintf("$pdkdf2-sha512$%s$%s", salt, encodedPwd)

	result = global.DB.Create(&user)
	if result.Error != nil {
		return nil, status.Errorf(codes.Internal, result.Error.Error())
	}

	userInfoRsp := ModelToResponse(user)
	return &userInfoRsp, nil

	// fmt.Println(newPassword)
	// passwordInfo := strings.Split(newPassword, "$")
	// fmt.Println(passwordInfo)
	// check := password.Verify("generic password", salt, encodedPwd, options)
	// fmt.Println(check) // true
}

// UpdateUser(context.Context, *UpdateUserInfo) (*google_protobuf.Empty, error)
// CheckPassWord(context.Context, *CheckPassWordInfo) (*CheckResponse, error)
// Go to Method Specifications
func (s *UserServer) UpdateUser(ctx context.Context, req *__proto.UpdateUserInfo) (*empty.Empty, error) {
	// 个人中心更新用户
	var user model.User
	result := global.DB.First(&user, req.Id)
	if result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "用户不存在")
	}

	birthDay := time.Unix(int64(req.BirthDay), 0)
	user.NickName = req.NickName
	user.Birthday = &birthDay
	user.Gender = req.Gender

	result = global.DB.Save(&user)
	if result.Error != nil {
		return nil, status.Errorf(codes.Internal, result.Error.Error())
	}

	return &empty.Empty{}, nil

}

// 检查密码
func (s *UserServer) CheckPassWord(ctx context.Context, req *__proto.CheckPassWordInfo) (*__proto.CheckResponse, error) {
	options := &password.Options{16, 100, 32, sha512.New}
	passwordInfo := strings.Split(req.EncryptedPassword, "$")
	check := password.Verify(req.Password, passwordInfo[2], passwordInfo[3], options)
	return &__proto.CheckResponse{Success: check}, nil
}
