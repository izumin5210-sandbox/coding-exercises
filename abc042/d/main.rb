H, W, A, B = gets.split(" ").map(&:to_i)

=begin
cache = Hash.new { |h, k| h[k] = Hash.new }

def routes(i, j, cache)
  puts "#{i}, #{j}"
  if i == 0 && j == 0
    return 1
  else
    c = 0
    c += cache[i - 1][j] || routes(i - 1, j, cache) if i - 1 >= 0 && j >= 0 && !(i > (H - A - 1) && j < B)
    c += cache[i][j - 1] || routes(i, j - 1, cache) if i >= 0 && j - 1 >= 0 && !(i > (H - A - 1) && j < B)
    cache[i][j] = c
    return c
  end
end

puts routes(H - 1, W - 1, cache)
=end

cache = Hash.new { |h, k| h[k] = Hash.new(0) }
cache[0][0] = 1

H.times do |i|
  (i > (H - A - 1) ? B.upto(W) : W.times).each do |j|
    next if i > (H - A - 1) && j < B
    cache[i][j] += cache[i - 0][j - 1] if i - 0 >= 0 && j - 1 >= 0
    cache[i][j] += cache[i - 1][j - 0] if i - 1 >= 0 && j - 0 >= 0
  end
end

puts cache[H - 1][W - 1]
