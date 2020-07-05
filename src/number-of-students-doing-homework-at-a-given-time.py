class Solution:
    def busyStudent(self, startTime: List[int], endTime: List[int], queryTime: int) -> int:
        result = 0
        for index in  range(len(startTime)):
            if queryTime >= startTime[index] and queryTime <= endTime[index]:
                result = result + 1
        return result
