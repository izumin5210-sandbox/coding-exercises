n = gets.to_i

s = Math.sqrt(n)

if s.integer?
  puts s
else
  s = s.to_i
end

