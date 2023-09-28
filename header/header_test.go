package header

import (
	"context"
	"testing"
	"time"

	"github.com/micro/go-micro/v2/metadata"
)

func TestVersionToInt(t *testing.T) {
	t.Logf("%d", VersionToInt("1.5.4"))
}

func TestGetClientIP(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.TODO(), time.Second)
	defer cancel()

	md := metadata.Metadata{}
	md[ForwardIP] = "127.0.0.1,127.0.2.1"
	cx := metadata.NewContext(ctx, md)

	t.Log(GetClientIP(cx))
}
