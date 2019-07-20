N, D = gets.split(' ').map(&:to_i)
pts = N.times.map { gets.split(' ').map(&:to_i) }

cnt = 0

pts.combination(2) do |p1, p2|
  d = Math.sqrt(D.times.inject(0) { |s, i| s + ((p1[i] - p2[i]) ** 2) })
  cnt += 1 if d == d.floor
end

puts cnt
