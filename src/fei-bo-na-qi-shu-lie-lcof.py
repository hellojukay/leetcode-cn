#https://leetcode-cn.com/problems/fei-bo-na-qi-shu-lie-lcof/
class Solution:
    def fib(self, n: int) -> int:
        arr = []
        for i in range(n+1):
            if i <= 1:
                arr.append(i)
                continue
            if i == 2:
                arr.append(1)
                continue
            arr.append(arr[i-1] + arr[i-2])
        return arr[-1] % 1000000007


