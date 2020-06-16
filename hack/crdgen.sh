#!/bin/bash
set -eu -o pipefail

cd "$(dirname "$0")/.."

del() {
  yq delete $1 $2 >tmp
  mv tmp "$1"
}

add_header() {
  cat "$1" | ./hack/auto-gen-msg.sh >tmp
  mv tmp "$1"
}

if [ "$(command -v controller-gen)" = "" ]; then
  go install sigs.k8s.io/controller-tools/cmd/controller-gen
fi

if [ "$(command -v yq)" = "" ]; then
  brew install yq
fi

echo "Generating CRDs"
controller-gen crd:trivialVersions=true,maxDescLen=0 paths=./pkg/apis/... output:dir=manifests/base/crds

find manifests/base/crds -name 'argoproj.io*.yaml' | while read -r file; do
  echo "Patching ${file}"
  # remove junk fields
  del "$file" metadata.annotations
  del "$file" metadata.creationTimestamp
  del "$file" spec.validation
  del "$file" status
  add_header "$file"
done
