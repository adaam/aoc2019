for i in $(cat input); do ./main $i; done| awk 'a1 += $1 END{ print a1 }'
