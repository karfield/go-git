package server_test

import (
	"github.com/karfield/go-git/plumbing/transport"
	"github.com/karfield/go-git/plumbing/transport/test"

	. "github.com/karfield/go-check"
)

type UploadPackSuite struct {
	BaseSuite
	test.UploadPackSuite
}

var _ = Suite(&UploadPackSuite{})

func (s *UploadPackSuite) SetUpSuite(c *C) {
	s.BaseSuite.SetUpSuite(c)
	s.UploadPackSuite.Client = s.client
}

func (s *UploadPackSuite) SetUpTest(c *C) {
	s.prepareRepositories(c, &s.Endpoint, &s.EmptyEndpoint, &s.NonExistentEndpoint)
}

// Overwritten, it's not an error in server-side.
func (s *UploadPackSuite) TestAdvertisedReferencesEmpty(c *C) {
	r, err := s.Client.NewUploadPackSession(s.EmptyEndpoint, s.EmptyAuth)
	c.Assert(err, IsNil)
	ar, err := r.AdvertisedReferences()
	c.Assert(err, IsNil)
	c.Assert(len(ar.References), Equals, 0)
}

// Overwritten, server returns error earlier.
func (s *UploadPackSuite) TestAdvertisedReferencesNotExists(c *C) {
	r, err := s.Client.NewUploadPackSession(s.NonExistentEndpoint, s.EmptyAuth)
	c.Assert(err, Equals, transport.ErrRepositoryNotFound)
	c.Assert(r, IsNil)
}
