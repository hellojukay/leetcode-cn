#https://leetcode-cn.com/problems/long-pressed-name/
class Solution:
    def isLongPressedName(self, name: str, typed: str) -> bool:
        nameList = []
        pre = ''
        count = 0
        for ch in name:
            if ch == pre:
                count = count + 1
                pre = ch
            else:
                nameList.append((pre,count))
                count = 1
                pre = ch
        nameList.append((ch,count))
        pre = ''
        count = 0
        typedList = []
        for ch in typed:
            if ch == pre:
                count = count + 1
                pre = ch
            else:
                typedList.append((pre,count))
                count = 1
                pre = ch
        typedList.append((pre,count))
        if len(nameList) != len(typedList):
            return False
        for i in range(len(typedList)):
            typeCh,typeCount = typedList[i]
            nameCh,nameCount = nameList[i]
            if typeCh != nameCh:
                return False
            if typeCount < nameCount:
                return False
        return True

