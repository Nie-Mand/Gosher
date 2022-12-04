// @generated
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct Request {
    #[prost(string, tag="1")]
    pub msg: ::prost::alloc::string::String,
    #[prost(string, tag="2")]
    pub destination: ::prost::alloc::string::String,
}
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct Identity {
    #[prost(string, tag="1")]
    pub id: ::prost::alloc::string::String,
}
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct Response {
    #[prost(string, tag="1")]
    pub msg: ::prost::alloc::string::String,
}
include!("schemas.tonic.rs");
// @@protoc_insertion_point(module)
