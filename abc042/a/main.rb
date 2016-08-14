puts gets.split(" ").map(&:to_i).tap { |a| break a.count(5) == 2 && a.count(7) == 1 } ? "YES" : "NO"

