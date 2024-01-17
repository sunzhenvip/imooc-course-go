package api

import (
	"context"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin/binding"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"mxshop-api/user-web/middlewares"
	"mxshop-api/user-web/models"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"mxshop-api/user-web/forms"
	"mxshop-api/user-web/global"
	"mxshop-api/user-web/global/reponse"
	"net/http"
	"strconv"
	"strings"
	"time"

	myvalidator "mxshop-api/user-web/validator"

	"mxshop-api/user-web/proto"
)

func removeTopStruct(fileds map[string]string) map[string]string {
	rsp := map[string]string{}
	for field, err := range fileds {
		rsp[field[strings.Index(field, ".")+1:]] = err
	}
	return rsp
}

func HandleGrpcErrorToHttp(err error, c *gin.Context) {
	// 将grpc都code转换成http都状态码
	if err != nil {
		if e, ok := status.FromError(err); ok {
			switch e.Code() {
			case codes.NotFound:
				c.JSON(http.StatusNotFound, gin.H{
					"msg": e.Message(),
				})
			case codes.Internal:
				c.JSON(http.StatusInternalServerError, gin.H{
					"msg": "内部错误",
				})
			case codes.InvalidArgument:
				c.JSON(http.StatusBadRequest, gin.H{
					"msg": "参数错误",
				})
			case codes.Unavailable:
				c.JSON(http.StatusInternalServerError, gin.H{
					"msg": "用户服务不可用" + e.Message(),
				})
			default:
				c.JSON(http.StatusInternalServerError, gin.H{
					"msg": "其他错误" + e.Message(),
				})
			}
			return
		}
	}
}

func HandleValidatorError(c *gin.Context, err error) {
	errs, ok := err.(validator.ValidationErrors)
	if !ok {
		c.JSON(http.StatusOK, gin.H{
			"msg": err.Error(),
		})
	}
	c.JSON(http.StatusBadRequest, gin.H{
		"error": removeTopStruct(errs.Translate(global.Trans)),
	})
	return
}

func GetUserList(ctx *gin.Context) {
	// ip := "127.0.0.1"
	// port := 50051
	// cfg := consulApi.DefaultConfig()
	// consulInfo := fmt.Sprintf("%s:%d", global.ServerConfig.ConsulInfo.Host, global.ServerConfig.ConsulInfo.Port)
	// cfg.Address = consulInfo
	// userSrvHost := ""
	// userSrvPort := 0
	// client, err := consulApi.NewClient(cfg)
	// 查找数据
	// data, err := client.Agent().ServicesWithFilter(fmt.Sprintf(`Service == "%s"`, global.ServerConfig.UserSrvInfo.Name))
	// if err != nil {
	// 	panic(err)
	// }
	// 目前只获取一个
	// for _, val := range data {
	// 	userSrvHost = val.Address
	// 	userSrvPort = val.Port
	// 	break
	// }
	// 拨号连接
	// userConn, err := grpc.Dial(fmt.Sprintf("%s:%d", userSrvHost, userSrvPort), grpc.WithInsecure())
	// if err != nil {
	// 	zap.S().Errorw("[GteUserList]连接 用户服务失败", "msg", err.Error())
	// }
	claims, _ := ctx.Get("claims")
	currentUser, _ := claims.(*models.CustomClaims)
	zap.S().Infof("访问用户：%d", currentUser.ID)
	// 调用接口
	// userSrvClient := __proto.NewUserClient(userConn)

	pn := ctx.DefaultQuery("pn", "0")
	pnInt, _ := strconv.Atoi(pn)
	psize := ctx.DefaultQuery("psize", "10")
	psizeInt, _ := strconv.Atoi(psize)
	rsp, err := global.UserSrvClient.GetUserList(context.Background(), &__proto.PageInfo{
		Pn:    uint32(pnInt),
		PSize: uint32(psizeInt),
	})
	if err != nil {
		zap.S().Errorw("[GteUserList]查询用户列表失败", "msg", err.Error())
		HandleGrpcErrorToHttp(err, ctx)
		return
	}
	result := make([]interface{}, 0)
	for _, val := range rsp.Data {
		// data := make(map[string]interface{})

		user := reponse.UserResponse{
			Id:       val.Id,
			NickName: val.NickName,
			// BirthDay: time.Time(time.Unix(int64(val.BirthDay), 0)).Format("2006-01-02"),
			BirthDay: reponse.JsonTime(time.Unix(int64(val.BirthDay), 0)),
			Gender:   val.Gender,
			Mobile:   val.Mobile,
		}

		// data["id"] = val.Id
		// data["name"] = val.NickName
		// data["birthday"] = val.BirthDay
		// data["gender"] = val.Gender
		// data["mobile"] = val.Mobile

		// result = append(result, data)

		result = append(result, user)
	}
	ctx.JSON(http.StatusOK, result)
	// zap.S().Debug("获取列表页")
}

