=begin
#commit.proto

#No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)

OpenAPI spec version: version not set

Generated by: https://github.com/swagger-api/swagger-codegen.git

=end

require 'spec_helper'
require 'json'

# Unit tests for SwaggerClient::DiffApi
# Automatically generated by swagger-codegen (github.com/swagger-api/swagger-codegen)
# Please update as you see appropriate
describe 'DiffApi' do
  before do
    # run before each test
    @instance = SwaggerClient::DiffApi.new
  end

  after do
    # run after each test
  end

  describe 'test an instance of DiffApi' do
    it 'should create an instact of DiffApi' do
      expect(@instance).to be_instance_of(SwaggerClient::DiffApi)
    end
  end

  # unit tests for commit_diff
  # Returns stream of CommitDiffResponse: 1 per changed file
  # 
  # @param body 
  # @param [Hash] opts the optional parameters
  # @return [GitalyCommitDiffResponse]
  describe 'commit_diff test' do
    it "should work" do
      # assertion here. ref: https://www.relishapp.com/rspec/rspec-expectations/docs/built-in-matchers
    end
  end

end
