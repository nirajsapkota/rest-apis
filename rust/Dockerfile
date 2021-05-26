FROM rust:1.52.1

WORKDIR /usr/src/myapp

COPY . .

RUN cargo build --release

RUN cargo install --path .

EXPOSE 8080

CMD [ "/usr/local/cargo/bin/rust" ]