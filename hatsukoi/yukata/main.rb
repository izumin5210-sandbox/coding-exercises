n = gets.to_i
operations = n.times.map { gets.chomp.split(" ") }.map { |k, v| [k.to_i, v] }.to_h

deg = 0

cost = 24.times.inject(0) { |cost, t|
  op = operations[t]

  unless op.nil?
    deg +=
      case op
      when "in" then 5
      when "out" then 3
      end
  end

  if deg > 0
    deg -= 1
    cost += 1
  end

  cost + 1
}

p cost
