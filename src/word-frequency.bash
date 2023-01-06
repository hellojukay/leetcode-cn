cat words.txt |  tr " " "\n" | sed '/^$/d' | sort | uniq -c  | sort -r | awk '{print $2, $1}'
