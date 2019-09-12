class Solution {
public:
    string longestCommonPrefix(vector<string>& strs) {
        string prefix = "";
        if (strs.size() == 0) {
            return prefix;
        }
        
        int len = strs[0].size();
        for (int i=0;i<strs.size();i++){
            if (len > strs[i].size() ) {
                len = strs[i].size();
            }
        }
        
        
        for (int i=0; i < len; i++) {
            char c = strs[0][i];
            bool same = true;
            for (int j = 0; j < strs.size(); j++) {
                if (c != strs[j][i]) {
                    same = false;
                    break;
                }
            }
            
            if (same) {
                prefix += c;
            } else {
               return prefix;
            }
                
        }
        return prefix;
    }
};