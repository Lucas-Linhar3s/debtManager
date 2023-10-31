package auth

import "time"

type UserCredentials struct {
	Email    string
	Password string
	Data     interface{}
}

type User struct {
	ID                 string                    `json:"id"`
	Aud                string                    `json:"aud"`
	Role               string                    `json:"role"`
	Email              string                    `json:"email"`
	InvitedAt          time.Time                 `json:"invited_at"`
	ConfirmedAt        time.Time                 `json:"confirmed_at"`
	ConfirmationSentAt time.Time                 `json:"confirmation_sent_at"`
	AppMetadata        struct{ provider string } `json:"app_metadata"`
	UserMetadata       map[string]interface{}    `json:"user_metadata"`
	CreatedAt          time.Time                 `json:"created_at"`
	UpdatedAt          time.Time                 `json:"updated_at"`
}

type AuthenticatedDetails struct {
	AccessToken          string `json:"access_token"`
	TokenType            string `json:"token_type"`
	ExpiresIn            int    `json:"expires_in"`
	RefreshToken         string `json:"refresh_token"`
	User                 User   `json:"user"`
	ProviderToken        string `json:"provider_token"`
	ProviderRefreshToken string `json:"provider_refresh_token"`
}

type ExchangeCodeOpts struct {
	AuthCode     string `json:"auth_code"`
	CodeVerifier string `json:"code_verifier"`
}

type ProviderSignInOptions struct {
	Provider   string   `url:"provider"`
	RedirectTo string   `url:"redirect_to"`
	Scopes     []string `url:"scopes"`
	FlowType   FlowType
}

type FlowType string

const (
	Implicit FlowType = "implicit"
	PKCE     FlowType = "pkce"
)

type ProviderSignInDetails struct {
	URL          string `json:"url"`
	Provider     string `json:"provider"`
	CodeVerifier string `json:"code_verifier"`
}

type AppMetadata struct {
	Provider  *string   `json:"provider"`
	Providers *[]string `json:"providers"`
}

type IdentityData struct {
	Email *string `json:"email"`
	Sub   *string `json:"sub"`
}

type Identities struct {
	CreatedAt    *time.Time   `json:"created_at"`
	ID           *string      `json:"id"`
	IdentityData IdentityData `json:"identity_data"`
	LastSignInAt *time.Time   `json:"last_sign_in_at"`
	Provider     *string      `json:"provider"`
	UpdatedAt    *time.Time   `json:"updated_at"`
	UserID       *string      `json:"user_id"`
}

type Res struct {
	AppMetadata        *AppMetadata            `json:"app_metadata"`
	Aud                *string                 `json:"aud"`
	ConfirmationSentAt *time.Time              `json:"confirmation_sent_at"`
	Email              *string                 `json:"email"`
	ID                 *string                 `json:"id"`
	Identities         *[]Identities           `json:"identities"`
	Phone              *string                 `json:"phone"`
	Role               *string                 `json:"role"`
	UpdatedAt          *time.Time              `json:"updated_at"`
	UserMetadata       *map[string]interface{} `json:"user_metadata"`
}
