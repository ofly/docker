// Code generated by protoc-gen-go.
// source: google.golang.org/genproto/googleapis/api/serviceconfig/auth.proto
// DO NOT EDIT!

package google_api // import "google.golang.org/genproto/googleapis/api/serviceconfig"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// `Authentication` defines the authentication configuration for an API.
//
// Example for an API targeted for external use:
//
//     name: calendar.googleapis.com
//     authentication:
//       rules:
//       - selector: "*"
//         oauth:
//           canonical_scopes: https://www.googleapis.com/auth/calendar
//
//       - selector: google.calendar.Delegate
//         oauth:
//           canonical_scopes: https://www.googleapis.com/auth/calendar.read
type Authentication struct {
	// A list of authentication rules that apply to individual API methods.
	//
	// **NOTE:** All service configuration rules follow "last one wins" order.
	Rules []*AuthenticationRule `protobuf:"bytes,3,rep,name=rules" json:"rules,omitempty"`
	// Defines a set of authentication providers that a service supports.
	Providers []*AuthProvider `protobuf:"bytes,4,rep,name=providers" json:"providers,omitempty"`
}

func (m *Authentication) Reset()                    { *m = Authentication{} }
func (m *Authentication) String() string            { return proto.CompactTextString(m) }
func (*Authentication) ProtoMessage()               {}
func (*Authentication) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *Authentication) GetRules() []*AuthenticationRule {
	if m != nil {
		return m.Rules
	}
	return nil
}

func (m *Authentication) GetProviders() []*AuthProvider {
	if m != nil {
		return m.Providers
	}
	return nil
}

// Authentication rules for the service.
//
// By default, if a method has any authentication requirements, every request
// must include a valid credential matching one of the requirements.
// It's an error to include more than one kind of credential in a single
// request.
//
// If a method doesn't have any auth requirements, request credentials will be
// ignored.
type AuthenticationRule struct {
	// Selects the methods to which this rule applies.
	//
	// Refer to [selector][google.api.DocumentationRule.selector] for syntax details.
	Selector string `protobuf:"bytes,1,opt,name=selector" json:"selector,omitempty"`
	// The requirements for OAuth credentials.
	Oauth *OAuthRequirements `protobuf:"bytes,2,opt,name=oauth" json:"oauth,omitempty"`
	// Whether to allow requests without a credential. The credential can be
	// an OAuth token, Google cookies (first-party auth) or EndUserCreds.
	//
	// For requests without credentials, if the service control environment is
	// specified, each incoming request **must** be associated with a service
	// consumer. This can be done by passing an API key that belongs to a consumer
	// project.
	AllowWithoutCredential bool `protobuf:"varint,5,opt,name=allow_without_credential,json=allowWithoutCredential" json:"allow_without_credential,omitempty"`
	// Requirements for additional authentication providers.
	Requirements []*AuthRequirement `protobuf:"bytes,7,rep,name=requirements" json:"requirements,omitempty"`
}

func (m *AuthenticationRule) Reset()                    { *m = AuthenticationRule{} }
func (m *AuthenticationRule) String() string            { return proto.CompactTextString(m) }
func (*AuthenticationRule) ProtoMessage()               {}
func (*AuthenticationRule) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *AuthenticationRule) GetOauth() *OAuthRequirements {
	if m != nil {
		return m.Oauth
	}
	return nil
}

func (m *AuthenticationRule) GetRequirements() []*AuthRequirement {
	if m != nil {
		return m.Requirements
	}
	return nil
}

// Configuration for an anthentication provider, including support for
// [JSON Web Token (JWT)](https://tools.ietf.org/html/draft-ietf-oauth-json-web-token-32).
type AuthProvider struct {
	// The unique identifier of the auth provider. It will be referred to by
	// `AuthRequirement.provider_id`.
	//
	// Example: "bookstore_auth".
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	// Identifies the principal that issued the JWT. See
	// https://tools.ietf.org/html/draft-ietf-oauth-json-web-token-32#section-4.1.1
	// Usually a URL or an email address.
	//
	// Example: https://securetoken.google.com
	// Example: 1234567-compute@developer.gserviceaccount.com
	Issuer string `protobuf:"bytes,2,opt,name=issuer" json:"issuer,omitempty"`
	// URL of the provider's public key set to validate signature of the JWT. See
	// [OpenID Discovery](https://openid.net/specs/openid-connect-discovery-1_0.html#ProviderMetadata).
	// Optional if the key set document:
	//  - can be retrieved from
	//    [OpenID Discovery](https://openid.net/specs/openid-connect-discovery-1_0.html
	//    of the issuer.
	//  - can be inferred from the email domain of the issuer (e.g. a Google service account).
	//
	// Example: https://www.googleapis.com/oauth2/v1/certs
	JwksUri string `protobuf:"bytes,3,opt,name=jwks_uri,json=jwksUri" json:"jwks_uri,omitempty"`
}

