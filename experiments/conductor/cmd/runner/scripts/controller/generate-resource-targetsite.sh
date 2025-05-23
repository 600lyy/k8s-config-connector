#!/bin/bash
# Copyright 2025 Google LLC
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.


set -o errexit
set -o nounset
set -o pipefail

    # - name: discoveryengine-targetsite
    #   local: resource-discoveryengine-targetsite
    #   remote: resource-discoveryengine-targetsite
    #   command: ""
    #   group: discoveryengine
    #   resource: targetsite
    #   controller: Unknown
    #   kind: DiscoveryEngineTargetSite
    #   package: google.cloud.discoveryengine.v1
    #   proto: TargetSite
    #   proto-path: google/cloud/discoveryengine/v1/completion_service.proto
    #   proto-service: google.cloud.discoveryengine.v1.SiteSearchEngineService
    #   proto-msg: google.cloud.discoveryengine.v1.TargetSite
    #   host-name: ""
    #   notes:
    #     - No gcloud command found

export SERVICE=discoveryengine
export HTTP_HOST=discoveryengine.googleapis.com
export PROTO_PACKAGE=google.cloud.discoveryengine.v1
export PROTO_SERVICE=google.cloud.discoveryengine.v1.SiteSearchEngineService
export CRD_GROUP=${SERVICE}.cnrm.cloud.google.com

export PROTO_MESSAGE=google.cloud.discoveryengine.v1.TargetSite
export PROTO_RESOURCE=TargetSite
export RESOURCE=datastoretargetsite
export CRD_KIND=DiscoveryEngineDataStoreTargetSite
export CRD_VERSION=v1alpha1
./_generate-resource.sh || true
