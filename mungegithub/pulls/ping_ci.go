/*
Copyright 2015 The Kubernetes Authors All rights reserved.

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

package pulls

import (
	"time"

	github_util "k8s.io/contrib/github"
	"k8s.io/contrib/mungegithub/config"

	"github.com/golang/glog"
	"github.com/google/go-github/github"
	"github.com/spf13/cobra"
)

// PingCIMunger looks for situations CI (Travis | Shippable) has flaked for some
// reason and we want to re-run them.  Achieves this by closing and re-opening the pr
type PingCIMunger struct{}

func init() {
	RegisterMungerOrDie(PingCIMunger{})
}

func (PingCIMunger) Name() string { return "ping-ci" }

func (PingCIMunger) AddFlags(cmd *cobra.Command) {}

func (PingCIMunger) MungePullRequest(config *config.MungeConfig, pr *github.PullRequest, issue *github.Issue, commits []github.RepositoryCommit, events []github.IssueEvent) {
	if !github_util.HasLabel(issue.Labels, "lgtm") {
		return
	}
	if mergeable, err := config.IsPRMergeable(pr); err != nil {
		glog.V(2).Infof("Skipping %d - problem determining mergeability", *pr.Number)
	} else if !mergeable {
		glog.V(2).Infof("Skipping %d - not mergeable", *pr.Number)
	}
	status, err := config.GetStatus(pr, []string{"Shippable", "continuous-integration/travis-ci/pr"})
	if err != nil {
		glog.Errorf("unexpected error getting status: %v", err)
		return
	}
	if status == "incomplete" {
		glog.V(2).Infof("status is incomplete, closing and re-opening")
		msg := "Continuous integration appears to have missed, closing and re-opening to trigger it"
		config.WriteComment(*pr.Number, msg)

		config.ClosePR(pr)
		time.Sleep(5 * time.Second)
		config.OpenPR(pr, 10)
	}
}
