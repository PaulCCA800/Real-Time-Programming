use std::net::UdpSocket;
use std::borrow::Cow;

mod tcp; 
mod udp;

fn main(){

    udp::tcp_listener();

}
