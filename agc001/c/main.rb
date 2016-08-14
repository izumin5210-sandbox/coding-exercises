require 'set'

N, K = gets.split(" ").map(&:to_i)

TREE = Hash.new { |h, k| h[k] = [] }

(N-1).times do
  a, b = gets.split(" ").map(&:to_i)
  TREE[a - 1] << b - 1
  TREE[b - 1] << a - 1
end

@excluded = Set.new

def depth(i, depth = 0, visited = [])
  return [i] if visited.include?(i)
  visited << i
  paths = [[i]]
  TREE[i].each do |j|
    next if visited.include?(j) || @excluded.include?(j)
    paths |= depth(j, depth + 1, visited).map { |path| [i, *path] }
  end
  @excluded.add(i) if depth > K
  paths
end

paths = Set.new

N.times do |i|
  paths += depth(i)
end

paths.each do |path|
  p path if path.size > K + 1
end

p @excluded

