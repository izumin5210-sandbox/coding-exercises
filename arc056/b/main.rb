require 'set'

n, m, s = gets.split(' ').map(&:to_i)

routes = Hash.new { |h, k| h[k] = [] }

m.times do
  u, v = gets.split(' ').map(&:to_i)
  routes[u] << v
  routes[v] << u
end

used = Set.new
users = []

1.upto(n) do |n_i|
  i = s
  visited = Set.new
  candidates = Set.new
  candidates.add(i) unless used.include?(i)

  until candidates.empty?
    visited.merge(candidates)

    candidates = candidates.inject(Set.new) { |s, c| s.merge(routes[c].select { |j| !visited.include?(j) && !used.include?(j) }) }
  end

  if visited.include?(n_i)
    users << n_i
    used.add(n_i)
  end
end

puts users
