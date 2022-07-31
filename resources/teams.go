package resources

import (
	"context"
	"encoding/json"
	"strconv"
	"strings"

	"github.com/cloudquery/cq-provider-github/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/google/go-github/v45/github"
)

func Teams() *schema.Table {
	return &schema.Table{
		Name:        "github_teams",
		Description: "Team represents a team within a GitHub organization",
		Resolver:    fetchTeams,
		Multiplex:   client.OrgMultiplex,
		Options:     schema.TableCreationOptions{PrimaryKeys: []string{"id"}},
		Columns: []schema.Column{
			{
				Name:        "org",
				Description: "The Github Organization of the resource.",
				Type:        schema.TypeString,
				Resolver:    client.ResolveOrg,
			},
			{
				Name:     "id",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("ID"),
			},
			{
				Name:     "node_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("NodeID"),
			},
			{
				Name: "name",
				Type: schema.TypeString,
			},
			{
				Name: "description",
				Type: schema.TypeString,
			},
			{
				Name:     "url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("URL"),
			},
			{
				Name: "slug",
				Type: schema.TypeString,
			},
			{
				Name:        "permission",
				Description: "Permission specifies the default permission for repositories owned by the team.",
				Type:        schema.TypeString,
			},
			{
				Name:        "permissions",
				Description: "Permissions identifies the permissions that a team has on a given repository",
				Type:        schema.TypeJSON,
			},
			{
				Name:        "privacy",
				Description: "Privacy identifies the level of privacy this team should have. Possible values are:     secret - only visible to organization owners and members of this team     closed - visible to all members of this organization Default is \"secret\".",
				Type:        schema.TypeString,
			},
			{
				Name: "members_count",
				Type: schema.TypeBigInt,
			},
			{
				Name: "repos_count",
				Type: schema.TypeBigInt,
			},
			{
				Name:     "html_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("HTMLURL"),
			},
			{
				Name:     "members_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("MembersURL"),
			},
			{
				Name:     "repositories_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("RepositoriesURL"),
			},
			{
				Name:     "parent",
				Type:     schema.TypeBigInt,
				Resolver: resolveTeamsParent,
			},
			{
				Name:        "ldapdn",
				Description: "LDAPDN is only available in GitHub Enterprise and when the team membership is synchronized with LDAP.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("LDAPDN"),
			},
		},
		Relations: []*schema.Table{
			{
				Name:        "github_team_members",
				Description: "User represents a GitHub user.",
				Resolver:    fetchTeamMembers,
				Options:     schema.TableCreationOptions{PrimaryKeys: []string{"team_id", "id"}},
				Columns: []schema.Column{
					{
						Name:        "team_cq_id",
						Description: "Unique CloudQuery ID of github_teams table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "team_id",
						Description: "The id of the name",
						Type:        schema.TypeBigInt,
						Resolver:    schema.ParentResourceFieldResolver("id"),
					},
					{
						Name:        "org",
						Description: "The Github Organization of the resource.",
						Type:        schema.TypeString,
						Resolver:    client.ResolveOrg,
					},
					{
						Name: "login",
						Type: schema.TypeString,
					},
					{
						Name:     "id",
						Type:     schema.TypeBigInt,
						Resolver: schema.PathResolver("ID"),
					},
					{
						Name:     "node_id",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("NodeID"),
					},
					{
						Name:     "avatar_url",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("AvatarURL"),
					},
					{
						Name:     "html_url",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("HTMLURL"),
					},
					{
						Name:     "gravatar_id",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("GravatarID"),
					},
					{
						Name: "name",
						Type: schema.TypeString,
					},
					{
						Name: "company",
						Type: schema.TypeString,
					},
					{
						Name: "blog",
						Type: schema.TypeString,
					},
					{
						Name: "location",
						Type: schema.TypeString,
					},
					{
						Name: "email",
						Type: schema.TypeString,
					},
					{
						Name: "hireable",
						Type: schema.TypeBool,
					},
					{
						Name: "bio",
						Type: schema.TypeString,
					},
					{
						Name: "twitter_username",
						Type: schema.TypeString,
					},
					{
						Name: "public_repos",
						Type: schema.TypeBigInt,
					},
					{
						Name: "public_gists",
						Type: schema.TypeBigInt,
					},
					{
						Name: "followers",
						Type: schema.TypeBigInt,
					},
					{
						Name: "following",
						Type: schema.TypeBigInt,
					},
					{
						Name:     "created_at_time",
						Type:     schema.TypeTimestamp,
						Resolver: schema.PathResolver("CreatedAt.Time"),
					},
					{
						Name:     "updated_at_time",
						Type:     schema.TypeTimestamp,
						Resolver: schema.PathResolver("UpdatedAt.Time"),
					},
					{
						Name:     "suspended_at_time",
						Type:     schema.TypeTimestamp,
						Resolver: schema.PathResolver("SuspendedAt.Time"),
					},
					{
						Name: "type",
						Type: schema.TypeString,
					},
					{
						Name: "site_admin",
						Type: schema.TypeBool,
					},
					{
						Name: "total_private_repos",
						Type: schema.TypeBigInt,
					},
					{
						Name: "owned_private_repos",
						Type: schema.TypeBigInt,
					},
					{
						Name: "private_gists",
						Type: schema.TypeBigInt,
					},
					{
						Name: "disk_usage",
						Type: schema.TypeBigInt,
					},
					{
						Name: "collaborators",
						Type: schema.TypeBigInt,
					},
					{
						Name: "two_factor_authentication",
						Type: schema.TypeBool,
					},
					{
						Name:     "plan_name",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("Plan.Name"),
					},
					{
						Name:     "plan_space",
						Type:     schema.TypeBigInt,
						Resolver: schema.PathResolver("Plan.Space"),
					},
					{
						Name:     "plan_collaborators",
						Type:     schema.TypeBigInt,
						Resolver: schema.PathResolver("Plan.Collaborators"),
					},
					{
						Name:     "plan_private_repos",
						Type:     schema.TypeBigInt,
						Resolver: schema.PathResolver("Plan.PrivateRepos"),
					},
					{
						Name:     "plan_filled_seats",
						Type:     schema.TypeBigInt,
						Resolver: schema.PathResolver("Plan.FilledSeats"),
					},
					{
						Name:     "plan_seats",
						Type:     schema.TypeBigInt,
						Resolver: schema.PathResolver("Plan.Seats"),
					},
					{
						Name: "ldap_dn",
						Type: schema.TypeString,
					},
					{
						Name:        "url",
						Description: "API URLs",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("URL"),
					},
					{
						Name:     "events_url",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("EventsURL"),
					},
					{
						Name:     "following_url",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("FollowingURL"),
					},
					{
						Name:     "followers_url",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("FollowersURL"),
					},
					{
						Name:     "gists_url",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("GistsURL"),
					},
					{
						Name:     "organizations_url",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("OrganizationsURL"),
					},
					{
						Name:     "received_events_url",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("ReceivedEventsURL"),
					},
					{
						Name:     "repos_url",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("ReposURL"),
					},
					{
						Name:     "starred_url",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("StarredURL"),
					},
					{
						Name:     "subscriptions_url",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("SubscriptionsURL"),
					},
					{
						Name:        "text_matches",
						Description: "TextMatches is only populated from search results that request text matches See: search.go and https://docs.github.com/en/rest/search/#text-match-metadata",
						Type:        schema.TypeJSON,
						Resolver:    resolveTeamMembersTextMatches,
					},
					{
						Name:        "permissions",
						Description: "Permissions and RoleName identify the permissions and role that a user has on a given repository",
						Type:        schema.TypeJSON,
					},
					{
						Name: "role_name",
						Type: schema.TypeString,
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchTeams(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	opts := &github.ListOptions{
		Page:    0,
		PerPage: 100,
	}
	for {
		repos, resp, err := c.Github.Teams.ListTeams(ctx, c.Org, opts)
		if err != nil {
			return err
		}
		res <- repos
		opts.Page = resp.NextPage
		if opts.Page == resp.LastPage {
			break
		}
	}
	return nil
}
func resolveTeamsParent(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	u := resource.Item.(*github.Team)
	if u.Parent == nil {
		return nil
	}
	return resource.Set(c.Name, u.Parent.ID)
}
func fetchTeamMembers(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	t := parent.Item.(*github.Team)
	c := meta.(*client.Client)
	opts := &github.TeamListTeamMembersOptions{
		ListOptions: github.ListOptions{
			Page:    0,
			PerPage: 100,
		},
	}
	orgId, err := strconv.Atoi(strings.Split(*t.MembersURL, "/")[4])
	if err != nil {
		return err
	}
	for {
		repos, resp, err := c.Github.Teams.ListTeamMembersByID(ctx, int64(orgId), *t.ID, opts)
		if err != nil {
			return err
		}
		res <- repos
		opts.Page = resp.NextPage
		if opts.Page == resp.LastPage {
			break
		}
	}
	return nil
}
func resolveTeamMembersTextMatches(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	u := resource.Item.(*github.User)
	j, err := json.Marshal(u.TextMatches)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, j)
}
