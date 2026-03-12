#!/usr/bin/env bash
curl --data '{"url":"http://www.netflix.com/$i"}' http://localhost:8080/api/add/ --header 'Content-Type: application/json' -sS | jq .short
