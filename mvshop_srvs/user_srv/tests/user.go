package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	__proto "mvshop_srvs/user_srv/proto"
)

var userClient __proto.UserClient

var conn *grpc.ClientConn

func Init() {
	var err error
	conn, err = grpc.Dial("127.0.0.1:50051", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	userClient = __proto.NewUserClient(conn)
}

func TestGetUserList() {

	rsp, err := userClient.GetUserList(context.Background(), &__proto.PageInfo{
		Pn:    1,
		PSize: 4,
	})
	if err != nil {
		panic(err)
	}
	for _, user := range rsp.Data {
		fmt.Println(user.Mobile, user.NickName, user.PassWord)
		checkRsp, err := userClient.CheckPassWord(context.Background(), &__proto.CheckPassWordInfo{
			Password:          "admin123",
			EncryptedPassword: user.PassWord,
		})
		if err != nil {
			panic(err)
		}
		fmt.Println(checkRsp)
	}
}

func TestCreateUser() {
	// options := &password.Options{16, 100, 32, sha512.New}
	// salt, encodedPwd := password.Encode("admin123", options)
	// newPassword := fmt.Sprintf("$pdkdf2-sha512$%s$%s", salt, encodedPwd)
	// 18782222220 admin123
	for i := 0; i < 10; i++ {
		rsp, err := userClient.CreateUser(context.Background(), &__proto.CreateUserInfo{
			NickName: fmt.Sprintf("bobby%d", i),
			Mobile:   fmt.Sprintf("1878222222%d", i),
			PassWord: "admin123",
		})
		if err != nil {
			panic(err)
		}
		fmt.Println(rsp.Id)
	}
}

func main() {
	Init()
	defer conn.Close()
	TestGetUserList()
	// TestCreateUser()
}
