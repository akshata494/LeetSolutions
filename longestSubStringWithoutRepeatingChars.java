import java.util.Hashtable;

class Solution {
    public int lengthOfLongestSubstring(String s) {
        char[] sa;
        int i=0, j=0, maxlen=0, curlen=0;

        if(s != null)
            sa = s.toCharArray();
        else
            return 0;

        if(sa.length == 1)
            return 1;

        Hashtable table = new Hashtable(26);
        for(i=0; i<sa.length ; i++) {
            table.clear();
            for(j=i ; j<sa.length; j++)
            {
                if (table.contains(sa[j])) {
                    //we encountered a duplicate should stop parsing and calculate the length
                    //of unique substring encountered so far -> which is the length of the hash table
                    //break from inner loop go to next iteration of outer loop
                    break;
                } else
                    table.put(sa[j], sa[j]);
            }
            if(table.size() > maxlen)
                maxlen = table.size();
        }
        //if s was null, we would have returned 0
        //otherwise maxlen should atleast be 1
        //if it's passed all iterations and yet 0, then all elements were unique
        //which is why it never went in the "contains" block and maxlen remainied 0
        if(maxlen == 0 )
            maxlen = sa.length;
        return maxlen;
    }
}
