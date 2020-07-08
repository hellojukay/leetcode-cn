class Solution(object):
    def missingNumber(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        # 0到n的和
        sum = 0
        for i in range(len(nums)+1):
            sum = sum + i
        # 因为缺少某个数，所有nums所有数字的肯定比 sum 小，插值就是所有的数
        for i in nums:
            sum = sum -i
        return sum

