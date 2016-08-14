n = gets.to_i
a = gets.split(' ').map(&:to_i)

i = 0

fixed = [0]
costs = {}
costs[0] = 0
nearest = {}

visited = {}

loop do
  c = a[i]

  if i + 1 < n
    c1 = (c - a[i + 1]).abs
    if !costs.key?(i + 1) || (c1 + costs[i]) < costs[i + 1]
      costs[i + 1] = c1 + costs[i]
      nearest[i + 1] = i
      visited[i + 1] = costs[i + 1]
    end
  end

  if i + 2 < n
    c2 = (c - a[i + 2]).abs
    if !costs.key?(i + 2) || (c2 + costs[i]) < costs[i + 2]
      costs[i + 2] = c2 + costs[i]
      nearest[i + 2] = i
      visited[i + 2] = costs[i + 2]
    end
  end

  break if visited.empty?

  i = visited.min { |(k1, v1), (k2, v2)| v1 <=> v2 }[0]

  fixed << i
  visited.delete(i)
end

puts costs[n - 1]
