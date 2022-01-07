class Solution:
    def frequencySort(self, s: str) -> str:
        def mysort(word: str, m: dict[str, int]) -> str:
            if len(word) <= 1:
                return word
            if len(word) == 2:
                if m.get(word[1],0) < m.get(word[0],0):
                    return word[1] + word[0]
                return word
            left = ""
            mid = ""
            right = ""
            index = int(len(word) / 2)
            i = 0
            while i < len(word):
                if m.get(word[i], 0) > m.get(word[index], 0):
                    left = left + (word[i])
                    i = i + 1
                    continue
                if m.get(word[i], 0) == m.get(word[index], 0):
                    mid = mid + (word[i])
                else:
                    right = right + (word[i])
                i = i + 1
            return mysort(left, m) + "".join(sorted(mid)) + mysort(right, m)

        m = {}
        for c in s:
            n = m.get(c, 0)
            m[c] = n + 1
        return mysort(s, m)
