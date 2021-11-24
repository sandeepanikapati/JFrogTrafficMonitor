cat artifactory*.log | awk -F '|' '{ sum += $7 }; END { print sum }'


