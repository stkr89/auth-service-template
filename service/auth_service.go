package service

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	cognito "github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	"github.com/go-kit/kit/log"
	"github.com/stkr89/auth-service-template/common"
	"github.com/stkr89/auth-service-template/types"
	"os"
)

// AuthService interface
type AuthService interface {
	SignUp(ctx context.Context, request *types.CreateUserRequest) (*types.CreateUserResponse, error)
}

type AuthServiceImpl struct {
	logger log.Logger
	client *cognito.CognitoIdentityProvider
}

func NewAuthServiceImpl() *AuthServiceImpl {
	return &AuthServiceImpl{
		logger: common.NewLogger(),
		client: common.NewAWSCognitoClient(),
	}
}

func (s AuthServiceImpl) SignUp(ctx context.Context, request *types.CreateUserRequest) (*types.CreateUserResponse, error) {
	awsUser := &cognito.SignUpInput{
		Username: aws.String(request.Email),
		Password: aws.String(request.Password),
		ClientId: s.getClientId(),
		UserAttributes: []*cognito.AttributeType{
			{
				Name:  common.StringToPtr("custom:firstName"),
				Value: &request.FirstName,
			},
			{
				Name:  common.StringToPtr("custom:lastName"),
				Value: &request.LastName,
			},
		},
	}

	signUpOutput, err := s.client.SignUp(awsUser)
	if err != nil {
		s.logger.Log("message", "unable to signup user", "error", err)
		return nil, common.SignUpFailed
	}

	s.logger.Log("message", "sign up successful", "email", request.Email)

	return &types.CreateUserResponse{
		ID:        *signUpOutput.UserSub,
		FirstName: request.FirstName,
		LastName:  request.LastName,
		Email:     request.Email,
	}, nil
}

func (s AuthServiceImpl) getClientId() *string {
	return aws.String(os.Getenv("AWS_COGNITO_CLIENT_ID"))
}
