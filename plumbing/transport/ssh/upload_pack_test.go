package ssh

import (
	"os"

	"github.com/karfield/go-git/plumbing/transport"
	"github.com/karfield/go-git/plumbing/transport/test"

	. "github.com/karfield/go-check"
)

type UploadPackSuite struct {
	test.UploadPackSuite
}

var _ = Suite(&UploadPackSuite{})

func (s *UploadPackSuite) SetUpSuite(c *C) {
	if os.Getenv("SSH_AUTH_SOCK") == "" {
		c.Skip("SSH_AUTH_SOCK is not set")
	}

	s.UploadPackSuite.Client = DefaultClient

	ep, err := transport.NewEndpoint("git@github.com:git-fixtures/basic.git")
	c.Assert(err, IsNil)
	s.UploadPackSuite.Endpoint = ep

	ep, err = transport.NewEndpoint("git@github.com:git-fixtures/empty.git")
	c.Assert(err, IsNil)
	s.UploadPackSuite.EmptyEndpoint = ep

	ep, err = transport.NewEndpoint("git@github.com:git-fixtures/non-existent.git")
	c.Assert(err, IsNil)
	s.UploadPackSuite.NonExistentEndpoint = ep

}
