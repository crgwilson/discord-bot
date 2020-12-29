package bot

import (
	"testing"
)

const (
	filterPrefix   string = "!"
	blockedUser    string = "user1"
	allowedUser    string = "user3"
	blockedChannel string = "channel2"
	allowedChannel string = "channel1"
)

func TestFilter(t *testing.T) {
	usersToIgnore := []string{blockedUser}
	channelsToIgnore := []string{blockedChannel}
	filter := NewMessageFilter(usersToIgnore, channelsToIgnore)

	cases := []struct {
		Name           string
		UserId         string
		ChannelId      string
		Message        string
		ExpectedResult bool
	}{
		{
			"no filter",
			allowedUser,
			allowedChannel,
			"!this should not be filtered",
			true,
		},
		{
			"filter user",
			blockedUser,
			allowedChannel,
			"!this should be filtered",
			false,
		},
		{
			"filter channel",
			allowedUser,
			blockedChannel,
			"!this should be filtered",
			false,
		},
		{
			"filter user and channel",
			blockedUser,
			blockedChannel,
			"!this should be filtered",
			false,
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			filterResult := filter.Filter(test.UserId, test.ChannelId, testPrefix, test.Message)
			if filterResult != test.ExpectedResult {
				t.Errorf("got %v, want %v", filterResult, test.ExpectedResult)
			}
		})
	}
}
