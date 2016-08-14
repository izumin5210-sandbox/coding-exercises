require 'bigdecimal'

a, b, c = gets.split(' ').map(&:to_i)

i = 10 ** 9 + 7

puts a % i
puts b % i
puts c % i
puts ((((a % i) * (b % i)) % i) * (c % i)) % i
