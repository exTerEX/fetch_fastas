# Fetch

A tool to fetch FASTA-files from RCSBs database. As of now this is a small side project to create a useful tool to fetch a local copy of the RCSB fasta database. It is a simple tool only able to downloading FASTA-files into a folder, but the goal is to make this a more powerful tool to keep a local database of RCSB data synchronized.

## Example

```sh
$ docker pull ghcr.io/exterex/fetch:latest
```

`docker-compose.yml` example:

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