func (m *AuthProvider) Reset()                    { *m = AuthProvider{} }
func (m *AuthProvider) String() string            { return proto.CompactTextString(m) }
func (*AuthProvider) ProtoMessage()               {}
func (*AuthProvider) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

// OAuth scopes are a way to define data and permissions on data. For example,
// there are scopes defined for "Read-only access to Google Calendar" and
// "Access to Cloud Platform". Users can consent to a scope for an application,
// giving it permission to access that data on their behalf.
//
// OAuth scope specifications should be fairly coarse grained; a user will need
// to see and understand the text description of what your scope means.
//
// In most cases: use one or at most two OAuth scopes for an entire family of
// products. If your product has multiple APIs, you should probably be sharing
// the OAuth scope across all of those APIs.
//
// When you need finer grained OAuth consent screens: talk with your product
// management about how developers will use them in practice.
//
// Please note that even though each of the canonical scopes is enough for a
// request to be accepted and passed to the backend, a request can still fail
// due to the backend requiring additional scopes or permissions.
type OAuthRequirements struct {
	// The list of publicly documented OAuth scopes that are allowed access. An
	// OAuth token containing any of these scopes will be accepted.
	//
	// Example:
	//
	//      canonical_scopes: https://www.googleapis.com/auth/calendar,
	//                        https://www.googleapis.com/auth/calendar.read
	CanonicalScopes string `protobuf:"bytes,1,opt,name=canonical_scopes,json=canonicalScopes" json:"canonical_scopes,omitempty"`
}

func (m *OAuthRequirements) Reset()                    { *m = OAuthRequirements{} }
func (m *OAuthRequirements) String() string            { return proto.CompactTextString(m) }
func (*OAuthRequirements) ProtoMessage()               {}
func (*OAuthRequirements) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{3} }

// User-defined authentication requirements, including support for
// [JSON Web Token (JWT)](https://tools.ietf.org/html/draft-ietf-oauth-json-web-token-32).
type AuthRequirement struct {
	// [id][google.api.AuthProvider.id] from authentication provider.
	//
	// Example:
	//
	//     provider_id: bookstore_auth
	ProviderId string `protobuf:"bytes,1,opt,name=provider_id,json=providerId" json:"provider_id,omitempty"`
	// The list of JWT
	// [audiences](https://tools.ietf.org/html/draft-ietf-oauth-json-web-token-32#section-4.1.3).
	// that are allowed to access. A JWT containing any of these audiences will
	// be accepted. When this setting is absent, only JWTs with audience
	// "https://[Service_name][google.api.Service.name]/[API_name][google.protobuf.Api.name]"
	// will be accepted. For example, if no audiences are in the setting,
	// LibraryService API will only accept JWTs with the following audience
	// "https://library-example.googleapis.com/google.example.library.v1.LibraryService".
	//
	// Example:
	//
	//     audiences: bookstore_android.apps.googleusercontent.com,
	//                bookstore_web.apps.googleusercontent.com
	Audiences string `protobuf:"bytes,2,opt,name=audiences" json:"audiences,omitempty"`
}

func (m *AuthRequirement) Reset()                    { *m = AuthRequirement{} }
func (m *AuthRequirement) String() string            { return proto.CompactTextString(m) }
func (*AuthRequirement) ProtoMessage()               {}
func (*AuthRequirement) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{4} }

func init() {
	proto.RegisterType((*Authentication)(nil), "google.api.Authentication")
	proto.RegisterType((*AuthenticationRule)(nil), "google.api.AuthenticationRule")
	proto.RegisterType((*AuthProvider)(nil), "google.api.AuthProvider")
	proto.RegisterType((*OAuthRequirements)(nil), "google.api.OAuthRequirements")
	proto.RegisterType((*AuthRequirement)(nil), "google.api.AuthRequirement")
}

func init() {
	proto.RegisterFile("google.golang.org/genproto/googleapis/api/serviceconfig/auth.proto", fileDescriptor1)
}

