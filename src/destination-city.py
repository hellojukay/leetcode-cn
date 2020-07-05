class Solution:
    def destCity(self, paths: List[List[str]]) -> str:
        dic = {}
        citys = []
        for arr in paths:
            dic[arr[0]] = arr[1]
            citys.append(arr[0])
            citys.append(arr[1])
        for city in citys:
            if dic.get(city,None) == None:
                return city
        return None
