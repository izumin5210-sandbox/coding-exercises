require 'set'

n, k = gets.split(' ').map(&:to_i)

r = {}

n.times do |i|
  r[i] = gets.split(' ').map(&:to_i)
end

deps = []
