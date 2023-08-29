#!/bin/bash
RED='\033[0;31m'
GREEN='\033[0;32m'
NC='\033[0m'

for ((i=1;i<=20;i++))
do
math1=$(./linear-stats)
math2=$(./lc_c)

echo -e "---Linear-Stats---"
echo -e "$math1\n"

echo -e "-----my_program------"
echo -e "$math2\n"

if [[ $math1 == $math2 ]]
  then
    printf "${GREEN}------EQUAL------\n\n${NC}"
    else 
    printf "${RED}----NOT-EQUAL----\n\n${NC}"
    break
fi
done
