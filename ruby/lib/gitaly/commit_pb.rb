# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: commit.proto

require 'google/protobuf'

require 'shared_pb'
require 'google/protobuf/timestamp_pb'
Google::Protobuf::DescriptorPool.generated_pool.build do
  add_message "gitaly.CommitStatsRequest" do
    optional :repository, :message, 1, "gitaly.Repository"
    optional :revision, :bytes, 2
  end
  add_message "gitaly.CommitStatsResponse" do
    optional :oid, :string, 1
    optional :additions, :int32, 2
    optional :deletions, :int32, 3
  end
  add_message "gitaly.CommitIsAncestorRequest" do
    optional :repository, :message, 1, "gitaly.Repository"
    optional :ancestor_id, :string, 2
    optional :child_id, :string, 3
  end
  add_message "gitaly.CommitIsAncestorResponse" do
    optional :value, :bool, 1
  end
  add_message "gitaly.TreeEntryRequest" do
    optional :repository, :message, 1, "gitaly.Repository"
    optional :revision, :bytes, 2
    optional :path, :bytes, 3
    optional :limit, :int64, 4
  end
  add_message "gitaly.TreeEntryResponse" do
    optional :type, :enum, 1, "gitaly.TreeEntryResponse.ObjectType"
    optional :oid, :string, 2
    optional :size, :int64, 3
    optional :mode, :int32, 4
    optional :data, :bytes, 5
  end
  add_enum "gitaly.TreeEntryResponse.ObjectType" do
    value :COMMIT, 0
    value :BLOB, 1
    value :TREE, 2
    value :TAG, 3
  end
  add_message "gitaly.CommitsBetweenRequest" do
    optional :repository, :message, 1, "gitaly.Repository"
    optional :from, :bytes, 2
    optional :to, :bytes, 3
  end
  add_message "gitaly.CommitsBetweenResponse" do
    repeated :commits, :message, 1, "gitaly.GitCommit"
  end
  add_message "gitaly.CountCommitsRequest" do
    optional :repository, :message, 1, "gitaly.Repository"
    optional :revision, :bytes, 2
    optional :after, :message, 3, "google.protobuf.Timestamp"
    optional :before, :message, 4, "google.protobuf.Timestamp"
    optional :path, :bytes, 5
  end
  add_message "gitaly.CountCommitsResponse" do
    optional :count, :int32, 1
  end
  add_message "gitaly.TreeEntry" do
    optional :oid, :string, 1
    optional :root_oid, :string, 2
    optional :path, :bytes, 3
    optional :type, :enum, 4, "gitaly.TreeEntry.EntryType"
    optional :mode, :int32, 5
    optional :commit_oid, :string, 6
    optional :flat_path, :bytes, 7
  end
  add_enum "gitaly.TreeEntry.EntryType" do
    value :BLOB, 0
    value :TREE, 1
    value :COMMIT, 3
  end
  add_message "gitaly.GetTreeEntriesRequest" do
    optional :repository, :message, 1, "gitaly.Repository"
    optional :revision, :bytes, 2
    optional :path, :bytes, 3
  end
  add_message "gitaly.GetTreeEntriesResponse" do
    repeated :entries, :message, 1, "gitaly.TreeEntry"
  end
  add_message "gitaly.ListFilesRequest" do
    optional :repository, :message, 1, "gitaly.Repository"
    optional :revision, :bytes, 2
  end
  add_message "gitaly.ListFilesResponse" do
    repeated :paths, :bytes, 1
  end
  add_message "gitaly.FindCommitRequest" do
    optional :repository, :message, 1, "gitaly.Repository"
    optional :revision, :bytes, 2
  end
  add_message "gitaly.FindCommitResponse" do
    optional :commit, :message, 1, "gitaly.GitCommit"
  end
  add_message "gitaly.ListCommitsByOidRequest" do
    optional :repository, :message, 1, "gitaly.Repository"
    repeated :oid, :string, 2
  end
  add_message "gitaly.ListCommitsByOidResponse" do
    repeated :commits, :message, 1, "gitaly.GitCommit"
  end
  add_message "gitaly.FindAllCommitsRequest" do
    optional :repository, :message, 1, "gitaly.Repository"
    optional :revision, :bytes, 2
    optional :max_count, :int32, 3
    optional :skip, :int32, 4
    optional :order, :enum, 5, "gitaly.FindAllCommitsRequest.Order"
  end
  add_enum "gitaly.FindAllCommitsRequest.Order" do
    value :NONE, 0
    value :TOPO, 1
    value :DATE, 2
  end
  add_message "gitaly.FindAllCommitsResponse" do
    repeated :commits, :message, 1, "gitaly.GitCommit"
  end
  add_message "gitaly.FindCommitsRequest" do
    optional :repository, :message, 1, "gitaly.Repository"
    optional :revision, :bytes, 2
    optional :limit, :int32, 3
    optional :offset, :int32, 4
    repeated :paths, :bytes, 5
    optional :follow, :bool, 6
    optional :skip_merges, :bool, 7
    optional :disable_walk, :bool, 8
    optional :after, :message, 9, "google.protobuf.Timestamp"
    optional :before, :message, 10, "google.protobuf.Timestamp"
  end
  add_message "gitaly.FindCommitsResponse" do
    repeated :commits, :message, 1, "gitaly.GitCommit"
  end
  add_message "gitaly.CommitLanguagesRequest" do
    optional :repository, :message, 1, "gitaly.Repository"
    optional :revision, :bytes, 2
  end
  add_message "gitaly.CommitLanguagesResponse" do
    repeated :languages, :message, 1, "gitaly.CommitLanguagesResponse.Language"
  end
  add_message "gitaly.CommitLanguagesResponse.Language" do
    optional :name, :string, 1
    optional :share, :float, 2
    optional :color, :string, 3
  end
  add_message "gitaly.RawBlameRequest" do
    optional :repository, :message, 1, "gitaly.Repository"
    optional :revision, :bytes, 2
    optional :path, :bytes, 3
  end
  add_message "gitaly.RawBlameResponse" do
    optional :data, :bytes, 1
  end
  add_message "gitaly.LastCommitForPathRequest" do
    optional :repository, :message, 1, "gitaly.Repository"
    optional :revision, :bytes, 2
    optional :path, :bytes, 3
  end
  add_message "gitaly.LastCommitForPathResponse" do
    optional :commit, :message, 1, "gitaly.GitCommit"
  end
  add_message "gitaly.CommitsByMessageRequest" do
    optional :repository, :message, 1, "gitaly.Repository"
    optional :revision, :bytes, 2
    optional :offset, :int32, 3
    optional :limit, :int32, 4
    optional :path, :bytes, 5
    optional :query, :string, 6
  end
  add_message "gitaly.CommitsByMessageResponse" do
    repeated :commits, :message, 1, "gitaly.GitCommit"
  end
  add_message "gitaly.FilterShasWithSignaturesRequest" do
    optional :repository, :message, 1, "gitaly.Repository"
    repeated :shas, :bytes, 2
  end
  add_message "gitaly.FilterShasWithSignaturesResponse" do
    repeated :shas, :bytes, 1
  end
  add_message "gitaly.ListTreeEntriesRequest" do
    optional :repository, :message, 1, "gitaly.Repository"
    optional :commit_blob, :message, 2, "gitaly.ListTreeEntriesRequest.RevisionPath"
    optional :limit, :int64, 3
  end
  add_message "gitaly.ListTreeEntriesRequest.RevisionPath" do
    optional :oid, :bytes, 1
    optional :path, :bytes, 2
  end
  add_message "gitaly.ListTreeEntriesResponse" do
    optional :size, :int64, 1
    optional :data, :bytes, 2
    optional :oid, :string, 3
  end
