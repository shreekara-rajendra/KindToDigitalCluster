#! /bin/bash
source /home/shreekara-rajendra/go/pkg/mod/k8s.io/code-generator@v0.32.2/kube_codegen.sh

BOILERPLATE="/home/shreekara-rajendra/KindToDigitalOcean/hack/boilerplate.go.txt"

kube::codegen::gen_helpers /home/shreekara-rajendra/KindToDigitalOcean/pkg/apis/shreekararajendra.dev/v1alpha1 \
--boilerplate "$BOILERPLATE"

echo "helper completed"

kube::codegen::gen_openapi /home/shreekara-rajendra/KindToDigitalOcean/pkg/apis/shreekararajendra.dev/v1alpha1 \
--boilerplate "$BOILERPLATE" \
--output-dir /home/shreekara-rajendra/KindToDigitalOcean/pkg/client \
--output-pkg github.com/shreekara-rajendra/KindToDigitalOcean/pkg/client \

echo "gen_openapi completed"

kube::codegen::gen_client /home/shreekara-rajendra/KindToDigitalOcean/pkg/apis/ \
--output-dir /home/shreekara-rajendra/KindToDigitalOcean/pkg/client \
--output-pkg github.com/shreekara-rajendra/KindToDigitalOcean/pkg/client \
--boilerplate "$BOILERPLATE" \
--with-watch

echo "gen_client completed"

kube::codegen::gen_register /home/shreekara-rajendra/KindToDigitalOcean/pkg/apis/shreekararajendra.dev/v1alpha1 \
--boilerplate "$BOILERPLATE"

echo "gen_register completed"