package hook

import "github.com/linrongjian/cavy/common/api"

var hooks = []func(s *api.Conf){}

// AddHook 添加钩子
func AddHook(hook func(s *api.Conf)) {
	hooks = append(hooks, hook)
}

// CallHook 调用钩子
func CallHook(s *api.Conf) {
	for _, h := range hooks {
		h(s)
	}
}
