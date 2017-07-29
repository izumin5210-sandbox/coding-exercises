n = gets.to_i

B, W = ?B, ?W
player_by_color = {}
player_by_color[B] = "Fennec"
player_by_color[W] = "Snuke"

Node = Struct.new(:id, :color)
Edge = Struct.new(:src, :dest)
class Graph
  attr_reader :nodes, :edges

  def initialize(n)
    @nodes = 1.upto(n).map { |i| [i, Node.new(i)] }.to_h
    nodes[1].color = B
    nodes[n].color = W
    @edges = Hash.new { |h, k| h[k] = [] }
  end

  def add_edge(a, b)
    edges[a] << b
    edges[b] << a
  end

  def nodes_by_color(color)
    nodes.values.select { |n| n.color == color }
  end

  def available_nodes_by_nodes(node)
    edges[node.id].map { |i| nodes[i] }.select { |n| n.color.nil? }
  end

  def available_nodes_by_color(color)
    nodes_by_color(color).flat_map do |n|
      available_nodes_by_nodes(n)
    end
  end

  def colorize_by_id(id, color)
    nodes[id].color = color
  end
end

graph = Graph.new(n)

(n - 1).times do
  graph.add_edge(*gets.split(" ").map(&:to_i))
end

[B, W].cycle.each do |c|
  nodes = graph.available_nodes_by_color(c)
  if nodes.empty?
    puts player_by_color[c == B ? W : B]
    break
  end
  id = nodes.sort_by { |n| graph.available_nodes_by_nodes(n).size }.last.id
  graph.colorize_by_id(id, c)
end
