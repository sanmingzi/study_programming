# @param {String} s
# @param {Integer} num_rows
# @return {String}
def convert(s, num_rows)
  length = s.length
  return s if length <= 1

  rtn = []
  (0..(num_rows - 1)).each { |i| rtn[i] = [] }

  index = 0
  while index < length
    row = 0
    while index < length && row < num_rows
      rtn[row] << s[index]
      index += 1
      row += 1
    end

    row = num_rows - 2
    while index < length && row > 0
      rtn[row] << s[index]
      index += 1
      row -= 1
    end
  end

  rtn_str = ''
  (0..(num_rows - 1)).each do |i|
    rtn[i].each { |c| rtn_str += c }
  end
  # puts rtn_str
  rtn_str
end

# convert("PAYPALISHIRING", 3)