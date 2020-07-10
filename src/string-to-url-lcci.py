class Solution:
    def replaceSpaces(self, S: str, length: int) -> str:
        result = ""
        index = 0
        for ch in S:
            if index >= length :
                break
            if ch.isspace():
                result = result + "%20"
            else:
                result = result + ch
            index = index + 1
        return result

