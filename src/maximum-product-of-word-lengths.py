from typing import List


class Solution:
    def take(self, s1: str, s2: str) -> int:
        l1 = len(s1)
        l2 = len(s2)
        for c in s1:
            if c in s2:
                return 0
        return l1 * l2

    def maxProduct(self, words: List[str]) -> int:
        m = 0
        if len(words) <= 1:
            return 0
        i = 0
        length = len(words)
        while i < length:
            j = i + 1
            while j < length:
                tmp = self.take(words[i], words[j])
                print(words[i], words[j], tmp)
                j = j + 1
                if tmp > m:
                    m = tmp
            i = i + 1
        return m


s = Solution()
print(s.maxProduct(["eae", "ea", "aaf", "bda",
      "fcf", "dc", "ac", "ce", "cefde", "dabae"]))
