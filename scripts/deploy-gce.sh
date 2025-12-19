#!/bin/bash
# RobusTest Web - GCE Deployment Script
# Usage: ./deploy-gce.sh <instance-name> <zone> <project>

set -e

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[0;33m'
NC='\033[0m' # No Color

# Configuration
APP_NAME="robustest-web"
REMOTE_DIR="/opt/robustest-web"
SERVICE_NAME="robustest-web"

# Parse arguments
GCE_INSTANCE="${1:-}"
GCE_ZONE="${2:-us-west1-b}"
GCE_PROJECT="${3:-}"

# Validate arguments
if [ -z "$GCE_INSTANCE" ]; then
    echo -e "${RED}Error: Instance name is required${NC}"
    echo "Usage: $0 <instance-name> [zone] [project]"
    echo "Example: $0 prod-instance us-west1-b my-project"
    exit 1
fi

# Build project ID argument
PROJECT_ARG=""
if [ -n "$GCE_PROJECT" ]; then
    PROJECT_ARG="--project=$GCE_PROJECT"
fi

echo -e "${GREEN}========================================${NC}"
echo -e "${GREEN}RobusTest Web - GCE Deployment${NC}"
echo -e "${GREEN}========================================${NC}"
echo ""
echo -e "Instance: ${YELLOW}$GCE_INSTANCE${NC}"
echo -e "Zone:     ${YELLOW}$GCE_ZONE${NC}"
echo -e "Project:  ${YELLOW}${GCE_PROJECT:-default}${NC}"
echo ""

# Get the latest release tarball
DIST_DIR="./dist"
TARBALL=$(ls -t ${DIST_DIR}/${APP_NAME}-*.tar.gz 2>/dev/null | head -1)

if [ -z "$TARBALL" ]; then
    echo -e "${RED}Error: No release tarball found in ${DIST_DIR}${NC}"
    echo "Run 'make release' first"
    exit 1
fi

VERSION=$(basename "$TARBALL" | sed "s/${APP_NAME}-//" | sed 's/.tar.gz//')
echo -e "Version:  ${YELLOW}$VERSION${NC}"
echo ""

# Confirm deployment
read -p "Deploy to $GCE_INSTANCE? (y/n) " -n 1 -r
echo
if [[ ! $REPLY =~ ^[Yy]$ ]]; then
    echo -e "${YELLOW}Deployment cancelled${NC}"
    exit 0
fi

echo ""
echo -e "${GREEN}Step 1: Uploading release tarball...${NC}"
gcloud compute scp "$TARBALL" "${GCE_INSTANCE}:/tmp/${APP_NAME}.tar.gz" \
    --zone="$GCE_ZONE" \
    $PROJECT_ARG

echo ""
echo -e "${GREEN}Step 2: Deploying on remote server...${NC}"
gcloud compute ssh "$GCE_INSTANCE" \
    --zone="$GCE_ZONE" \
    $PROJECT_ARG \
    --command="
        set -e

        echo 'Creating backup...'
        sudo mkdir -p ${REMOTE_DIR}/backups
        if [ -f ${REMOTE_DIR}/${APP_NAME}-linux ]; then
            sudo cp ${REMOTE_DIR}/${APP_NAME}-linux ${REMOTE_DIR}/backups/${APP_NAME}-linux.\$(date +%Y%m%d_%H%M%S)
        fi

        echo 'Extracting new release...'
        sudo mkdir -p ${REMOTE_DIR}
        cd ${REMOTE_DIR}
        sudo tar -xzf /tmp/${APP_NAME}.tar.gz

        echo 'Setting permissions...'
        sudo chmod +x ${REMOTE_DIR}/${APP_NAME}-linux

        echo 'Restarting service...'
        if sudo systemctl is-active --quiet ${SERVICE_NAME}; then
            sudo systemctl restart ${SERVICE_NAME}
        else
            echo 'Service not running, starting...'
            sudo systemctl start ${SERVICE_NAME} || echo 'Note: Service may need to be configured'
        fi

        echo 'Cleaning up...'
        rm -f /tmp/${APP_NAME}.tar.gz

        echo 'Deployment complete!'
    "

echo ""
echo -e "${GREEN}Step 3: Verifying deployment...${NC}"
sleep 3

gcloud compute ssh "$GCE_INSTANCE" \
    --zone="$GCE_ZONE" \
    $PROJECT_ARG \
    --command="
        if curl -s http://localhost:3000/health | grep -q 'healthy'; then
            echo 'Health check: PASSED'
        else
            echo 'Health check: FAILED'
            exit 1
        fi
    "

echo ""
echo -e "${GREEN}========================================${NC}"
echo -e "${GREEN}Deployment successful!${NC}"
echo -e "${GREEN}========================================${NC}"
echo ""
echo -e "Version ${YELLOW}$VERSION${NC} deployed to ${YELLOW}$GCE_INSTANCE${NC}"