end

module Gitaly
  CommitStatsRequest = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.CommitStatsRequest").msgclass
  CommitStatsResponse = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.CommitStatsResponse").msgclass
  CommitIsAncestorRequest = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.CommitIsAncestorRequest").msgclass
  CommitIsAncestorResponse = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.CommitIsAncestorResponse").msgclass
  TreeEntryRequest = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.TreeEntryRequest").msgclass
  TreeEntryResponse = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.TreeEntryResponse").msgclass
  TreeEntryResponse::ObjectType = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.TreeEntryResponse.ObjectType").enummodule
  CommitsBetweenRequest = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.CommitsBetweenRequest").msgclass
  CommitsBetweenResponse = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.CommitsBetweenResponse").msgclass
  CountCommitsRequest = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.CountCommitsRequest").msgclass
  CountCommitsResponse = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.CountCommitsResponse").msgclass
  TreeEntry = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.TreeEntry").msgclass
  TreeEntry::EntryType = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.TreeEntry.EntryType").enummodule
  GetTreeEntriesRequest = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.GetTreeEntriesRequest").msgclass
  GetTreeEntriesResponse = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.GetTreeEntriesResponse").msgclass
  ListFilesRequest = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.ListFilesRequest").msgclass
  ListFilesResponse = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.ListFilesResponse").msgclass
  FindCommitRequest = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.FindCommitRequest").msgclass
  FindCommitResponse = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.FindCommitResponse").msgclass
  ListCommitsByOidRequest = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.ListCommitsByOidRequest").msgclass
  ListCommitsByOidResponse = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.ListCommitsByOidResponse").msgclass
  FindAllCommitsRequest = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.FindAllCommitsRequest").msgclass
  FindAllCommitsRequest::Order = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.FindAllCommitsRequest.Order").enummodule
  FindAllCommitsResponse = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.FindAllCommitsResponse").msgclass
  FindCommitsRequest = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.FindCommitsRequest").msgclass
  FindCommitsResponse = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.FindCommitsResponse").msgclass
  CommitLanguagesRequest = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.CommitLanguagesRequest").msgclass
  CommitLanguagesResponse = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.CommitLanguagesResponse").msgclass
  CommitLanguagesResponse::Language = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.CommitLanguagesResponse.Language").msgclass
  RawBlameRequest = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.RawBlameRequest").msgclass
  RawBlameResponse = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.RawBlameResponse").msgclass
  LastCommitForPathRequest = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.LastCommitForPathRequest").msgclass
  LastCommitForPathResponse = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.LastCommitForPathResponse").msgclass
  CommitsByMessageRequest = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.CommitsByMessageRequest").msgclass
  CommitsByMessageResponse = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.CommitsByMessageResponse").msgclass
  FilterShasWithSignaturesRequest = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.FilterShasWithSignaturesRequest").msgclass
  FilterShasWithSignaturesResponse = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.FilterShasWithSignaturesResponse").msgclass
  ListTreeEntriesRequest = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.ListTreeEntriesRequest").msgclass
  ListTreeEntriesRequest::RevisionPath = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.ListTreeEntriesRequest.RevisionPath").msgclass
  ListTreeEntriesResponse = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.ListTreeEntriesResponse").msgclass
end
