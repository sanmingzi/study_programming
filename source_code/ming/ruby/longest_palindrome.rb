# @param {String} s
# @return {String}
def longest_palindrome(s)
    return s if s.length <= 1

    ss = '$'
    s.chars.each do |c|
        ss += c
        ss += '$'
    end

    dp = []
    dp[0] = id = max_right = rtn = 0
    length = ss.length

    for i in (1..(length - 1))
        dp[i] = i < max_right ? [dp[id * 2 - i], max_right - i].min : 0

        while i - dp[i] - 1 >= 0 && i + dp[i] + 1 < length &&  ss[i - dp[i] - 1] == ss[i + dp[i] + 1]
            dp[i] += 1
        end

        if i + dp[i] > max_right
            id = i
            max_right = i + dp[i]
        end

        if dp[i] > dp[rtn]
            rtn = i
        end

        # puts "i: #{i}, dp[i]: #{dp[i]}"
    end

    longest_str = ss[rtn - dp[rtn], 2 * dp[rtn] + 1].delete('$')
    # puts longest_str
    longest_str
end

# longest_palindrome("tattarrattat")
# longest_palindrome('babad')
# longest_palindrome('cbbd')
# longest_palindrome('bb')
# longest_palindrome("babadada")