#!/bin/bash

set -exo pipefail

curl -H 'Cache-Control: no-cache' https://raw.githubusercontent.com/fossas/fossa-cli/v1.0.22/install.sh | bash -s -- -b ~/ v1.0.22

~/fossa init
~/fossa analyze

# Capture the exit status
EXIT_STATUS=$?

echo "fossa script exits with status $EXIT_STATUS"
exit $EXIT_STATUS