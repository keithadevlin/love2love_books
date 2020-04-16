//go:generate mockgen -package awssecret -source=secretmanager.go -destination secretmanager_mock.go

package awssecret

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/client"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
	"github.com/aws/aws-sdk-go/service/secretsmanager/secretsmanageriface"
	"github.com/aws/aws-xray-sdk-go/xray"
	"github.com/pkg/errors"
)

type SecretManagerClient interface {
	secretsmanageriface.SecretsManagerAPI
}

// NewSecretManager configures and return an instance of a Kinesis client
func NewSecretManager(provider client.ConfigProvider) *SecretManager {
	sc := secretsmanager.New(provider)
	xray.AWS(sc.Client)
	return NewSecretManagerWithClient(sc)
}

// NewSecretManagerWithClient configures and return an instance of a Kinesis client
func NewSecretManagerWithClient(c SecretManagerClient) *SecretManager {
	return &SecretManager{
		client: c,
	}
}

// SecretManager is a simple abstraction to help dealing with Kinesis streams
type SecretManager struct {
	client SecretManagerClient
}

// Get returns a string secret from secrets manager
func (c SecretManager) Get(ctx context.Context, id string) (string, error) {
	//ctx, span := trace.StartSpan(ctx, "awssecret.Get")
	//defer span.End()

	record := &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(id),
	}

	res, err := c.client.GetSecretValueWithContext(ctx, record)
	if err != nil {
		return "", errors.Wrap(err, `error getting value from secret manager`)
	}

	if res.SecretString == nil {
		return "", fmt.Errorf("empty secret string found")
	}
	return *res.SecretString, nil
}

// Update value for a secret in secrets manager
func (c SecretManager) Update(ctx context.Context, id string, value string) error {
	//ctx, span := trace.StartSpan(ctx, "awssecret.Update")
	//defer span.End()

	record := &secretsmanager.PutSecretValueInput{
		SecretId:     aws.String(id),
		SecretString: aws.String(value),
	}

	_, err := c.client.PutSecretValueWithContext(ctx, record)
	if err != nil {
		return errors.Wrap(err, `error updating value in secret manager`)
	}

	return nil
}
