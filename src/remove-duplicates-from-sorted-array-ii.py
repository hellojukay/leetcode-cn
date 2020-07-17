class Solution:
    def removeDuplicates(self, nums: List[int]) -> int:
        if not nums:
            return 0
        count = 0
        pre = nums[0]
        index = 0
        for n in nums:
            if n == pre:
                count = count + 1
            else:
                # 归位 pre 
                m = count
                if count > 2:
                    m = 2
                for i in range(m):
                    nums[index] = pre
                    index = index + 1
                count = 1
            pre = n
        if count > 2:
            count = 2
        for i in range(count):
            nums[index] = pre
            index = index + 1
        return index