func Test(c *gin.Context) {
	// 表单验证
	passwordLoginForm := forms.PassWordLoginForm{}
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		_ = v.RegisterValidation("mobile", myvalidator.ValidateMobile)
		// 翻译自定义为中文描述
		_ = v.RegisterTranslation("mobile", global.Trans, func(ut ut.Translator) error {
			return ut.Add("mobile", "{0} 非法都手机号码", true)
		}, func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T("mobile", fe.Field())
			return t
		})
	}
	if err := c.ShouldBind(&passwordLoginForm); err != nil {
		// 如何返回错误信息
		HandleValidatorError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "测试",
	})
	return
}

// 验证登陆
func PassWordLogin(c *gin.Context) {
	// 表单验证
	passwordLoginForm := forms.PassWordLoginForm{}
	if err := c.ShouldBind(&passwordLoginForm); err != nil {
		// 如何返回错误信息
		HandleValidatorError(c, err)
		return
	}
	//
	// userServerAddress := fmt.Sprintf("%s:%d", global.ServerConfig.UserSrvInfo.Host, global.ServerConfig.UserSrvInfo.Port)
	// // 拨号连接
	// userConn, err := grpc.Dial(userServerAddress, grpc.WithInsecure())
	// if err != nil {
	// 	zap.S().Errorw("[GteUserList]连接 用户服务失败", "msg", err.Error())
	// }
	// 生成grpc都client并调用接口
	// userSrvClient := __proto.NewUserClient(userConn)

	// 登陆都逻辑
	rsp, err := global.UserSrvClient.GetUserByMobile(context.Background(), &__proto.MobileRequest{
		Mobile: passwordLoginForm.Mobile,
	})
	if err != nil {
		if e, ok := status.FromError(err); ok {
			switch e.Code() {
			case codes.NotFound:
				c.JSON(http.StatusBadRequest, map[string]string{
					"mobile": "用户不存在",
				})
			default:
				c.JSON(http.StatusBadRequest, map[string]string{
					"mobile": "登陆失败",
				})
			}
			return
		}
	}

	passRsp, pasErr := global.UserSrvClient.CheckPassWord(context.Background(), &__proto.CheckPassWordInfo{
		Password:          passwordLoginForm.PassWord,
		EncryptedPassword: rsp.PassWord,
	})

	fmt.Println(passRsp)
	fmt.Println(12121)

	if pasErr != nil {
		c.JSON(http.StatusBadRequest, map[string]string{
			"mobile": "密码错误",
		})
		return
	} else {
		if passRsp.Success {
			j := middlewares.NewJWT()
			claims := models.CustomClaims{
				ID:          uint(rsp.Id),
				NickName:    rsp.NickName,
				AuthorityId: uint(rsp.Role),
				StandardClaims: jwt.StandardClaims{
					NotBefore: time.Now().Unix(),
					ExpiresAt: time.Now().Unix() + 60*60*24*30, // 30天过期
					Issuer:    "imooc",
				},
			}
			token, err := j.CreateToken(claims)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"msg": "生成token失败",
				})
				return
			}
			c.JSON(http.StatusOK, gin.H{
				"id":         rsp.Id,
				"nick_name":  rsp.NickName,
				"token":      token,
				"expired_at": (time.Now().Unix() + 60*60*24*30) * 1000,
			})
		} else {
			c.JSON(http.StatusBadRequest, map[string]string{
				"mobile": "密码错误",
			})
			return
		}
		// c.JSON(http.StatusBadRequest, map[string]string{
		// 	"mobile": "登陆成功",
		// })
	}

}
