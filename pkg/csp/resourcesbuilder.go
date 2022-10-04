package csp

import (
	"fmt"
	core "vmc-go-client/pkg/core"
	"vmc-go-client/pkg/utils"
)

const name = "CSP"
const scheme = "https"
const url = "console.cloud.vmware.com"

var resources = map[string]core.Endpoint{
	"PublicKey":                    {"/csp/gateway/am/api/auth/token-public-key", "GET", initType(new(PublicKeyDto)), queryparams["PublicKey"]},
	"CspOpenIdConfiguration":       {"/csp/gateway/am/api/auth/.well-known/openid-configuration", "GET", initType(new(CspOpenIdConfigurationDto)), queryparams["CspOpenIdConfiguration"]},
	"OpenIdUserInfo":               {"/csp/gateway/am/api/userinfo", "GET", initType(new(OidcUserInfoDto)), queryparams["OpenIdUserInfo"]},
	"LoggedInUser":                 {"/csp/gateway/am/api/loggedin/user", "GET", initType(new(UserDto)), queryparams["LoggedInUser"]},
	"LoggedInUserInfo":             {"/csp/gateway/am/api/loggedin/user/orgs/{orgId}/info", "GET", initType(new(UserInfoDto)), queryparams["LoggedInUserInfo"]},
	"LoggedInUserGroups":           {"/csp/gateway/am/api/loggedin/user/orgs/{orgId}/groups", "GET", initType(new(UserGroupsResponse)), queryparams["LoggedInUserGroups"]},
	"LoggedInUserRoles":            {"/csp/gateway/am/api/loggedin/user/orgs/{orgId}/roles", "GET", initType(new(RolesDto)), queryparams["LoggedInUserRoles"]},
	"LoggedInUserServiceRoles":     {"/csp/gateway/am/api/loggedin/user/orgs/{orgId}/service-roles", "GET", initType(new(UserServiceRolesResponse)), queryparams["LoggedInUserServiceRoles"]},
	"LoggedInUserOrgs":             {"/csp/gateway/am/api/v2/loggedin/user/orgs", "GET", initType(new(PagedResponseOrganizationDto)), queryparams["LoggedInUserOrgs"]},
	"LoggedInMfaStatus":            {"/csp/gateway/iam/vmwid/api/principal/mfa/status", "GET", initType(new(UserMfaStatusDto)), queryparams["LoggedInMfaStatus"]},
	"User":                         {"/csp/gateway/am/api/v2/users/{userId}/orgs/{orgId}/info", "GET", initType(new(UserInfoDto)), queryparams["User"]},
	"UserCustomRoles":              {"/csp/gateway/am/api/users/{userId}/orgs/{orgId}/custom-roles", "GET", initType(new(UserCustomRolesResponse)), queryparams["UserCustomRoles"]},
	"UserGroups":                   {"/csp/gateway/am/api/v2/users/{userId}/orgs/{orgId}/groups", "GET", initType(new(UserGroupsResponse)), queryparams["UserGroups"]},
	"UserGroupsByEmail":            {"/csp/gateway/am/api/users/{acct}/orgs/{orgId}/groups", "GET", initType(new(UserGroupsResponse)), queryparams["UserGroupsByEmail"]},
	"UserRoles":                    {"/csp/gateway/am/api/v2/users/{userId}/orgs/{orgId}/roles", "GET", initType(new(RolesDto)), queryparams["UserRoles"]},
	"UserServiceRoles":             {"/csp/gateway/am/api/v2/users/{userId}/orgs/{orgId}/service-roles", "GET", initType(new(UserServiceRolesResponse)), queryparams["UserServiceRoles"]},
	"OrgOAuthApps":                 {"/csp/gateway/am/api/orgs/{orgId}/oauth-apps", "GET", initType(new(PagedResponseOrgOAuthAppResponse)), queryparams["OrgOAuthApps"]},
	"OrgOAuthAppById":              {"/csp/gateway/am/api/orgs/{orgId}/oauth-apps/{oauthAppId}", "GET", initType(new(OrgOAuthAppResponse)), queryparams["OrgOAuthAppById"]},
	"OrgSearchGroups":              {"/csp/gateway/am/api/orgs/{orgId}/groups-search", "GET", initType(new(SearchGroupsResponse)), queryparams["OrgSearchGroups"]},
	"OrgLinkedIdp":                 {"/csp/gateway/am/api/orgs/{orgId}/idp", "GET", initType(new(OrgLinkedIdpResponse)), queryparams["OrgLinkedIdp"]},
	"Org":                          {"/csp/gateway/am/api/orgs/{orgId}", "GET", initType(new(OrganizationDto)), queryparams["Org"]},
	"OrgClients":                   {"/csp/gateway/am/api/orgs/{orgId}/clients", "GET", initType(new(PagedResponseExpandedAuthClientDto)), queryparams["OrgClients"]},
	"OrgInvitations":               {"/csp/gateway/am/api/orgs/{orgId}/invitations", "GET", initType(new(UserOrganizationInvitationResponse)), queryparams["OrgInvitations"]},
	"OrgRoles":                     {"/csp/gateway/am/api/orgs/{orgId}/roles", "GET", initType(new(OrganizationRolesResponse)), queryparams["OrgRoles"]},
	"OrgTrusts":                    {"/csp/gateway/am/api/orgs/{orgId}/trusts", "GET", initType(new(PagedResponseOrganizationTrustDto)), queryparams["OrgTrusts"]},
	"OrgUsers":                     {"/csp/gateway/am/api/v2/orgs/{orgId}/users", "GET", initType(new(PagedResponseExpandedTypedUserDto)), queryparams["OrgUsers"]},
	"OrgSubOrgs":                   {"/csp/gateway/am/api/orgs/{orgId}/sub-orgs ", "GET", initType(new(PagedResponseOrganizationDto)), queryparams["OrgSubOrgs"]},
	"OrgMemberOrgs":                {"/csp/gateway/am/api/orgs/{orgId}/member-orgs", "GET", initType(new(PagedResponseOrganizationMemberForTrustDto)), queryparams["OrgMemberOrgs"]},
	"OrgInvitationById":            {"/csp/gateway/am/api/orgs/{orgId}/invitations/{userInvitationId}", "GET", initType(new(UserOrganizationInvitationResponse)), queryparams["OrgInvitationById"]},
	"OrgGroupById":                 {"/csp/gateway/am/api/orgs/{orgId}/groups/{groupId}", "GET", initType(new(ExpandedGroupDto)), queryparams["OrgGroupById"]},
	"OrgGroupRolesById":            {"/csp/gateway/am/api/orgs/{orgId}/groups/{groupId}/roles", "GET", initType(new(RolesDto)), queryparams["OrgGroupRolesById"]},
	"OrgNestedEnterpriseGroups":    {"/csp/gateway/am/api/orgs/{orgId}/groups/{groupId}/groups", "GET", initType(new(PagedResponseGroupDto)), queryparams["OrgNestedEnterpriseGroups"]},
	"OrgUsersInGroup":              {"/csp/gateway/am/api/orgs/{orgId}/groups/{groupId}/users", "GET", initType(new(PagedResponseBaseUserWithProfileDto)), queryparams["OrgUsersInGroup"]},
	"OrgGroups":                    {"/csp/gateway/am/api/orgs/{orgId}/groups", "GET", initType(new(PagedResponseExpandedGroupDto)), queryparams["OrgGroups"]},
	"AutoEntitlementPolicies":      {"/csp/gateway/am/api/orgs/{orgId}/auto-entitlement-policies", "GET", initType(new(PagedResponseAutoEntitlementPoliciesDto)), queryparams["AutoEntitlementPolicies"]},
	"AutoEntitlementPoliciesByOrg": {"/csp/gateway/am/api/orgs/{orgId}/auto-entitlement-policies/{policyId}", "GET", initType(new(ExpandedAutoEntitlementPoliciesDto)), queryparams["AutoEntitlementPoliciesByOrg"]},
	"CustomRole":                   {"/csp/gateway/iam-roles-mgmt/api/orgs/{orgId}/custom-roles/{roleName}", "GET", initType(new(CustomRoleDto)), queryparams["CustomRole"]},
	"CustomRoles":                  {"/csp/gateway/iam-roles-mgmt/api/orgs/{orgId}/custom-roles", "GET", initType(new(ResultsDtoCustomRoleDto)), queryparams["CustomRoles"]},
}

