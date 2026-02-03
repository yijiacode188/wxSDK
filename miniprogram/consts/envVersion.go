package consts

type EnvVersion string

const (
	EnvVersionRelease EnvVersion = "release" // 正式版
	EnvVersionTrial   EnvVersion = "trial"   // 体验版
	EnvVersionDevelop EnvVersion = "develop" // 开发版
)
