package ding_bot

type (
	Param func(config *DingConfig)
)

func WithSecret(secret string) Param {
	return func(dt *DingConfig) {
		dt.secret = secret
	}
}
