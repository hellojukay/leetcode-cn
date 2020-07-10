class Solution:
    def numWays(self, n: int) -> int:
        nums = []
        for i in range(101):
            if i <= 1:
                nums.append(1)
                continue
            if i == 2:
                nums.append(2)
                continue
            nums.append(nums[i-1]+nums[i-2])
        return nums[n] % 1000000007
