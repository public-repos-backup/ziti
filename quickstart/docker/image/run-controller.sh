#!/bin/bash
ziti_controller_cfg="${ZITI_HOME}/ziti-edge-controller.yaml"

if [[ "${ZITI_CTRL_EDGE_ADVERTISED_ADDRESS-}" == "" ]]; then export ZITI_CTRL_EDGE_ADVERTISED_ADDRESS="ziti-edge-controller"; fi
if [[ "${ZITI_CTRL_NAME-}" == "" ]]; then export ZITI_CTRL_NAME="${ZITI_NETWORK}-controller"; fi

. "${ZITI_SCRIPTS}/ziti-cli-functions.sh"

if [ ! -f "${ZITI_HOME}/access-control.init" ]; then
  setupEnvironment
  persistEnvironmentValues
fi
. ${ZITI_HOME}/ziti.env

# create pki
createPki
if [ ! -f "${ziti_controller_cfg}" ]; then
  echo " "
  echo "${ziti_controller_cfg} doesn't exist. Generating config file"
  echo " "
  # generates the config file for the controller
  createControllerConfig
else
  echo " "
  echo "${ziti_controller_cfg} exists. Not overwriting"
  echo " "
fi

# initialize the database with the admin user:
"${ZITI_BIN_DIR}/ziti" controller edge init "${ZITI_HOME}/${ZITI_CTRL_NAME}.yaml" -u "${ZITI_USER}" -p "${ZITI_PWD}"

echo "controller initialized. unsetting ZITI_USER/ZITI_PWD from env"
unset ZITI_USER
unset ZITI_PWD

# create a place for the internal db
mkdir -p $ZITI_HOME/db

"${ZITI_BIN_DIR}/ziti" controller run "${ZITI_HOME}/${ZITI_CTRL_NAME}.yaml"
