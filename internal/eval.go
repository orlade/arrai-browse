package internal

import (
	"context"
	"io/ioutil"
	"path/filepath"

	"github.com/arr-ai/arrai/pkg/arraictx"
	"github.com/arr-ai/arrai/rel"
	"github.com/arr-ai/arrai/syntax"
	"github.com/sirupsen/logrus"
)

func Eval(path string, args ...string) (string, error) {
	val, err := eval(path, args...)
	if err != nil {
		return "", err
	}
	return val.String(), nil
}

// eval loads an arr.ai script and executes it, returning its output Value.
func eval(path string, args ...string) (rel.Value, error) {
	path, err := filepath.Abs(path)
	if err != nil {
		return nil, err
	}
	bs, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	logrus.Debugf("read %s from disk", path)
	script := string(bs)

	ctx := arraictx.InitRunCtx(context.Background())
	ctx = arraictx.WithArgs(ctx, args...)
	out, err := syntax.EvaluateExpr(ctx, path, script)
	if err != nil || out == nil {
		return nil, err
	}
	return out, nil
}