var fileDescriptor1 = []byte{
	// 425 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x52, 0x4d, 0x6f, 0x13, 0x31,
	0x10, 0x55, 0xd2, 0xa6, 0xcd, 0x4e, 0xaa, 0x14, 0x7c, 0xa8, 0x4c, 0xf9, 0xaa, 0x56, 0x1c, 0xca,
	0x65, 0x57, 0x6a, 0x11, 0xe2, 0x04, 0x6a, 0x38, 0xa0, 0x9c, 0x08, 0x46, 0x88, 0xe3, 0xca, 0x78,
	0xcd, 0xd6, 0xe0, 0x7a, 0x82, 0xed, 0x6d, 0x6e, 0xfc, 0x18, 0x7e, 0x19, 0x3f, 0x05, 0xaf, 0x77,
	0x9b, 0x6c, 0x93, 0x1b, 0x97, 0x28, 0x33, 0xef, 0xcd, 0x7b, 0x7e, 0x33, 0x0b, 0xb3, 0x0a, 0xb1,
	0xd2, 0x32, 0xab, 0x50, 0x73, 0x53, 0x65, 0x68, 0xab, 0xbc, 0x92, 0x66, 0x69, 0xd1, 0x63, 0xde,
	0x42, 0x7c, 0xa9, 0x5c, 0x1e, 0x7e, 0x72, 0x27, 0xed, 0xad, 0x12, 0x52, 0xa0, 0xf9, 0xae, 0xaa,
	0x9c, 0xd7, 0xfe, 0x3a, 0x8b, 0x3c, 0x02, 0x9d, 0x46, 0x20, 0x9d, 0xce, 0xff, 0x5b, 0xcf, 0x18,
	0xf4, 0xdc, 0x2b, 0x34, 0xae, 0x95, 0x4d, 0x7f, 0xc3, 0xf4, 0x2a, 0x98, 0x48, 0xe3, 0x95, 0x88,
	0x00, 0x79, 0x05, 0x23, 0x5b, 0x6b, 0xe9, 0xe8, 0xde, 0xd9, 0xde, 0xf9, 0xe4, 0xe2, 0x59, 0xb6,
	0x31, 0xce, 0xee, 0x53, 0x59, 0xa0, 0xb1, 0x96, 0x4c, 0x5e, 0x43, 0x12, 0x04, 0x6f, 0x55, 0x29,
	0xad, 0xa3, 0xfb, 0x71, 0x92, 0x6e, 0x4f, 0x2e, 0x3a, 0x02, 0xdb, 0x50, 0xd3, 0xbf, 0x03, 0x20,
	0xbb, 0xaa, 0xe4, 0x14, 0xc6, 0x4e, 0x6a, 0x29, 0x3c, 0x5a, 0x3a, 0x38, 0x1b, 0x9c, 0x27, 0x6c,
	0x5d, 0x93, 0x4b, 0x18, 0x61, 0xb3, 0x18, 0x3a, 0x0c, 0xc0, 0xe4, 0xe2, 0x69, 0xdf, 0xe6, 0x63,
	0xa3, 0xc5, 0xe4, 0xaf, 0x5a, 0x59, 0x79, 0x13, 0x34, 0x1d, 0x6b, 0xb9, 0xe4, 0x0d, 0x50, 0xae,
	0x35, 0xae, 0x8a, 0x95, 0xf2, 0xd7, 0x58, 0xfb, 0x42, 0x58, 0x59, 0x36, 0xa6, 0x5c, 0xd3, 0x51,
	0xd0, 0x19, 0xb3, 0x93, 0x88, 0x7f, 0x6d, 0xe1, 0xf7, 0x6b, 0x94, 0xbc, 0x83, 0x23, 0xdb, 0x13,
	0xa4, 0x87, 0x31, 0xdc, 0xe3, 0xed, 0x70, 0x3d, 0x53, 0x76, 0x6f, 0x20, 0xfd, 0x04, 0x47, 0xfd,
	0xf4, 0x64, 0x0a, 0x43, 0x55, 0x76, 0xa9, 0xc2, 0x3f, 0x72, 0x02, 0x07, 0xca, 0xb9, 0x5a, 0xda,
	0x18, 0x28, 0x61, 0x5d, 0x45, 0x1e, 0xc1, 0xf8, 0xc7, 0xea, 0xa7, 0x2b, 0x6a, 0xab, 0xc2, 0x2d,
	0x1a, 0xe4, 0xb0, 0xa9, 0xbf, 0x58, 0x95, 0xbe, 0x85, 0x87, 0x3b, 0x49, 0xc9, 0x4b, 0x78, 0x20,
	0xb8, 0x41, 0x13, 0xf6, 0xa8, 0x0b, 0x27, 0x70, 0x19, 0x6e, 0xd8, 0xba, 0x1c, 0xaf, 0xfb, 0x9f,
	0x63, 0x3b, 0x5d, 0xc0, 0xf1, 0xd6, 0x38, 0x79, 0x0e, 0x93, 0xbb, 0xab, 0x14, 0xeb, 0xe7, 0xc1,
	0x5d, 0x6b, 0x5e, 0x92, 0x27, 0x90, 0xf0, 0xba, 0x54, 0xd2, 0x88, 0xa0, 0xdb, 0xbe, 0x74, 0xd3,
	0x98, 0xbd, 0x80, 0xa9, 0xc0, 0x9b, 0xde, 0x52, 0x66, 0x49, 0x17, 0xda, 0xe3, 0x62, 0xf0, 0x67,
	0xb8, 0xff, 0xe1, 0x6a, 0x31, 0xff, 0x76, 0x10, 0x3f, 0xba, 0xcb, 0x7f, 0x01, 0x00, 0x00, 0xff,
	0xff, 0x0d, 0x41, 0xfd, 0x7a, 0x11, 0x03, 0x00, 0x00,
}
