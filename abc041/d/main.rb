require 'set'
N, M = gets.split(' ').map(&:to_i)

=begin

LIST = []

N.times do |i|
  LIST[i] = N.times.to_a
  LIST[i].delete(i)
end

M.times do
  pair = gets.split(' ').map(&:to_i)
  LIST[pair[1] - 1].delete(pair[0] - 1)
end


def patterns(i, visited)
  result = 0
  if visited.size == N
    result = 1
  else
    LIST[i].each do |j|
      if !visited.include?(j) && visited.all? { |k| LIST[k].include?(j) }
        result += patterns(j, visited + Set.new([j]))
      end
    end
  end
  return result
end

result = 0

LIST.each.with_index do |arr, i|
  result += patterns(i, Set.new([i]))
end

puts result

=end

PAIRS = Hash.new { |h, k| h[k] = [] }

M.times do
  pair = gets.split(' ').map { |i| i.to_i - 1 }
  PAIRS[pair[1]] << pair[0]
end

LIST = N.times.to_a

def count(visited = [])
  if visited.size == N
    return 1
  else
    result = 0
    LIST.each do |i|
      if !visited.include?(i) && visited.all? { |j| !PAIRS[j].include?(i) }
        result += count(visited + [i])
      end
    end
    return result
  end
end

p count
