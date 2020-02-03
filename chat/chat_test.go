package chat_test

import (
	"github.com/fx2y/hands-on-software-engineering-with-golang/chat"
	"reflect"
	"testing"
)

func TestChatRoomBroadcast(t *testing.T) {
	pub := new(spyPublisher)
	room := chat.NewRoom(pub)
	room.AddUser("bob")
	room.AddUser("alice")
	_ = room.Broadcast("hi")

	// Check published entries
	exp := []entry{
		{user: "bob", message: "hi"},
		{user: "alice", message: "hi"},
	}

	if got := pub.published; !reflect.DeepEqual(got, exp) {
		t.Fatalf("expected the following message:\n%#+v\ngot:\n%#+v", exp, got)
	}
}

type entry struct {
	user    string
	message string
}

type spyPublisher struct {
	published []entry
}

func (p *spyPublisher) Publish(user, message string) error {
	p.published = append(p.published, entry{user: user, message: message})
	return nil
}
