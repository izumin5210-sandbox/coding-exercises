require "set"

N, M = gets.split(" ").map(&:to_i)
graph = Hash.new { |h, k| h[k] = {} }


M.times do |i|
  a, b = gets.split(" ").map(&:to_i)
  graph[a][b] = i + 1
  graph[b][a] = i + 1
end

Q = gets.to_i
pairs = Q.times.map { gets.split(" ").map(&:to_i) }

pairs.each do |(x, y, z)|
  visited = Set.new
  visited.add(x)
  visited.add(y)
  score = -1
  ca = Hash.new { |h, k| h[k] = [] }
  until visited.size >= z
    ca[:x] += graph[x].sort_by { |k, v| v }
    ca[:y] += graph[y].sort_by { |k, v| v }

    ca[:x] = ca[:x].reject { |(k, v)| visited.include?(k) }.sort_by { |(k, v)| v }.uniq { |(k, v)| k }
    ca[:y] = ca[:y].reject { |(k, v)| visited.include?(k) }.sort_by { |(k, v)| v }.uniq { |(k, v)| k }

    x_min = ca[:x][0]
    y_min = ca[:y][0]

    if y_min.nil? || x_min[1] <= y_min[1]
      x = x_min[0]
      score = [x_min[1], score].max
      visited.add(x)
    elsif x_min.nil? || y_min[1] <= y_min[1]
      y = y_min[0]
      score = [y_min[1], score].max
      visited.add(y)
    end
  end

  puts score
end


