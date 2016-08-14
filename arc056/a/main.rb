a, b, k, l = gets.split(' ').map(&:to_i)

n_b = k / l
n_a = k - n_b * l

puts [(n_b * b) + (n_a * a), (n_b + 1) * b].min
