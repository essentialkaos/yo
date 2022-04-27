#!/bin/bash

################################################################################

NORM=0
BOLD=1
UNLN=4
RED=31
GREEN=32
BROWN=33
BLUE=34
MAG=35
CYAN=36
GREY=37
DARK=90

CL_NORM="\e[${NORM}m"
CL_BOLD="\e[${BOLD}m"
CL_UNLN="\e[${UNLN}m"
CL_RED="\e[${RED}m"
CL_GREEN="\e[${GREEN}m"
CL_BROWN="\e[${BROWN}m"
CL_BLUE="\e[${BLUE}m"
CL_MAG="\e[${MAG}m"
CL_CYAN="\e[${CYAN}m"
CL_GREY="\e[${GREY}m"
CL_DARK="\e[${DARK}m"

################################################################################

TEST_DATA=".scripts/data.yaml"
BINARY="yo"

################################################################################

has_errors=""

################################################################################

main() {
  if [[ ! -e $BINARY ]] ; then
    echo -e "${CL_RED}Can't find yo binary${CL_NORM}"
    exit 1
  fi

  runTest
}

runTest() {
  header "Basic selectors"

  check ".name"    "John Doe"
  check ".age"     "35"
  check ".balance" "45.89"
  check ".admin"   "true"

  header "Map selectors"

  check ".meta.uid" "120"
  check ".meta.gid" "350"

  header "Array selectors"

  check ".categories[0]" "category1"
  check ".categories[0:1]" "category1"
  check ".categories[:1]" "category1"
  check ".categories[]" "category1 category2"
  check ".categories[0,1]" "category1 category2"
  check ".categories[1,0]" "category2 category1"
  check ".categories[0,1,2,3,4]" "category1 category2"
  check ".categories[:]" "category1 category2"
  check ".categories[0:2]" "category1 category2"
  check ".array2" "- file: test1   size: 100 - file: test2   size: 200"
  check ".array2[0].file" "test1"

  header "Processors"

  check ".name | length"        "8"
  check ".categories | length"  "2"
  check ".meta | length"        "2"
  check ".array2 | length"      "2"
  check ".array2[] | length"    "2 2"
  check ".meta | keys | length" "2"
  check ".meta | keys | sort | length" "2"
  check ".meta | keys | sort"   "gid uid"

  echo -e "$CL_NORM"

  if [[ -n $has_errors ]] ; then
    exit 1
  fi
}

check() {
  local query="$1"
  local result="$2"
  local output

  output=$($BINARY -f $TEST_DATA "$query" | tr '\n' ' ' | sed 's/ $//')

  if [[ "$result" == "$output" ]] ; then
    echo -e "${CL_GREEN}✓ ${CL_NORM}${query}${CL_DARK} → \"$output\"${CL_NORM}"
  else
    echo -e "${CL_RED}✕ ${CL_NORM}${query}"
    echo -e "${CL_GREY}  \"$output\" ≠ \"$result\"${CL_NORM}"
    has_errors=true
  fi
}

header() {
  echo -e "\n${CL_BOLD}▾ ${1}${CL_NORM}\n"
}

################################################################################

main "$@"
