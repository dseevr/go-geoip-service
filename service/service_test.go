package service

import (
	. "testing"

	. "gopkg.in/check.v1"
)

type serviceSuite struct {
}

var _ = Suite(&serviceSuite{})

func TestService(t *T) { TestingT(t) }

func setup() {
	LoadMaxmindDB("../test.mmdb")
}

func (s *serviceSuite) TestLoadingMultipleTimes(c *C) {
	for i := 0; i < 10; i++ {
		setup()
	}
}

func (s *serviceSuite) TestLoadUnloadLoad(c *C) {
	setup()

	for i := 0; i < 10; i++ {
		UnloadMaxmindDB()
		setup()
	}
}

func (s *serviceSuite) TestDbNotLoaded(c *C) {
	UnloadMaxmindDB()

	_, err := LookupIP("1.2.3.4")
	c.Assert(err, NotNil)
	c.Assert(err, ErrorMatches, ".* not loaded")
}

func (s serviceSuite) TestDbLoaded(c *C) {
	setup()

	_, err := LookupIP("1.2.3.4")
	c.Assert(err, IsNil)
}

func (s *serviceSuite) TestBadInputs(c *C) {

	setup()

	var err error

	// no ip
	_, err = LookupIP("")
	c.Assert(err, NotNil)

	// invalid IPs
	_, err = LookupIP("asdf")
	c.Assert(err, NotNil)

	_, err = LookupIP("999.999.999.999")
	c.Assert(err, NotNil)
}

func (s *serviceSuite) TestGoodInputs(c *C) {

	setup()

	var err error

	// valid IP
	_, err = LookupIP("127.0.0.1")
	c.Assert(err, IsNil)

	// test country lookups
	var resp *Response

	resp, err = LookupIP("1.2.3.4")
	c.Assert(err, IsNil)
	c.Assert(resp.Country, Equals, "AU")

	resp, err = LookupIP("2.3.4.5")
	c.Assert(err, IsNil)
	c.Assert(resp.Country, Equals, "FR")

	resp, err = LookupIP("3.4.5.6")
	c.Assert(err, IsNil)
	c.Assert(resp.Country, Equals, "US")
}
