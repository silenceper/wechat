package credential

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/h2non/gock.v1"
)

// TestGetTicketFromServerContext 测试 GetTicketFromServerContext 函数
func TestGetTicketFromServerContext(t *testing.T) {
	defer gock.Off()
	gock.New(fmt.Sprintf(getTicketURL, "arg-ak")).Reply(200).JSON(&ResTicket{Ticket: "mock-ticket", ExpiresIn: 10})

	ticket, err := GetTicketFromServerContext(context.Background(), "arg-ak")
	assert.Nil(t, err)
	assert.Equal(t, int64(0), ticket.ErrCode)
	assert.Equal(t, "mock-ticket", ticket.Ticket, "they should be equal")
	assert.Equal(t, int64(10), ticket.ExpiresIn, "they should be equal")
}
