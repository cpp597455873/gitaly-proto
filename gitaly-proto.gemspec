# coding: utf-8
prefix = 'ruby'
lib = File.expand_path(File.join('..', prefix, 'lib'), __FILE__)
$LOAD_PATH.unshift(lib) unless $LOAD_PATH.include?(lib)
require 'gitaly/version'

Gem::Specification.new do |spec|
  spec.name          = "gitaly-proto"
  spec.version       = Gitaly::VERSION
  spec.authors       = ["Jacob Vosmaer"]
  spec.email         = ["jacob@gitlab.com"]

  spec.summary       = %q{Auto-generated gRPC client for gitaly}
  spec.description   = %q{Auto-generated gRPC client for gitaly.}
  spec.homepage      = "https://gitlab.com/gitlab-org/gitaly-proto"
  spec.license       = "MIT"

  spec.files         = `git ls-files -z #{prefix}`.split("\x0").reject { |f| f.match(%r{^#{prefix}/(test|spec|features)/}) }
  spec.require_paths = ["#{prefix}/lib"]

  spec.add_dependency "google-protobuf", "~> 3.1"
  spec.add_dependency "grpc", "~> 1.10"
end
