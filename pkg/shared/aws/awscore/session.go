package awscore

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/kelseyhightower/envconfig"
)

type config struct {
	Region string `envconfig:"REGION" required:"true"`
}

// NewSession returns an new instance of a default AWS session
func NewSession() (*session.Session, error) {
	var cfg config
	if err := envconfig.Process("", &cfg); err != nil {
		return nil, err
	}
	return session.NewSession(&aws.Config{Region: aws.String(cfg.Region)})
}

// NewSessionWithRegion returns an new instance of an AWS session for the specified region
func NewSessionWithRegion(region string) (*session.Session, error) {
	return session.NewSession(&aws.Config{Region: aws.String(region)})
}

// NewSessionWithRegion returns an new instance of an AWS session with the specified configuration
func NewSessionWithConfig(cfg *aws.Config) (*session.Session, error) {
	return session.NewSession(cfg)
}
