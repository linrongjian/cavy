package util

import "net/http"

var (
	defaultLang = "zh-cn"
	Lang        = map[string]*Status{
		"en-us": NewStatus("en-us"),
		"zh-cn": NewStatus("zh-cn"),
	}
)

// SetDefaultLang 设置默认语言环境
func SetDefaultLang(lang string) {
	defaultLang = lang
}

// StatusText 获取默认语言环境
func StatusText(code int32) string {
	return Lang[defaultLang].Text(code)
}

// GetText 获取语言环境信息
func GetText(lang string, code int32) string {
	if s, ok := Lang[lang]; ok {
		return s.Text(code)
	}
	return "unknown language"
}

// AddStatus 注册语言环境
func AddStatus(status *Status) {
	Lang[status.Lang()] = status
}

// GetStatus 注册语言环境
func GetStatus(lang string) *Status {
	return Lang[lang]
}

// AddStatusText 注册语言环境信息
func AddStatusText(lang string, code int32, text string) {
	Lang[lang].Add(code, text)
}

// Status 语言包
type Status struct {
	lang   string           // 语言类型
	status map[int32]string // 错误码提示内容注册
}

func (s *Status) Lang() string {
	return s.lang
}

func (s *Status) Add(code int32, text string) {
	s.status[code] = text
}

func (s *Status) Text(code int32) string {
	if text, ok := s.status[code]; ok {
		return text
	}
	if text := http.StatusText(int(code)); text != "" {
		return text
	}
	return "unknown"
}

func (s *Status) Status() map[int32]string {
	return s.status
}

func NewStatus(lang string) *Status {
	return &Status{
		lang:   lang,
		status: make(map[int32]string),
	}
}
