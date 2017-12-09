package cache_test

import (
	"testing"

	. "gopkg.in/check.v1"
	"github.com/vlamug/scfg/cache"
)

func TestMem(t *testing.T) { TestingT(t) }

type TestMemSuite struct{}

var (
	_ = Suite(&TestMemSuite{})
)

func (s *TestMemSuite) TestLoad(c *C) {
	ch := cache.NewMem()

	ch.Set("test_key", "test_value")
	value := ch.Get("test_key")

	c.Assert(value, Equals, "test_value")

	ch.Set("second_test_key", "second_test_value")
	value = ch.Get("second_test_key")

	c.Assert(value, Equals, "second_test_value")
}