//params to set variables and construct API resources
var params = map[string]string{
	"{orgId}":            "",
	"{userId}":           "",
	"{acct}":             "",
	"{oauthAppId}":       "",
	"{policyId}":         "",
	"{roleName}":         "",
	"{userInvitationId}": "",
	"{groupId}":          "",
}

//queryparams for specific resource types
var queryparams = map[string]map[string]string{
	"LoggedInUserOrgs":          {"pageStart": "1", "pageLimit": "15"},
	"OrgOAuthApps":              {"pageStart": "", "pageLimit": ""},
	"OrgClients":                {"pageStart": "", "pageLimit": ""},
	"OrgTrusts":                 {"pageStart": "", "pageLimit": ""},
	"OrgUsers":                  {"pageStart": "", "pageLimit": ""},
	"OrgMemberOrgs":             {"pageStart": "", "pageLimit": ""},
	"OrgSubOrgs":                {"pageStart": "", "pageLimit": ""},
	"OrgNestedEnterpriseGroups": {"pageStart": "", "pageLimit": ""},
	"OrgUsersInGroup":           {"pageStart": "", "pageLimit": ""},
	"OrgGroups":                 {"pageStart": "", "pageLimit": ""},
	"AutoEntitlementPolicies":   {"pageStart": "", "pageLimit": ""},
	"OrgSearchGroups":           {"groupSearchTerm": ""},
}

func BuildConfig() {
	config := core.ResourceFactory{
		ApiName:   name,
		ApiScheme: scheme,
		ApiUrl:    url,
		Resources: resources,
		Params:    params,
	}
	fmt.Println("#### Auto-Generated Config file for ", name, " API ####")
	fmt.Println(utils.Marshal(config))
}
