package request

import (
	"github.com/LinuxSploit/SendMe/internal/resource"
	"github.com/LinuxSploit/SendMe/internal/user"
)

type Request struct {
	*user.User
	*resource.Resource
	RequestStatus bool
}

func NewRequest(user *user.User, resource *resource.Resource) *Request {

	return &Request{
		User:          user,
		Resource:      resource,
		RequestStatus: false,
	}

}
