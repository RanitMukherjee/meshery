// Copyright Meshery Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package root

import "github.com/meshery/meshkit/errors"

// Please reference the following before contributing an error code:
// https://docs.meshery.io/project/contributing/contributing-error
// https://github.com/meshery/meshkit/blob/master/errors/errors.go
const (
	ErrProcessingConfigCode        = "mesheryctl-1010"
	ErrCreatingConfigFileCode      = "mesheryctl-1011"
	ErrAddingTokenToConfigCode     = "mesheryctl-1012"
	ErrAddingContextToConfigCode   = "mesheryctl-1013"
	ErrUnmarshallingConfigFileCode = "mesheryctl-1014"
	ErrGettingRequestContextCode   = "mesheryctl-1015"
	ErrUnmarshallingAPIDataCode    = "mesheryctl-1016"
	ErrConnectingToServerCode      = "mesheryctl-1017"
)

var (
	ErrCreatingConfigFile = errors.New(ErrCreatingConfigFileCode, errors.Alert, []string{"Unable to create config file"}, []string{"Unable to create config file"}, []string{}, []string{})

	ErrAddingTokenToConfig = errors.New(ErrAddingTokenToConfigCode, errors.Alert, []string{"Unable to add token to config"}, []string{"Unable to add token to config"}, []string{}, []string{})

	ErrAddingContextToConfig = errors.New(ErrAddingContextToConfigCode, errors.Alert, []string{"Unable to add context to config"}, []string{"Unable to add context to config"}, []string{}, []string{})

	ErrUnmarshallingConfigFile = errors.New(ErrUnmarshallingConfigFileCode, errors.Alert, []string{"Error processing json in config file"}, []string{"Error processing json in config file"}, []string{}, []string{})
)

func ErrProcessingConfig(err error) error {
	return errors.New(ErrProcessingConfigCode, errors.Alert, []string{"Error processing config"}, []string{"Error processing config", err.Error()}, []string{}, []string{})
}

func ErrConnectingToServer(err error) error {
	return errors.New(ErrConnectingToServerCode, errors.Fatal, []string{"Unable to communicate with Meshery server"}, []string{"Unable to communicate with Meshery server", err.Error(), "See https://docs.meshery.io for help getting started with Meshery"}, []string{}, []string{"See https://docs.meshery.io for help getting started with Meshery"})
}

func ErrGettingRequestContext(err error) error {
	return errors.New(ErrGettingRequestContextCode, errors.Fatal, []string{"Unable to add token to config"}, []string{"Unable to add token to config", err.Error()}, []string{}, []string{})
}

func ErrUnmarshallingAPIData(err error) error {
	return errors.New(ErrUnmarshallingAPIDataCode, errors.Fatal, []string{"Error processing json API data"}, []string{"Error processing json API data", err.Error()}, []string{}, []string{})
}
