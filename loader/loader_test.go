package loader_test

import (
	"testing"

	"github.com/vlamug/scfg/model"

	. "gopkg.in/check.v1"
	"github.com/vlamug/scfg/loader"
	"github.com/vlamug/scfg/request"
)

func TestLoaer(t *testing.T) { TestingT(t) }

type TestLoaderSuite struct{}

var (
	_ = Suite(&TestLoaderSuite{})

	storageUsed = false
	cacheUsed   = false
)

func (s *TestLoaderSuite) TestLoad(c *C) {
	st := &storageMock{}
	ch := &cacheMock{}

	storageUsed = false
	cacheUsed = false

	ld := loader.NewLoader(st, ch)
	value := ld.Load(request.GetRequest{Type: "type", Data: "data"})
	c.Assert(value, Equals, "test_set")
	c.Assert(storageUsed, Equals, true)
	c.Assert(cacheUsed, Equals, true)

	storageUsed = false
	cacheUsed = false

	value = ld.Load(request.GetRequest{Type: "type", Data: "data"})
	c.Assert(value, Equals, "test_set")
	c.Assert(storageUsed, Equals, false)
	c.Assert(cacheUsed, Equals, true)
}

type storageMock struct{}

func (st *storageMock) Get(ckey string) (model.Config) {
	storageUsed = true
	return model.Config{
		CKey: "type_data",
		CSet: "test_set",
	}
}

type cacheMock struct {
	value string
}

func (ch *cacheMock) Get(key string) string {
	cacheUsed = true
	return ch.value
}

func (ch *cacheMock) Set(key, value string) {
	ch.value = value
}
