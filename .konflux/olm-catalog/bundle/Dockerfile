# use non-scratch parent image as dummy for https://issues.redhat.com/browse/KFLUXBUGS-10
FROM registry.redhat.io/ubi8/ubi-micro:latest
FROM scratch

COPY .konflux/olm-catalog/bundle/kodata /kodata
COPY .konflux/olm-catalog/bundle/manifests /manifests
COPY .konflux/olm-catalog/bundle/metadata/annotations.yaml /metadata/annotations.yaml

LABEL operators.operatorframework.io.bundle.mediatype.v1=registry+v1
LABEL operators.operatorframework.io.bundle.manifests.v1=manifests/
LABEL operators.operatorframework.io.bundle.metadata.v1=metadata/
LABEL operators.operatorframework.io.bundle.package.v1=openshift-pipelines-operator-rh
LABEL operators.operatorframework.io.bundle.channel.default.v1="pipelines-main"
LABEL operators.operatorframework.io.bundle.channels.v1="latest,pipelines-5.0"

LABEL \
      com.redhat.component="openshift-pipelines-operator-bundle-container" \
      name="openshift-pipelines/pipelines-operator-bundle-container" \
      version="5.0.5-594" \
      summary="Red Hat OpenShift Pipelines Operator Bundle" \
      maintainer="pipelines-extcomm@redhat.com" \
      description="Red Hat OpenShift Pipelines Operator Bundle" \
      io.k8s.display-name="Red Hat OpenShift Pipelines Operator Bundle" \
      io.k8s.description="Red Hat OpenShift Pipelines Operator Bundle" \
      com.redhat.delivery.operator.bundle=true \
      com.redhat.delivery.backport=false \
      io.openshift.tags="pipelines,tekton,openshift" \
      vendor="Red Hat, Inc." \
      distribution-scope="public" \
      url="https://access.redhat.com/containers/#/registry.access.redhat.com/ubi9-minimal/images/9.4-1227.1725849298" \
      release="1227.1725849298"

USER 65532
