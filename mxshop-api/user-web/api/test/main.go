package main

import (
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
)

// AccessKey ID LTAI5tAqbguUjTvFtxF7UvJa
// AccessKey Secret rFZLaazArJ6Yqjm72xLvOEx7AhLgCt

func main() {
	client, err := dysmsapi.NewClientWithAccessKey("cn-hangzhou", "LTAI5tAqbguUjTvFtxF7UvJa", "rFZLaazArJ6Yqjm72xLvOEx7AhLgCt")
	if err != nil {
		panic(err)
	}
	request := requests.NewCommonRequest()
	request.Method = "POST"
	request.Scheme = "https" // https | http
	request.Domain = "dysmsapi.aliyuncs.com"
	request.Version = "2017-05-25"
	request.ApiName = "SendSms"
	request.QueryParams["RegionId"] = "cn-hangzhou"
	request.QueryParams["PhoneNumbers"] = "18511903994"                  // 手机号
	request.QueryParams["SignName"] = "东方网信"                             // 阿里云验证过的项目名 自己设置
	request.QueryParams["TemplateCode"] = "SMS_162630385"                // 阿里云的短信模板号 自己设置
	// request.QueryParams["TemplateParam"] = "{\"code\":" + "777777" + "}" // 短信模板中的验证码内容 自己生成 之前试过直接返回，但是失败，加上code成功。
	request.QueryParams["TemplateParam"] = "{\"code\":" + "1234" + "}" // 短信模板中的验证码内容 自己生成 之前试过直接返回，但是失败，加上code成功。
	response, err := client.ProcessCommonRequest(request)
	fmt.Printf("DoAction is %#v\n", client.DoAction(request, response))
	//  fmt.Print(response)
	if err != nil {
		fmt.Print(err.Error())
	}
	fmt.Printf("response is %#v\n", response)
	// json数据解析
}
