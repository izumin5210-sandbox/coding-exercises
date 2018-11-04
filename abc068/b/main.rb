n = gets.to_i

i = 0

loop do
  break if n < 2 ** i
  i += 1
end

puts 2 ** (i - 1)
