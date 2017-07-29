a, b = gets.split(' ').map(&:to_i)

def possible?(v)
  v / 3 > 0 && v % 3 == 0
end

puts possible?(a) || possible?(b) || possible?(a + b) ? "Possible" : "Impossible"
