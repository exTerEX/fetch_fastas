# Fetch fasta RCSB

# Example

```sh
$ docker pull <TO-BE-DONE>
```

docker-compose.yml

```yaml
version: "3"
services:
  fetchfastas:
    container_name: fetch
    image: fetch
    volumes:
      - ./fasta:/go/src/app/fasta
```

## License

This project is licensed under `EUPL-1.2`. See [LICENSE](LICENSE) for more information.
