def big_random(digits = 100)
  rtn = ''
  digits.times { rtn += "#{rand(10)}" }
  p rtn
  rtn
end

def add_string(str1, str2)
  str1.to_i + str2.to_i
end

def run
  r1 = big_random
  r2 = big_random
  rtn = add_string(r1, r2)
  p rtn
end

run
