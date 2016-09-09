#/bin/bash

DOCKER_IMAGE="loxonegoweather"

GIT_COMMIT=`git rev-parse --short HEAD`
CODE_VERSION=`head -n 1 VERSION`

VERSION_COMMIT=`git rev-list ${CODE_VERSION} -n 1 | cut -c1-7`
echo $VERSION_COMMIT

if [[ "${VERSION_COMMIT}" != "${GIT_COMMIT}" ]]; then
    echo "You are trying to push a build based on commit ${GIT_COMMIT} but the tagged release version is ${VERSION_COMMIT}"
    exit
fi

DOCKER_TAG=${CODE_VERSION}-${GIT_COMMIT}

docker build \
      --tag ${DOCKER_IMAGE}:${DOCKER_TAG} \
      --build-arg VCS_REF=${GIT_COMMIT} \
      --build-arg BUILD_DATE=`date -u +"%Y-%m-%dT%H:%M:%SZ"` .
      
