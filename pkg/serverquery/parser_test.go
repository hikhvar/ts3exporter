package serverquery

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const channellist = `cid=1 pid=0 channel_name=Default\sChannel|cid=8 pid=0 channel_name=Weiß\snicht\swie\sman\sTS\sbeendet\s(AFK)`
const equalInValue = `virtualserver_unique_identifier=asdasdlASlkasd\/asdfgASdax=`

func TestParse(t *testing.T) {
	type args struct {
		in string
	}
	tests := []struct {
		name    string
		args    args
		want    Result
		wantErr bool
	}{
		{
			name: "channellist",
			args: args{
				in: channellist,
			},
			want: Result{
				Raw: channellist,
				Items: []Parsed{
					{
						"cid":          "1",
						"pid":          "0",
						"channel_name": "Default Channel",
					},
					{
						"cid":          "8",
						"pid":          "0",
						"channel_name": "Weiß nicht wie man TS beendet (AFK)",
					},
				},
			},
		},
		{
			name: "= in value",
			args: args{
				in: equalInValue,
			},
			want: Result{
				Raw: equalInValue,
				Items: []Parsed{
					{
						"virtualserver_unique_identifier": "asdasdlASlkasd/asdfgASdax=",
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Parse(tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestParsed_ReadInto(t *testing.T) {
	type testValue struct {
		ChannelID   int    `sq:"cid"`
		ChannelName string `sq:"channel_name"`
	}
	type args struct {
		v interface{}
	}
	tests := []struct {
		name      string
		p         Parsed
		args      args
		wantErr   bool
		wantAfter interface{}
	}{
		{
			name:      "normal decoding",
			p:         Parsed{"cid": "1", "channel_name": "foobar"},
			args:      args{v: &testValue{}},
			wantErr:   false,
			wantAfter: &testValue{ChannelID: 1, ChannelName: "foobar"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.p.ReadInto(tt.args.v); (err != nil) != tt.wantErr {
				t.Errorf("ReadInto() error = %v, wantErr %v", err, tt.wantErr)
			}
			assert.Equal(t, tt.wantAfter, tt.args.v)
		})
	}
}
