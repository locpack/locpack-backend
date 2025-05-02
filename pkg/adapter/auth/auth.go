package auth

import (
	"context"

	"github.com/Nerzal/gocloak/v13"
	"github.com/google/uuid"
	"locpack-backend/pkg/adapter"
	"locpack-backend/pkg/cfg"
	"locpack-backend/pkg/types"
)

const userRole = "user"

type authImpl struct {
	cfg    *cfg.Auth
	client *gocloak.GoCloak
}

func New(cfg *cfg.Auth) adapter.Auth {
	client := gocloak.NewClient(cfg.URL)
	return &authImpl{cfg, client}
}

func (a *authImpl) Register(username string, email string, password string) (uuid.UUID, error) {
	ctx := context.Background()

	token, err := a.client.LoginAdmin(ctx, a.cfg.AdminUsername, a.cfg.AdminPassword, a.cfg.Realm)
	if err != nil {
		return uuid.UUID{}, err
	}

	user := gocloak.User{
		Username:      gocloak.StringP(username),
		Email:         gocloak.StringP(email),
		Enabled:       gocloak.BoolP(true),
		EmailVerified: gocloak.BoolP(true),
		Credentials: &[]gocloak.CredentialRepresentation{
			{
				Type:      gocloak.StringP("password"),
				Value:     gocloak.StringP(password),
				Temporary: gocloak.BoolP(false),
			},
		},
	}

	userID, err := a.client.CreateUser(ctx, token.AccessToken, a.cfg.Realm, user)
	if err != nil {
		return uuid.UUID{}, err
	}

	role, err := a.client.GetRealmRole(ctx, token.AccessToken, a.cfg.Realm, userRole)
	if err != nil {
		return uuid.UUID{}, err
	}

	err = a.client.AddRealmRoleToUser(ctx, token.AccessToken, a.cfg.Realm, userID, []gocloak.Role{*role})
	if err != nil {
		return uuid.UUID{}, err
	}

	userUUID, err := uuid.Parse(userID)
	if err != nil {
		return uuid.UUID{}, err
	}

	return userUUID, err
}

func (a *authImpl) Login(username string, password string) (types.AccessToken, error) {
	ctx := context.Background()

	token, err := a.client.Login(ctx, a.cfg.ClientID, a.cfg.ClientSecret, a.cfg.Realm, username, password)
	if err != nil {
		return types.AccessToken{}, err
	}

	resultToken := types.AccessToken{
		Value:        token.AccessToken,
		RefreshToken: token.RefreshToken,
		ExpiresIn:    float64(token.ExpiresIn),
	}

	return resultToken, err
}

func (a *authImpl) Refresh(value string) (types.AccessToken, error) {
	ctx := context.Background()

	token, err := a.client.RefreshToken(ctx, value, a.cfg.ClientID, a.cfg.ClientSecret, a.cfg.Realm)
	if err != nil {
		return types.AccessToken{}, err
	}

	resultToken := types.AccessToken{
		Value:        token.AccessToken,
		RefreshToken: token.RefreshToken,
		ExpiresIn:    float64(token.ExpiresIn),
	}

	return resultToken, err
}

func (a *authImpl) DecodeToken(accessToken string) (types.Token, error) {
	ctx := context.Background()

	token, jwtClaims, err := a.client.DecodeAccessToken(ctx, accessToken, a.cfg.Realm)
	if err != nil {
		return types.Token{}, err
	}

	claims := *jwtClaims
	realmAccess := claims["realm_access"].(map[string]interface{})

	var roles []string
	for _, role := range realmAccess["roles"].([]interface{}) {
		role := role.(string)
		roles = append(roles, role)
	}

	tokenInsight := types.Token{
		Valid:     token.Valid,
		Username:  claims["preferred_username"].(string),
		Email:     claims["email"].(string),
		ExpiresIn: claims["exp"].(float64),
		Roles:     roles,
	}

	return tokenInsight, err
}
