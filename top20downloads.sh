grep 'DOWNLOAD|' artifactory-* | awk -F '|' '{ print $6 "|" $7 "|"$5 }' | sort | uniq -c | sort -nr | head -n 20 
