FROM scratch


ADD utron-docs /docs

EXPOSE 8090

CMD ["/docs/untron-docs"]
