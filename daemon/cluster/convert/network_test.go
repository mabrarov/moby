package convert

import (
	"testing"
	"time"

	gogotypes "github.com/gogo/protobuf/types"
	swarmapi "github.com/moby/swarmkit/v2/api"
)

func TestNetworkConvertBasicNetworkFromGRPCCreatedAt(t *testing.T) {
	expected, err := time.Parse("Jan 2, 2006 at 3:04pm (MST)", "Jan 10, 2018 at 7:54pm (PST)")
	if err != nil {
		t.Fatal(err)
	}
	createdAt, err := gogotypes.TimestampProto(expected)
	if err != nil {
		t.Fatal(err)
	}

	nw := swarmapi.Network{
		Meta: swarmapi.Meta{
			Version: swarmapi.Version{
				Index: 1,
			},
			CreatedAt: createdAt,
		},
	}

	n := BasicNetworkFromGRPC(nw)
	if !n.Created.Equal(expected) {
		t.Fatalf("expected time %s; received %s", expected, n.Created)
	}
}
