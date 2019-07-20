L, R = gets.split(' ').map(&:to_i)

if R - L >= 2019
  puts 0
  exit
end

min = { n1: -1, n2: -1, mod: 2019 }

(L..R).to_a.combination(2) do |(n1, n2)|
  mod = n1 * n2 % 2019
  min = { n1: n1, n2: n2, mod: mod } if min[:mod] > mod
  break if mod == 0
end

puts min[:mod]
