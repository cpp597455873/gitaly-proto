=begin
#commit.proto

#No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)

OpenAPI spec version: version not set

Generated by: https://github.com/swagger-api/swagger-codegen.git

=end

require 'spec_helper'
require 'json'
require 'date'

# Unit tests for SwaggerClient::GitalyFindLocalBranchesRequest
# Automatically generated by swagger-codegen (github.com/swagger-api/swagger-codegen)
# Please update as you see appropriate
describe 'GitalyFindLocalBranchesRequest' do
  before do
    # run before each test
    @instance = SwaggerClient::GitalyFindLocalBranchesRequest.new
  end

  after do
    # run after each test
  end

  describe 'test an instance of GitalyFindLocalBranchesRequest' do
    it 'should create an instact of GitalyFindLocalBranchesRequest' do
      expect(@instance).to be_instance_of(SwaggerClient::GitalyFindLocalBranchesRequest)
    end
  end
  describe 'test attribute "repository"' do
    it 'should work' do
       # assertion here. ref: https://www.relishapp.com/rspec/rspec-expectations/docs/built-in-matchers
    end
  end

  describe 'test attribute "sort_by"' do
    it 'should work' do
       # assertion here. ref: https://www.relishapp.com/rspec/rspec-expectations/docs/built-in-matchers
    end
  end

end

