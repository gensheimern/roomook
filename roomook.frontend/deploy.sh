#!/usr/bin/env bash

ACTION="update"

POSITIONAL=()
while [[ $# -gt 0 ]]
    do
    key="$1"

    case $key in
        -v|--version)
        VERSION="$2"
        shift # past argument
        shift # past value
        ;;
        -s|--scale)
        SCALE="$2"
        shift
        shift
        ;;
        -c|--create)
        ACTION="create"
        shift
        ;;
        -e|--env)
        ENVIRONMENT="$2"
        shift
        shift
        ;;
        -p|--profile)
        PROFILE="--profile=$2"
        shift
        shift
        ;;
        *)    # unknown option
        echo "Please add at least a version (--version|-v) to deploy"
        POSITIONAL+=("$1") # save it in an array for later
        shift # past argument
        ;;
    esac
done
set -- "${POSITIONAL[@]}" # restore positional parameters

echo ${ACTION}

IMAGE_NAME="312583731641.dkr.ecr.eu-central-1.amazonaws.com/roomook/roomook-frontend"
DOCKER_IMAGE="${IMAGE_NAME}:${VERSION}"

echo "Deploying Version ${VERSION}"

sed <deployment/${ENVIRONMENT}/task.json.template -e "s,@VERSION@,${VERSION}," -e "s,@IMAGE@,${IMAGE_NAME}," -e "s,@ENV@,${ENVIRONMENT}," >deployment/${ENVIRONMENT}/last/task-definition-generated.json
aws ${PROFILE} ecs register-task-definition --region=eu-central-1 --cli-input-json file://deployment/${ENVIRONMENT}/last/task-definition-generated.json > deployment/${ENVIRONMENT}/last/task-definition-output.json

if [ ${ENVIRONMENT} == "prod" ]; then
    TARGETGROUP_ARN=$(aws ${PROFILE} elbv2 describe-target-groups --names shop-product-manager-fe-prod --query 'TargetGroups[0].TargetGroupArn')
else
    TARGETGROUP_ARN=$(aws ${PROFILE} elbv2 describe-target-groups --names shop-product-manager-fe-dev --query 'TargetGroups[0].TargetGroupArn')
fi


TASKDEFINITION_ARN=$( < deployment/${ENVIRONMENT}/last/task-definition-output.json jq .taskDefinition.taskDefinitionArn )

sed -e "s,@TASKDEFINITION_ARN@,$TASKDEFINITION_ARN," -e "s,@SCALE@,${SCALE}," -e "s,@TARGETGROUP_ARN@,${TARGETGROUP_ARN},"  -e "s,@ENV@,${ENVIRONMENT}," <deployment/${ENVIRONMENT}/service-${ACTION}.json.template >deployment/${ENVIRONMENT}/last/service-definition-generated.json

aws ${PROFILE} ecs ${ACTION}-service --region=eu-central-1 --cli-input-json file://deployment/${ENVIRONMENT}/last/service-definition-generated.json | tee deployment/${ENVIRONMENT}/last/service-definition-output.json
