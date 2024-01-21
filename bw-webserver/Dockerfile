# step === base
FROM oven/bun:1 as base
WORKDIR /usr/src/app

# step === install
FROM base as install
RUN mkdir /tmp/dev
COPY package.json /tmp/dev/
RUN cd /tmp/dev && bun install

# # install with --production (exclude devDependencies)
# RUN mkdir -p /temp/prod
# COPY package.json bun.lockb /temp/prod/
# RUN cd /temp/prod && bun install --frozen-lockfile --production

# copy node_modules from temp directory
# then copy all (non-ignored) project files into the image
FROM base AS prerelease
COPY --from=install /tmp/dev/node_modules node_modules
COPY . .

# [optional] tests & build
ENV NODE_ENV=development
# RUN bun test
RUN bun run build

# # copy production dependencies and source code into final image
# FROM base AS release
# COPY --from=install /temp/prod/node_modules node_modules
# COPY --from=prerelease /usr/src/app/index.ts .
# COPY --from=prerelease /usr/src/app/package.json .

# run the app
USER bun
EXPOSE 3000/tcp