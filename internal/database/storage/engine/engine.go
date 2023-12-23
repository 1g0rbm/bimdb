package engine

type IEngine interface {
	Set(key string, value string)
	Get(key string) (string, bool)
	Del(key string)
}
