// https://rfdonnelly.github.io/posts/tauri-async-rust-process/
#![cfg_attr(
    all(not(debug_assertions), target_os = "windows"),
    windows_subsystem = "windows"
)]

use tonic::transport::Channel;
mod schemas;
use schemas::gosher_client::GosherClient;


async fn get_client() -> GosherClient<Channel> {
    let channel = Channel::builder("grpc://[::]:50051".parse().unwrap())
        .connect()
        .await
        .unwrap();
    // .accept_compressed(CompressionEncoding::Gzip); ?
    let client = GosherClient::new(channel);
    return client;
}

async fn rpc_say_hi(name: &str) -> String {
    let mut client = get_client().await;
    let request = tonic::Request::new(schemas::Request {
        msg: name.to_string(),
    });

    let response = client.say_hi(request)
        .await
        .unwrap();

    response.into_inner().msg
}

#[tauri::command]
async fn say_hi(name: &str) -> Result<String, ()> {
    let response = rpc_say_hi(name).await;
    return Ok(format!("Response from gRPC Server: {}, FUCK RUST", response));
}

fn main() {
    tauri::Builder::default()
        .invoke_handler(tauri::generate_handler![say_hi])
        .run(tauri::generate_context!())
        .expect("error while running tauri application");
}
