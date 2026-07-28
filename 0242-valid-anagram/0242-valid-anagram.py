class Solution(object):
    def isAnagram(self, s, t):
        """
        :type s: str
        :type t: str
        :rtype: bool
        """
        s_dict = {}
        for i in range(len(s)):
            if s[i] in s_dict:
                s_dict[s[i]] = s_dict[s[i]] + 1
            else :
                s_dict[s[i]] = 1
        t_dict = {}
        for i in range(len(t)):
            if t[i] in t_dict:
                t_dict[t[i]] = t_dict[t[i]] + 1
            else :
                t_dict[t[i]] = 1
    
        if len(s) != len(t):
            return False
        for i in s_dict:
            if i not in t_dict:
                return False

            if s_dict[i] != t_dict[i]:
                return False

        return True      