package resources

import (
	"testing"

	"github.com/cloudquery/cq-provider-github/client"
	"github.com/cloudquery/cq-provider-github/client/mocks"
	"github.com/google/go-github/v45/github"

	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
)

func buildOrganizations(t *testing.T, ctrl *gomock.Controller) client.GithubServices {
	mock := mocks.NewMockOrganizationsService(ctrl)

	var cs *github.Organization
	if err := faker.FakeData(&cs); err != nil {
		t.Fatal(err)
	}
	mock.EXPECT().Get(gomock.Any(), "testorg").Return(cs, &github.Response{}, nil)
	return client.GithubServices{Organizations: mock}
}

func TestOrganizations(t *testing.T) {
	client.GithubMockTestHelper(t, Organizations(), buildOrganizations, client.TestOptions{})
}
