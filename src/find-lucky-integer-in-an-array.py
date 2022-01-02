class Solution:
    def findLucky(self, arr: List[int]) -> int:
        m = {}
        for num in arr:
            n = m.get(num, 0)
            m[num] = n + 1
        count = 0
        result = []
        for key in m.keys():
            if key == m.get(key):
                result.append(key)
        if len(result) > 0:
            x = result[0]
            for n in result:
                if n > x:
                    x = n
            return x
        else:
            return -1
