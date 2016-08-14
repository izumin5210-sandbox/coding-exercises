S = gets.chomp

Key = Struct.new(:label, :color)

keys = [
  Key.new("Do", "W"),
  Key.new("Do", "B"),
  Key.new("Re", "W"),
  Key.new("Re", "B"),
  Key.new("Mi", "W"),
  Key.new("Fa", "W"),
  Key.new("Fa", "B"),
  Key.new("So", "W"),
  Key.new("So", "B"),
  Key.new("La", "W"),
  Key.new("La", "B"),
  Key.new("Si", "W")
]

i = 0

head, sep, tail = S.partition("WW")
case tail.partition("WW")[0].size
when 3
  puts keys[11 - head.size].label
when 5
  puts keys[4 - head.size].label
end
