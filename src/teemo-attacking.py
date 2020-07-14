#https://leetcode-cn.com/problems/teemo-attacking/
class Solution:
    def findPoisonedDuration(self, timeSeries: List[int], duration: int) -> int:
        t = 0 # 上次中毒结束时间
        count = 0
        for time in timeSeries:
            if time > t :
                count = count + duration
                t = time + duration
            else:
                count = count + duration - (t - time)
                t = time + duration
        return count
s = Solution()
s.findPoisonedDuration([1,2,3,4,5,6,7,8,9],1)
