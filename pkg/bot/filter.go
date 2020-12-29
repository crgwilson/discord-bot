package bot

import (
	"errors"
)

var ErrUserAlreadyIgnored = errors.New("Provided user ID is already present in the ignore list")
var ErrChannelAlreadyIgnored = errors.New("Provided channel ID is already present in the ignore list")

type MessageFilter struct {
	IgnoredUserList    []string
	IgnoredChannelList []string
}

func (f *MessageFilter) UserIsIgnored(userId string) bool {
	for _, u := range f.IgnoredUserList {
		if u == userId {
			return true
		}
	}
	return false
}

func (f *MessageFilter) IgnoreUser(userId string) error {
	if f.UserIsIgnored(userId) {
		return ErrUserAlreadyIgnored
	}

	f.IgnoredUserList = append(f.IgnoredUserList, userId)
	return nil
}

func (f *MessageFilter) ChannelIsIgnored(channelId string) bool {
	for _, c := range f.IgnoredChannelList {
		if c == channelId {
			return true
		}
	}
	return false
}

func (f *MessageFilter) IgnoreChannel(channelId string) error {
	if f.ChannelIsIgnored(channelId) {
		return ErrChannelAlreadyIgnored
	}

	f.IgnoredChannelList = append(f.IgnoredChannelList, channelId)
	return nil
}

func (f *MessageFilter) Filter(authorId, channelId, expectedPrefix, message string) bool {
	if f.UserIsIgnored(authorId) {
		return false
	}

	if f.ChannelIsIgnored(channelId) {
		return false
	}

	prefixLength := len(expectedPrefix)
	if message[:prefixLength] != expectedPrefix {
		return false
	}

	return true
}

func NewMessageFilter(usersToIgnore, channelsToIgnore []string) *MessageFilter {
	filter := MessageFilter{
		IgnoredUserList:    usersToIgnore,
		IgnoredChannelList: channelsToIgnore,
	}
	return &filter
}
