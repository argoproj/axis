/*
Copyright 2018 KompiTech GmbH
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
	http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package github

import (
	"context"
	"fmt"
	"github.com/argoproj/argo-events/gateways/common/webhook"

	"github.com/argoproj/argo-events/gateways"
	gwcommon "github.com/argoproj/argo-events/gateways/common"
	"github.com/argoproj/argo-events/pkg/apis/eventsources/v1alpha1"
	"github.com/ghodss/yaml"
)

// Validate validates an event source
func (listener *EventListener) ValidateEventSource(ctx context.Context, eventSource *gateways.EventSource) (*gateways.ValidEventSource, error) {
	var githubEventSource *v1alpha1.GithubEventSource
	if err := yaml.Unmarshal(eventSource.Value, &githubEventSource); err != nil {
		return &gateways.ValidEventSource{
			IsValid: false,
			Reason:  err.Error(),
		}, err
	}

	if err := validateGithubEventSource(githubEventSource); err != nil {
		return &gateways.ValidEventSource{
			Reason:  err.Error(),
			IsValid: false,
		}, err
	}

	return &gateways.ValidEventSource{
		IsValid: true,
	}, nil
}

// validateGithubEventSource checks if github event source is valid
func validateGithubEventSource(githubEventSource *v1alpha1.GithubEventSource) error {
	if githubEventSource == nil {
		return gwcommon.ErrNilEventSource
	}
	if githubEventSource.Repository == "" {
		return fmt.Errorf("repository cannot be empty")
	}
	if githubEventSource.Owner == "" {
		return fmt.Errorf("owner cannot be empty")
	}
	if githubEventSource.APIToken == nil {
		return fmt.Errorf("api token can't be empty")
	}
	if githubEventSource.Events == nil || len(githubEventSource.Events) < 1 {
		return fmt.Errorf("events must be defined")
	}
	if githubEventSource.ContentType != "" {
		if !(githubEventSource.ContentType == "json" || githubEventSource.ContentType == "form") {
			return fmt.Errorf("content type must be \"json\" or \"form\"")
		}
	}
	return webhook.ValidateWebhookContext(githubEventSource.Webhook)
}
