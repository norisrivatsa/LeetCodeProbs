class Solution(object):
    def groupAnagrams(self, strs):
        """
        :type strs: List[str]
        :rtype: List[List[str]]
        """
        result_arr = []
        str_dict = {}
        for i in strs :
            sorted_i = "".join(sorted(i))
            if sorted_i in str_dict :
                str_dict[sorted_i].append(i)
            else :
                str_dict[sorted_i] = [i]
        for i in str_dict :
            result_arr.append(str_dict[i])
        return result_arr
