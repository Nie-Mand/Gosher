// @generated
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct FileSendRequest {
    #[prost(string, tag="1")]
    pub file_bytes: ::prost::alloc::string::String,
    #[prost(string, tag="2")]
    pub destination: ::prost::alloc::string::String,
}
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct Request {
    #[prost(string, tag="1")]
    pub msg: ::prost::alloc::string::String,
    #[prost(string, tag="2")]
    pub destination: ::prost::alloc::string::String,
}
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct Response {
    #[prost(string, tag="1")]
    pub msg: ::prost::alloc::string::String,
}
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct Identity {
}
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct PingForFileRequest {
    #[prost(string, tag="1")]
    pub description: ::prost::alloc::string::String,
}
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct PingForFileResponse {
    #[prost(string, tag="1")]
    pub who: ::prost::alloc::string::String,
    #[prost(string, tag="2")]
    pub metadata: ::prost::alloc::string::String,
}
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct ListenForFilePingsRequest {
    #[prost(string, tag="1")]
    pub description: ::prost::alloc::string::String,
}
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct ListenForFilePingsResponse {
    #[prost(string, tag="1")]
    pub metadata: ::prost::alloc::string::String,
}
include!("schemas.tonic.rs");
// @@protoc_insertion_point(module)
