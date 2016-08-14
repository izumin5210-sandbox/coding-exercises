require "set"

class City
  attr_reader :id, :routes

  def initialize(id)
    @id = id
    @routes = {}
  end

  def add(dest_id, created_at)
    if !@routes.key?(dest_id) || @routes[dest_id] < created_at
      @routes[dest_id] = created_at
    end
  end
end

class Citizen
  attr_reader :id, :city_id, :min_year

  def initialize(id, city_id, min_year)
    @id = id
    @city_id = city_id
    @min_year = min_year
  end
end

n, m = gets.split(' ').map(&:to_i)

cities = {}

m.times do
  a_id, b_id, year = gets.split(' ').map(&:to_i)

  a_i = cities[a_id]
  if a_i.nil?
    a_i = City.new(a_id)
    cities[a_id] = a_i
  end
  b_i = cities[b_id]
  if b_i.nil?
    b_i = City.new(b_id)
    cities[b_id] = b_i
  end

  a_i.add(b_i.id, year)
  b_i.add(a_i.id, year)
end

q = gets.to_i

citizens = []

q.times do |i|
  vi, wi = gets.split(' ').map(&:to_i)

  citizens << Citizen.new(i, vi, wi)
end

def arrivable_cities(cities, city_id, min_year, visited = [])
  city = cities[city_id]
  visited.add(city_id)
  result = Set[city_id]
  city.routes.each do |c_id, y|
    if !visited.include?(c_id) && y > min_year
      result += arrivable_cities(cities, c_id, min_year, visited)
    end
  end
  result
end


citizens.each do |citizen|
  city = cities[citizen.city_id]

  if city.nil?
    puts 1
  else
    puts arrivable_cities(cities, city.id, citizen.min_year, Set.new).size
  end
end

