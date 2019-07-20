N = gets.to_i
dams = gets.split(' ').map(&:to_i)

# x0/2 + x1/2 = d0
# x1/2 + x2/2 = d1
# x2/2 + x0/2 = d2

# x0/2 + x1/2 = d0
# x1/2 + x2/2 = d1
# x2/2 + x3/2 = d2
# x3/2 + x4/2 = d3
# x4/2 + x0/2 = d4

# x0 = d0 - d1 + d2 - d3 + d4
# x1 = d1 - d2 + d3 - d4 + d0
# x2 = d2 - d3 + d4 - d0 + d1

ms = Array.new(N, 0)

N.times do |i|
  ms[0] += (i.odd? ? -1 : 1) * dams[i]
end

N.times do |i|
  next unless ms[i].zero?

  ms[i] = 2 * (dams[i - 1] - ms[i - 1] / 2)
end

puts ms.join(' ')
