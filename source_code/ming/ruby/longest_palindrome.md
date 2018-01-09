# @param {String} s
# @return {String}
def longest_palindrome(s)
    length = s.length
    return s if length <= 1

    longest = 1
    rtn = s[0]
    dp = continu = []
    dp[0] = continu[0] = 1

    for i in 1..(length - 1)
        dp[i] = continu[i] = 1

        if s[i] == s[i - 1]
            continu[i] = continu[i - 1] + 1
        end

        before_index = i - dp[i -1] - 1
        if before_index >= 0 and s[i] == s[before_index]
            dp[i] = dp[i - 1] + 2
        end

        puts continu[i]
        puts dp[i]

        if continu[i] > longest
            longest = continu[i]
            rtn = s[i - longest, longest]
        elsif dp[i] > longest
            longest = dp[i]
            rtn = s[i - longest, longest]
        end
    end

    rtn
end

longest_palindrome('ababd')