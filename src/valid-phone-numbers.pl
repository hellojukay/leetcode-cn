perl -ne 'if (/^\d{3}-\d{3}-\d{4}$/){print;} if(/^\(\d{3}\)\s\d{3}-\d{4}$/){print;}' file.txt
