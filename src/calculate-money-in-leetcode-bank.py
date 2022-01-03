class Solution:
    def totalMoney(self, n: int) -> int:
        if n == 1:
            return 1
        result = [0, 1]
        i = 2
        while i <= n:
            if i > 7 and i % 7 == 1:
                result.append(result[i - 7] + 1)
            else:
                result.append(result[i - 1] + 1)
            i = i + 1
        return sum(result)
