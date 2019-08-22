#!/bin/bash

BLUE='\E[1;34m'
YELLOW='\E[1;33m'
GREEN='\E[1;32m'
RESET='\E[0m'
RESULT=0

# go files: incomplete comment check
INCOMPLETE_COMMENTS=$(find . -iname "*.go*" | grep -v " " |  xargs grep --colour -H -n -E "^\s*\/\/\s*[A-Z]\w+ \.{3}" 2>/dev/null)
if [[ -n "$INCOMPLETE_COMMENTS" ]]; then
	echo -e "${BLUE}❯ ${YELLOW}WARN: Please fix incomplete exported comments:${RESET}"
	echo -e "${RED}${INCOMPLETE_COMMENTS}${RESET}"
	echo
	RESULT=1
fi


if ! golangci-lint run ./...; then
	exit 1
fi

exit "$RESULT"
