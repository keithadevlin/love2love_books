// +build unit

package awssecret_test

import (
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

type testSecretManager struct {
	ctrl   *gomock.Controller
	client *awssecret.MockSecretManagerClient
	sm     *awssecret.SecretManager
}

func (s *testSecretManager) Finish() {
	s.ctrl.Finish()
}

func newTestSecretManager(t *testing.T) *testSecretManager {
	ctrl := gomock.NewController(t)
	c := awssecret.NewMockSecretManagerClient(ctrl)

	h := &testSecretManager{
		ctrl:   ctrl,
		client: c,
		sm:     awssecret.NewSecretManagerWithClient(c),
	}
	return h
}

func TestSecretManager_Get(t *testing.T) {
	tests := []struct {
		Desc           string
		SecretKey      string
		ClientResult   *secretsmanager.GetSecretValueOutput
		ClientErr      error
		ExpectedResult string
		ExpectedErr    error
	}{
		{
			Desc:           "should fail because of secretsmanager error",
			ClientResult:   nil,
			ClientErr:      fmt.Errorf("cannot get secret"),
			ExpectedResult: "",
			ExpectedErr:    fmt.Errorf("error getting value from secret manager: cannot get secret"),
		},
		{
			Desc:           "should fail because of empty secrets returned",
			ClientResult:   &secretsmanager.GetSecretValueOutput{},
			ClientErr:      nil,
			ExpectedResult: "",
			ExpectedErr:    fmt.Errorf("empty secret string found"),
		},
		{
			Desc: "should succeed",
			ClientResult: &secretsmanager.GetSecretValueOutput{
				SecretString: aws.String("mysecret"),
			},
			ClientErr:      nil,
			ExpectedResult: "mysecret",
		},
	}

	for _, tt := range tests {
		t.Run(tt.Desc, func(t *testing.T) {
			unit := newTestSecretManager(t)
			defer unit.Finish()

			expectedRecord := &secretsmanager.GetSecretValueInput{
				SecretId: aws.String(tt.SecretKey),
			}

			unit.client.EXPECT().
				GetSecretValueWithContext(gomock.Any(), expectedRecord).
				Return(tt.ClientResult, tt.ClientErr)

			result, err := unit.sm.Get(testutil.Context(), tt.SecretKey)

			require.Equal(t, tt.ExpectedResult, result)

			if tt.ExpectedErr != nil {
				require.Error(t, err)
				require.EqualError(t, err, tt.ExpectedErr.Error())
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestSecretManager_Update(t *testing.T) {
	tests := []struct {
		Desc           string
		SecretKey      string
		SecretValue    string
		ClientErr      error
		ExpectedResult string
		ExpectedErr    error
	}{
		{
			Desc:           "should fail because of secretsmanager error",
			ClientErr:      fmt.Errorf("cannot update secret"),
			ExpectedResult: "",
			ExpectedErr:    fmt.Errorf("error updating value in secret manager: cannot update secret"),
		},
		{
			Desc:           "should succeed",
			SecretKey:      "some-key",
			SecretValue:    "some-value",
			ClientErr:      nil,
			ExpectedResult: "mysecret",
		},
	}

	for _, tt := range tests {
		t.Run(tt.Desc, func(t *testing.T) {
			unit := newTestSecretManager(t)
			defer unit.Finish()

			expectedRecord := &secretsmanager.PutSecretValueInput{
				SecretId:     aws.String(tt.SecretKey),
				SecretString: aws.String(tt.SecretValue),
			}

			unit.client.EXPECT().
				PutSecretValueWithContext(gomock.Any(), expectedRecord).
				Return(nil, tt.ClientErr)

			err := unit.sm.Update(testutil.Context(), tt.SecretKey, tt.SecretValue)

			if tt.ExpectedErr != nil {
				require.Error(t, err)
				require.EqualError(t, err, tt.ExpectedErr.Error())
			} else {
				require.NoError(t, err)
			}
		})
	}
}
