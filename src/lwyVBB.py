from typing import List


class Solution:
    def isAlienSorted(self, words: List[str], order: str) -> bool:
        m = {}
        i = 0
        while i < len(order):
            ch = order[i]
            m[ch] = i
            i = i + 1
        length = len(words)
        if length <= 1:
            return True

        def compare(s1: str, s2: str) -> int:
            length = min(len(s1), len(s2))
            i = 0
            while i < length:
                c1 = s1[i]
                c2 = s2[i]
                if c1 == c2:
                    i = i + 1
                    continue
                n1 = m.get(c1)
                n2 = m.get(c2)
                return n1-n2
            if len(s1) > len(s2):
                return 1
            else:
                return -1
        i = 1
        while i < length:
            if compare(words[i-1], words[i]) <= 0:
                i = i + 1
                continue
            return False
        return True


s = Solution()
print(s.isAlienSorted(["kuvp", "q"], "ngxlkthsjuoqcpavbfdermiywz"))
