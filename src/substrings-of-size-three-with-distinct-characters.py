class Solution:
    def countGoodSubstrings(self, s: str) -> int:
        def check(ss: str) -> bool:
            return len(set(ss)) == 3

        count = 0
        i = 0
        j = i + 2
        while j < len(s):
            if check(s[i : j + 1]):
                count = count + 1
            i = i + 1
            j = i + 2
        return count
