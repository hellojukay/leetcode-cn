#https://leetcode-cn.com/problems/compare-version-numbers/
class Solution:
    def compareVersion(self, version1: str, version2: str) -> int:
        if version1 == version2:
            return 0
        list1 = []
        list2 = []
        for version in   version1.split("."):
            version = version.lstrip("0")
            if len(version) == 0:
                version = "0"
            list1.append(version)
        for version in version2.split("."):
            version = version.lstrip("0")
            if len(version) == 0:
                version = "0"
            list2.append(version)
        length1 = len(list1)
        length2 = len(list2)
        minLength = min(length1,length2)
        for i in range(minLength):
            if int(list1[i]) == int(list2[i]):
                continue
            if int(list1[i]) <  int(list2[i]):
                return -1
            if int(list1[i]) >  int(list2[i]):
                return 1
        if length1 == length2:
            return 0
        if length1 < length2:
            flag = False
            for i in range(length1,length2):
                if int(list2[i]) != 0:
                    flag = True
            if flag:
                return -1
            return 0
        if length2 < length1:
            flag = False
            for i in range(length2,length1):
                if int(list1[i]) != 0:
                    flag = True
            if flag:
                return 1
            return 0
