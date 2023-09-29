package header

import (
	"context"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2/metadata"
)

const (
	UserId            = "X-User-Id"
	UserToken         = "X-User-Token"
	XParams           = "X-Params"
	ForwardIP         = "X-Forwarded-For"
	RealIP            = "X-Real-Ip"
	XDevice           = "X-Device"
	XAuthorizationOut = "X-Authorization-Out"

	HVersion    = "Version"     // APP版本
	AppVersion  = "AppVersion"  // APP版本
	OSVersion   = "OSVersion"   // OS版本
	Platform    = "Platform"    // 平台 android/iOs
	Brand       = "Brand"       // 手机品牌
	Model       = "Model"       // 手机型号
	PackageName = "PackageName" // 包名称
	DeviceId    = "DeviceId"    // 设备ID
	Mac         = "Mac"         // mac地址
	Idfa        = "Idfa"        // idfa地址
)

// GetHeaderVal 从header中获取数据
func GetHeaderVal(ctx context.Context, key string) string {
	v, ok := metadata.Get(ctx, key)
	if !ok {
		return ""
	}

	return v
}

func VersionToInt(version string) int64 {
	v := strings.ReplaceAll(version, ".", "")
	parseInt, _ := strconv.ParseInt(v, 10, 64)
	return parseInt
}

// IsAllowVideo 是否展示视频 true:展示 false:不展示
func IsAllowVideo(ctx context.Context, version string) bool {
	platform := GetHeaderVal(ctx, Platform)
	if strings.ToLower(platform) != "ios" {
		return true
	}

	curVersion := GetHeaderVal(ctx, HVersion)
	cvi := VersionToInt(curVersion)
	bvi := VersionToInt(version)
	if cvi < bvi {
		return false
	}

	return true
}

// GetAppVersion 获取app版本
func GetAppVersion(ctx context.Context) string {
	return parseHeader(ctx, AppVersion)
}

// GetOsVersion 运行版本
func GetOsVersion(ctx context.Context) string {
	return parseHeader(ctx, OSVersion)
}

// GetPlatform iOS / android
func GetPlatform(ctx context.Context) string {
	return parseHeader(ctx, Platform)
}

// GetBrand 手机品牌
func GetBrand(ctx context.Context) string {
	return parseHeader(ctx, Brand)
}

// GetModel 手机型号
func GetModel(ctx context.Context) string {
	return parseHeader(ctx, Model)
}

// GetPackage 安装包名
func GetPackage(ctx context.Context) string {
	return parseHeader(ctx, PackageName)
}

// GetDeviceId 设备ID
func GetDeviceId(ctx context.Context) string {
	return parseHeader(ctx, DeviceId)
}

// GetMac mac地址
func GetMac(ctx context.Context) string {
	return parseHeader(ctx, Mac)
}

// GetIdfa Idfa地址
func GetIdfa(ctx context.Context) string {
	return parseHeader(ctx, Idfa)
}

// GetUserIdFromCtx 从context中获取用户ID
func GetUserIdFromCtx(ctx context.Context) int64 {
	var userId int64
	v, ok := metadata.Get(ctx, UserId)
	if !ok {
		return 0
	}

	userId, _ = strconv.ParseInt(v, 10, 64)
	if userId <= 0 {
		userId = 0
	}

	return userId
}

// GetUserIdFromGinCtx 从context中获取用户ID
func GetUserIdFromGinCtx(ctx *gin.Context) int64 {
	var userId int64
	v := ctx.Request.Header.Get(UserId)
	userId, _ = strconv.ParseInt(v, 10, 64)
	if userId <= 0 {
		userId = 0
	}
	return userId
}

// GetXAuthorizationOutFromCtx 从context中获取XAuthorizationOut
func GetXAuthorizationOutFromCtx(ctx context.Context) string {
	v, ok := metadata.Get(ctx, XAuthorizationOut)
	if !ok {
		return ""
	}
	return v
}

// GetClientIP 从context中获取ip
func GetClientIP(ctx context.Context) string {
	v, ok := metadata.Get(ctx, ForwardIP)
	if !ok || len(v) == 0 {
		if cv, cok := metadata.Get(ctx, RealIP); cok && len(cv) > 0 {
			v = cv
		}
	}
	return strings.TrimSpace(strings.SplitN(v, ",", 2)[0])
}

// 参数格式：AppVersion=1.0.0;OSVersion=1.0.0;Platform=iOS
func parseHeader(ctx context.Context, key string) string {
	s, ok := metadata.Get(ctx, XParams)
	if !ok {
		return s
	}
	arr := strings.Split(s, ";")
	if len(arr) > 0 {
		for _, item := range arr {
			kv := strings.Split(item, "=")
			if len(kv) != 2 {
				continue
			}
			if strings.ToLower(kv[0]) == strings.ToLower(key) {
				return kv[1]
			}
		}
	}
	return ""
}
