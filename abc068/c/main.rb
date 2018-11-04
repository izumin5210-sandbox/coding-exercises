require 'set'

n , m = gets.split(" ").map(&:to_i)

map = Hash.new { |h, k| h[k] = [] }

m.times do
  a , b = gets.split(" ").map(&:to_i)
  map[a] << b
  map[b] << a
end

q = Set.new([*1..n])
d = {}
d[1] = 0
routes = {}

until q.empty?
  u = nil
  d.keys.each do |k|
    u = k if q.include?(k) && (u.nil? || u < d[k])
  end
  q.delete(u)

  break if u.nil?

  map[u].each do |v|
    if d.key?(u) || (d[u] + 1 < d[v])
      d[v] = d[u] + 1
      routes[v] = u
    end
  end
end

prev = 1
cnt = 0

while cnt < 2
  prev = routes[prev]
  cnt += 1
  break if prev == n
end

puts prev == n ? "POSSIBLE" : "IMPOSSIBLE"
