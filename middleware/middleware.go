package middleware

import "github.com/dengzii/blog/com.dengzii.blog/bootstrap"

type Middleware interface {
	Attach(bootstrap *bootstrap.Bootstrapper)
}

func WithBootstrap(bootstrap *bootstrap.Bootstrapper) {

}
