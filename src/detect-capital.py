import re
class Solution:
    def detectCapitalUse(self, word: str) -> bool:
        if re.match("^[A-Z]*$",word):
            return True
        if re.match("^[A-Z]?[a-z]*$",word):
            return True
        return False
