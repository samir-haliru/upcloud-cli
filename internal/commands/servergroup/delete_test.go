package servergroup

import (
	"testing"

	"github.com/UpCloudLtd/upcloud-cli/v2/internal/commands"
	"github.com/UpCloudLtd/upcloud-cli/v2/internal/config"
	smock "github.com/UpCloudLtd/upcloud-cli/v2/internal/mock"

	"github.com/UpCloudLtd/upcloud-go-api/v6/upcloud"
	"github.com/UpCloudLtd/upcloud-go-api/v6/upcloud/request"
	"github.com/gemalto/flume"
	"github.com/stretchr/testify/assert"
)

func TestDeleteCommand(t *testing.T) {
	targetMethod := "DeleteServerGroup"

	serverGroup := upcloud.ServerGroup{
		Title: "test-server-group",
		UUID:  "17fbd082-30b0-11eb-adc1-0242ac120003",
	}

	for _, test := range []struct {
		name  string
		arg   string
		error string
		req   request.DeleteServerGroupRequest
	}{
		{
			name: "delete with UUID",
			arg:  serverGroup.UUID,
			req:  request.DeleteServerGroupRequest{UUID: serverGroup.UUID},
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			mService := smock.Service{}
			mService.On(targetMethod, &test.req).Return(nil)

			conf := config.New()
			c := commands.BuildCommand(DeleteCommand(), nil, conf)

			_, err := c.(commands.MultipleArgumentCommand).Execute(commands.NewExecutor(conf, &mService, flume.New("test")), test.arg)

			if test.error != "" {
				assert.EqualError(t, err, test.error)
			} else {
				assert.NoError(t, err)
				mService.AssertNumberOfCalls(t, targetMethod, 1)
			}
		})
	}
}
