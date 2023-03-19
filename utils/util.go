package utils

import (
	"context"
	"github.com/gin-gonic/gin"
	"time"
)

type ValueOnlyContext struct{ context.Context }

func (ValueOnlyContext) Deadline() (deadline time.Time, ok bool) { return }
func (ValueOnlyContext) Done() <-chan struct{}                   { return nil }
func (ValueOnlyContext) Err() error                              { return nil }
func GetValueOnlyRequestContext(c *gin.Context) ValueOnlyContext {
	return ValueOnlyContext{Context: c.Request.Context()}
}
