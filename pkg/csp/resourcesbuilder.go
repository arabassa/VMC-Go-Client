package csp

import (
	"fmt"
	core "vmc-go-client/pkg/core"
	"vmc-go-client/pkg/utils"
)

const name = "CSP"
const scheme = "https"
const url = "console.cloud.vmware.com"

var readMethods = []string{"GET"}

var resources = map[string]core.Endpoint{
	"PublicKey":                    {"/csp/gateway/am/api/auth/token-public-key", readMethods, initType(new(PublicKeyDto)), queryparams["PublicKey"]},
	"CspOpenIdConfiguration":       {"/csp/gateway/am/api/auth/.well-known/openid-configuration", readMethods, initType(new(CspOpenIdConfigurationDto)), queryparams["CspOpenIdConfiguration"]},
	"OpenIdUserInfo":               {"/csp/gateway/am/api/userinfo", readMethods, initType(new(OidcUserInfoDto)), queryparams["OpenIdUserInfo"]},
	"LoggedInUser":                 {"/csp/gateway/am/api/loggedin/user", readMethods, initType(new(UserDto)), queryparams["LoggedInUser"]},
	"LoggedInUserInfo":             {"/csp/gateway/am/api/loggedin/user/orgs/{orgId}/info", readMethods, initType(new(UserInfoDto)), queryparams["LoggedInUserInfo"]},
	"LoggedInUserGroups":           {"/csp/gateway/am/api/loggedin/user/orgs/{orgId}/groups", readMethods, initType(new(UserGroupsResponse)), queryparams["LoggedInUserGroups"]},
	"LoggedInUserRoles":            {"/csp/gateway/am/api/loggedin/user/orgs/{orgId}/roles", readMethods, initType(new(RolesDto)), queryparams["LoggedInUserRoles"]},
	"LoggedInUserServiceRoles":     {"/csp/gateway/am/api/loggedin/user/orgs/{orgId}/service-roles", readMethods, initType(new(UserServiceRolesResponse)), queryparams["LoggedInUserServiceRoles"]},
	"LoggedInUserOrgs":             {"/csp/gateway/am/api/v2/loggedin/user/orgs", readMethods, initType(new(PagedResponseOrganizationDto)), queryparams["LoggedInUserOrgs"]},
	"LoggedInMfaStatus":            {"/csp/gateway/iam/vmwid/api/principal/mfa/status", readMethods, initType(new(UserMfaStatusDto)), queryparams["LoggedInMfaStatus"]},
	"User":                         {"/csp/gateway/am/api/v2/users/{userId}/orgs/{orgId}/info", readMethods, initType(new(UserInfoDto)), queryparams["User"]},
	"UserCustomRoles":              {"/csp/gateway/am/api/users/{userId}/orgs/{orgId}/custom-roles", readMethods, initType(new(UserCustomRolesResponse)), queryparams["UserCustomRoles"]},
	"UserGroups":                   {"/csp/gateway/am/api/v2/users/{userId}/orgs/{orgId}/groups", readMethods, initType(new(UserGroupsResponse)), queryparams["UserGroups"]},
	"UserGroupsByEmail":            {"/csp/gateway/am/api/users/{acct}/orgs/{orgId}/groups", readMethods, initType(new(UserGroupsResponse)), queryparams["UserGroupsByEmail"]},
	"UserRoles":                    {"/csp/gateway/am/api/v2/users/{userId}/orgs/{orgId}/roles", readMethods, initType(new(RolesDto)), queryparams["UserRoles"]},
	"UserServiceRoles":             {"/csp/gateway/am/api/v2/users/{userId}/orgs/{orgId}/service-roles", readMethods, initType(new(UserServiceRolesResponse)), queryparams["UserServiceRoles"]},
	"OrgOAuthApps":                 {"/csp/gateway/am/api/orgs/{orgId}/oauth-apps", readMethods, initType(new(PagedResponseOrgOAuthAppResponse)), queryparams["OrgOAuthApps"]},
	"OrgOAuthAppById":              {"/csp/gateway/am/api/orgs/{orgId}/oauth-apps/{oauthAppId}", readMethods, initType(new(OrgOAuthAppResponse)), queryparams["OrgOAuthAppById"]},
	"OrgSearchGroups":              {"/csp/gateway/am/api/orgs/{orgId}/groups-search", readMethods, initType(new(SearchGroupsResponse)), queryparams["OrgSearchGroups"]},
	"OrgLinkedIdp":                 {"/csp/gateway/am/api/orgs/{orgId}/idp", readMethods, initType(new(OrgLinkedIdpResponse)), queryparams["OrgLinkedIdp"]},
	"Org":                          {"/csp/gateway/am/api/orgs/{orgId}", readMethods, initType(new(OrganizationDto)), queryparams["Org"]},
	"OrgClients":                   {"/csp/gateway/am/api/orgs/{orgId}/clients", readMethods, initType(new(PagedResponseExpandedAuthClientDto)), queryparams["OrgClients"]},
	"OrgInvitations":               {"/csp/gateway/am/api/orgs/{orgId}/invitations", readMethods, initType(new(UserOrganizationInvitationResponse)), queryparams["OrgInvitations"]},
	"OrgRoles":                     {"/csp/gateway/am/api/orgs/{orgId}/roles", readMethods, initType(new(OrganizationRolesResponse)), queryparams["OrgRoles"]},
	"OrgTrusts":                    {"/csp/gateway/am/api/orgs/{orgId}/trusts", readMethods, initType(new(PagedResponseOrganizationTrustDto)), queryparams["OrgTrusts"]},
	"OrgUsers":                     {"/csp/gateway/am/api/v2/orgs/{orgId}/users", readMethods, initType(new(PagedResponseExpandedTypedUserDto)), queryparams["OrgUsers"]},
	"OrgSubOrgs":                   {"/csp/gateway/am/api/orgs/{orgId}/sub-orgs ", readMethods, initType(new(PagedResponseOrganizationDto)), queryparams["OrgSubOrgs"]},
	"OrgMemberOrgs":                {"/csp/gateway/am/api/orgs/{orgId}/member-orgs", readMethods, initType(new(PagedResponseOrganizationMemberForTrustDto)), queryparams["OrgMemberOrgs"]},
	"OrgInvitationById":            {"/csp/gateway/am/api/orgs/{orgId}/invitations/{userInvitationId}", readMethods, initType(new(UserOrganizationInvitationResponse)), queryparams["OrgInvitationById"]},
	"OrgGroupById":                 {"/csp/gateway/am/api/orgs/{orgId}/groups/{groupId}", readMethods, initType(new(ExpandedGroupDto)), queryparams["OrgGroupById"]},
	"OrgGroupRolesById":            {"/csp/gateway/am/api/orgs/{orgId}/groups/{groupId}/roles", readMethods, initType(new(RolesDto)), queryparams["OrgGroupRolesById"]},
	"OrgNestedEnterpriseGroups":    {"/csp/gateway/am/api/orgs/{orgId}/groups/{groupId}/groups", readMethods, initType(new(PagedResponseGroupDto)), queryparams["OrgNestedEnterpriseGroups"]},
	"OrgUsersInGroup":              {"/csp/gateway/am/api/orgs/{orgId}/groups/{groupId}/users", readMethods, initType(new(PagedResponseBaseUserWithProfileDto)), queryparams["OrgUsersInGroup"]},
	"OrgGroups":                    {"/csp/gateway/am/api/orgs/{orgId}/groups", readMethods, initType(new(PagedResponseExpandedGroupDto)), queryparams["OrgGroups"]},
	"AutoEntitlementPolicies":      {"/csp/gateway/am/api/orgs/{orgId}/auto-entitlement-policies", readMethods, initType(new(PagedResponseAutoEntitlementPoliciesDto)), queryparams["AutoEntitlementPolicies"]},
	"AutoEntitlementPoliciesByOrg": {"/csp/gateway/am/api/orgs/{orgId}/auto-entitlement-policies/{policyId}", readMethods, initType(new(ExpandedAutoEntitlementPoliciesDto)), queryparams["AutoEntitlementPoliciesByOrg"]},
	"CustomRole":                   {"/csp/gateway/iam-roles-mgmt/api/orgs/{orgId}/custom-roles/{roleName}", readMethods, initType(new(CustomRoleDto)), queryparams["CustomRole"]},
	"CustomRoles":                  {"/csp/gateway/iam-roles-mgmt/api/orgs/{orgId}/custom-roles", readMethods, initType(new(ResultsDtoCustomRoleDto)), queryparams["CustomRoles"]},
}

// params to set variables and construct API resources
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

// queryparams for specific resource types
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
