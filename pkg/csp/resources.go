package csp

import (
	"reflect"
)

var CSPResourceMap = map[string]reflect.Type{
	"PublicKey":                    initType(new(PublicKeyDto)),
	"CspOpenIdConfiguration":       initType(new(CspOpenIdConfigurationDto)),
	"OpenIdUserInfo":               initType(new(OidcUserInfoDto)),
	"LoggedInUser":                 initType(new(UserDto)),
	"LoggedInUserInfo":             initType(new(UserInfoDto)),
	"LoggedInUserGroups":           initType(new(UserGroupsResponse)),
	"LoggedInUserRoles":            initType(new(RolesDto)),
	"LoggedInUserServiceRoles":     initType(new(UserServiceRolesResponse)),
	"LoggedInUserOrgs":             initType(new(PagedResponseOrganizationDto)),
	"LoggedInMfaStatus":            initType(new(UserMfaStatusDto)),
	"User":                         initType(new(UserInfoDto)),
	"UserCustomRoles":              initType(new(UserCustomRolesResponse)),
	"UserGroups":                   initType(new(UserGroupsResponse)),
	"UserGroupsByEmail":            initType(new(UserGroupsResponse)),
	"UserRoles":                    initType(new(RolesDto)),
	"UserServiceRoles":             initType(new(UserServiceRolesResponse)),
	"OrgOAuthApps":                 initType(new(PagedResponseOrgOAuthAppResponse)),
	"OrgOAuthAppById":              initType(new(OrgOAuthAppResponse)),
	"OrgSearchGroups":              initType(new(SearchGroupsResponse)),
	"OrgLinkedIdp":                 initType(new(OrgLinkedIdpResponse)),
	"Org":                          initType(new(OrganizationDto)),
	"OrgClients":                   initType(new(PagedResponseExpandedAuthClientDto)),
	"OrgInvitations":               initType(new(UserOrganizationInvitationResponse)),
	"OrgRoles":                     initType(new(OrganizationRolesResponse)),
	"OrgTrusts":                    initType(new(PagedResponseOrganizationTrustDto)),
	"OrgUsers":                     initType(new(PagedResponseExpandedTypedUserDto)),
	"OrgSubOrgs":                   initType(new(PagedResponseOrganizationDto)),
	"OrgMemberOrgs":                initType(new(PagedResponseOrganizationMemberForTrustDto)),
	"OrgInvitationById":            initType(new(UserOrganizationInvitationResponse)),
	"OrgGroupById":                 initType(new(ExpandedGroupDto)),
	"OrgGroupRolesById":            initType(new(RolesDto)),
	"OrgNestedEnterpriseGroups":    initType(new(PagedResponseGroupDto)),
	"OrgUsersInGroup":              initType(new(PagedResponseBaseUserWithProfileDto)),
	"OrgGroups":                    initType(new(PagedResponseExpandedGroupDto)),
	"AutoEntitlementPolicies":      initType(new(PagedResponseAutoEntitlementPoliciesDto)),
	"AutoEntitlementPoliciesByOrg": initType(new(ExpandedAutoEntitlementPoliciesDto)),
	"CustomRole":                   initType(new(CustomRoleDto)),
	"CustomRoles":                  initType(new(ResultsDtoCustomRoleDto)),
}

func initType(i interface{}) reflect.Type {
	return reflect.TypeOf(i)
}
