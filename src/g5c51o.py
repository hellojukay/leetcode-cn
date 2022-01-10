from types import resolve_bases


class Solution:
    def topKFrequent(self, words: List[int], k: int) -> List[int]:
        def mysort(strs: List[int], m: dict[int, int]) -> List[int]:
            if len(strs) <= 1:
                return strs
            if len(strs) == 2:
                if m.get(strs[1], 0) > m.get(strs[0], 0):
                    return [strs[1], strs[0]]
                if m.get(strs[1], 0) == m.get(strs[0], 0):
                    return sorted(strs)
                return strs
            left = []
            mid = []
            right = []
            index = int(len(strs) / 2)
            i = 0
            while i < len(strs):
                if m.get(strs[i], 0) == m.get(strs[index], 0):
                    mid.append(strs[i])
                    i = i + 1
                    continue
                if m.get(strs[i], 0) > m.get(strs[index], 0):
                    left.append(strs[i])
                    i = i + 1
                    continue
                right.append(strs[i])
                i = i + 1
            return mysort(left, m) + sorted(mid) + mysort(right, m)

        m = {}
        for word in words:
            m[word] = m.get(word, 0) + 1
        l = mysort(words, m)
        i = 0
        j = 0
        result = []
        while i < len(l) and i + 1 < len(l):
            if l[i] != l[i + 1]:
                result.append(l[i])
                j = j + 1
            if j == k:
                break
            i = i + 1
        if len(result) < k:
            result.append(l[i])
        return result

