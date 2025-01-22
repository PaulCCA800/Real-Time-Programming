use std::net::UdpSocket;
use std::borrow::Cow;
use std::io::{Read,Write};
use std::net::{TcpListener, TcpStream};
use std::thread::spawn; 

fn handle_clinet(mut stream : TcpStream){
    let mut buffer = [0; 1024];
    stream.read(&mut buffer).expect("Failed to read");

    let request = String::from_utf8_lossy(&buffer[..]);
    println!("Recived request: {}", request);

    let response = "Hello, Client".as_bytes();
    stream.write(response).expect("Failed to wirte");

}


pub fn tcp_listener(){

    let listener:TcpListener= TcpListener::bind("localIP").expect("Failed to bind to port");
    println!("Listening on port");

    for stream in listener.incoming(){
    match stream {
    Ok(stream) => {
    std::thread::spawn(||handle_clinet(stream));
    }Err(e) => {
    eprintln!("Failed to establish connection: {}", e)
}
}
}

}