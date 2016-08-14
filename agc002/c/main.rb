N, L = gets.split(" ").map(&:to_i)
ropes = []

ropes << gets.split(" ").map(&:to_i)

history = []

i = 0

loop do
  break if ropes.all? { |r| r.size == 1 || r.inject(0, :+) < L }

  rope = ropes[i]
  j = 1

  loop do
    length = rope.inject(0, :+)
    left_length = rope[0...j].inject(0, :+)
    right_length = rope[j...rope.size].inject(0, :+)

    if length >= L && ((j == 1 && right_length >= L) || (left_length >= L && right_length >= L) || (left_length >= L && j == rope.size - 1) || (rope.size == 2))
      ropes.delete(rope)
      ropes << rope[0...j]
      ropes << rope[j...rope.size] if j < rope.size
      history << ropes[0...i].inject(0) { |s, r| s + r.size } + j
      i = 0
      break
    end

    if j > rope.size - 1
      i += 1
      break
    end

    j += 1
  end
end


if ropes.all? { |r| r.size == 1 }
  puts "Possible"
  puts history 
else
  puts "Impossible"
end

